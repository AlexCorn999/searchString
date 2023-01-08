package main

import (
	"flag"
	"fmt"
)

func searchString(str []rune, substr []rune) bool {
	flag := false
	j := 0
	for i := 0; i < len(str); i++ {

		if j < len(substr) {

			if str[i] == substr[j] && flag == false {
				flag = true
			}

			if flag == true {
				if j == len(substr)-1 {
					break
				} else if str[i] == substr[j] {
					flag = true
				} else {
					flag = false
					j = 0
					continue
				}
				j++
			}

		}
		if i == len(str)-1 && j < len(substr) {
			flag = false
		}
	}
	return flag
}

func main() {
	str := new(string)
	substr := new(string)
	flag.StringVar(str, "str", "default string", "set string")
	flag.StringVar(substr, "substr", "default substring", "set substring")
	flag.Parse()
	fmt.Println(*str, *substr)

	strS := []rune(*str)
	substrS := []rune(*substr)
	fmt.Println(searchString(strS, substrS))
}
