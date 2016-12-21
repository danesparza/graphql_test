package datastores_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danesparza/graphql_test/datastores"
)

//	Gets the database connection information from the environment
func getDBConnection() datastores.MySQLDB {

	//	Set this information from environment variables?
	return datastores.MySQLDB{
		Address:  os.Getenv("graphql_test_mysql_server"), /* Ex: test-server:3306 If this is blank, it assumes a local database on port 3306 */
		Database: os.Getenv("graphql_test_mysql_database"),
		User:     os.Getenv("graphql_test_mysql_user"),
		Password: os.Getenv("graphql_test_mysql_password")}
}

// TestMysql_GetHuman_ItemExists_Successful should fetch the human with the correct id
func TestMysql_GetHuman_ValidId_Successful(t *testing.T) {

	//	Arrange
	db := getDBConnection()
	query := 1000

	//	Act
	response, err := db.GetHuman(query)

	//	Assert
	if err != nil {
		t.Errorf("GetHuman failed: Should have returned a dataset without error: %s", err)
	}

	if response.ID != "1000" {
		t.Errorf("GetHuman failed: Shouldn't have returned the value '%s'", response.ID)
	}

}

// TestMysql_GetFriends_ValidId_CorrectNumberReturned should fetch the correct friends
func TestMysql_GetFriends_ValidId_CorrectFriendsReturned(t *testing.T) {

	//	Arrange
	db := getDBConnection()
	query := 1000

	//	Act
	friends, err := db.GetFriends(query)

	//	Assert
	if err != nil {
		t.Errorf("GetHuman failed: Should have returned a dataset without error: %s", err)
	}

	if len(friends) != 4 {
		t.Errorf("GetFriends failed: Expecting 4 friends but got '%v'", len(friends))
	}

	if friends[0].Name != "C-3PO" {
		t.Errorf("GetFriends failed: first friend should be C-3PO but got '%v'", friends[0].Name)
	}

}
