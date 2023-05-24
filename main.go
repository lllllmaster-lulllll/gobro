package main

import (
	"fmt"
	"time"
)

type OmitEmptyTest struct {
	Name   string `json:"name,omitempty"`
	Age    int64  `json:"age,omitempty"`
	Number string `json:"number"`
}

const s string = "constant"

const (
	num1 = iota
	num2
	num3
	num4
)

func main() {
	// rule := factory.NewIRuleConfigParserFactory("json")
	// fmt.Println(rule)
	// ot := OmitEmptyTest{}
	// testStr := `{"age":8, "number":""}`
	// json.Unmarshal([]byte(testStr), &ot)

	// fmt.Printf("-%s-%d-%s-\n", ot.Name, ot.Age, ot.Number)
	// data, err := json.Marshal(ot)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(data))
	// fmt.Println(fmt.Sprintf("proxy_gw_%s_%s", "bjmc", "s1"))

	// image := "Windows-Server-2016-x64-20200826"
	// imageName := strings.Split(image, "-")
	// fmt.Println(strings.Join(imageName[:len(imageName)-1], "-"))
	// fmt.Println("kangnan")

	fmt.Println(time.Now().UnixNano())

}
