package colors

type Colors struct {
	Reset   []byte
	Red     []byte
	Green   []byte
	Yellow  []byte
	Blue    []byte
	Magenta []byte
	Cyan    []byte
	White   []byte
}


func InitColours() *Colors {
	return &Colors{
		Reset:   []byte("\033[0m"),
		Red:     []byte("\033[31m"),
		Green:   []byte("\033[32m"),
		Yellow:  []byte("\033[33m"),
		Blue:    []byte("\033[34m"),
		Magenta: []byte("\033[35m"),
		Cyan:    []byte("\033[36m"),
		White:   []byte("\033[37m"),
	}
}
