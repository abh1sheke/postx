package colors

import (
	"fmt"
)

type Color uint8

const (
	Reset Color = 0
	Red   Color = iota + 30
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) toBytes() []byte {
	return []byte(fmt.Sprintf("\033[%vm", c))
}
