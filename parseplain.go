package manout

import "fmt"

type PlainParse struct {
}

func (c PlainParse) Message(i ...interface{}) string {
	return fmt.Sprint(i...)
}

func (c PlainParse) Enable(mo *MOut) bool {
	return true
}
