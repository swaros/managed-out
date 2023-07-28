// Copyright (c) 2022 Thomas Ziegler <thomas.zglr@googlemail.com>. All rights reserved.
//
// Licensed under the MIT License
//
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package manout

import (
	"fmt"
	"strings"
)

// ColorEnabled enables or disables the color usage
var ColorEnabled = true

const (
	resetCode = "\033[0m"
	// CleanTag is the tag to reset to default
	CleanTag = "</>"
	// ForeBlack black foreground color
	ForeBlack = "<f:black>"
	// ForeRed red foreground color
	ForeRed = "<f:red>"
	// ForeGreen red foreground color
	ForeGreen = "<f:green>"
	// ForeYellow red foreground color
	ForeYellow = "<f:yellow>"
	// ForeBlue red foreground color
	ForeBlue = "<f:blue>"
	// ForeMagenta red foreground color
	ForeMagenta = "<f:magenta>"
	// ForeCyan red foreground color
	ForeCyan = "<f:cyan>"
	// ForeLightGrey red foreground color
	ForeLightGrey = "<f:light-grey>"
	// ForeDarkGrey red foreground color
	ForeDarkGrey = "<f:dark-grey>"
	// ForeLightRed red foreground color
	ForeLightRed = "<f:light-red>"
	// ForeLightGreen red foreground color
	ForeLightGreen = "<f:light-green>"
	// ForeLightYellow red foreground color
	ForeLightYellow = "<f:light-yellow>"
	// ForeLightBlue red foreground color
	ForeLightBlue = "<f:light-blue>"
	// ForeLightCyan red foreground color
	ForeLightCyan = "<f:light-cyan>"
	// ForeLightMagenta red foreground color
	ForeLightMagenta = "<f:light-magenta>"
	// ForeWhite red foreground color
	ForeWhite = "<f:white>"

	// BoldTag writes bolder text
	BoldTag = "<b>"
	// Dim dim
	Dim = "<dim>"
	// Underlined tag
	Underlined = "<u>"
	// Reverse tag
	Reverse = "<r>"
	// Hidden tag
	Hidden = "<hide>"
	// ResetBold tag
	ResetBold = "</b>"
	// ResetDim tag
	ResetDim = "</dim>"
	// ResetUnderline tag
	ResetUnderline = "</u>"
	//ResetReverse tag
	ResetReverse = "</r>"
	//ResetHidden tag
	ResetHidden = "</hide>"

	// BackBlack black Background color
	BackBlack = "<b:black>"
	// BackRed red Background color
	BackRed = "<b:red>"
	// BackGreen red Background color
	BackGreen = "<b:green>"
	// BackYellow red Background color
	BackYellow = "<b:yellow>"
	// BackBlue red Background color
	BackBlue = "<b:blue>"
	// BackMagenta red Background color
	BackMagenta = "<b:magenta>"
	// BackCyan red Background color
	BackCyan = "<b:cyan>"
	// BackLightGrey red Background color
	BackLightGrey = "<b:light-grey>"
	// BackDarkGrey red Background color
	BackDarkGrey = "<b:dark-grey>"
	// BackLightRed red Background color
	BackLightRed = "<b:light-red>"
	// BackLightGreen red Background color
	BackLightGreen = "<b:light-green>"
	// BackLightYellow red Background color
	BackLightYellow = "<b:light-yellow>"
	// BackLightBlue red Background color
	BackLightBlue = "<b:light-blue>"
	// BackLightCyan red Background color
	BackLightCyan = "<b:light-cyan>"
	// BackLightMagenta red Background color
	BackLightMagenta = "<b:light-magenta>"
	// BackWhite red Background color
	BackWhite = "<b:white>"
)

var tagMap = map[string]string{
	ForeBlack:        "30",
	ForeRed:          "31",
	ForeGreen:        "32",
	ForeYellow:       "33",
	ForeBlue:         "34",
	ForeMagenta:      "35",
	ForeCyan:         "36",
	ForeLightGrey:    "37",
	ForeDarkGrey:     "90",
	ForeLightRed:     "91",
	ForeLightGreen:   "92",
	ForeLightYellow:  "93",
	ForeLightBlue:    "94",
	ForeLightMagenta: "95",
	ForeLightCyan:    "96",
	ForeWhite:        "97",
	BoldTag:          "1",
	Dim:              "2",
	Underlined:       "4",
	Reverse:          "7",
	Hidden:           "8",
	ResetBold:        "21",
	ResetDim:         "22",
	ResetUnderline:   "24",
	ResetReverse:     "27",
	ResetHidden:      "28",
	BackBlack:        "40",
	BackRed:          "41",
	BackGreen:        "42",
	BackYellow:       "43",
	BackBlue:         "44",
	BackMagenta:      "45",
	BackCyan:         "46",
	BackLightGrey:    "47",
	BackDarkGrey:     "100",
	BackLightRed:     "101",
	BackLightGreen:   "102",
	BackLightYellow:  "103",
	BackLightBlue:    "104",
	BackLightMagenta: "105",
	BackLightCyan:    "106",
	BackWhite:        "107",
}

// Error Print Error Message to stdout
func Error(label string, a ...interface{}) {
	labelMessage := Message(ForeRed, Dim, BackLightRed, " ", label, " ", CleanTag, " ! ")

	stringResult := MessageCln(a...)
	fmt.Println(labelMessage, stringResult)
}

// MessageCln converts arguemnst to a fomated string and adding cleanup and newline code
func MessageCln(a ...interface{}) string {
	result := Message(a...)
	needToDo := strings.Contains(result, "\033[")
	if needToDo {
		result = Message(result, CleanTag)
	}
	return result
}

// Message get the message an handle the formating of them
func Message(a ...interface{}) string {
	stringResult := fmt.Sprint(a...)
	needToDo := strings.Contains(stringResult, "<")
	if needToDo {
		replaceed := buildColored(stringResult)
		return replaceed
	}
	return stringResult
}

func buildColored(origin string) string {

	for key, code := range tagMap {
		colCde := "\033[" + code + "m"
		if !ColorEnabled {
			colCde = ""
		}
		if strings.Contains(origin, key) {
			origin = strings.ReplaceAll(origin, key, colCde)
		}

		if strings.Contains(origin, CleanTag) {
			if !ColorEnabled {
				origin = strings.ReplaceAll(origin, CleanTag, "")
			} else {
				origin = strings.ReplaceAll(origin, CleanTag, resetCode)
			}
		}

	}

	return origin
}
