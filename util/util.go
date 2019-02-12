package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func EntryExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
func FileExists(file string) bool {
	return EntryExists(file)
}
func DirExists(dir string) bool {
	return EntryExists(dir)
}
func GetProjectDir() string {
	return filepath.Dir(GetExecutableDir())
}

func GetRelativeDir(dir string) string {
	fp := fmt.Sprintf("%s/%s", GetExecutableDir(), dir)
	return filepath.Clean(fp)
}

func GetExecutableFilepath() string {
	fp, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func GetExecutableDir() string {
	return filepath.Dir(GetExecutableFilepath())
}

func Error(msg interface{}, args ...interface{}) {
	var err error
	_msg, ok := msg.(string)
	if !ok {
		err, ok = msg.(error)
		if ok {
			_msg = err.Error()
		}
	}
	if !ok {
		panic(err)
	}
	if len(args) > 0 {
		_msg = fmt.Sprintf(_msg, args...)
	}
	parts := strings.Split(_msg, " ")
	_msg = strings.Title(parts[0])
	if len(parts) > 1 {
		_msg += " " + strings.Join(parts[1:], " ")
	}
	fmt.Println(_msg + ".")
	os.Exit(1)
}
