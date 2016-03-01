package main

import (
	// "errors"
	"container/list"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"user/git_sum/go_shell"

	"gopkg.in/fatih/set.v0"
	"gopkg.in/mgo.v2"
)

/**
 * @api {get} /user/:id Request User information
 * @apiName GetUser
 * @apiGroup User
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 */

type Commit struct {
	Name        string
	Email       string
	Date        time.Time
	Change      string
	FileChanged int16
	Insertions  int16
	Delettions  int16
}

var dirList = list.New()
var pathSet = set.New()

func main() {
	// GetFilesInPath("/Users/tiger/start.sh")
	filepath.Walk("/Users/tiger/Company/MrWind/", walkFunc)

	for it := dirList.Front(); it != nil; it = it.Next() {
		// fmt.Printf("PATH IS %s", fmt.Sprint(it.Value))
		// fmt.Printf("\n")
		// fmt.Printf("%v", GetCommit(fmt.Sprint(it.Value)))

		commitSet := GetCommit(fmt.Sprint(it.Value))
		for coit := commitSet.Pop(); coit != nil; coit = commitSet.Pop() {
			var commit Commit
			json.Unmarshal([]byte(coit.(string)), &commit)
			// fmt.Printf("%v", commit)
			// fmt.Printf("\n")
			storeInMongo(commit, getProjectName(it.Value.(string)))
		}

	}

}

func getProjectName(path string) string {
	folderNames := strings.Split(path, "/")
	return folderNames[len(folderNames)-2]
}

func storeInMongo(commit Commit, project string) {
	fmt.Println(project)
	fmt.Println(commit)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(project).C("Commit")
	err = c.Insert(commit)
	if err != nil {
		log.Fatal(err)
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
		if !pathSet.Has(md5.Sum([]byte(paths))) {
			pathSet.Add(md5.Sum([]byte(paths)))
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

// Author: ws <王松@123feng.com>
//GetCommit 用来提取commit list
func GetCommit(path string) *set.Set {
	var authorSet = set.New()
	commitList, _ := goshell.GetCommandMessage(gitLogCommand(path))
	reg := regexp.MustCompile(`(?U){.+}`)
	// reg := regexp.MustCompile(`(?U)<.+>`)
	authorlist := reg.FindAllString(commitList, -1)
	for i := range authorlist {
		authorSet.Add(authorlist[i])
	}
	// for it := authorSet.Pop(); it != nil; it = authorSet.Pop() {
	// 	fmt.Println(it.(string))
	// }
	return authorSet
}

//GitLogCommand 用来返回获取commit log的具体命令
func gitLogCommand(path string) string {
	// git log | grep @123feng.com | uniq -d -i | sort
	// return "cd " + path + "; git log | grep Author: | uniq -d -i | sort"
	// return "cd " + path + "; git log --shortstat --since='2016-01-01' --no-merges  --pretty=format:'Author:%an,Email:%ae,date:%aI'"
	return "cd " + path + "; git log --shortstat --no-merges  --since='2016-01-15' --pretty=format:\"\\\"}{\\\"Name\\\":\\\"%an\\\",\\\"Email\\\":\\\"%ae\\\",\\\"Date\\\":\\\"%aI\\\",\\\"Change\\\":\\\"\" | sed 'H;$!d;g;s/\\n//g' >.tmp.count ; echo \"\\\"}\">>.tmp.count; echo {\\\";cat .tmp.count"

}
