# Assembler test
![](https://github.com/HagaSpa/Assembler/workflows/go%20test%20&%20go%20build/badge.svg)

assembler for nand2tetris

## Requirements　！！
* Docker 19.03 or later

You need a docker environment to build with Docker.

Also, since it uses Docker BuildKit, please using enabled the [Docker BuildKit flag](https://docs.docker.com/develop/develop-images/build_enhancements/) by using Docker 19.03 or later.


## Build
This Assembler can cross compile the binary for the host operating system by adding arguments to make command. 
(using cross compile in golang)

```
// compile the binary for you host operating system.
$ make

// macOS
$ make PLATFORM=darwin/amd64 

// linux
$ make PLATFORM=linux/amd64

// Windows
& make PLATFORM=windows/amd64

```

## Run
```
$ make
$ ls bin 
main
$ ./bin/main test.asm 

$ cat test.hack 
0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
```

`Note: Confirmed to work only on macOS now`
