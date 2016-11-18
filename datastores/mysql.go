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
	stmt, err := db.Prepare("insert into starwarschar(id, name, home_planet, type_id) values(?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec("1000", "Luke Skywalker", "Tatooine", 1000)
	if err != nil {
		return err
	}

	return nil
}
