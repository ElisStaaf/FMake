package main

import (
	"fmake/utils"
	"fmt"
	"os"
	"strings"
)

var usage = `
usage: fmake [-S] [-h, --help]
flags:
    -h, --help: Show this help screen.
    -S: Save all tmp files.
`

func main() {
    var fmake utils.FMakeObject 
    cwd, err := os.ReadDir("./")
    if err != nil {
        fmt.Println("[ERROR]: Error getting files in current working directory.")
        os.Exit(1)
    }
    for _, file := range cwd {
        if strings.ToLower(file.Name()) == "fmakefile" {
            fmake.Name = file.Name()
        }
    }
    if len(fmake.Name) == 0 {
        fmt.Println("[ERROR]: Couldn't find an FMakefile in current directory.")
        os.Exit(1)
    }
    if len(os.Args) == 1 {
        fmake.Compile()
        os.Remove("tmp.sh")
        os.Remove("tmp.m4")
        os.Exit(0)
    }
    if os.Args[1] == "--help" || os.Args[1] == "-h" {
        fmt.Println(usage)
        os.Exit(0)
    } else if os.Args[1] == "-S" {
        fmake.Compile()
        os.Exit(0)
    }
}
