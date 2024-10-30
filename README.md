# FMake: Build software for idiots
[![Build](https://img.shields.io/badge/Build%20(Fedora)-passing-2a7fd5?logo=fedora&logoColor=2a7fd5&style=for-the-badge)](https://github.com/ElisStaaf/vine)
[![Version](https://img.shields.io/badge/Version-1.0.0-38c747?style=for-the-badge)](https://github.com/ElisStaaf/vine)
[![Lang](https://img.shields.io/badge/Language-Python-2d6cd2?logo=python&style=for-the-badge)](https://github.com/ElisStaaf/vine)  
FMake is build software focused on working. Unlike other build software this one doesn't work half of the
time. It also (right now at least) doesn't include anything special. You can only just... Build stuff.

## Requirements
* python3
* pyinstaller
* m4
* git or gh

## Install
To install, firstly, clone the git repo:
```bash
# git
git clone https://github.com/ElisStaaf/FMake ~/fmake

# gh
gh repo clone ElisStaaf/Fmake ~/fmake
```
Then compile an executable:
```bash
make
```
The above compiling is done by pyinstaller, that's why you need it. Then you can unpack the executable,
and go on your merry way.

## Introduction to the FMakefile
The `FMakeFile` is a layer of abstraction, so that you don't have to compile with M4, the FMake compiler
does that for you. Say you have a project with a file called `main.rs`, you can create an `FMakefile`
and write this into it:
```
rust-build main.rs main
```
I'm not going to go *too* far into the low level interface of M4, but this is how your code expands
in the M4 compiled file.
```
_rust_build(`main', `main.rs')
```
And *that* expands to *this* in shell language:
```bash
rustc -o main main.rs
```
Comments in FMake start with `--`:
```
-- This is a comment, and it is awesome!
rust-build main.rs main
```
There are other compilers you can use in FMake, here's a showcase:
```
-- This is the rust compiler, the one I showed earlier:
rust-build main.rs main

-- This is the GCC compiler:
gcc-build main.c main

-- This is the G++ compiler:
g++-build main.cpp main

-- And this is the Go compiler:
go-build main.go main
```
And that's it for now! Feel free to contribute to the project, help is much appreciated!
And with that, enjoy!
