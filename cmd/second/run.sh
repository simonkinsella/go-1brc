#!/bin/zsh
set -e
INPUT=measurements-1k.txt

go build -o calculate-temps .
go version
echo input: ${INPUT}

/usr/bin/time -p sh -c './calculate-temps -in ../../datasets/'${INPUT}' > results.txt' | grep real
