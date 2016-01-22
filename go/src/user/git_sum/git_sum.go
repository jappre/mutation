package main

import (
	// "errors"
	"container/list"
	"crypto/md5"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"user/git_sum/go_shell"

	"gopkg.in/fatih/set.v0"
)

var dirList = list.New()
var s = set.New()

func main() {
	// GetFilesInPath("/Users/tiger/start.sh")
	filepath.Walk("/Users/tiger/Company/MrWind/server/new/Dispatcher", walkFunc)

	for it := dirList.Front(); it != nil; it = it.Next() {
		fmt.Printf("PATH IS %s", fmt.Sprint(it.Value))
		fmt.Printf("\n")
		fmt.Println(GetCommitGuy(fmt.Sprint(it.Value)))
	}

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
		// return errors.New("skip this directory")
		if !s.Has(md5.Sum([]byte(paths))) {
			s.Add(md5.Sum([]byte(paths)))
			dirList.PushFront(strings.TrimSuffix(paths, ".git"))
		}
		return nil
	}

	// fmt.Printf("path is %s", path.Base(paths))
	// fmt.Printf("\n")
	// fmt.Printf("info is %d", info.Mode())
	// fmt.Printf("\n")

	return nil
}

//GetCommitGuy 用来提取commit的人员名单列表
func GetCommitGuy(path string) []string {
	guyList, _ := goshell.GetCommandMessage(gitLogCommand(path))
	fmt.Println(guyList)
	return nil
}

//GitLogCommand 用来返回获取commit log的具体命令
func gitLogCommand(path string) string {
	// git log | grep @123feng.com | uniq -d -i | sort
	return "cd " + path + "; git log | grep Author: | uniq -d -i | sort"
}
