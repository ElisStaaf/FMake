.PHONY: install remove

SRC="fmake.go"
BIN="/usr/local/bin/fmake"

all: install

install:
	go build -o ${BIN} ${SRC}

remove:
	go clean
	rm ${BIN}
