package manout

import (
	"fmt"
	"strings"

	"github.com/swaros/outinject"
)

type Colored struct {
	enabled bool
}

var ColParser Colored

func NewColoredOut() (*outinject.MOut, *Colored) {
	mo := outinject.NewStdout()
	mo.SetParser(&ColParser)
	return mo, &ColParser
}

func (c *Colored) Enable(mo *outinject.MOut) bool {
	c.enabled = mo.IsTerminal
	return c.enabled
}

func (c *Colored) Parse(i ...interface{}) string {
	stringResult := fmt.Sprint(i...)
	needToDo := strings.Contains(stringResult, "<")
	if needToDo {
		replaceed := c.buildColored(stringResult)
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

func (c *Colored) buildColored(origin string) string {

	for key, code := range tagMap {
		colCde := "\033[" + code + "m"
		if !c.enabled {
			colCde = ""
		}
		if strings.Contains(origin, key) {
			origin = strings.ReplaceAll(origin, key, colCde)
		}

		if strings.Contains(origin, CleanTag) {
			if !c.enabled {
				origin = strings.ReplaceAll(origin, CleanTag, "")
			} else {
				origin = strings.ReplaceAll(origin, CleanTag, resetCode)
			}
		}

	}

	return origin
}
