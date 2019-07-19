package dlldb

type IRepository interface {
	One(model interface{}, query interface{}, args ...interface{}) error
	List(model interface{}, order string, limit, offset int, query interface{}, args ...interface{}) error
	Count(model interface{}, query interface{}, args ...interface{}) int
	Save(model interface{}) error
	Delete(model interface{}, query interface{}, args ...interface{}) error
	Update(model interface{}, attrs interface{}, query interface{}, args ...interface{}) error
	ExecSql(model interface{}, sql string, args ...interface{}) error
}
