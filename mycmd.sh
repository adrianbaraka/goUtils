#!/usr/bin/env bash

for i in {1..10}; do
    (( i % 2 == 0 )) && echo "Stderr $i" >&2 && continue

    echo "Time $i"
    sleep 1
done
