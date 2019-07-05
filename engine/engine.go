package engine

import (
	"bufio"
	"fmt"
	"github.com/scott-x/gcase/model"
	"github.com/scott-x/gcase/parse"
	"github.com/scott-x/gcase/utils"
	"os"
	"strings"
)

func Run() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your job type (1-6):")
	fmt.Println(" 1. USA ")
	fmt.Println(" 2. CAN")
	fmt.Println(" 3. CAB")
	fmt.Println(" 4. LNC")
	fmt.Println(" 5. PRINT")
	fmt.Println(" 6. OTHER")

	fmt.Printf("My Selection is : ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.Trim(input, "\n")
	switch input {
	case "1":
		Copy(parse.Write_Data(model.GetCase("USA")))
	case "2":
		Copy(parse.Write_Data(model.GetCase("CAN")))
	case "3":
		Copy(parse.Write_CAB(model.GetCase("CAB")))
	case "4":
		Copy(parse.Write_Data(model.GetCase("LNC")))
	case "5":
		Copy(parse.Write_Print(model.GetCase("PRINT")))
	case "6":
		Copy(parse.Write_Data(model.GetCase("OTHER")))
	default:
		fmt.Println("Type Error: Please input numer 1-5")
	}

}

func Copy(s1, s2 string) {
	copyError := utils.CopyFolder(s1, s2)
	if copyError != nil {
		panic(copyError)
	}
	err := os.RemoveAll("./temp")
	if err != nil {
		panic(err)
	}
	fmt.Println("Job: "+strings.Trim(s2,"./")+ " was created successfully ^__^")
}
