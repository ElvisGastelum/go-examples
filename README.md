# go-examples
Here are some examples in golang

## Web server
To use the simple web server you need import the package and run the func Serve.
example
```go
package main

import "github.com/elvisgastelum/go-examples/webserver"

func main(){
  webserver.Serve(":5504")
}
```