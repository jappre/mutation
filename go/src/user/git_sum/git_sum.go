package main

import (
	// "errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	// "strings"
)

func main() {
	// GetFilesInPath("/Users/tiger/start.sh")
	filepath.Walk("/Users/tiger/private/git-flow-cheatsheet", walkFunc)
}

//GetFilesInPath 用来获取某个路径下的所有文件或文件夹
func GetFilesInPath(path string) (files []string) {
	file, error := os.Open(path)

	if error != nil {
		fmt.Printf("An error happened")
		return
	}
	fmt.Printf(file.Name())
	defer file.Close()
	return
}

// GitRepository 用来判断该目录是否是一个git仓库
func GitRepository(path string) (is bool) {
	return false
}

// GetCommitSummary 用来获取某个目录下某个人的commit概要
func GetCommitSummary(path, user string) (insertion, deletion int) {
	return insertion, deletion
}

func walkFunc(paths string, info os.FileInfo, err error) error {
	fmt.Printf("path is %s", path.Base(paths))
	fmt.Printf("\n")
	fmt.Printf("info is %d", info.Mode())

	// if strings.Contains(paths, ".git") {
	// 	return errors.New("skip this directory")
	// }

	return nil
}
