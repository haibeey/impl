cmake_minimum_required(VERSION 3.16)
project(ImageBlurSimd LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

include(FetchContent)


FetchContent_Declare(
  simde
  GIT_REPOSITORY https://github.com/simd-everywhere/simde.git
  GIT_TAG        master
)

if (NOT simde_POPULATED)
  FetchContent_Populate(simde)
endif()


add_library(imageblur STATIC
    ImageBlur.cpp
)


target_include_directories(imageblur
    PUBLIC
        ${CMAKE_CURRENT_SOURCE_DIR}
        ${simde_SOURCE_DIR}
        ${simde_SOURCE_DIR}/simde
)
