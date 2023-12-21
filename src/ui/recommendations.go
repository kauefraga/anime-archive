/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/list"
)

func ShowRecommendations() {
	l := list.NewWriter()
	l.SetOutputMirror(os.Stdout)

	l.AppendItem("Rakudai Kishi")
	l.AppendItem("Kimetsu no Yaiba")
	l.AppendItem("Yamada-kun and The Seven Witches")
	l.AppendItem("Ragna Crimson")
	l.AppendItem("Seishun no Buta Yarou")
	l.AppendItem("Akashic Records")
	l.AppendItem("Bocchi The Rock")

	l.Render()
}
