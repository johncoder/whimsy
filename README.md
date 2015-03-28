### Whimsy

A simple Go library for generating phrases.

```bash
go get github.com/johncoder/whimsy
```

Example:

```go
package main

import (
	"fmt"

	"github.com/johncoder/whimsy"
)

func main() {
	fmt.Prinln("A phrase of 4 words:")
	fmt.Println(whimsy.NewPhrase(4))
}
```

Output:

```
A phrase of 4 words:
gopher death punch shampoo
```