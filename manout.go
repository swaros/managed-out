/*MIT License

Copyright (c) 2022 Thomas Ziegler, <thomas.zglr@googlemail.com>. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// manout package wraps a couple of most used fmt functions.
// the mapping is not nessesary to get colored output. it makes
// the usage just a little bit easier.
// as an example:
//     mo, parser := manout.NewColoredOut()
//     manout.Om.SetOM(mo)
//
//     manout.Om.Println("peter", manout.ForeLightGreen, "jackson")
//     manout.Om.Fprintf(os.Stdout, "Hello <f:white>%s</> \n", "samurai")
//
package manout

import (
	"fmt"
	"io"

	"github.com/swaros/outinject"
)

// list of writers accessible by the keyname
type OutWrapper struct {
	mo outinject.OutputManager
}

var Om *OutWrapper = &OutWrapper{mo: outinject.NewStdout()}

func (Ow *OutWrapper) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(Ow.mo.ToString(a...))
}

func (Ow *OutWrapper) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(Ow.mo.ToString(a...))
}

// Printf maps the fmt.Printf func
func (Ow *OutWrapper) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(Ow.pString(format), Ow.parseInterfaces(a...)...)
}

func (Ow *OutWrapper) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, Ow.pString(format), Ow.parseInterfaces(a...)...)
}

func (Ow *OutWrapper) SetOM(mo outinject.OutputManager) {
	Ow.mo = mo
}

func (Ow *OutWrapper) pString(s string) string {
	prs := *Ow.mo.GetParser()
	return prs.Parse(s)
}

func (Ow *OutWrapper) parseInterfaces(a ...interface{}) []interface{} {
	var b []interface{}
	prs := *Ow.mo.GetParser()
	for _, v := range a {
		switch v.(type) {
		case string:
			n := prs.Parse(v)
			b = append(b, n)

		default:
			b = append(b, v)
		}

	}
	return b
}
