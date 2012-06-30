package goappdata

import "os"

const (
	format = "%s/Local/%s"
)

var (
	root = os.Getenv("APPDATA")
)
