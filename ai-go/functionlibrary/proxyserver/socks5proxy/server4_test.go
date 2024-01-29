/*
@File   : server4_test.go
@Author : pan
@Time   : 2023-11-09 14:48:46
*/
package socket5proxy

import (
	"testing"
)

func TestProxyServer4(t *testing.T) {
	StartProxyServer4("localhost", "1080")
}
