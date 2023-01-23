TPress
======

TPress was made for simulating keyboard input on Windows systems. This simulate keyboard input, is very low level. There are C++ and Go bindings provided.

For C/C++ see examples

And for golang do this


<code>go mod init $(PROJECT_NAME)</code>


<code>go get github.com/TevesManuel/TPress/golang</code>


in your workspace you need download libTPress.dll file, this is very important.
You can do this with this commands


<code>wget https://github.com/TevesManuel/TPress/raw/master/lib/libTPress.dll</code>


or


<code>curl https://github.com/TevesManuel/TPress/raw/master/lib/libTPress.dll</code>


Or simply download libTPress.dll file on https://github.com/TevesManuel/TPress/blob/master/lib/libTPress.dll


You can use the library writing this in your go file

``` Go
  package main

  import TPress "github.com/TevesManuel/TPress/golang"

  func main() {
    TPress.TypeText("Hello world")
  }
  ```
