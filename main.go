package main

import (
	"design_pattern/design_pattern/factory"
	"fmt"
)

func main() {
	rule := factory.NewIRuleConfigParserFactory("json")
	fmt.Println(rule)
}
