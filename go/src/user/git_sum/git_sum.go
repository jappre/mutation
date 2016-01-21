package main

import (
	// "errors"
	"container/list"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"user/git_sum/go_shell"
)

var dirList = list.New()

func main() {
	// GetFilesInPath("/Users/tiger/start.sh")
	filepath.Walk("/Users/tiger/private/", walkFunc)

	for it := dirList.Front(); it != nil; it = it.Next() {
		fmt.Printf("PATH IS %s", it.Value)
		fmt.Printf("\n")
	}
	goshell.GetCommandMessage("cd ~;ls")
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
	if strings.HasSuffix(path.Base(paths), ".git") && info.Mode() > 1420 {
		fmt.Printf("path is %s", paths)
		fmt.Printf("\n")
		// return errors.New("skip this directory")
		dirList.PushFront(paths)
		return nil
	}

	// fmt.Printf("path is %s", path.Base(paths))
	// fmt.Printf("\n")
	// fmt.Printf("info is %d", info.Mode())
	// fmt.Printf("\n")

	return nil
}
