package dlldb

import (
	"errors"
)

type IServices interface {
	One(model interface{}, query interface{}, args ...interface{}) error
	List(model interface{}, order string, limit, offset int, query interface{}, args ...interface{}) (paging Paging, err error)
	Count(model interface{}, query interface{}, args ...interface{}) int
	Save(model interface{}) error
	Delete(model interface{}, query interface{}, args ...interface{}) error
	Update(model interface{}, attrs interface{}, query interface{}, args ...interface{}) error
	ExecSql(model interface{}, sql string, args ...interface{}) error
}
type Services struct {
	Repo IRepository
}

func NewServices() *Services {
	return &Services{
		Repo: NewGormRepository(),
	}
}
func (srv *Services) One(model interface{}, query interface{}, args ...interface{}) error {
	return srv.Repo.One(model, query, args...)
}
func (srv *Services) List(model interface{}, order string, limit, offset int, query interface{}, args ...interface{}) (paging Paging, err error) {
	page := Paging{}
	page.Offset = offset
	page.Limit = limit
	page.TotalCount = srv.Repo.Count(model, query, args...)
	if page.TotalCount == 0 {
		return page, errors.New("no data")
	}
	page.TotalPage = page.TotalPages()
	if err := srv.Repo.List(model, order, limit, offset, query, args...); err != nil {
		return page, err
	}
	return page, nil
}
func (srv *Services) Count(model interface{}, query interface{}, args ...interface{}) int {
	return srv.Repo.Count(model, query, args...)
}
func (srv *Services) Save(model interface{}) error {
	return srv.Repo.Save(model)
}
func (srv *Services) Delete(model interface{}, query interface{}, args ...interface{}) error {
	return srv.Repo.Delete(model, query, args...)
}
func (srv *Services) Update(model interface{}, attrs interface{}, query interface{}, args ...interface{}) error {
	return srv.Repo.Update(model, attrs, query, args...)
}
func (srv *Services) ExecSql(model interface{}, sql string, args ...interface{}) error {
	return srv.Repo.ExecSql(model, sql, args...)
}
