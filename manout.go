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
package manout

import (
	"fmt"
	"io"
	"os"
)

// list of writers accessible by the keyname
var (
	writers map[string]io.Writer = make(map[string]io.Writer)
)

// interface for parsers
type OutParser interface {
	Message(i ...interface{}) string
}

type MOut struct {
	io          io.Writer
	namedWriter string
	parser      OutParser
}

// set stdout as writer
func (m *MOut) Std() *MOut {
	m.io = os.Stdout
	return m
}

// set stderr as writer
func (m *MOut) Err() *MOut {
	m.io = os.Stderr
	return m
}

// sets the parser they is responsible for formatting
func (m *MOut) SetParser(parser OutParser) *MOut {
	m.parser = parser
	return m
}

// sets an named writer if exists.
// or ignores if not.
// writer have to be set with SetNamedWriter before.
func (m *MOut) Named(key string) *MOut {
	if io, exists := writers[key]; exists {
		m.io = io
		m.namedWriter = key
	}
	return m
}

// register or overidde a io.Writer by key-name.
// also it will be set as the current writer
func (m *MOut) SetNamedWriter(key string, io io.Writer) *MOut {
	if key == "" {
		key = "default"
	}
	writers[key] = io
	m.namedWriter = key
	m.io = io
	return m
}

// ToString get the formated string, depending on the used formatter
func (m *MOut) ToString(i ...interface{}) string {
	if m.parser == nil {
		var plain PlainParse
		m.parser = plain
	}
	return m.parser.Message(i...)

}

// Out print the formatted content by using fmt.Fprint
// have the same return values.
func (m *MOut) Out(i ...interface{}) (n int, err error) {
	out := m.ToString(i...)
	if m.io == nil {
		m.Std()
	}
	return fmt.Fprint(m.io, out)
}

// OutLn print the formatted content by using fmt.Fprintln
// it have the same return values like fmt.Fprintln
func (m *MOut) OutLn(i ...interface{}) (n int, err error) {
	out := m.ToString(i...)
	if m.io == nil {
		m.Std()
	}
	return fmt.Fprintln(m.io, out)
}
