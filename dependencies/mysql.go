package dependencies

import (
	"database/sql"

	"github.com/addione/web-service/helper"
	_ "github.com/go-sql-driver/mysql"
)

type commonMysql struct {
	uri string
}

func newMysql() *commonMysql {

	uri, _ := helper.GetEnvVariable(helper.MYSQL_URI)

	return &commonMysql{
		uri: uri,
	}
}

func (m *commonMysql) getMysqlClient(dbName string) *sql.DB {
	db, err := sql.Open("mysql", m.uri+"/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
