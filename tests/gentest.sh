#!/bin/sh

WORKDIR=$1

if [ -d $WORKDIR ]; then
    exit "[ERROR]: Directory \"$WORKDIR\" already exists!"
fi

mkdir "$WORKDIR"

echo "[INFO]: Generating $WORKDIR/FMakefile"
echo "-- FMakefile [$HOSTTYPE]" >> "$WORKDIR/FMakefile"

for param in $@; do
    if [ $param == $1 ]; then
        continue
    fi
    echo "[INFO]: Generating $WORKDIR/$param..."
    touch "$WORKDIR/$param"
done

echo "[INFO]: Test generation was a success!"
