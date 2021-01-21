package main

import (
	utils "github.com/salman-trimulabs/go-util"
)

func main(){
	x :=  utils.Init{
		"A", "B",
	}
	x.CopyImageFromDockerRepository()
}