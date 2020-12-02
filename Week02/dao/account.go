package dao

import (
	"database/sql"
	e "errors"
	"learngo/week02/dao/mysql"
	"learngo/week02/model"

	"github.com/pkg/errors"
)

type AccountDao struct {
}

func (dao *AccountDao) GetAccounts(id int) (*model.Account, error) {
	stmt, err := mysql.DBConn().Prepare("select id, username, password " +
		"from account where id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare account error: ")
	}
	defer stmt.Close()
	_, err = stmt.Query(id)
	if err != nil {
		if e.Is(err, sql.ErrNoRows) {
			//吞掉
			return nil, nil
		}
		return nil, errors.Wrap(err, "query account error : ")
	}
	account := model.Account{}
	//处理rows TODO 这里不处理了直接返回
	return &account, nil
}

func NewAccountDao() *AccountDao {
	return &AccountDao{}
}
