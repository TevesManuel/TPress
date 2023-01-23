TPress
======

TPress was made for simulating keyboard input on Windows systems. This simulate keyboard input, is very low level. There are C++ and Go bindings provided.

For C/C++ see examples

And for golang do this


<code>go mod init $(PROJECT_NAME)</code>


<code>go get github.com/TevesManuel/TPress/golang</code>


You can use the library writing this in your go file

``` Go
  package main

  import TPress "github.com/TevesManuel/TPress/golang"

  func main() {
    TPress.TypeText("Hello world")
  }
  ```
