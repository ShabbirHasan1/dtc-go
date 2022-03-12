cmake_minimum_required(VERSION 3.21)
project(testdata)

set(CMAKE_CXX_STANDARD 20)

add_executable(testdata DTCProtocol.h DTCProtocolVLS.h DTCProtocol_NonStandard.h main.cpp)
