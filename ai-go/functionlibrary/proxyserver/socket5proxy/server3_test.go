/*
@File   : server3_test.go
@Author : pan
@Time   : 2023-11-09 14:34:19
*/
package socket5proxy

import (
	"testing"
)

func TestProxyServer3(t *testing.T) {
	StartProxyServer3("localhost", "1080")
}
