#!/bin/sh
go get github.com/sah4ez/gosvm

command -v vgo || go get -u golang.org/x/vgo;

make -f $GOPATH/src/github.com/sah4ez/gosvm/Makefile 
