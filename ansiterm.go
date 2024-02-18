package colo

import (
	"fmt"
)

// TODO stringify all attributes
type TxFormat []TxAttr // multiple eg. 1;4

func NewComposeTxFormat(a ...TxAttr) TxFormat {
	for i, v := range a {
		fmt.Println(i, v)
	}
	// could just return a
	// res := make([]TxAttr, 0, len(a))
	// res = append(res, a...)
	return TxFormat(a)
}

func NewTxFormat() *TxFormat {
	return &TxFormat{BgBlack, White}
}

func (f TxFormat) Fmt(val any) string {
	var attr string
	for i, a := range f {
		attr += fmt.Sprintf("%d", a)
		if i != len(f)-1 {
			attr += ";"
		}
	}
	fmt.Println("attr: ", attr)
	return fmt.Sprintf("\033[%sm%v\033[0m", attr, val)
}

func (c TxAttr) Fmt(val any) string {
	return fmt.Sprintf("\033[%dm%v\033[0m", c, val)
}

// Reset value?
func Clear() string {
	return "\033[0m"
}

// // TODO generic (?)
// func fmtColor(s string, color int) string {
// 	attr := 1
// 	return fmt.Sprintf("\033[%d;%dm%s\033[0m", color, attr, s)
// }
