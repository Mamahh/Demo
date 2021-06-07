#### usage:

```shell
$ tree
.
├── add.c
├── CMakeLists.txt
├── include
│   └── main.h
├── lib
│   └── libadd.a
└── main.c

2 directories, 5 files

$ cmake .
-- The C compiler identification is GNU 8.3.0
-- The CXX compiler identification is GNU 8.3.0
-- Check for working C compiler: /usr/bin/cc
-- Check for working C compiler: /usr/bin/cc -- works
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
-- Detecting C compile features
-- Detecting C compile features - done
-- Check for working CXX compiler: /usr/bin/c++
-- Check for working CXX compiler: /usr/bin/c++ -- works
-- Detecting CXX compiler ABI info
-- Detecting CXX compiler ABI info - done
-- Detecting CXX compile features
-- Detecting CXX compile features - done
-- THis is source dirmain.c
-- Configuring done
-- Generating done
-- Build files have been written to: /home/lkh/test/C/cmake/cmake-demo/hello_world-12

$ make
Scanning dependencies of target demo
[ 50%] Building C object CMakeFiles/demo.dir/main.c.o
[100%] Linking C executable demo
[100%] Built target demo

$ ./demo 
Hello World
r = 5
```



#### Explanation:

```cmake
CMAKE_MINIMUM_REQUIRED(VERSION 3.5) #规定cmake最低版本
PROJECT (HELLO) #工程名
LINK_DIRECTORIES(lib) #库文件路径
set(MYSRC_SRC main.c ) #用变量代替值
FILE(GLOB MYSRC_SRC  #GLOB 遍历指定文件 
     	${PROJECT_SOURCE_DIR}/*.c
     	${PROJECT_SOURCE_DIR}/file/*.c)
include_directories(include) #头文件路径
MESSAGE(STATUS "THis is source dir" ${MYSRC_SRC}) #打印消息，可dbug用
ADD_EXECUTABLE(demo  ${MYSRC_SRC} )  #将${MYSRC_SRC}的代码编译成名为demo的可执行文件
target_link_libraries(demo libadd.a) #添加链接库
```



