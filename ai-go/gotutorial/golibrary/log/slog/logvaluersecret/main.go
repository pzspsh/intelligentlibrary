/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:31:54
*/
package main

import (
	"log/slog"
	"os"
)

// A token is a secret value that grants permissions.
type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

func RemoveTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}

// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
func main() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: RemoveTime}))
	logger.Info("permission granted", "user", "Perry", "token", t)
}
