
# GofeR
A **Go** reactive library inspired by [Vue.js](https://vuejs.org/guide/extras/reactivity-in-depth.html).


## Install
use `go get` to install this package 

```shell
go get github.com/nasan016/gofer
```

## Getting Started
To start, define the GofeR package and import the library.

```go
package main

import(
    gf "github.com/nasan016/gofer"
)
```

## Examples

```go
package main

import (
    "fmt"
    gf "github.com/nasan016/gofer"
)

func main() {
    x := gf.Ref("Hello")

    gf.WatchEffect(func() {
        fmt.Println(x.GetValue())
    })

    x.SetValue("World!")
}
```

**Output**

```shell
Hello
World!
```