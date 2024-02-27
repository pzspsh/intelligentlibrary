package healthcheck

import (
	"errors"

	fileutil "gotutorial/golibrary/utils/file"
)

type PathPermission struct {
	path       string
	isReadable bool
	isWritable bool
	Error      error
}

// CheckPathPermission checks the permissions of the given file or directory.
func CheckPathPermission(path string) (pathPermission PathPermission) {
	pathPermission.path = path
	if !fileutil.FileExists(path) {
		pathPermission.Error = errors.New("file or directory doesn't exist at " + path)
		return
	}

	pathPermission.isReadable, _ = fileutil.IsReadable(path)
	pathPermission.isWritable, _ = fileutil.IsWriteable(path)

	return
}
