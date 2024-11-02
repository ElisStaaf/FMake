#!/bin/sh

WORKDIR=$1

if [ -d $WORKDIR ]; then
    exit "[ERROR]: Directory \"$WORKDIR\" already exists!"
fi

mkdir "$WORKDIR"

echo "[INFO]: Generating $WORKDIR/FMakefile"
echo "require $2" >> "$WORKDIR/FMakefile"
echo "set PAKG_VERSION \"1.0.0\"" >> "$WORKDIR/FMakefile"
echo "set PAKG_NAME $(basename "$WORKDIR")" >> "$WORKDIR/FMakefile"
echo "println \"\$PAKG_NAME -- Version \$PAKG_VERSION\"" >> "$WORKDIR/FMakefile"

echo "[INFO]: Project generation was a success!"
