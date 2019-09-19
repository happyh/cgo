#!/usr/bin/env bash
set -x

swig -go -cgo -intgosize 64 -c++ ./go_test_cpp.swigcxx

cp ./go_test_cpp.go ./../go_src/go_test_cpp/
