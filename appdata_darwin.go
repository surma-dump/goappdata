package goappdata

import "os"

const (
	format = "%s/Library/ApplicationSupport/%s"
)

var (
	root = os.Getenv("HOME")
)
