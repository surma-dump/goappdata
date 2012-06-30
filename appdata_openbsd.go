package goappdata

import "os"

const (
	format = "%s/.config/%s"
)

var (
	root = os.Getenv("HOME")
)
