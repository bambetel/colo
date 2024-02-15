package colo

import (
	"fmt"
	"strings"
)

type TxColor int

const (
	Black TxColor = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BgBlack TxColor = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BriBlack TxColor = iota + 90
	BriRed
	BriGreen
	BriYellow
	BriBlue
	BriMagenta
	BriCyan
	BriWhite
)

const (
	BgBriBlack TxColor = iota + 100
	BgBriRed
	BgBriGreen
	BgBriYellow
	BgBriBlue
	BgBriMagenta
	BgBriCyan
	BgBriWhite
)

type TxAttr int

const (
	Reset TxAttr = iota
	Bold
	Faint
	Italic
	Underline
	Blink
	_
	Reverse // Invert
	Conceal
	Strike
	FontPrimary
	Font1
	Font2
	Font3
	Font4
	Font5
	Font6
	Font7
	Font8
	Font9
	Font10
)

// TODO stringify all attributes
type TxFormat struct {
	fg   TxColor
	bg   TxColor
	attr TxAttr // multiple eg. 1;4
}

func NewComposeTxFormat(a ...TxAttr) {
	if true {
		m := strings.Join([]string{"33", "1"}, ";")
		(fmt.Printf("\033[%s\033[0m", m))
	} else {
		fmt.Println("Empty format.")
	}
}

func NewTxFormat() *TxFormat {
	return &TxFormat{fg: Black, bg: White, attr: 1}
}

func (f *TxFormat) Fmt(val any) string {
	return fmt.Sprintf("\033[%d;%dm%v\033[0m", f.fg, f.attr, val)
}

func (c TxColor) Fmt(val any) string {
	return fmt.Sprintf("\033[%d;%dm%v\033[0m", c, 1, val)
}

// // TODO generic (?)
// func fmtColor(s string, color int) string {
// 	attr := 1
// 	return fmt.Sprintf("\033[%d;%dm%s\033[0m", color, attr, s)
// }
