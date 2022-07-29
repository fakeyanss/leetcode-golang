#! /bin/bash

go test -coverprofile=cov.data -coverpkg=./lc ./lc

go tool cover -html=cov.data -o cov.html

open cov.html
