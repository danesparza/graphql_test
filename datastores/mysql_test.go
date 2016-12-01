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

//	MySQL get should return successfully even if the item doesn't exist
func TestMysql_Get_ItemDoesntExist_Successful(t *testing.T) {

	//	Arrange
	db := getDBConnection()
	query := 1000

	//	Act
	response, err := datastores.GetHuman(query)

	//	Assert
	if err != nil {
		t.Errorf("GetHuman failed: Should have returned a dataset without error: %s", err)
	}

	if response.ID != "1000" {
		t.Errorf("GetHuman failed: Shouldn't have returned the value %s", response.ID)
	}
}
