package parse

import (
  "github.com/scott-x/gcase/model"
  "github.com/scott-x/gcase/utils"
  "github.com/tealeg/xlsx"
  "path"
  "strings"
  "time"
)

func Write_Data(c model.Case) (s1, s2 string) {
  //project_folder := utils.Dir() //返回的是main方法的目录 也是当前命令行的执行目录，所以这里需写死
  //fmt.Printf("project_folder: %s", project_folder)
  // s,err :=exec.Command("pwd").Output()
  // if err!=nil{
  //   panic(err)
  // }
  project_folder := "/Users/scottxiong/go/src/github.com/scott-x/gcase"

  from := path.Join(project_folder, "xlsx_templates/"+c.Folder)
  to := path.Join(project_folder, "temp/"+c.Folder)
  err := utils.CopyFolder(from, to)
  if err != nil {
    panic(err)
  }

  //fmt.Printf("copy folder from %s to %s", from, to)
  job := Receive_data()
  //fmt.Println(job)
  // //rename file
  oldFile := project_folder + "/temp/" + c.Folder + "/" + c.Pf
  newFile := project_folder + "/temp/" + c.Folder + "/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  error := utils.Rename(oldFile, newFile)
  if error != nil {
    panic(error)
  }
  //fmt.Printf("rename file %s done", oldFile)
  //rename folder
  oldFolder := project_folder + "/temp/" + c.Folder
  newFolder := project_folder + "/temp/" + strings.ToUpper(job.job_number) + " 做稿"
  error1 := utils.Rename(oldFolder, newFolder)
  //fmt.Printf("rename folder %s done", oldFolder)
  if error1 != nil {
    panic(error1)
  }
  // 再新建2个当天日期的空目录
  todayDateFoler := time.Now().Format("0102")
  folder1 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+" 做稿/1 intake sheet & order/"+string(todayDateFoler))
  folder2 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+" 做稿/2 raw client files/"+string(todayDateFoler))
  utils.CreateDirIfNotExist(folder1)
  utils.CreateDirIfNotExist(folder2)

  //开始写数据
  newFile = project_folder + "/temp/" + strings.ToUpper(job.job_number) + " 做稿/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  file, error := xlsx.OpenFile(newFile)

  if error != nil {
    panic(error)
  }
  defer file.Save(newFile)
  // fmt.Println(job.job_number)
  // fmt.Println(reflect.TypeOf(job.job_number))
  /*
     Row 从0开始
     Col 从1开始
  */
  //job_number string
  file.Sheet["任务单"].Rows[1].Cells[1].SetString("Job #: " + strings.ToUpper(job.job_number))
  //short_brand string
  file.Sheet["任务单"].Rows[2].Cells[1].SetString(strings.ToUpper(job.short_brand) + "/" + job.country)
  //author
  file.Sheet["任务单"].Rows[2].Cells[3].SetString("业务制稿人: Scott")
  //job#
  file.Sheet["任务单"].Rows[3].Cells[7].SetString("Job #: " + strings.ToUpper(job.job_number))
  //create_date string
  file.Sheet["任务单"].Rows[4].Cells[0].SetString(string(time.Now().Format("01/02/2006")))
  //addtional_notes string
  file.Sheet["任务单"].Rows[5].Cells[1].SetString(job.addtional_notes)

  //file.Sheet["任务单"].Rows[5].Cells[1].SetStyle(font)
  //program
  file.Sheet["任务单"].Rows[4].Cells[7].SetString("Program: " + job.program)
  //supplier string
  file.Sheet["任务单"].Rows[5].Cells[7].SetString("Supplier: " + utils.FirstLtterToUpper(job.supplier))
  //buyer string
  file.Sheet["任务单"].Rows[6].Cells[7].SetString("Buyer: " + utils.FirstLtterToUpper(job.buyer) + " (" + strings.ToUpper(job.dep) + ") ")
  //due string
  file.Sheet["任务单"].Rows[7].Cells[7].SetString("Artwork due date: " + job.due)
  //packout string
  file.Sheet["任务单"].Rows[8].Cells[7].SetString("Packout date: " + job.packout)
  //ship string
  file.Sheet["任务单"].Rows[9].Cells[7].SetString("Shipdate: " + job.ship)
  //instore string
  file.Sheet["任务单"].Rows[10].Cells[7].SetString("In-store date: " + job.instore)
  //contact string
  file.Sheet["任务单"].Rows[12].Cells[7].SetString("联系人: " + job.contact)

  //copy
  return newFolder, "./" + strings.ToUpper(job.job_number) + " 做稿"
}

