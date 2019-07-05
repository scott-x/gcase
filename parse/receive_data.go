package parse

import (
    "bufio"
    "os"
    "strings"
    "github.com/fatih/color"
    "regexp"
)

type Data struct {
    job_number      string
    short_brand     string
    program         string
    country         string
    addtional_notes string
    supplier        string
    buyer           string
    dep             string
    due             string
    packout         string
    ship            string
    instore         string
    contact         string
}

type Data_CAB struct {
    job_number string
}

type Data_Print struct {
    job_number string
    po         string
}

func check_job(re,s string) bool{
  Re := regexp.MustCompile(re)
  match := Re.FindString(s)
  if len(match) > 0 {
    return true
  }else{
    return false
  }
}

func Receive_data(n int) *Data {
    job_number := question("请输入工单号: ")
    switch n{
    case 1:
        re :=`[Uu][12][980][01][0-9][0-9a-zA-Z][0-9]_[a-zA-z]{3}`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
    case 2:
        re :=`[Cc][12][980][01][0-9][0-9a-zA-Z][0-9]_[a-zA-z]{3}`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
    case 4:
        re :=`[Bb][12][980][01][0-9][0-9a-zA-Z][0-9]_[Ll][Nn][Cc]`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
    case 6:
        re :=`[Bb][12][980][01][0-9][0-9a-zA-Z][0-9]_[a-zA-z]{3}`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
                
    }

    short_brand := question("请输入系列简称: ")
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    for{
        _, ok := GetBrand()[strings.ToUpper(short_brand)] /*如果确定是真实的,则存在,否则不存在 */
        if ok {
            break
        } else {
            short_brand = question("请输入正确的系列简称: ")
        }
    }

    country := question("国家（eg USA, India）: ")
    addtional_notes := question("注意事项: ")
    supplier := question("Supplier: ")
    buyer := question("buyer: ")
    dep := question("dep: ")
    due := question("Artwork due date: ")
    packout := question("Artwork packout date: ")
    ship := question("Artwork ship date: ")
    instore := question("Artwork in store date: ")
    contact := question("contact: ")
    job := &Data{
        job_number,
        short_brand,
        GetBrand()[strings.ToUpper(short_brand)],
        country,
        addtional_notes,
        supplier,
        buyer,
        dep,
        due,
        packout,
        ship,
        instore,
        contact,
    }
    return job
}

func question(q string) string {
    cyan := color.New(color.FgCyan)
    boldCyan := cyan.Add(color.Bold)
    inputReader := bufio.NewReader(os.Stdin)
 
    boldCyan.Printf(q)
    inputData, err := inputReader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    inputData = strings.Trim(inputData, "\n")
    return inputData
}

func Receive_Print(n int) *Data_Print {
    job_number := question("请输入工单号: ")
    
    switch n {
    case 5:
        re :=`[Pp][12][980][01][0-9][0-9a-zA-Z][0-9]_[Ll][Nn][Cc]`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
        
    }
    po := question("请输入订单(PO): ")
    job := &Data_Print{
        job_number,
        po,
    }
    return job
}

func Receive_CAB(n int) *Data_CAB {
    job_number := question("请输入工单号: ")
    switch n {
    case 3:
        re :=`[B][12][980][01][0-9][0-9a-zA-Z][0-9]_[Cc][Aa][Bb]`
        for {
            if check_job(re,job_number) {
               break
            }else{
              job_number = question("请输入正确的工单号: ")
            }
        }
        
    }
    job := &Data_CAB{
        job_number,
    }
    return job
}
