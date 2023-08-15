package mystrings

import (
	"fmt"
	"strings"
)

func Start() {
	str := "Hello World Hello Hi and Bye Hi-Bye Hello $+ and Hey Hi %3"

	fmt.Println(strings.Contains(str, "Hello"))

	fmt.Println(strings.Compare("Bigger in Length", "Small"))
	fmt.Println(strings.Compare("Big", "Smaller in Length"))
	fmt.Println(strings.Compare("Same", "Same"))

	fmt.Println(strings.HasPrefix(str, "Hello"))
	fmt.Println(strings.HasPrefix(str, "No No No"))

	fmt.Println(strings.HasSuffix(str, "%3"))
	fmt.Println(strings.HasSuffix(str, "Nope"))

	fmt.Println(strings.Split(str, "and"))

	str1 := "A_B"
	str2 := "B_C"
	str3 := "C_D"
	fmt.Println(strings.Join([]string{str1, str2, str3}, "$$$$"))

	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}
