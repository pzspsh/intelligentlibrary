/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:27:35
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"crypto/md5"
	"encoding/hex"
)

func main() {
	bodyBuffer := &bytes.Buffer{}
	bodyBuffer.WriteString(`-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="needShow"

0
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docreplyable"

0
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="usertype"

1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="from"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="userCategory"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="userId"

114
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="userType"

1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docstatus"

0
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="doccode"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docedition"

-1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="doceditionid"

-1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="maincategory"

15
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="subcategory"

49
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="seccategory"

48
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="ownerid"

114
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docdepartmentid"

10
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="doclangurage"

7
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="maindoc"

-1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="topage"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="operation"

addsave
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="SecId"

48
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="imageidsExt"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="imagenamesExt"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="delImageidsExt"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="namerepeated"

0
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docsubject"

1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="doccontent"

1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="readoptercanprint"

1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="selectCategory"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="tempDocModule"

-1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docmodule"

-1
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="keyword"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="selectMainDocument"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="invalidationdate"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="dummycata"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="hrmresid"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="crmid"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="projectid"


-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="imgType"

2
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="imgUrl_doccontent"

http://
-----------------------------7e431d37a30abc
Content-Disposition: form-data; name="docimages_num"

0
-----------------------------7e431d37a30abc--`)

	headers := `Accept: text/html, application/xhtml+xml, */*
Referer: http://192.168.132.80/docs/docs/DocAdd.jsp?mainid=15&subid=49&secid=48&showsubmit=1&coworkid=&prjid=&isExpDiscussion=&crmid=&hrmid=&topage=
Accept-Language: zh-CN
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko
Content-Type: multipart/form-data; boundary=---------------------------7e431d37a30abc
Accept-Encoding: gzip, deflate
Host: 192.168.132.80
Content-Length: 4212
Connection: Keep-Alive
Pragma: no-cache
Cookie: testBanCookie=test; JSESSIONID=abcIswHnk9uU49ql9MP2w; loginfileweaver=%2Fwui%2Ftheme%2Fecology7%2Fpage%2Flogin.jsp%3FtemplateId%3D6%26logintype%3D1%26gopage%3D; loginidweaver=114; languageidweaver=7`

	uri := `http://192.168.132.80/docs/docs/DocDsp.jsp?fromFlowDoc=&id=803038&blnOsp=false&topage=&pstate=sub`
	req, err := http.NewRequest("POST", uri, io.NopCloser(bodyBuffer))
	if err != nil {
		log.Printf("Cannot NewRequest: %s , err: %v", uri, err)
		return
	}
	AddHeaders(req, headers)
	fmt.Println(req.Header)
	//fmt.Println(req.Body)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Cannot client.Do, err: %v", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(len(string(body)))

}

func AttachField(bodyWriter *multipart.Writer, keyname, keyvalue string) error {
	if err := bodyWriter.WriteField(keyname, keyvalue); err != nil {
		log.Printf("Cannot WriteField: %s, err: %v", keyname, err)
		return err
	}
	return nil
}

func AttachFile(bodyWriter *multipart.Writer, formname, filename string) error {
	fullname := filepath.Join(".", filename)
	file, err := os.Open(fullname)
	if err != nil {
		log.Printf("Cannot open file: %s , err: %v", fullname, err)
		return err
	}
	defer file.Close()

	// MD5
	md5hash := md5.New()
	if _, err = io.Copy(md5hash, file); err != nil {
		log.Printf("Cannot open md5 hash: %s , err: %v", fullname, err)
		return err
	}

	keyname := filename + ".md5cksum"
	keyvalue := hex.EncodeToString(md5hash.Sum(nil)[:16])
	if err = AttachField(bodyWriter, keyname, keyvalue); err != nil {
		log.Printf("Cannot WriteField: %s, err: %v", keyname, err)
		return err
	}

	// file
	part, err := bodyWriter.CreateFormFile(formname, filename)
	if err != nil {
		log.Printf("Cannot CreateFormFile for: %s , err: %v", filename, err)
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Printf("Cannot Copy file: %s , err: %v", fullname, err)
		return err
	}

	return nil
}

func AddHeaders(req *http.Request, headers string) *http.Request {
	//将传入的Header分割成[]ak和[]av
	a := strings.Split(headers, "\n")
	ak := make([]string, len(a[:]))
	av := make([]string, len(a[:]))
	//要用copy复制值；若用等号仅表示指针，会造成修改ak也就是修改了av
	copy(ak, a[:])
	copy(av, a[:])
	//fmt.Println(ak[0], av[0])
	for k, v := range ak {
		i := strings.Index(v, ":")
		j := i + 1
		ak[k] = v[:i]
		av[k] = v[j:]
		//设置Header
		req.Header.Set(ak[k], av[k])
	}
	return req
}
