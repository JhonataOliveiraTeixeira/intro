package db

import "intro/domain"

type DB map[int]domain.User


func Init()*map[int]domain.User{
	db := make(map[int]domain.User)
	return &db
}
