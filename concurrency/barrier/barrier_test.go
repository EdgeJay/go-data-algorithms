package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"https://httpbin.org/headers", "https://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "user-agent") {
			t.Error("Missing content")
		}

		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"https://malformed-url", "https://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Error("Expected output to contain ERROR")
		}

		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"https://httpbin.org/headers", "https://httpbin.org/user-agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Error("Expected output to contain Timeout")
		}

		t.Log(result)
	})
}
