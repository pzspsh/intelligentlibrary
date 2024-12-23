/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:31:54
*/
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	r, _ := http.NewRequest("GET", "localhost", nil)
	// ...

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: RemoveTime}))
	logger.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))

}

func RemoveTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}
