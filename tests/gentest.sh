#!/bin/sh

function print () {
    echo -e "\e[$1;0m$2"
}

WORKDIR=$1
mkdir "$WORKDIR"

print "33" "[INFO]: Generating $WORKDIR/FMakefile"
print "33" "-- FMakefile [$HOSTTYPE]" >> "$WORKDIR/FMakefile"

for param in ${@[@]:1}; do
    print "33" "[INFO]: Generating $WORKDIR/$param..."
    touch "$WORKDIR/$param"

print "32" "[INFO]: Test generation was a success!"
