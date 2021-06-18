package services

import (
	"rpcServer/dao"
	"rpcServer/model"
)

func GetArticleById(id int) *model.Article {
	res := dao.GetArticleById(id)
	return res
}
