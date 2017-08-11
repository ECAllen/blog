
### func declartion
func ( argname type ) return type { ... }

### list dir files/dir
```Go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```

### recursively list files
```Go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func printFile(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Print(err)
        return nil
    }
    fmt.Println(path)
    return nil
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    err := filepath.Walk(dir, printFile)
    if err != nil {
        log.Fatal(err)
    }
}
```
