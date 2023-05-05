# goDie
Simple dice macros for creating systems

# Example
```golang
package main

import (
  "fmt"
  roll "github.com/StephanieSunshine/godie"
)

func main() {
  fmt.Println(roll.d20())
}
```

# Extras
You can also extend by defining your own die
```golang
func d256() uint {
  return Roll(256)
}
```
MIT License (c) 2023 Stephanie Sunshine
