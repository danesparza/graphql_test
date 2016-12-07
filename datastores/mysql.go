package datastores

import (
	"fmt"
	"strconv"

	"database/sql"
)

// MySQLDB is the database information
type MySQLDB struct {
	Protocol string
	Address  string
	Database string
	User     string
	Password string
}

// GetHuman gets the human with the given id from the database
func (store MySQLDB) GetHuman(humanID int) (StarWarsChar, error) {
	//	Our return item:
	retval := StarWarsChar{}

	//	Open the database:
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", store.User, store.Password, store.Protocol, store.Address, store.Database))
	defer db.Close()
	if err != nil {
		return retval, err
	}

	//	Prepare our query
	stmt, err := db.Prepare("select id, name, home_planet, COALESCE(primary_function, '') as primary_function, type_id from starwarschar where id=? order by name")
	defer stmt.Close()
	if err != nil {
		return retval, err
	}

	//	Get the human with the given id
	stringHuman := strconv.Itoa(humanID)

	rows, err := stmt.Query(stringHuman)
	defer rows.Close()
	if err != nil {
		return retval, err
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
			return retval, err
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
	return retval, nil
}
