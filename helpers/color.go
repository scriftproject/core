package helpers

import "fmt"

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorPurple Color = "\u001b[35m"
	ColorCyan   Color = "\u001b[36m"
	ColorWhite  Color = "\u001b[37m"
	ColorReset  Color = "\u001b[0m"
)

func ColorPrint(color Color, message string) string {
	return fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset))
}