func Write_Print(c model.Case) (s1, s2 string) {
  project_folder := utils.Dir() //返回的是main方法的目录
  //fmt.Printf("project_folder: %s", project_folder)
  from := path.Join(project_folder, "xlsx_templates/"+c.Folder)
  to := path.Join(project_folder, "temp/"+strings.Trim(c.Folder,"印刷/"))
  err := utils.CopyFolder(from, to)
  if err != nil {
    panic(err)
  }

  //fmt.Printf("copy folder from %s to %s", from, to)
  job := Receive_Print()
  //fmt.Println(job)
  // //rename file
  oldFile := project_folder + "/temp/" + strings.Trim(c.Folder,"印刷/")+ "/" + c.Pf
  newFile := project_folder + "/temp/" + strings.Trim(c.Folder,"印刷/") + "/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  error := utils.Rename(oldFile, newFile)
  if error != nil {
    panic(error)
  }
  //fmt.Printf("rename file %s done", oldFile)
  //rename folder
  oldFolder := project_folder + "/temp/" + strings.Trim(c.Folder,"印刷/")
  newFolder := project_folder + "/temp/" + strings.ToUpper(job.job_number)
  error1 := utils.Rename(oldFolder, newFolder)
  //fmt.Printf("rename folder %s done", oldFolder)
  if error1 != nil {
    panic(error1)
  }
  // 再新建2个当天日期的空目录
  todayDateFoler := time.Now().Format("0102")
  //folder1 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+" 做稿/1 intake sheet & order/"+string(todayDateFoler))
  folder2 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+"/3报价单或订单/"+string(todayDateFoler))
  //utils.CreateDirIfNotExist(folder1)
  utils.CreateDirIfNotExist(folder2)

  //开始写数据
  newFile = project_folder + "/temp/" + strings.ToUpper(job.job_number) +"/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  file, error := xlsx.OpenFile(newFile)

  if error != nil {
    panic(error)
  }
  defer file.Save(newFile)
  // fmt.Println(job.job_number)
  // fmt.Println(reflect.TypeOf(job.job_number))
  /*
     Row 从0开始
     Col 从1开始
  */
  //job_number string
  file.Sheet["任务单"].Rows[1].Cells[1].SetString("Job #: " + strings.ToUpper(job.job_number))
  //short_brand string
  file.Sheet["任务单"].Rows[2].Cells[1].SetString(job.po)
  //author
  file.Sheet["任务单"].Rows[2].Cells[3].SetString("业务制稿人: Scott")
  //job#
  file.Sheet["任务单"].Rows[3].Cells[7].SetString("Job #: " + strings.ToUpper(job.job_number))
  //create_date string
  file.Sheet["任务单"].Rows[4].Cells[0].SetString(string(time.Now().Format("01/02/2006")))

  //copy
  return newFolder, "./" + strings.ToUpper(job.job_number) + " 做稿"
}

