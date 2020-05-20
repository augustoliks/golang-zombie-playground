package main

import (
	"fmt"
	"parentpack"
	"parentpack/childpack"
)

func main() {
	fmt.Println(parentpack.PublicVariable01)
	fmt.Println(parentpack.PublicVariable02)
	fmt.Println(childpack.PublicVariable01)
}
