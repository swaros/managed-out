package manout

import (
	"fmt"
	"io"
	"os"
)

var (
	writers map[string]io.Writer = make(map[string]io.Writer)
)

type OutParser interface {
	Message(i ...interface{}) string
}

type MOut struct {
	io          io.Writer
	namedWriter string
	parser      OutParser
}

func (m *MOut) Std() *MOut {
	m.io = os.Stdout
	return m
}

func (m *MOut) Err() *MOut {
	m.io = os.Stderr
	return m
}

func (m *MOut) SetParser(parser OutParser) *MOut {
	m.parser = parser
	return m
}

func (m *MOut) Named(key string) *MOut {
	if io, exists := writers[key]; exists {
		m.io = io
		m.namedWriter = key
	}
	return m
}

func (m *MOut) SetNamedWriter(key string, io io.Writer) *MOut {
	if key == "" {
		key = "default"
	}
	writers[key] = io
	m.namedWriter = key
	m.io = io
	return m
}

func (m *MOut) ToString(i ...interface{}) string {
	if m.parser == nil {
		var plain PlainParse
		m.parser = plain
	}
	return m.parser.Message(i...)

}

func (m *MOut) Out(i ...interface{}) {
	out := m.ToString(i...)
	if m.io == nil {
		m.Std()
	}
	fmt.Fprint(m.io, out)
}

func (m *MOut) OutLn(i ...interface{}) {
	out := m.ToString(i...)
	if m.io == nil {
		m.Std()
	}
	fmt.Fprintln(m.io, out)
}
