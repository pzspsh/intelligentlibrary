/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:08:04
*/
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

/*
	需求：实现基于椭圆曲线ECDSA的数字签名和签名的验证
	基本步骤
	1.生成ECDSA密匙对并写入磁盘
	2.读取ecdsa私匙并对明文散列值进行签名
	3.读取ecdsa公匙验证数字签名的正确性
*/
//------------1.获取ECDSA密匙对
func GenerateEccKey() {
	//----------获取私匙并写入磁盘----------
	/*
		 	Step1.生成ecc私匙
		 	函数：func GenerateKey(c elliptic.Curve, rand io.Reader) (priv *PrivateKey, err error)
		 	作用：生成ecdsa的密匙对
		 	返回参数1：私匙
		 	type PrivateKey struct {
		    				PublicKey
							D   *big.Int
						}
		 	其中，
		 	type PublicKey struct {
		 					elliptic.Curve
		 					X, Y *big.Int
						}
		 	返回参数2：error
		 	参数1：elliptic包中实现了几条覆盖素数有限域的标准椭圆曲线，选择椭圆曲线
		 	参数2：rand.Reader
	*/
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//Step2:采用x509序列化
	derText, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		panic(err)
	}
	//Step3:组织一个pem的block结构体
	block := pem.Block{
		Type:  "ECDSA Private Key",
		Bytes: derText,
	}
	//Step4:进行pem编码
	file, err := os.Create("eccPrivate.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)
	//----------获取公匙并写入磁盘----------
	//Step1:获取公匙
	publicKey := privKey.PublicKey
	//Step2:采用x509序列化
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//Step3:组织一个pem的block结构体
	block = pem.Block{
		Type:  "ECDSA Public Key",
		Bytes: derText,
	}
	//Step4:进行pem编码
	file, err = os.Create("eccPublic.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)
}

// ------------ECDSA私匙执行数字签名------------
func EccSignature(plainText []byte, privName string) (rText, sText []byte) {
	//------1.获取私匙------
	//Step1:打开文件获取原始私匙
	file, err := os.Open(privName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	//Step2:私匙的反pem编码化
	block, _ := pem.Decode(buf)
	//Step3:私匙的反x509序列化
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//------求取明文的散列值------
	//Step1:创建基于sha256哈希函数的hash接口
	myHash := sha256.New()
	//Step2:写入数据
	myHash.Write(plainText)
	//Step:求出明文的散列值
	hashText := myHash.Sum(nil)
	//------对明文的散列值进行数字签名
	/*
			Step1:基于ECDSA实现数字签名
			函数：func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
			作用：使用私钥对任意长度的hash值（必须是较大信息的hash结果）进行签名，返回签名结果（一对大整数），
			私钥的安全性取决于密码读取器的熵度（随机程度）
			返回参数1：一对大整数,表示点的x和y轴坐标值，需要转换为[]byte
				转换函数:func (z *Int) MarshalText() (text []byte, err error)  本方法实现了encoding.TextMarshaler接口
						type TextMarshaler interface {
		  					MarshalText() (text []byte, err error)
								}
						实现了TextMarshaler接口的类型可以将自身序列化为utf-8编码的textual格式
			返回参数2：error
			参数1：rand.Reader
			参数2：私匙
			参数3：明文的散列值
	*/
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hashText)
	if err != nil {
		panic(err)
	}
	//Step2:序列化
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}
	return
}

// ------------ECDSA公匙验证数字签名------------
func EccVerify(plainText, rText, sText []byte, pubFile string) bool {
	//------1.获取公钥------
	//Step1:打开文件获取公匙
	file, err := os.Open(pubFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	//Step2：将公匙反pem码化
	block, _ := pem.Decode(buf)
	//Step3:将公匙反x509序列化
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//Step4:执行公匙的类型断言
	publicKey := pubInterface.(*ecdsa.PublicKey)
	//------2.获取明文的散列值------
	//Step1:创建hash接口，指定采用的哈希函数
	myHash := sha256.New()
	//Step2:向myHash中写入内容
	myHash.Write(plainText)
	//Step3:生成明文的散列值
	hashText := myHash.Sum(nil)
	//------3.对数字签名后的内容进行解密------
	/*
		Step1:对rText和sText解序列化
		函数：func (z *Int) UnmarshalText(text []byte) error
	*/
	var r, s big.Int //空指针，指向了系统的0x00地址，不能对这个地址进行操作
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	/*
		Step2:采用ECDSA的公匙验证数字签名的正确性
		函数：func Verify(pub *PublicKey, hash []abyte, r, s *big.Int) bool
		作用：使用公钥验证hash值和两个大整数r、s构成的签名，并返回签名是否合法
		返回参数：验证数字签名是否正确
		参数1：公匙
		参数2：明文的散列值
		参数3,4:rText、sText的序列化
	*/
	return ecdsa.Verify(publicKey, hashText, &r, &s)
}

// 主函数
func main() {
	GenerateEccKey()
	src := []byte("使用x509对pem.Block中的Bytes变量中的数据进行解析 ->  得到一接口")
	rText, sText := EccSignature(src, "eccPrivate.pem")
	bl := EccVerify(src, rText, sText, "eccPublic.pem")
	fmt.Println(bl)
}
