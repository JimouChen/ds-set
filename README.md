# ds-set
🚀 A set data structure implemented with go.

# 🔥 Usage && Quick Start
- setup: 
`import "github.com/JimouChen/ds-set/easyset"`
- `go mod tidy` in terminal or use `go get` to download package
- just code

# 🌰 Example

```go
package main

import (
	"fmt"
	"github.com/JimouChen/ds-set/easyset"
)

func main() {
	s := easyset.Set{}
	s.InitSet()
	s.Add("2")
	s.Add(1)
	s.Add(456)
	s.Show()
	_ = s.Delete(1)
	s.Show()
	fmt.Println(s.Length())
	fmt.Println(s.Contain("22"))
	s.Clear()
	fmt.Println(s.Length())
	s.Show()
}
```
- res

```shell
2 1 456 
2 456 
2
false
0


Process finished with the exit code 0
```
