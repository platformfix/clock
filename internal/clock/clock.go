// Package clock prints the current time, once a second, until told to stop.
package clock

import (
	"context"
	"fmt"
	"io"
	"time"
)

// Run writes the current time to w, formatted as RFC3339, immediately and
// then once every interval, until ctx is cancelled.
func Run(ctx context.Context, w io.Writer, interval time.Duration) {
	fmt.Fprintln(w, time.Now().UTC().Format(time.RFC3339))

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ticker.C:
			fmt.Fprintln(w, t.UTC().Format(time.RFC3339))
		}
	}
}
