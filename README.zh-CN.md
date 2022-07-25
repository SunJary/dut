# dut

一个跨平台的，统计文件及目录大小的命令行工具，开发这个工具的目的是解决Windows下没有du命令，无法快速统计各个目录大小，以便清理占用磁盘的大文件。
以文件树的形式打印文件及目录大小

dut du 命令 加 tree 命令 

## Install

使用 [go](http://go.dev) 语言开发 

```sh
$ go install github.com/SunJary/dut
```

## Usage


```sh
$ dut -l 2
# 打印目录及文件大小，最多打印两层目录
```

```sh
$ dut -l 2 /root /data
# 打印指定目录
```

## Maintainers

[@SunJary](https://github.com/SunJary).
