package dao

import (
	"fmt"
	"rpcServer/model"
	"rpcServer/utils"
)

func GetArticleById(id int) *model.Article {
	db := utils.GetDB()
	var result model.Article
	db.Raw("SELECT * FROM article WHERE id = ?", id).Scan(&result)
	fmt.Println(result)
	return &result
}
