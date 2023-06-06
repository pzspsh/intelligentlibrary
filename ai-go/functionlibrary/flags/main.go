/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 17:35:36
*/
package main

import (
	"fmt"
	"log"

	"function/flags/flags"
)

type Options struct {
	name     string
	Email    flags.StringSlice
	Phone    string
	Address  flags.StringSlice
	fileSize flags.Size
}

func main() {
	testOptions := &Options{}
	CheckUpdate := func() {
		fmt.Println("checking if new version is available")
		fmt.Println("updating tool....")
	}

	flagSet := flags.NewFlagSet()
	flagSet.CreateGroup("info", "Info",
		flagSet.StringVarP(&testOptions.name, "name", "n", "", "name of the user"),
		flagSet.StringSliceVarP(&testOptions.Email, "email", "e", nil, "email of the user", goflags.CommaSeparatedStringSliceOptions),
	)
	flagSet.CreateGroup("additional", "Additional",
		flagSet.StringVarP(&testOptions.Phone, "phone", "ph", "", "phone of the user"),
		flagSet.StringSliceVarP(&testOptions.Address, "address", "add", nil, "address of the user", goflags.StringSliceOptions),
		flagSet.CallbackVarP(CheckUpdate, "update", "ut", "update this tool to latest version"),
		flagSet.SizeVarP(&testOptions.fileSize, "max-size", "ms", "", "max file size"),
	)

	if err := flagSet.Parse(); err != nil {
		log.Fatal(err)
	}
}
