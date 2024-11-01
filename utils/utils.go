package utils

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "os/exec"
    "path/filepath"
    "runtime"
)

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

func PackagePath() string {
    _, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)
    splitpath := strings.Split(basepath, "/")
    path := strings.Join(splitpath[:len(splitpath)-1], "/")

    return path
}

type FMakeObject struct {
    Name string;
    body []string;
}

func M4Rule(name string, params []string) string {
    return strings.Join([]string{name, "(`", strings.Join(params, "', `"), "')"}, "")
}

func (fmake *FMakeObject) Compile() {
    lines, err := ReadLines(fmake.Name)
    if err != nil {
        fmt.Println("[ERROR]: Couldn't read FMakefile.")
        os.Exit(1)
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
            fmake.body = append(fmake.body, M4Rule("_println", nodelist[1:]))
        } else if strings.HasPrefix(startnode, "--") {
            continue
        } else {
            fmt.Println("[ERROR]: FMakefile syntax error.")
            os.Exit(1)
        }
        WriteLines("tmp.m4", fmake.body)
        out, err := exec.Command("m4", PackagePath() + "/m4/build.m4", "tmp.m4").Output()
        fmt.Println(out)
        if err != nil {
            fmt.Println("[ERROR]: M4 compilation failed.")
            os.Exit(1)
        }
        WriteLines("tmp.sh", strings.Split(string(out), "\n"))
        exec.Command("sh", "tmp.sh").Run()
    }
}
