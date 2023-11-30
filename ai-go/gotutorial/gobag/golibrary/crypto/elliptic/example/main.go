/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:28:44
*/
package main

import (
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"crypto/elliptic"
	"crypto/rand"
)

func main() {
	src := []byte(string("少壮不努力,活该你单身23333"))
	//src1 := []byte(string("少壮不努力,老大徒伤悲3344"))
	mysrc := myHash(src)
	//mysrc1 := myHash(src1)

	prk, puk, _ := genePriPubKey()
	mystring := sign(prk, mysrc)

	r, s := getSign(mystring)

	result := verifySign(&r, &s, mysrc, puk)
	fmt.Print(result)

}

func genePriPubKey() (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	var err error
	var pubkey ecdsa.PublicKey
	var prikey *ecdsa.PrivateKey
	var curve elliptic.Curve

	curve = elliptic.P384()
	prikey, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return prikey, pubkey, err
	}
	pubkey = prikey.PublicKey

	return prikey, pubkey, err
}
func myHash(src []byte) []byte {
	myhash := md5.New()
	myhash.Write(src)
	return myhash.Sum(nil)
}

func sign(key *ecdsa.PrivateKey, myhash []byte) string {
	r, s, _ := ecdsa.Sign(rand.Reader, key, myhash)
	rm, _ := r.MarshalText()
	sm, _ := s.MarshalText()

	return hex.EncodeToString([]byte(string(rm) + "+" + string(sm)))
}

func getSign(hexrs string) (rint, sint big.Int) {
	st, _ := hex.DecodeString(hexrs)
	str := strings.Split(string(st), "+")
	_ = rint.UnmarshalText([]byte(str[0]))
	//rint是指针:error: invalid memory address or nil pointer dereference
	_ = sint.UnmarshalText([]byte(str[1]))

	return
}

func verifySign(rint, sint *big.Int, myhash []byte, pubkey ecdsa.PublicKey) bool {
	result := ecdsa.Verify(&pubkey, myhash, rint, sint)
	return result
}
