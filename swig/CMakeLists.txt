cmake_minimum_required(VERSION 2.8)

project(go_test_cpp C CXX)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../cpp)
set(SRC_LISTS go_test_cpp_wrap.cxx)
add_library(go_test_cpp STATIC ${SRC_LISTS})
