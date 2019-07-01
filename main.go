package main
import (
    "fmt"
    "bufio"
    "os"
    "reflect"
)

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
    if err == nil {
    	fmt.Println(reflect.TypeOf(input))
    	switch input{
    	   case "1":
    	     fmt.Println("USA") 
    	     return
    	   case "2":
    	     fmt.Println("CAN")   
    	   default:
    	     fmt.Println("default") 
       
    	}
     
    }
}

