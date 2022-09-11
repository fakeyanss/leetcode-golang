#!/usr/bin/env bash

echo "" > coverage.txt

modules=(helper leetcode)

for mod in ${modules[@]}; do
    cd $mod
    go test -covermode=count -coverprofile=profile.out ./... | grep 'no test files'
    if [ -f profile.out ]; then
        go tool cover -func=profile.out >> ../coverage.txt
        rm profile.out
    fi
    cd ..
done

cat coverage.txt|grep -v '100.0%'