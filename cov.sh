#! /bin/bash

# go test -coverprofile=cov.out -coverpkg=./lc ./lc
go test -covermode=count -coverprofile=count.out -coverpkg=./lc ./lc

# go tool cover -html=cov.data -o cov.html
go tool cover -func=count.out|grep -v '100.0%'

# open cov.html
