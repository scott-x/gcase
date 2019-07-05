package main

import (
	"fmt"
	"path"
)

func main(){
	fmt.Printf("the dir is %s",path.Dir('./'))
}