# go-exec-js
- [go-exec-js](#go-exec-js)
  - [Introduction](#introduction)
  - [Requirement](#requirement)
  - [Intall](#intall)
  - [Usage](#usage)
  - [Thanks to](#thanks-to)
## Introduction
参考PyExecJS，提供了在go语言里执行js的能力。由于go是强类型语言，所以在迁移的时候使用字符串作为替代传递任意类型的参数。

## Requirement
需要本地path里面配置了node。

## Intall
```
go get -u github.com/cokeBeer/execjs
```
## Usage
可以使用Eval方法获取表达式的值，这将输出`12`
```go
output, err := execjs.Eval(`1+"2"`)
if err != nil {
    log.Fatal(err)
}
fmt.Println(output)
```
可以使用Compile方法编译一个Context，然后调用。这将输出`3`
```go
c, err := execjs.Compile(`function add(x, y) {
    return x + y;
    }`)
if err != nil {
    log.Fatal(err)
}
output, err := c.Call("add", `1`, `2`)
if err != nil {
    log.Fatal(err)
}
fmt.Println(output)
```
注意参数一定要用反引号包裹
## Thanks to
[PyExecJS](https://github.com/doloopwhile/PyExecJS)