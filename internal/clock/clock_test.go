package clock

import (
	"bufio"
	"bytes"
	"context"
	"strings"
	"testing"
	"time"
)

func TestRunPrintsTimestampsUntilCancelled(t *testing.T) {
	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		Run(ctx, &buf, 10*time.Millisecond)
		close(done)
	}()

	time.Sleep(55 * time.Millisecond)
	cancel()
	<-done

	scanner := bufio.NewScanner(strings.NewReader(buf.String()))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 3 {
		t.Fatalf("expected at least 3 timestamp lines in 55ms at a 10ms interval, got %d: %v", len(lines), lines)
	}

	for _, line := range lines {
		if _, err := time.Parse(time.RFC3339, line); err != nil {
			t.Fatalf("line %q is not valid RFC3339: %v", line, err)
		}
	}
}

func TestRunStopsWhenContextCancelled(t *testing.T) {
	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := make(chan struct{})
	go func() {
		Run(ctx, &buf, time.Second)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Run did not return promptly after context cancellation")
	}

	if buf.Len() == 0 {
		t.Fatal("expected at least the immediate timestamp to be written before cancellation")
	}
}
