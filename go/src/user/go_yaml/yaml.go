package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
---
:object_kind: push
:event_name: push
:before: bdcb6f2f0d6e00747d281700062f72f1213ffc79
:after: b73b450dfc0e3e5e3a2539310d12d9f2941c7c72
:ref: refs/heads/new_dev
:checkout_sha: b73b450dfc0e3e5e3a2539310d12d9f2941c7c72
:message:
:user_id: 61
:user_name: "王政清"
:user_email: wangzhengqing@123feng.com
:user_avatar:
:project_id: 148
:project:
  :name: fengyun
  :description: ''
  :web_url: http://code.gomrwind.com/MrWind-Java-Server/fengyun
  :avatar_url:
  :git_ssh_url: git@code.gomrwind.com:MrWind-Java-Server/fengyun.git
  :git_http_url: http://code.gomrwind.com/MrWind-Java-Server/fengyun.git
  :namespace: MrWind-Java-Server
  :visibility_level: 0
  :path_with_namespace: MrWind-Java-Server/fengyun
  :default_branch: master
  :homepage: http://code.gomrwind.com/MrWind-Java-Server/fengyun
  :url: git@code.gomrwind.com:MrWind-Java-Server/fengyun.git
  :ssh_url: git@code.gomrwind.com:MrWind-Java-Server/fengyun.git
  :http_url: http://code.gomrwind.com/MrWind-Java-Server/fengyun.git
:commits:
- :id: 0eb159a79eac729b6f92d159782cc39aadf632c1
  :message: |
    feat:客户管理品类接口
  :timestamp: '2016-08-13T09:36:12+08:00'
  :url: http://code.gomrwind.com/MrWind-Java-Server/fengyun/commit/0eb159a79eac729b6f92d159782cc39aadf632c1
  :author:
    :name: "王政清"
    :email: wangzhengqing@123feng.com
  :added: []
  :modified:
  - WindCloud/src/main/java/com/mrwind/windCloud/controller/account/BsShopInfoController.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/entity/account/BsUserShopInfo.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/account/BsUserShopInfoMapper.java
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/account/BsUserShopInfoMapper.xml
  - WindService/src/main/java/com/mrwind/busiService/account/BsShopInfoService.java
  - WindService/src/main/java/com/mrwind/daoService/account/BsAccountInfoDaoService.java
  - WindService/src/main/java/com/mrwind/daoService/account/BsShopInfoDaoService.java
  :removed: []
- :id: b73b450dfc0e3e5e3a2539310d12d9f2941c7c72
  :message: "Merge branch 'new_dev' of code.gomrwind.com:MrWind-Java-Server/fengyun
    into new_dev\n\n# Conflicts:\n#\tWindService/src/main/java/com/mrwind/daoService/account/BsAccountInfoDaoService.java\n"
  :timestamp: '2016-08-13T09:46:12+08:00'
  :url: http://code.gomrwind.com/MrWind-Java-Server/fengyun/commit/b73b450dfc0e3e5e3a2539310d12d9f2941c7c72
  :author:
    :name: "王政清"
    :email: wangzhengqing@123feng.com
  :added:
  - WindCloud/src/main/java/com/mrwind/windCloud/scheduler/NotifyFramework.java
  - WindCloud/src/main/java/com/mrwind/windCloud/scheduler/NotifyHumanResource.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/Data/UserDataAnalyzeMapper.java
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/data/UserDateAnalyzeMapper.xml
  - WindService/src/main/java/com/mrwind/daoService/data/UserDataAnalyzeDaoService.java
  :modified:
  - WindCloud/src/main/java/com/mrwind/windCloud/controller/account/BsAccountPermissionController.java
  - WindCloud/src/main/java/com/mrwind/windCloud/controller/account/BsAskResignController.java
  - WindCloud/src/main/java/com/mrwind/windCloud/controller/daily/FyDailyController.java
  - WindCloud/src/main/java/com/mrwind/windCloud/controller/framework/FyFrameworkStaffController.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/entity/account/BsUserPermission.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/account/BsUserAccountInfoMapper.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/account/BsUserPermissionMapper.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/daily/FyDailyMasterMapper.java
  - WindPgsql/src/main/java/com/mrwind/pgsql/orm/mapper/framework/FyFrameworkStaffMapper.java
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/account/BsUserAccountInfoMapper.xml
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/account/BsUserPermissionMapper.xml
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/daily/FyDailyMasterMapper.xml
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/framework/FyFrameworkFocusMapper.xml
  - WindPgsql/src/main/resources/mybatis-mapper/pgsql/framework/FyFrameworkStaffMapper.xml
  - WindService/src/main/java/com/mrwind/busiService/account/BsAccountInfoService.java
  - WindService/src/main/java/com/mrwind/busiService/account/UserCommonService.java
  - WindService/src/main/java/com/mrwind/busiService/daily/FyDailyService.java
  - WindService/src/main/java/com/mrwind/busiService/framework/FyFrameworkBusService.java
  - WindService/src/main/java/com/mrwind/daoService/account/BsAccountInfoDaoService.java
  - WindService/src/main/java/com/mrwind/daoService/account/BsUserPermissionDaoService.java
  - WindService/src/main/java/com/mrwind/daoService/daily/FyDailyMasterDaoService.java
  - WindService/src/main/java/com/mrwind/daoService/framework/FyFrameworkStaffDaoService.java
  - WindService/src/main/test/com/mrwind/service/test/man/CommonTest.java
  - WindService/src/main/test/com/mrwind/service/test/man/TestMan.java
  :removed: []
:total_commits_count: 2
:repository:
  :name: fengyun
  :url: git@code.gomrwind.com:MrWind-Java-Server/fengyun.git
  :description: ''
  :homepage: http://code.gomrwind.com/MrWind-Java-Server/fengyun
  :git_http_url: http://code.gomrwind.com/MrWind-Java-Server/fengyun.git
  :git_ssh_url: git@code.gomrwind.com:MrWind-Java-Server/fengyun.git
  :visibility_level: 0
`

type T struct {
	ObjectKind string `yaml:":object_kind"`
	UserName   string `yaml:":user_name"`
	Project    struct {
		Path_with_namespace string `yaml:":path_with_namespace"`
	} `yaml:":project"`
	Commits []struct {
		Timestamp string `yaml:":timestamp"`
		Author    struct {
			Name string `yaml:":name"`
		} `yaml:":author"`
		Added    []string `yaml:":added"`
		Modified []string `yaml:":modified"`
		Removed  []string `yaml:":removed"`
	} `yaml:":commits"`
	TotalCommitsCount int `yaml:":total_commits_count"`
}

// var data = `
// abse: Easy!
// b:
//   c: 2
//   d: [3, 4]
// `
//
// type T struct {
// 	AbSe string
// 	BdfF struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	} `yaml:"b"`
// }

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t.TotalCommitsCount)

	// d, err := yaml.Marshal(&t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//
	// m := make(map[interface{}]interface{})
	//
	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)
	//
	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
