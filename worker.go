package main

import (
	"Mustang/pkg/machinery"
	"Mustang/utils"
)

func main() {
	utils.InitDB()
	server := machinery.CreateMachinery()
	worker := server.NewWorker("mustang", 1)
	err := worker.Launch()
	if err != nil {
		panic(err)
	}
}
