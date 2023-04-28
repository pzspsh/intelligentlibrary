package telnet

import (
	"fmt"
	"net"
)

func TelnetProtocolConn(conn net.Conn) (bool, error) {
	var IsAuthentication = false
	var buf [4096]byte
	n, err := conn.Read(buf[0:])
	if nil != err {
		return false, fmt.Errorf("method: conn.Read, errInfo:%v", err)
	}

	buf[1] = 252
	buf[4] = 252
	buf[7] = 252
	buf[10] = 252
	_, err = conn.Write(buf[0:n])
	if nil != err {
		return false, fmt.Errorf("method: conn.Write, errInfo:%v", err)
	}

	n, err = conn.Read(buf[0:])
	if nil != err {
		return false, fmt.Errorf("method: conn.Read, errInfo:%v", err)
	}

	buf[1] = 252
	buf[4] = 251
	buf[7] = 252
	buf[10] = 254
	buf[13] = 252
	_, err = conn.Write(buf[0:n])
	if nil != err {
		return false, fmt.Errorf("method: conn.Write, errInfo:%v", err)
	}

	n, err = conn.Read(buf[0:])
	if nil != err {
		return false, fmt.Errorf("method: conn.Read, errInfo:%v", err)
	}

	buf[1] = 252
	buf[4] = 252
	_, err = conn.Write(buf[0:n])
	if nil != err {
		return false, fmt.Errorf("method: conn.Write, errInfo:%v", err)
	}

	_, err = conn.Read(buf[0:])
	if nil != err {
		return false, fmt.Errorf("method: conn.Read, errInfo:%v", err)
	}
	if !IsAuthentication {
		return true, nil
	} else {
		return false, fmt.Errorf("method: conn.Read, errInfo:%v", err)
	}
}
