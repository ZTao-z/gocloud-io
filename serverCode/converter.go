package service

import (
	"github.com/axgle/mahonia"
)

func ConvertToString(src, srcCode, tarCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tarCoder := mahonia.NewDecoder(tarCode)
	_, cdata, _ := tarCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
