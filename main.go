package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var timestamp atomic.Value
	timestamp.Store(time.Now().In(time.Local).Format(time.RFC3339))

	var t = timestamp.Load().(string)
	fmt.Printf("%T %v\n", t,t)
	d, _ := time.Parse(time.RFC3339, t)

	fmt.Printf("%T %v\n", d,d)


	
}