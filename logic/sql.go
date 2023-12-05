package logic

import (
	"fmt"
	"strings"
)

func ReplaceDollarWithData(sql string, inputArgs []string) string {
	replaceVal := sql
	var replace *strings.Replacer
	incre := 1
	for _, data := range inputArgs {
		val := fmt.Sprintf("%d", incre)
		val1 := "$" + val
		replace = strings.NewReplacer(val1, data)
		replaceVal = replace.Replace(replaceVal)
		incre++
	}

	return replaceVal
}
