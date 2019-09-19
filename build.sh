#!/bin/bash
#########################################################################
# File Name: build.sh
# Author: happyhe
# mail: heguang@qiyi.com
# Created Time: Thu 19 Sep 2019 03:38:40 PM CST
#########################################################################

cd cpp && mkdir -p cmake && cd cmake && cmake .. && make && cd ../..
cd swig && ./gen_go_cpp.sh && mkdir -p cmake && cd cmake && cmake .. && make && cd ../..
cd go_src && go run main.go

