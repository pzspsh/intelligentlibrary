/*
@File   : server2_test.go
@Author : pan
@Time   : 2023-11-09 14:23:09
*/
package socket5proxy

import (
	"testing"
)

func TestProxyServer2(t *testing.T) {
	StartProxyServer2("localhost", "1080")
}
