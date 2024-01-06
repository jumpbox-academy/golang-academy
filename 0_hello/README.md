# Hello World Golang

## Instruction

1. Create module via command `$ go mod init` 
```bash
go mod init hello
```

2. Create Entrypoint file `<name>.go`
```bash
touch main.go
```

Add then add content
```go
package main

import "fmt"

func main() {
   fmt.Println("Hello World golang")
}
```

3. Execute Command `$ go run <name>.go` for code translation and then run the executable file 
```bash
go run main.go
```