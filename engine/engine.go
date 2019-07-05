package engine

import (
	"bufio"
	"github.com/scott-x/gcase/model"
	"github.com/scott-x/gcase/parse"
	"github.com/scott-x/gcase/utils"
	"os"
	"strings"
	"github.com/fatih/color"
)

func Run() {
	inputReader := bufio.NewReader(os.Stdin)

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)

	boldGreen.Println("Please enter your job type (1-6):")
	color.Cyan(" 1. USA ")
	color.Cyan(" 2. CAN")
	color.Cyan(" 3. CAB")
	color.Cyan(" 4. LNC")
	color.Cyan(" 5. PRINT")
	color.Cyan(" 6. OTHER")

	boldRed.Printf("My Selection is : ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.Trim(input, "\n")
	switch input {
	case "1":
		Copy(parse.Write_Data(model.GetCase("USA"),1))
	case "2":
		Copy(parse.Write_Data(model.GetCase("CAN"),2))
	case "3":
		Copy(parse.Write_CAB(model.GetCase("CAB"),3))
	case "4":
		Copy(parse.Write_Data(model.GetCase("LNC"),4))
	case "5":
		Copy(parse.Write_Print(model.GetCase("PRINT"),5))
	case "6":
		Copy(parse.Write_Data(model.GetCase("OTHER"),6))
	default:
		boldRed.Printf("Type Error: Please input numer 1-5")
	}

}

func Copy(s1, s2 string) {
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	
	project_folder := "/Users/scottxiong/go/src/github.com/scott-x/gcase"
	copyError := utils.CopyFolder(s1, s2)
	if copyError != nil {
		panic(copyError)
	}
	err := os.RemoveAll(project_folder+"/temp")
	if err != nil {
		panic(err)
	}
	boldGreen.Printf("Job: ")
	boldRed.Printf(strings.Trim(s2,"./"))
	boldGreen.Printf(" was created successfully ^_^")
}
