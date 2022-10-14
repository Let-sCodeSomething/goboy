# GoBoy

⚠️**Currently, In Development**⚠️

Game boy emulation with golang


## Control Key

| Key          | Action      |
|:-------------|:------------|
| `ArrowUP`    | Up          |
| `ArrowDown`  | Down        |
| `ArrowLeft`  | Left        |
| `ArrowRight` | Right       |
| `a` OR `q`   | A           |
| `z` OR `w`   | B           |
| `Enter`      | Start       |
| `Space`      | Select      |
| `R`          | Change Game |


## Usage

for normal run

```bash
$ go run .
```

for debug run

```bash
$ go run . -d
```

## Create your Game Boy in Go

```go
package main

import (
    goboy "github.com/Let-sCodeSomething/goboy/gb"
)


func main() {
	gb := new(goboy.Goboy)
	// mode :
	//		"debug" = open window with debugging information
	// 		"normal" = open the simple game window
	err := gb.Init("you path of rom", "normal")
	
	if err != nil {
        panic(err)
    }	
	// start game
	go gb.Run()
	
	// for open window
	gb.WindowMode()
	// or for get the current frame
	// gb.APIMode()
}
```

## Collaborators

- [mkarten](https://github.com/mkarten)
- [tot0p](https://github.com/tot0p)