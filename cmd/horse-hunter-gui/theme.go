package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	_ "embed"
)

type myTheme struct{}
type myIcon struct{}

var (
	_ fyne.Theme = (*myTheme)(nil)

	//go:embed icon.png
	embedIcon []byte
)

func (m myIcon) Name() string {
	return "icon.png"
}

func (m myIcon) Content() []byte {
	return embedIcon
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	style.Monospace = true
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
