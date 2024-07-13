package main

import (
	"github.com/fonsecabc/go-basic-api/configs"
)

func main() {
	err := configs.LoadVariables()
	if err != nil {
		panic(err)
	}

}
