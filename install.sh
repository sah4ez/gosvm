#!/bin/sh
go get https://github.com/sah4ez/gosvm.git

command -v vgo || go get -u golang.org/x/vgo;

make -f $GOPATH/src/github.com/sah4ez/gosvm/Makefile 
