package controllers

import (
	"rpcServer/model"
	"rpcServer/services"
)

func GetArticleById(id int) *model.Article {
	res := services.GetArticleById(id)
	return res
}
