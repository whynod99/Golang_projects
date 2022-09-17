package main

import (
		"fmt"
		"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
	// var now time.Time = time.Now()
	// var year int = now.Year()
	// fmt.Println(year)
}
