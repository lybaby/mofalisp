package mofa

import (
	"io/ioutil"
)

func LoadFile(filepath string) string {
	code, err:=ioutil.ReadFile(filepath) 
	if err != nil {
		return ""
	}
	return string(code)
}
