
- Create a golang program to print the following things.

```
| Name     | Age | Score | Active |
|----------|-----|-------|--------|
| John Doe |  24 | 154.61| true   |
```

```go
package main

import (
	"fmt"
)

func main(){
    // name declaration
	var name string = "John Doe"
	fmt.Println(name)

    // age declaration
	var age int = 24
	fmt.Println(age)

    // score declaration
	var score float64 = 154.61
	fmt.Println(score)

    // boolean declaration
	var isActive bool = true
	fmt.Println(isActive)
}
```