package usercontext

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		var ctx context.Context

		if c.UserContext() == nil {
			ctx = context.Background()
		} else {
			ctx = c.UserContext()
		}

		if c.Locals("requestid") != nil {
			ctx = context.WithValue(ctx, "requestid", c.Locals("requestid"))
		}

		c.SetUserContext(ctx)

		// Continue stack
		return c.Next()
	}
}

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:       nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg
}

