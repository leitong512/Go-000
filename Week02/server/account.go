package server

import (
	"learngo/week02/dao"
	"learngo/week02/model"
)

type Server struct {
}

var accountDao = dao.NewAccountDao()

func (s *Server) GetAccount(id int) (*model.Account, error) {
	return accountDao.GetAccounts(id)
}
func NewServer() *Server {
	return &Server{}
}
