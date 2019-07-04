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
        if string(ch)==" " && strArry[i+1]>=96 && strArry[i+1]<=122 {
            
            strArry[i+1]-=32
        }
    }
    return string(strArry)
}
