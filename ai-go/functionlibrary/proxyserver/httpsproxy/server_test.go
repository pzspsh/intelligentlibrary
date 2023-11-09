/*
@File   : server2_test.go
@Author : pan
@Time   : 2023-11-09 15:17:47
*/
package httpsproxy

import (
	"testing"
)

func TestStartHttps2(t *testing.T) {
	SartHttps("localhost", "1080")
}
