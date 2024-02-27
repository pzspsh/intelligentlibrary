/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:29:45
*/
package main

import (
	"bytes"
	"compress/gzip" //实现了gzip格式压缩文件的读写
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex" //实现了16进制字符表示的编解码
	"errors"
	"fmt"
	"math/big" //实现了大数字的多精度计算
	"strings"
)

/*
io包提供了对I/O原语的基本接口。本包的基本任务是包装这些原语已有的实现（如os包里的原语），
使之成为共享的公共接口，这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作
*/
/**
  通过一个随机key创建公钥和私钥
  随机key至少为36位
*/

func getEcdsaKey() (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	var err error
	var prk *ecdsa.PrivateKey
	var puk ecdsa.PublicKey
	var curve elliptic.Curve
	curve = elliptic.P256()
	//func NewReader(s string) *Reader
	prk, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return prk, puk, err
	}
	//prk:私钥 puk:公钥
	puk = prk.PublicKey

	return prk, puk, err

}

/*
*

	对text加密，text必须是一个hash值，例如md5、sha1等
	使用私钥prk
	返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func sign(text []byte, prk *ecdsa.PrivateKey) (string, error) {
	//r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, text)
	r, s, err := ecdsa.Sign(rand.Reader, prk, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	//创建并返回一个Writer。写入返回值的数据都会在压缩后写入w
	w := gzip.NewWriter(&b)
	//内建函数close关闭信道，该通道必须为双向的或只发送的
	//defer通常用来释放函数内部变量。
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}

	w.Flush() //确保所有的缓存操作已写入底层写入器
	return hex.EncodeToString(b.Bytes()), nil
}

/*
证书分解
通过hex解码，分割成数字证书r，s
*/
func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error, " + err.Error())
		return
	}
	/*
	   gzip.NewReader(r io.Reader) (*Reader, error)
	   返回一个从r读取并解压数据的*Reader。其实现会缓冲输入流的数据，并可能从r中读取比需要的更多的数据。
	   调用者有责任在读取完毕后调用返回值的Close方法。
	   buffer.NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
	   使用buf作为初始内容创建并初始化一个Buffer。本函数用于创建一个用于读取已存在数据的buffer；
	   也用于指定用于写入的内部缓冲的大小，
	   此时，buf应为一个具有指定容量但长度为0的切片。buf会被作为返回值的底层缓冲切片。
	   大多数情况下，new(Buffer)（或只是声明一个Buffer类型变量）就足以初始化一个Buffer了。
	*/
	//Buffer是一个实现了读写方法的可变大小的字节缓冲
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	//Reader类型满足io.Reader接口，可以从gzip格式压缩文件读取并解压数据。
	//一般，一个gzip文件可以是多个gzip文件的串联，每一个都有自己的头域。从Reader读取数据会返回串联的每个文件的解压数据，
	// 但只有第一个文件的头域被记录在Reader的Header字段里。
	//gzip文件会保存未压缩数据的长度与校验和。当读取到未压缩数据的结尾时，如果数据的长度或者校验和不正确，
	//Reader会返回ErrCheckSum。因此，调用者应该将Read方法返回的数据视为暂定的，直到他们在数据结尾获得了一个io.EOF。
	count, err := r.Read(buf)
	//func (z *Reader) Read(p []byte) (n int, err error)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	//Split(s, sep string) []string //sep:步长
	rs := strings.Split(string(buf[:count]), "+")

	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	//实现了Marshaler接口的类型可以将自身序列化为合法的json描述。
	//UnmarshalText必须可以解码MarshalText生成的textual格式数据。
	//本函数可能会对data内容作出修改，所以如果要保持data的数据请事先进行拷贝
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

/*
校验文本内容是否与签名一致
使用公钥校验签名和文本内容
*/
func verify(text []byte, signature string, key ecdsa.PublicKey) (bool, error) {

	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(&key, text, &rint, &sint)

	return result, nil
}

/*
  hash加密
  使用md5加密
  msg+
*/
//func hashtext(text, salt string) []byte {
func hashtext(text string) []byte {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(text))
	//result := Md5Inst.Sum([]byte(salt))

	return Md5Inst.Sum(nil)
}

func main() {
	//创建公钥和私钥
	prk, puk, err := getEcdsaKey()
	if err != nil {
		fmt.Println(err)
	}

	//待加密的明文
	text := string("少壮不努力,活该你单身2333")

	//hash取值
	htext := hashtext(text)

	//hash值编码输出
	hex.EncodeToString(htext)

	//hash值+私钥进行签名
	result, err := sign(htext, prk)
	if err != nil {
		fmt.Println(err)
	}

	//签名与hash值进行校验
	//hash值+密文+公钥
	tmp, err := verify(htext, result, puk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmp)
}
