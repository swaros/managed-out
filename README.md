# managed-out
golang stdout console ansii

## usage
install
`go get -u github.com/swaros/manout`

## example

```golang
package main

import (
	"fmt"

	"github.com/swaros/manout"
)

func main() {
	fmt.Println("Hello, World!")

	// plain output with colorcodes, what should not work
	var mo manout.MOut
	mo.OutLn("hello", manout.ForeCyan, " world", manout.ForeLightGreen, " you should see the colorcodes but no colors ")
	// now with colored support
	var coloredParser manout.Colored
	mo.SetParser(&coloredParser).OutLn("hello", manout.ForeCyan, " world", manout.ForeLightGreen, " now it should work ")

}
```
### output
![example output](https://github.com/swaros/docu-asset-store/blob/main/demo-manout.png)
