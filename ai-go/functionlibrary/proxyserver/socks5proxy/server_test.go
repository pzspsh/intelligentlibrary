package socket5proxy

import (
	"testing"
)

func TestProxyServer(t *testing.T) {
	StartProxyServer("localhost", "1080")
}
