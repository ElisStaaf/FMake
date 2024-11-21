#!/bin/sh

WORKDIR=$1

if [ ! -d $WORKDIR ]; then
    mkdir "$WORKDIR"
fi

echo "fmake: Generating $WORKDIR/FMakefile"
echo "require $2" >> "$WORKDIR/FMakefile"
echo "set PAKG_VERSION \"1.0.0\"" >> "$WORKDIR/FMakefile"
echo "set PAKG_NAME $(basename "$WORKDIR")" >> "$WORKDIR/FMakefile"
echo "println \"\$PAKG_NAME -- Version \$PAKG_VERSION\"" >> "$WORKDIR/FMakefile"

echo "fmake: Project generation was a success!"
