package logger

import (
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewLogger(logger *zap.Logger, config ...LoggerConfig) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Get timezone location
	tz, err := time.LoadLocation(cfg.TimeZone)
	if err != nil || tz == nil {
		cfg.timeZoneLocation = time.Local
	} else {
		cfg.timeZoneLocation = tz
	}

	// Create correct timeformat
	var timestamp atomic.Value
	timestamp.Store(time.Now().In(cfg.timeZoneLocation).Format(cfg.TimeFormat))

	// Update date/time every 750 milliseconds in a separate go routine
	go func() {
		for {
			time.Sleep(cfg.TimeInterval)
			timestamp.Store(time.Now().In(cfg.timeZoneLocation).Format(cfg.TimeFormat))
		}
	}()

	// Set PID once
	pid := strconv.Itoa(os.Getpid())

	// Set variables
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Set error handler once
		once.Do(func() {
			// override error handler
			errHandler = c.App().ErrorHandler
		})

		var start, stop time.Time

		// Set latency start time
		if cfg.enableLatency {
			start = time.Now()
		}

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		chainErrStr := "-"
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
				chainErrStr = chainErr.Error()
			}
		}

		// Set latency stop time
		if cfg.enableLatency {
			stop = time.Now()
		}

		// Extract user context
		requestIdStr := "-"
		if ctx := c.UserContext(); ctx != nil {
			if requestId := ctx.Value("requestid"); requestId != nil {
				requestIdStr = requestId.(string)
			}
		}


		// Extract request body
		requestBody := "-"
		if request := c.Request(); request != nil {
			requestBody = string(c.Request().Body())
		}

		// Extract response body
		responseBody := "-"
		if response := c.Response(); response != nil {
			responseBody = string(c.Response().Body())
		}

		// Prepare fields
		fields := []zap.Field {
			zap.Time("timestamp", timestamp.Load().(time.Time)),
			zap.Duration("latency",  stop.Sub(start).Round(time.Millisecond)),
			zap.String("hostname", c.Hostname()),
			zap.String("ip", c.IP()),
			zap.Int("status_code", c.Response().StatusCode()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("error", chainErrStr),
			zap.String("url", c.OriginalURL()),
			zap.String("user_agent", c.Get(fiber.HeaderUserAgent)),
			zap.String("pid", pid),
			zap.String("request_id", requestIdStr),
			zap.String("request_body", requestBody),
			zap.String("response_body", responseBody),
		}

		n := c.Response().StatusCode()
		switch {
		case n >= 500:
			logger.With(zap.Error(err)).Error("Server error", fields...)
		case n >= 400:
			logger.With(zap.Error(err)).Warn("Client error", fields...)
		case n >= 300:
			logger.Info("Redirection", fields...)
		default:
			logger.Info("Success", fields...)
		}

		// End chain
		return nil
	}
}

// LoggerConfig defines the config for middleware.
type LoggerConfig struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// TimeFormat https://programming.guide/go/format-parse-string-time-date-example.html
	//
	// Optional. Default: 15:04:05
	TimeFormat string

	// TimeZone can be specified, such as "UTC" and "America/New_York" and "Asia/Chongqing", etc
	//
	// Optional. Default: "Local"
	TimeZone string

	// TimeInterval is the delay before the timestamp is updated
	//
	// Optional. Default: 500 * time.Millisecond
	TimeInterval time.Duration

	enableLatency    bool
	timeZoneLocation *time.Location
}

// ConfigDefault is the default config
var ConfigDefault = LoggerConfig{
	Next:         nil,
	TimeFormat:   "15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
}

// Helper function to set default values
func configDefault(config ...LoggerConfig) LoggerConfig {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.TimeZone == "" {
		cfg.TimeZone = ConfigDefault.TimeZone
	}
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = ConfigDefault.TimeFormat
	}
	if int(cfg.TimeInterval) <= 0 {
		cfg.TimeInterval = ConfigDefault.TimeInterval
	}
	return cfg
}
