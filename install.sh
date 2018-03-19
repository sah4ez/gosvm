#!/bin/sh
go get -d github.com/sah4ez/gosvm/cmd/gosvm

command -v vgo || go get -u golang.org/x/vgo;

make -f $GOPATH/src/github.com/sah4ez/gosvm/Makefile 
