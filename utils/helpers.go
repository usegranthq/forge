package utils

import "fmt"

// safeGo abstracts panic recovery for any goroutine
func SafeRoutine(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// Safely recover from any panic and log it
				fmt.Println("Recovered from panic:", r)
			}
		}()

		// Execute the provided function
		fn()
	}()
}
