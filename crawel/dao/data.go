package dao

import (
	"crawel/global"
	"crawel/model"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 爬去结果存入数据库
func InsertData(info *model.ArticleInfo) {
	db, err := sql.Open("mysql", "root:root@/android_blogs?charset=utf8")
	global.CheckErr(err)
	defer func() {
		fmt.Println("close")
		db.Close()
	}()
	sql := "INSERT INTO article_info(name,url) VALUES(?,?)"
	fmt.Println(strings.TrimSpace(info.Name))
	result, err := db.Exec(sql, strings.TrimSpace(info.Name), strings.TrimSpace(info.Url))
	global.CheckErr(err)
	id, _ := result.LastInsertId()
	fmt.Println(id)
}
