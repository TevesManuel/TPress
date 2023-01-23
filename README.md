TPress
======

TPress was made for simulating keyboard input on Windows systems. This simulate keyboard input, is very low level. There are C++ and Go bindings provided.

An C/C++ example


``` C

#include <stdio.h>
#include <TPress.h>

int main()
{
    const char * text;
    
    printf("Escribe el texto a simular: ");
    scanf("%[^\n]%*c", text);
    TypeText(text);
    return 0;
}

```


And for golang do this


<code>go mod init $(PROJECT_NAME)</code>


<code>go get github.com/TevesManuel/TPress/golang</code>


in your workspace you need download libTPress.dll file, this is very important.
You can do this with this commands


``` Bash
wget https://github.com/TevesManuel/TPress/raw/master/lib/libTPress.dll

```

or


``` Bash
curl https://github.com/TevesManuel/TPress/raw/master/lib/libTPress.dll

```


Or simply download libTPress.dll file on https://github.com/TevesManuel/TPress/blob/master/lib/libTPress.dll


You can use the library writing this in your go file

``` Go
  package main

  import TPress "github.com/TevesManuel/TPress/golang"

  func main() {
    TPress.TypeText("Hello world")
  }
  ```
