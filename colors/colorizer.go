package colors

import (
	"io"
	"regexp"
	"strings"
)

type Colorizer struct {
	Pattern *regexp.Regexp
	Color   Color
}

func insertColor(src *[]byte, color *[]byte, index int) {
	colorLen := len(*color)
	newLen := len(*src) + colorLen
	newSrc := make([]byte, newLen, newLen)
	for i := range newSrc {
		if i < index {
			newSrc[i] = (*src)[i]
		} else if i >= index && i < index+colorLen {
			newSrc[i] = (*color)[i-index]
		} else {
			newSrc[i] = (*src)[i-colorLen]
		}
	}
	*src = newSrc
}

func colorize(src *[]byte, colorizers *[]*Colorizer) {
	startLen := len(*src)
	reset := Reset.toBytes()
	for _, colorizer := range *colorizers {
		pattern, color := colorizer.Pattern, colorizer.Color
		matchIdx := pattern.FindAllIndex(*src, -1)
		for _, indices := range matchIdx {
			currentLen := len(*src)
			offset := currentLen - startLen
			start, end := indices[0], indices[1]
			colorb := color.toBytes()
			insertColor(src, &colorb, start+offset)
			offset += len(colorb)
			insertColor(src, &reset, end+offset)
		}
		startLen = len(*src)
	}
}

func ColorizeOutput(src *string) *string {
	protocol := Colorizer{Pattern: regexp.MustCompile(`HTTP/[\d\.]+`), Color: Blue}
	headerKeys := Colorizer{Pattern: regexp.MustCompile(`[\w\-]+:\ `), Color: Yellow}
	clientErr := Colorizer{Pattern: regexp.MustCompile(`4[\d]{2}\ [\w\ ]+`), Color: Red}
	serverErr := Colorizer{Pattern: regexp.MustCompile(`5[\d]{2}\ [\w\ ]+`), Color: Red}
	ok := Colorizer{Pattern: regexp.MustCompile(`2[\d]{2}\ [\w\ ]+`), Color: Cyan}

	colorizers := []*Colorizer{&protocol, &headerKeys, &clientErr, &serverErr, &ok}
	bytes, _ := io.ReadAll(strings.NewReader(*src))

	colorize(&bytes, &colorizers)
	toStr := string(bytes)

	return &toStr
}
