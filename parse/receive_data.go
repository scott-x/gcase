package parse

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    //"github.com/tealeg/xlsx"
    //"time"
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

func Receive_data() *Data {
    job_number := question("请输入工单号: ")

    short_brand := question("请输入系列简称: ")

    _, ok := GetBrand()[strings.ToUpper(short_brand)] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if ok {
    } else {
        fmt.Println("请输入正确简写格式的brand……")
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
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Printf(q)
    inputData, err := inputReader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    inputData = strings.Trim(inputData, "\n")
    return inputData
}
