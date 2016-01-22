package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "H<ello> 世界! 1<23> Go."

	// // 查找连续的小写字母
	// reg := regexp.MustCompile(`[a-z]+`)
	// fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// // ["ello" "o"]
	//
	// // 查找连续的非小写字母
	// reg = regexp.MustCompile(`[^a-z]+`)
	// fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// // ["H" " 世界！123 G" "."]

	// 查找行首以 H 开头，以空格结尾的字符串
	reg := regexp.MustCompile(`(?U)<.+>`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 "]

}
