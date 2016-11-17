package datastores

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MySqlDB is the database information
type MySqlDB struct {
	Protocol string
	Address  string
	Database string
	User     string
	Password string
}

// InsertTestData inserts our dataset into the database
func (store MySqlDB) InsertTestData() error {

	//	Open the database:
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", store.User, store.Password, store.Protocol, store.Address, store.Database))
	defer db.Close()
	if err != nil {
		return err
	}

	//	Insert our hero data:
	stmt, err := db.Prepare("insert into configitem(application, name, value, machine) values(?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	res, err := stmt.Exec(configItem.Application, configItem.Name, configItem.Value, configItem.Machine)
	if err != nil {
		return err
	}

	return nil
}
