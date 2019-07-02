package parse

import (
	"fmt"
  "os/exec"
	"github.com/tealeg/xlsx"
	"path"
	"reflect"
	"time"
  "strings"
  "github.com/scott-x/gcase/utils"
)

var prroject_folder string ="/Users/scottxiong/go/src/github.com/scott-x/gcase"
func Write_Data(){
  job := Receive_data()	
  //fmt.Println(job)
  
  from := path.Join(prroject_folder,"xlsx_templates/WMT_xxxxxx_xxx 做稿")
  to := path.Join(prroject_folder,"temp/")
  err :=CopyDir(from,to)
  if err==nil{
  	time.Sleep(10*time.Millisecond)
  	to=path.Join(to,"WMT_xxxxxx_xxx 做稿/U180XXX_XXX_DetailList_W.xlsx")
  	file,error := xlsx.OpenFile(to)
  	
  	if error!=nil{
  	    panic(error)
  	}
  	defer file.Save(to)
  	fmt.Println(job.job_number)
  	fmt.Println(reflect.TypeOf(job.job_number))
  	/*
  	  Row 从0开始
  	  Col 从1开始
  	*/
  	//job_number string
  	file.Sheet["任务单"].Rows[1].Cells[1].SetString(strings.ToUpper(job.job_number))
  	//short_brand string
  	file.Sheet["任务单"].Rows[2].Cells[1].SetString(strings.ToUpper(job.short_brand)+"/"+job.country)
  	//author
  	file.Sheet["任务单"].Rows[2].Cells[3].SetString("业务制稿人: Scott")
  	//job#
  	file.Sheet["任务单"].Rows[3].Cells[7].SetString(strings.ToUpper(job.job_number))
  	//create_date string
  	file.Sheet["任务单"].Rows[4].Cells[0].SetString(job.create_date)
  	//addtional_notes string
  	file.Sheet["任务单"].Rows[5].Cells[1].SetString(job.addtional_notes)
     
    //file.Sheet["任务单"].Rows[5].Cells[1].SetStyle(font)
    //program
    file.Sheet["任务单"].Rows[4].Cells[7].SetString("Program: "+job.program)
  	//supplier string
  	file.Sheet["任务单"].Rows[5].Cells[7].SetString("Supplier: "+utils.FirstLtterToUpper(job.supplier))
  	//buyer string
  	file.Sheet["任务单"].Rows[6].Cells[7].SetString("Buyer: "+utils.FirstLtterToUpper(job.buyer))
  	//due string
  	file.Sheet["任务单"].Rows[7].Cells[7].SetString("Artwork due date: "+job.due)
  	//packout string
  	file.Sheet["任务单"].Rows[8].Cells[7].SetString("Packout date: "+job.packout)
  	//ship string
  	file.Sheet["任务单"].Rows[9].Cells[7].SetString("Shipdate: "+job.ship)
  	//instore string
  	file.Sheet["任务单"].Rows[10].Cells[7].SetString("In-store date: "+job.instore)
  	//contact string
  	file.Sheet["任务单"].Rows[12].Cells[7].SetString("联系人: "+job.contact)
    
    //重命名

    
    //copy 
    err :=exec.Command("mv",path.Join(prroject_folder,"temp/*"),"./").Run()
    if err!=nil {
      panic(err)
    }
    


  	fmt.Println("done")

  }
}

func CopyDir(src, dst string) error {
	err :=exec.Command("mkdir", "-p", prroject_folder+"/temp").Run()
	if err!=nil {
		panic(err)
	}
    cmd := exec.Command("cp", "-r",src,dst)
    return cmd.Run()
}

