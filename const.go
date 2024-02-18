package colo

// type TxColor int // for what?

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

const (
	Black TxAttr = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BgBlack TxAttr = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BriBlack TxAttr = iota + 90
	BriRed
	BriGreen
	BriYellow
	BriBlue
	BriMagenta
	BriCyan
	BriWhite
)

const (
	BgBriBlack TxAttr = iota + 100
	BgBriRed
	BgBriGreen
	BgBriYellow
	BgBriBlue
	BgBriMagenta
	BgBriCyan
	BgBriWhite
)
