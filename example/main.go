package main

import (
	utils "github.com/hakeemsyd/go-util"
)

func main(){
	x :=  utils.Init{
		"A", "B",
	}
	x.CopyImageFromDockerRepository();
}