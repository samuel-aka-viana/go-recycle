package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	if s == "gopher" {
		fmt.Println("gopher is go")
	}

}
