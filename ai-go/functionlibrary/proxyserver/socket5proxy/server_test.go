/*
@File   : server_test.go
@Author : pan
@Time   : 2023-11-09 11:18:52
*/
package socket5proxy

import (
	"testing"
)

func TestProxyServer(t *testing.T) {
	StartProxyServer("localhost", "1080")
}
