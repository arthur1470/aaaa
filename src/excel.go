package main

import (
	"C"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {

}

//export WriteAndReturn
func WriteAndReturn(fileEncoded string, index *[]int, cell *[]string, value *[]string) string {
	decodedFile, err := base64.StdEncoding.DecodeString(fileEncoded)

	if err != nil {
		log.Panic(err)
		return ""
	}

	excel, err := func() (*excelize.File, error) {
		buffer := bytes.NewBuffer(decodedFile)
		return excelize.OpenReader(buffer)
	}()

	if err != nil {
		log.Panic(err)
		return ""
	}

	buf := new(bytes.Buffer)
	_, err = excel.WriteTo(buf)

	if err != nil {
		log.Panic(err)
		return ""
	}

	excel.SaveAs("excel.xlsx")
	return base64.StdEncoding.EncodeToString(buf.Bytes())
	return "bread"
}

//export Add
func Add(x, y int) int {
	fmt.Printf("Go says: adding %v and %v\n", x, y)
	return x + y
}
