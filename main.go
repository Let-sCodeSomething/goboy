package main

import (
	"flag"
	goboy "github.com/Let-sCodeSomething/goboy/gb"
	"log"
)

var debug = flag.Bool("d", false, "debug mode")

func init() {
	flag.Parse()
}

func main() {
	gb := new(goboy.Goboy)
	// mode :
	//		"debug" = open window with debugging information
	// 		"normal" = open the simple game window
	var err error
	if *debug {
		err = gb.Init("", "debug")
	} else {
		err = gb.Init("", "normal")
	}
	if err != nil {
		log.Fatal(err)
	}
	go gb.Run()
	gb.WindowMode()
}
