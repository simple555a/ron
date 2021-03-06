#!/bin/bash

clear
cat sep2bits.txt | ./sep2bits.pl > bitsep.go
go fmt bitsep.go
ragel -Z -G2 -e dfa.rl
go fmt dfa.go
go build
go test -timeout 5s
cd RDT
go test -timeout 5s
