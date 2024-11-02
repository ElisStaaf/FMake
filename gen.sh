#!/bin/sh

WORKDIR=$1

if [ -d $WORKDIR ]; then
    exit "[ERROR]: Directory \"$WORKDIR\" already exists!"
fi

mkdir "$WORKDIR"

echo "[INFO]: Generating $WORKDIR/FMakefile"
echo "set PAKG_VERSION \"1.0.0\"" >> "$WORKDIR/FMakefile"
echo "set PAKG_NAME $(basename "$WORKDIR")" >> "$WORKDIR/FMakefile"
echo "println \"\$PAKG_NAME -- Version \$PAKG_VERSION\"" >> "$WORKDIR/FMakefile"

for param in $@; do
    if [ $param == $1 ]; then
        continue
    fi
    echo "[INFO]: Generating $WORKDIR/$param..."
    touch "$WORKDIR/$param"
done

echo "[INFO]: Project generation was a success!"
