// Final Commit Done

package my_context

import (
	"context"
	"fmt"
	"time"
)

func StartContextDemo() {
	fmt.Println("\nInside Context Module")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	value := 1
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println(value)
			value++
		case <-ctx.Done():
			return
		}
	}
}
