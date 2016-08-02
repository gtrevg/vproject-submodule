package main

import (
	"fmt"
	"github.com/Comcast/golang-money"
	app "github.com/gtrevg/vproject-submodule/app"
)

func main() {
	fmt.Println("this is a test")
	m := money.Money{}
	fmt.Println(m.ToString())

	a := app.App{"ID-12345"}

	a.Identify()
}
