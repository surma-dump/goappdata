package goappdata

import "os"

const (
	format = "%s/Library/Application Support/%s"
)

var (
	root = os.Getenv("HOME")
)
