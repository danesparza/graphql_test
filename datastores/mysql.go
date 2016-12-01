package datastores

import (
	"fmt"

	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

// MySqlDB is the database information
type MySQLDB struct {
	Protocol string
	Address  string
	Database string
	User     string
	Password string
}

// GetHuman gets the human with the given id from the database
func (store MySQLDB) GetHuman(humanID int) StarWarsChar {
	//	Our return item:
	retval := StarWarsChar{}

	//	Open the database:
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", store.User, store.Password, store.Protocol, store.Address, store.Database))
	defer db.Close()
	if err != nil {
		return retval
	}

	//	Prepare our query
	stmt, err := db.Prepare("select id, name, home_planet, primary_function, type_id from starwarschar where id=? order by name")
	defer stmt.Close()
	if err != nil {
		return retval
	}

	//	Get the human with the given id
	rows, err := stmt.Query(humanID)
	defer rows.Close()
	if err != nil {
		return retval
	}

	for rows.Next() {
		var id string
		var name string
		var homePlanet string
		var primaryFunction string
		var typeID int

		//	Scan the row into our variables
		err = rows.Scan(&id, &name, &homePlanet, &primaryFunction, &typeID)

		if err != nil {
			return retval
		}

		//	Set our return value
		retval = StarWarsChar{
			ID:              id,
			Name:            name,
			HomePlanet:      homePlanet,
			PrimaryFunction: primaryFunction}

		break
	}

	//	Return it
	return retval
}
