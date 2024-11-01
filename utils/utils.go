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

/* We need this packagePath()
 * function to call the m4 files. */
func packagePath() string {
    _, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)
    splitpath := strings.Split(basepath, "/")
    path := strings.Join(splitpath[:len(splitpath)-1], "/")

    return path
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
    nodelist []string;
    inif bool;
}

/* This is a low level interface, refer to FMakeObject.AddRule() for 
 * a more high level interface (I will assume you're looking for that) */
func (fmake *FMakeObject) BuildRule(name string, params []string) string {
    if len(params) == 0 {
        return name + "()"
    }
    return strings.Join([]string{name, "(`", strings.Join(params, "', `"), "')"}, "")
}

func (fmake *FMakeObject) AddRule(name string, params []string) {
    if fmake.inif {
        fmake.body = append(fmake.body, "_indent()   " + fmake.BuildRule(name, params))
        return
    }
    fmake.body = append(fmake.body, fmake.BuildRule(name, params))
}

func (fmake *FMakeObject) Nodes(start int, end int) []string {
    return fmake.nodelist[start:end]
}

func (fmake *FMakeObject) Cmdn() []string {
    return []string{strings.Join(fmake.nodelist[1:], " ")}
}

func (fmake *FMakeObject) Compile() {
    lines, err := ReadLines(fmake.Name)

    if err != nil {
        Die("[ERROR]: Couldn't read FMakefile.")
    }

    for i := 0; i < len(lines); i++ {

        var line string = lines[i]
        fmake.nodelist = strings.Split(line, " ")
        var startnode string = fmake.nodelist[0]

        if strings.HasPrefix(startnode, "--") || len(strings.TrimSpace(startnode)) == 0 {
            continue
        }

        switch startnode {
            case "if":
                fmake.AddRule("_if", fmake.Cmdn())
                fmake.inif = true
            case "endif":
                fmake.inif = false
                fmake.AddRule("_endif", nil)
            case "gcc-build":
                fmake.AddRule("_gcc_build", fmake.Nodes(1, 2))
            case "rust-build":
                fmake.AddRule("_rust_build", fmake.Nodes(1, 2))
            case "go-build":
                fmake.AddRule("_go_build", fmake.Nodes(1, 2))
            case "g++-build":
                fmake.AddRule("_gpp_build", fmake.Nodes(1, 2))
            case "println":
                fmake.AddRule("_println", fmake.Cmdn())
            default:
                Die("[ERROR]: FMakefile syntax error.")
        }

    }
}

func (fmake *FMakeObject) Run() {
    WriteLines("tmp.m4", fmake.body)
    out, err := exec.Command("m4", packagePath() + "/m4/build.m4", "tmp.m4").Output()
    if err != nil {
        Die("[ERROR]: M4 compilation failed.")
    }
    WriteLines("tmp.sh", strings.Split(string(out), "\n"))
    exec.Command("sh", "tmp.sh").Run()
}
