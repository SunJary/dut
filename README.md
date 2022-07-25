# dut

A cross-platform command-line tool that counts the size of files and directories. The purpose of developing this tool is to solve the problem that there is no du command under Windows, and it is impossible to quickly count the size of each directory in order to clean up large files that occupy disks.
Print file and directory size as a file tree
## Install

This project writen by [go](http://go.dev)

```sh
$ go install github.com/SunJary/dut
```

## Usage


```sh
$ dut -l 2
# Prints dir and size in a tree. print 2 level
```

```sh
$ dut -l 2 /root /data
# Prints two dir (/root, /data) in a tree. print 2 level
```

## Maintainers

[@SunJary](https://github.com/SunJary).
