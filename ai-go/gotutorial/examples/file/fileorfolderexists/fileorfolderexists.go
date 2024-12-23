/*
@File   : fileorfolderexists.go
@Author : pan
@Time   : 2023-06-21 11:09:28
*/
package fileorfolderexists

import "os"

func FileOrFolderExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
