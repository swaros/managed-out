package manout_test

import (
	"fmt"
	"testing"

	"github.com/swaros/manout"
)

func TestOuthandler(t *testing.T) {

	manout.ColorEnabled = true
	outstr := manout.Message("hello world", " you sucks")
	outstr2 := manout.MessageCln("hello world", " you sucks")
	fmt.Println(outstr)
	if len(outstr) != 21 {
		t.Error("wrong result size", len(outstr))
	}

	if outstr != "hello world you sucks" {
		t.Error("unexpected result: ", outstr)
	}
	// string includes reset codes only if codes are in
	if len(outstr2) != 21 {
		t.Error("wrong result size", len(outstr2))
	}
}

func TestColorReplaced(t *testing.T) {
	manout.ColorEnabled = true
	outstr := manout.Message("<f:yellow>hello world in yellow</> this is back to regular")
	fmt.Println(outstr)
	if len(outstr) != 54 {
		t.Error("wrong result size", len(outstr))
	}

	outstr = manout.MessageCln("<f:yellow>auto reset this")
	fmt.Println(outstr, "resetted?")
	if len(outstr) != 24 {
		t.Error("wrong result size", len(outstr))
	}
}

func TestColorReplacedDisabled(t *testing.T) {
	manout.ColorEnabled = false
	outstr := manout.Message("<f:yellow>hello world in yellow</>")
	fmt.Println(outstr)
	if len(outstr) != 21 {
		t.Error("wrong result size", len(outstr))
	}
}
