package utils

import (
	"strings"
)

func FirstLtterToUpper(str string) string {
    if len(str) < 1 {
        return ""
    }
    strArry := []rune(strings.ToLower(strings.Trim(str, " ")))
    strArry[0]-=32
    for i,ch := range strArry{
        if string(ch)==" " && string(strArry[i+1])!="(" {
            
            strArry[i+1]-=32
        }
    }
    return string(strArry)
}
