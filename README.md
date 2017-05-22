## How to install
- Get the library using:
```
go get github.com/Tiked/FileEncryption
```
- Import it:
```go
import "github.com/Tiked/FileEncryption"
```

## How to use it

```go
package main
import "github.com/Tiked/FileEncryption"

func main() {
  // First initialize the chipher with your key. use a 32 bytes slice
  FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
  
  // Now encrypt a file with its path
  err := FileEncryption.Encrypter("/home/desktop/data.txt")
  if err != nil {
    panic(err.Error())
  }
  // Now you should see data.txt.enc in your desktop
  
  // To decrypt it
  err = FileEncryption.Decrypter("/home/desktop/data.txt.enc")
  if err != nil {
    panic(err.Error())
  }
  // Now you should see data.txt in your desktop
  
}

```

## Features
- Encrypts large files using stream encryption
- Per-file unique IV handling



