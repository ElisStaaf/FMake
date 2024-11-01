.PHONY: install remove

SRC="fmake.go"
BIN="/usr/bin/fmake"

install:
	go build -o ${BIN} ${SRC}

remove:
	go clean
	rm ${BIN}
