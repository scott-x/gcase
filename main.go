package main
import (
    "fmt"
    "bufio"
    "os"
    "github.com/scott-x/gcase/parse"
    "strings"
    "os/exec"
    "path"
)

var project_folder string = "/Users/scottxiong/go/src/github.com/scott-x/gcase"
func main(){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your job type (1-5):")
	fmt.Println(" 1. USA ")
	fmt.Println(" 2. CAN")
	fmt.Println(" 3. CAB")
	fmt.Println(" 4. LNC")
	fmt.Println(" 5. PRINT")
	
  fmt.Printf("My Selection is : ")
	input, err := inputReader.ReadString('\n')
    if err != nil {
    	panic(err)
    }
    input= strings.Trim(input,"\n")
    switch input{
      case "1":
        parse.Write_Data()
        copy()
      case "2":
        fmt.Println(2)
      case "3":
        fmt.Println(3)
      case "4":
        fmt.Println(4) 
      case "5":
        fmt.Println(5)    
      default:  
        fmt.Println("Type Error: Please input numer 1-5")    
    }
}

func copy(){
  //copy 
  err :=exec.Command("mv",path.Join(project_folder,"temp/*"),"./").Run()
  if err!=nil {
    panic(err)
  }
}

