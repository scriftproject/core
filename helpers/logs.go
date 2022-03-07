package helpers

import (
	"log"
	"os"
)

var ErrLog = log.New(os.Stderr, ColorPrint(ColorRed, "ERROR"), 0)
var WarnLog = log.New(os.Stderr, ColorPrint(ColorYellow, "WARNING"), 0)
var InfoLog = log.New(os.Stderr, ColorPrint(ColorBlue, "INFO"), 0)
