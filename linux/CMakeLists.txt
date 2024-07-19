cmake_minimum_required(VERSION 3.10)
SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -pthread")
set(PROJECT_NAME "wg_flutter_linux")
project(${PROJECT_NAME} LANGUAGES CXX)

set(PLUGIN_NAME "wg_flutter_linux_plugin")

list(APPEND PLUGIN_SOURCES
  "wg_flutter_linux_plugin.cc"
)

add_library(${PLUGIN_NAME} SHARED
  ${PLUGIN_SOURCES}
)

apply_standard_settings(${PLUGIN_NAME})

set_target_properties(${PLUGIN_NAME} PROPERTIES
  CXX_VISIBILITY_PRESET hidden)
target_compile_definitions(${PLUGIN_NAME} PRIVATE FLUTTER_PLUGIN_IMPL)

target_include_directories(${PLUGIN_NAME} INTERFACE
  "${CMAKE_CURRENT_SOURCE_DIR}/include")

# Add the Go shared library
add_library(wg_quick SHARED IMPORTED)
set_target_properties(wg_quick PROPERTIES
  IMPORTED_LOCATION "${CMAKE_CURRENT_SOURCE_DIR}/libwg_quick.so"
)

target_link_libraries(${PLUGIN_NAME} PRIVATE flutter)
target_link_libraries(${PLUGIN_NAME} PRIVATE PkgConfig::GTK)
target_link_libraries(${PLUGIN_NAME} PRIVATE wg_quick)

set(wg_flutter_linux_bundled_libraries
  "${CMAKE_CURRENT_SOURCE_DIR}/libwg_quick.so"
  PARENT_SCOPE
)

add_custom_command(TARGET ${PLUGIN_NAME} POST_BUILD
  COMMAND ${CMAKE_COMMAND} -E copy_if_different
  "${CMAKE_CURRENT_SOURCE_DIR}/libwg_quick.so"
  "$<TARGET_FILE_DIR:${PLUGIN_NAME}>/libwg_quick.so"
)
