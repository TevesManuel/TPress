# TPress
# library to emulate input keyboard maded in C, with bindings to C++ and Go

For C/C++ see examples

And for golang do this


<code>go mod init $(PROJECT_NAME)</code>


<code>go get github.com/TevesManuel/TPress/golang</code>


You can use the library writing this in your go file

<code>
  package main

  import TPress "github.com/TevesManuel/TPress/golang"

  func main() {
    TPress.TypeText("Hello world")
  }
</code>
