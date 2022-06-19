package manout

import (
	"fmt"
	"strings"
)

type Colored struct {
	enabled bool
}

func (c *Colored) Message(i ...interface{}) string {
	stringResult := fmt.Sprint(i...)
	needToDo := strings.Contains(stringResult, "<")
	if needToDo {
		replaceed := buildColored(stringResult)
		return replaceed
	}
	return stringResult
}

func (c *Colored) DisableColor() {
	c.enabled = false
}

func (c *Colored) EnableColor() {
	c.enabled = true
}
