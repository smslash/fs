package main

import (
	"os"

	"git/ssengerb/ascii-art-fs/logic"
)

func main() {
	length := len(os.Args)

	if length == 2 {
		if !logic.Characters(os.Args[1]) {
			logic.Error(2)
			return
		}
		logic.Default(os.Args[1])
	} else if length == 3 {
		if !logic.Characters(os.Args[1]) {
			logic.Error(2)
			return
		}
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			logic.Error(3)
			return
		}
		logic.Check(os.Args)
	} else {
		logic.Error(1)
		return
	}
}
