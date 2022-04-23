/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	_ "embed"
	"fmt"

	"github.com/imhshekhar47/ops-agent/cmd"
)

//go:embed LICENSE
var license string

func main() {
	fmt.Printf("\n%s\n\n", license)

	cmd.Execute()
}
