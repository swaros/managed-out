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
	mo, parser := manout.NewStdColored() // create a new managed and colored output handler instance

	// lets loop 2 times
	for i := 0; i < 2; i++ {

		// in the first loop, the output should be colored.
		// if not, then the stdout is not accepted as a terminal.
		// this will for exampe happens in launch visual-studio-code jobs.
		// even they are able to handle the colors.
		// on any console you should see the first output colored. 
		// (bash, zsh, fish, powershell >= v7)
		// you can also force color for the first itertion by uncomment the next 3 lines ...
		//if i == 0 {
		//	parser.EnableColor()
		//}

		mo.OutLn(
			manout.BoldTag, manout.ForeMagenta,
			"======================================\nhello world ",
			manout.CleanTag, manout.ForeDarkGrey,
			" ... you need more ",
			manout.ForeLightGreen, manout.BackGreen,
			" GREEN ",
			manout.CleanTag,
			"....",
		)
		mo.OutLn() // just some space

		// markup used inline. without color, the markup is also removed
		mo.OutLn(`.... using <b>markup</> <f:yellow>inline</> ... 
     it is <f:red> NOT </> html. <b:green><f:white>just looks similiar</>`)

		mo.OutLn() // just some space again
		parser.DisableColor() // disable color for the next loop
	}

}
```
### output
![example output](https://github.com/swaros/docu-asset-store/blob/main/manout-demo2.png)
