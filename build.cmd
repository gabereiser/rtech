@echo off
set BUILD_TYPE=Debug 

git submodule update --init --recursive

mkdir out
cd out
cmake -D CMAKE_INSTALL_PREFIX="./install" -DBGFX_BUILD_TOOLS=OFF ^
-DBGFX_BUILD_EXAMPLES=OFF -DBGFX_INSTALL=OFF -DCMAKE_BUILD_TYPE=$BUILD_TYPE ^
..
dir
msbuild rtech.vcxproj /m /p:configuration=%BUILD_TYPE% /p:Platform=x64
