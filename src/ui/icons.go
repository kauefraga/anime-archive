package ui

import "github.com/jedib0t/go-pretty/v6/text"

var (
	Plus          = text.FgGreen.Sprint("[") + "+" + text.FgGreen.Sprint("]")
	Minus         = text.FgRed.Sprint("[") + "-" + text.FgRed.Sprint("]")
	Interrogative = text.FgYellow.Sprint("[") + "?" + text.FgYellow.Sprint("]")
)
