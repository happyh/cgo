cmake_minimum_required(VERSION 2.8)

project(test_cpp C CXX)

set(SRC_LISTS test_cpp.cpp)
add_library(test_cpp SHARED ${SRC_LISTS})
