// Abstraction over the storage path for persistent app data
/*
Darwin:
	$HOME/Library/ApplicationSupport/<appname>

Linux/*BSD:
	$HOME/.config/<appname>

Windows:
	%APPDATA%/Local/<appname>
*/
package goappdata

import (
	"fmt"
	"os"
)

// Returns the path to folder in which the an app with the name `appname` is
// supposed to store its data.
func Path(appname string) string {
	return fmt.Sprintf(format, root, appname)
}

// Same as Path, but implicitly creates the folder, if it does not exist yet.
func CreatePath(appname string) (string, error) {
	var dir = Path(appname)
	return dir, os.MkdirAll(Path(appname), os.FileMode(0750))
}
