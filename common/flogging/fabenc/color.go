
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455951486816256>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package fabenc

import (
	"fmt"
)

type Color uint8

const ColorNone Color = 0

const (
	ColorBlack Color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func (c Color) Normal() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

func (c Color) Bold() string {
	if c == ColorNone {
		return c.Normal()
	}
	return fmt.Sprintf("\x1b[%d;1m", c)
}

func ResetColor() string { return ColorNone.Normal() }

