package log

import (
	"fmt"
	"runtime"
	"strings"
)

var rootPath string

func SetRootPath(path string) {
	rootPath = path
}

func withCaller(keysAndValues []interface{}) []interface{} {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return keysAndValues
	}

	var filename string
	if rootPath != "" {
		filename = strings.TrimPrefix(file, rootPath+"/")
	} else {
		filenameParts := strings.Split(file, "/")
		filename = filenameParts[len(filenameParts)-1]
	}

	return append([]interface{}{"caller", fmt.Sprintf("%s:%d", filename, line)}, keysAndValues...)
}
