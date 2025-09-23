package socket5proxy

import (
	"testing"
)

func TestProxyServer4(t *testing.T) {
	StartProxyServer4("localhost", "1080")
}
