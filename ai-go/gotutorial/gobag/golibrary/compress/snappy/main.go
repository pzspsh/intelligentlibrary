/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:42:17
*/
package main

import (
	"fmt"
	"os"

	"github.com/golang/snappy"
)

var (
	textMap = map[string]string{
		"a": `1234567890-=qwertyuiop[]\';lkjhgfdsazxcvbnm,./`,
		"b": `1234567890-=qwertyuiop[]\';lkjhgfdsazxcvbnm,./1234567890-=qwertyuiop[]\';lkjhgfdsazxcvbnm,./1234567890-=qwertyuiop[]\';lkjhgfdsazxcvbnm,./1234567890-=qwertyuiop[]\';lkjhgfdsazxcvbnm,./`,
		"c": `浕浉浄浀浂洉洡洣浐洘泚浌洼洽派洿浃浇浈浊测浍济浏浑浒浓浔泿洱涏洀洁洂洃洄洅洆洇洈洊洋洌洎洏洐洑洒洓洔洕洗洠洙洚洛洝洞洟洢洤津洦洧洨洩洪洫洬洭洮洲洳洴洵洶洷洸洹洺活涎`,
		"d": `浕浉浄浀浂洉洡洣浐洘泚浌洼洽派洿浃浇浈浊测浍济浏浑浒浓浔泿洱涏洀洁洂洃洄洅洆洇洈洊洋洌洎洏洐洑洒洓洔洕洗洠洙洚洛洝洞洟洢洤津洦洧洨洩洪洫洬洭洮洲洳洴洵洶洷洸洹洺活涎浕浉浄浀浂洉洡洣浐洘泚浌洼洽派洿浃浇浈浊测浍济浏浑浒浓浔泿洱涏洀洁洂洃洄洅洆洇洈洊洋洌洎洏洐洑洒洓洔洕洗洠洙洚洛洝洞洟洢洤津洦洧洨洩洪洫洬洭洮洲洳洴洵洶洷洸洹洺活涎浕浉浄浀浂洉洡洣浐洘泚浌洼洽派洿浃浇浈浊测浍济浏浑浒浓浔泿洱涏洀洁洂洃洄洅洆洇洈洊洋洌洎洏洐洑洒洓洔洕洗洠洙洚洛洝洞洟洢洤津洦洧洨洩洪洫洬洭洮洲洳洴洵洶洷洸洹洺活涎`,
	}
	imgSrc = []string{
		"1.jpg", "2.jpg", "3.jpg", "4.jpg",
	}
)

func main() {

	for k, v := range textMap {
		got := snappy.Encode(nil, []byte(v))
		fmt.Println("k:", k, "len:", len(v), len(got))
	}

	fmt.Println("snappy jpg")
	for _, v := range imgSrc {
		buf, err := os.ReadFile(v)
		if err == nil {
			got := snappy.Encode(nil, buf)
			fmt.Println("k:", v, "len:", len(buf), len(got))
		}
	}
}
