package gorm

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBService is a database engine object.
type DBService struct {
	Default *gorm.DB            // the default database engine
	List    map[string]*gorm.DB // database engine list
}

var dbService = func() (serv *DBService) {
	serv = &DBService{
		List: map[string]*gorm.DB{},
	}

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[gorm] " + strings.Join(errs, "\n"))
		}
		if serv.Default == nil {
			log.Panicf("[gorm] the `default` database engine must be configured and enabled")
		}
	}()

	err := loadDBConfig()
	if err != nil {
		log.Panicf("[gorm]", err.Error())
		return
	}

	for _, conf := range dbConfigs {
		if !conf.Enable {
			continue
		}
		engine, err := gorm.Open(conf.Driver, conf.Connstring)
		if err != nil {
			log.Panicf("[gorm]", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		engine.LogMode(conf.ShowSql)

		engine.DB().SetMaxOpenConns(conf.MaxOpenConns)
		engine.DB().SetMaxIdleConns(conf.MaxIdleConns)
		engine.SingularTable(true)
		serv.List[conf.Name] = engine
		if DEFAULTDB_NAME == conf.Name {
			serv.Default = engine
		}
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return conf.TableFix + defaultTableName
		}
	}
	return
}()
