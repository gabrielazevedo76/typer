package color

// ANSI color codes
const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	GRAY   = "\033[37m"
)

// Colorize applies color to the given text
func Colorize(text string, colorCode string) string {
	return colorCode + text + RESET
}
