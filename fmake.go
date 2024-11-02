/* ==================== FMAKE ====================
 * FMake is a build system, kind of like GNU Make,
 * but... Worse, in every single way. Or... I guess
 * not? They're just... Very different. FMake is a
 * build system written in Go and compiles to... M4??
 * It compiles to Shell script right after, so why
 * bother? Well, my friend; M4 is a macro processor,
 * Shell script is not. Hell, Shell script doesn't
 * even have macros! So M4 is more powerful when it
 * comes to my needs and dependency on macros.
 *
 * usage: fmake [-S] [--help, -h] [-v, --version]
 * flags:
 *     -h, --help: Show help screen
 *     -v, --version: Print version number.
 *     -S: Don't delete tmp files.
 * author: Elis Staaf
 * license: Apache-2.0
 */

package main

import (
	"fmake/utils"
	"fmt"
	"os"
	"strings"
)

var version string = "NET/1"
var usage = `
usage: fmake [-S] [-h, --help] [-v, --version]
flags:
    -h, --help: Show this help screen.
    -v, --version: Print version number.
    -S: Save all tmp files.
`

func main() {
    var fmake utils.FMakeObject 
    cwd, err := os.ReadDir("./")

    if err != nil {
        utils.Die("[ERROR]: Error getting files in current working directory.")
        os.Exit(1)
    }

    for _, file := range cwd {
        if strings.ToLower(file.Name()) == "fmakefile" {
            fmake.Name = file.Name()
        }
    }

    /* Pre compile-time flags. */
    if len(os.Args) >= 2 {
        switch os.Args[1] {
            case "--help", "-h":
                fmt.Println(usage)
                os.Exit(0)
            case "--version", "-v":
                fmt.Println("FMake", version, os.Getenv("HOSTTYPE"))
                os.Exit(0)
        }
    }

    if len(fmake.Name) == 0 {
        utils.Die("[ERROR]: Couldn't find an FMakefile in current directory.")
        os.Exit(1)
    }

    if len(os.Args) == 1 {
        fmake.Compile()
        fmake.Run()
        os.Remove("tmp.sh")
        os.Remove("tmp.m4")
        utils.Note("[INFO]: FMake compilation succeded. All tests pass!")
        os.Exit(0)
    }

    /* Compile-time flags. */
    switch os.Args[1] {
        case "-S":
            fmake.Compile()
            fmake.Run()
            utils.Note("[INFO]: FMake compilation succeded. All tests pass!")
            os.Exit(0)
        default:
            utils.Die("[ERROR]: Invalid flag")
            os.Exit(1)
    }
}
