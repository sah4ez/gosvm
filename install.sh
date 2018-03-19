#!/bin/sh
TMP=/tmp/gosvm-build-$(date +%s)
git clone https://github.com/sah4ez/gosvm.git $TMP

command -v vgo || go get -u golang.org/x/vgo;

make $TMP/Makefile 

rm -rf $TMP
