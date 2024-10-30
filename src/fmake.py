#!/bin/python3

"""
Convert "fmake" language to standard
macro "m4" macro processing.

author :: Elis Staaf
python :: >=3.0
license :: Apache 2.0 License
"""

import sys
import os

if sys.version_info < (3, 0):
    print("[ERROR]: Python3 is required to use Fmake.")
    exit(1)

class Fmake:
    def __init__(self, filename: str) -> None:
        with open(filename, "r") as file:
            self.lines: list[str] = file.readlines()
            file.close()
        self.program: list[str] = []

    def eval_expr(self) -> None:
        for line in self.lines:
            nodelist: list[str] = line.split(" ")
            startnode: str = nodelist[0]
            if startnode == "gcc-build":
                self.program.append("_gcc_build(`%s', `%s')" % nodelist[2], nodelist[1])
            elif startnode == "go-build":
                self.program.append("_go_build(`%s', `%s')" % nodelist[2], nodelist[1])
            elif startnode == "rust-build":
                self.program.append("_rust_build(`%s', `%s')" % nodelist[2], nodelist[1])
            elif startnode == "g++-build":
                self.program.append("_gpp_build(`%s', `%s')" % nodelist[2], nodelist[1])
            elif startnode.startswith("--") or startnode.strip() == "":
                continue

def main() -> None:
    if "FMakefile".lower() in os.listdir():
        fmake = Fmake("FMakefile".lower())
    elif len(sys.argv >= 2):
        fmake = Fmake(sys.argv[1])
    else:
        print("[ERROR]: No FMakefile provided.")
        exit(1)
    fmake.eval_expr()
    with open("tmp.m4", "w+") as macros:
        macros.write("\n".join(fmake.program))
        macros.close()
    os.system("m4 m4/build.m4 tmp.m4 > tmp.sh")
    os.system("sh tmp.sh")
