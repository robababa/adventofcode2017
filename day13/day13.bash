#!/bin/bash

INPUT=./input.txt

echo "Part 1:"
sed 's# ##g' $INPUT | awk -F: '{DEPTH=$1; RANGE=$2; CYCLE=2*(RANGE-1); if (DEPTH%CYCLE==0) SUM+=DEPTH*RANGE} END {print SUM}'

echo "Part 2:"
for (( i=0; ; i++ ))
do
  SEVERITY=$(
    sed 's# ##g' $INPUT |
    awk -F: -v WAIT=$i '{DEPTH=$1; RANGE=$2; CYCLE=2*(RANGE-1); if ((WAIT+DEPTH)%CYCLE==0) SUM+=DEPTH*RANGE} END {print SUM}'
  )
  #echo $i $SEVERITY
  if [[ $SEVERITY -eq 0 ]]
  then
    echo $i
    break
  fi
done

