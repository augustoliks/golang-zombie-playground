package main

import (
	"fmt"
	"parentpack"
	"parentpack/childpack"
)

func main() {
	fmt.Println("parentpack", parentpack.PublicVariable01)
	fmt.Println("parentpack", parentpack.PublicVariable02)
	fmt.Println("parentpack/childpack", childpack.PublicVariable01)
	fmt.Println("parentpack/childpack", childpack.PublicVariable02)
}
