package utils

import (
	"os"
	"strings"
)

func pathCombine(src string, des string) string {
	newSrcPath := strings.TrimRight(src, "\\")
	newSrcPath = strings.TrimRight(src, "/")
	newDestPath := strings.TrimLeft(des, "\\")
	newDestPath = strings.TrimLeft(des, "/")
	return newSrcPath + string(os.PathSeparator) + newDestPath
}

func PathCombine(src string, path ...string) string {
	result := src
	for _, p := range path {
		result = pathCombine(result, p)
	}
	return result
}