func Write_CAB(c model.Case) (s1, s2 string) {
  project_folder := utils.Dir() //返回的是main方法的目录
  //fmt.Printf("project_folder: %s", project_folder)
  from := path.Join(project_folder, "xlsx_templates/"+c.Folder)
  to := path.Join(project_folder, "temp/"+c.Folder)
  err := utils.CopyFolder(from, to)
  if err != nil {
    panic(err)
  }

  //fmt.Printf("copy folder from %s to %s", from, to)
  job := Receive_CAB()
  //fmt.Println(job)
  // //rename file
  oldFile := project_folder + "/temp/" + c.Folder + "/" + c.Pf
  newFile := project_folder + "/temp/" + c.Folder + "/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  error := utils.Rename(oldFile, newFile)
  if error != nil {
    panic(error)
  }
  //fmt.Printf("rename file %s done", oldFile)
  //rename folder
  oldFolder := project_folder + "/temp/" + c.Folder
  newFolder := project_folder + "/temp/" + strings.ToUpper(job.job_number) + " 做稿"
  error1 := utils.Rename(oldFolder, newFolder)
  //fmt.Printf("rename folder %s done", oldFolder)
  if error1 != nil {
    panic(error1)
  }
  // 再新建2个当天日期的空目录
  todayDateFoler := time.Now().Format("0102")
  //folder1 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+" 做稿/1 intake sheet & order/"+string(todayDateFoler))
  folder2 := path.Join(project_folder, "temp/"+strings.ToUpper(job.job_number)+" 做稿/2 raw client files/"+string(todayDateFoler))
  //utils.CreateDirIfNotExist(folder1)
  utils.CreateDirIfNotExist(folder2)

  //开始写数据
  newFile = project_folder + "/temp/" + strings.ToUpper(job.job_number) + " 做稿/" + strings.ToUpper(job.job_number) + "_DetailList_W.xlsx"
  file, error := xlsx.OpenFile(newFile)

  if error != nil {
    panic(error)
  }
  defer file.Save(newFile)
  // fmt.Println(job.job_number)
  // fmt.Println(reflect.TypeOf(job.job_number))
  /*
     Row 从0开始
     Col 从1开始
  */
  //job_number string
  file.Sheet["任务单"].Rows[1].Cells[1].SetString("Job #: " + strings.ToUpper(job.job_number))
  //short_brand string
  //file.Sheet["任务单"].Rows[2].Cells[1].SetString(strings.ToUpper(job.short_brand) + "/" + job.country)
  //author
  file.Sheet["任务单"].Rows[2].Cells[3].SetString("业务制稿人: Scott")
  //job#
  file.Sheet["任务单"].Rows[3].Cells[7].SetString("Job #: " + strings.ToUpper(job.job_number))
  //create_date string
  file.Sheet["任务单"].Rows[4].Cells[0].SetString(string(time.Now().Format("01/02/2006")))
  //addtional_notes string
  //file.Sheet["任务单"].Rows[5].Cells[1].SetString(job.addtional_notes)

  //file.Sheet["任务单"].Rows[5].Cells[1].SetStyle(font)
  //program
  //file.Sheet["任务单"].Rows[4].Cells[7].SetString("Program: " + job.program)
  //supplier string
  ///file.Sheet["任务单"].Rows[5].Cells[7].SetString("Supplier: " + utils.FirstLtterToUpper(job.supplier))
  //buyer string
  //file.Sheet["任务单"].Rows[6].Cells[7].SetString("Buyer: " + utils.FirstLtterToUpper(job.buyer) + " (" + strings.ToUpper(job.dep) + ") ")
  //due string
  //file.Sheet["任务单"].Rows[7].Cells[7].SetString("Artwork due date: " + job.due)
  //packout string
  //file.Sheet["任务单"].Rows[8].Cells[7].SetString("Packout date: " + job.packout)
  //ship string
  ///file.Sheet["任务单"].Rows[9].Cells[7].SetString("Shipdate: " + job.ship)
  //instore string
  //file.Sheet["任务单"].Rows[10].Cells[7].SetString("In-store date: " + job.instore)
  //contact string
  //file.Sheet["任务单"].Rows[12].Cells[7].SetString("联系人: " + job.contact)

  //copy
  return newFolder, "./" + strings.ToUpper(job.job_number) + " 做稿"
}
