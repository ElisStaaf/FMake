package utils

import (
	"bufio"
	"os"
    "strings"
    "os/exec"
    "path/filepath"
    "runtime"

    "github.com/fatih/color"
)

/*** File I/O ***/

func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func WriteLines(path string, lines []string) (error) {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    for _, line := range lines {
        _, err := file.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }
    return nil
}

/*** Misc ***/

/* We need this PackagePath()
 * function to call the m4 files. */
func PackagePath() string {
    _, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)
    splitpath := strings.Split(basepath, "/")
    path := strings.Join(splitpath[:len(splitpath)-1], "/")

    return path
}

/* This is just because I'm lazy. */
func M4Rule(name string, params []string) string {
    return strings.Join([]string{name, "(`", strings.Join(params, "', `"), "')"}, "")
}

func Note(msg string) {
    color.RGB(0, 255, 0).Println(msg)
    os.Exit(0)
}

func Die(msg string) {
    color.RGB(255, 0, 0).Println(msg)
    os.Exit(0)
}

/*** FMakeObject ***/

type FMakeObject struct {
    Name string;
    body []string;
}

func (fmake *FMakeObject) Compile() {
    lines, err := ReadLines(fmake.Name)
    if err != nil {
        Die("[ERROR]: Couldn't read FMakefile.")
    }
    for i := 0; i < len(lines); i++ {
        var line string = lines[i]
        var nodelist []string = strings.Split(line, " ")
        var startnode string = nodelist[0]
        if startnode == "gcc-build" {
            fmake.body = append(fmake.body, M4Rule("_gcc_build", []string{nodelist[1], nodelist[2]}))
        } else if startnode == "rust-build" {
            fmake.body = append(fmake.body, M4Rule("_rust_build", []string{nodelist[1], nodelist[2]}))
        } else if startnode == "go-build" {
            fmake.body = append(fmake.body, M4Rule("_go_build", []string{nodelist[1], nodelist[2]}))
        } else if startnode == "g++-build" {
            fmake.body = append(fmake.body, M4Rule("_gpp_build", []string{nodelist[1], nodelist[2]}))
        } else if startnode == "println" {
            fmake.body = append(fmake.body, M4Rule("_println", []string{strings.Join(nodelist[1:], " ")}))
        } else if strings.HasPrefix(startnode, "--") {
            continue
        } else {
            Die("[ERROR]: FMakefile syntax error.")
        }
        WriteLines("tmp.m4", fmake.body)
        out, err := exec.Command("m4", PackagePath() + "/m4/build.m4", "tmp.m4").Output()
        if err != nil {
            Die("[ERROR]: M4 compilation failed.")
        }
        WriteLines("tmp.sh", strings.Split(string(out), "\n"))
        exec.Command("sh", "tmp.sh").Run()
    }
}
