package datastores

import (
	"os"
	"strconv"

	"github.com/graphql-go/graphql"
)

/* Our package variables */
var (
	Luke           StarWarsChar
	Vader          StarWarsChar
	Han            StarWarsChar
	Leia           StarWarsChar
	Tarkin         StarWarsChar
	Threepio       StarWarsChar
	Artoo          StarWarsChar
	HumanData      map[int]StarWarsChar
	DroidData      map[int]StarWarsChar
	StarWarsSchema graphql.Schema

	humanType *graphql.Object
	droidType *graphql.Object
	db        MySQLDB
)

// StarWarsChar struct
type StarWarsChar struct {
	ID              string
	Name            string
	Friends         []StarWarsChar
	AppearsIn       []int
	HomePlanet      string
	PrimaryFunction string
}

/* Package initialization */
func init() {

	/* Initialize all the data */
	db := MySQLDB{
		Address:  os.Getenv("graphql_test_mysql_server"), /* Ex: test-server:3306 If this is blank, it assumes a local database on port 3306 */
		Database: os.Getenv("graphql_test_mysql_database"),
		User:     os.Getenv("graphql_test_mysql_user"),
		Password: os.Getenv("graphql_test_mysql_password")}

	Luke = StarWarsChar{
		ID:         "1000",
		Name:       "Luke Skywalker",
		AppearsIn:  []int{4, 5, 6},
		HomePlanet: "Tatooine",
	}
	Vader = StarWarsChar{
		ID:         "1001",
		Name:       "Darth Vader",
		AppearsIn:  []int{4, 5, 6},
		HomePlanet: "Tatooine",
	}
	Han = StarWarsChar{
		ID:        "1002",
		Name:      "Han Solo",
		AppearsIn: []int{4, 5, 6},
	}
	Leia = StarWarsChar{
		ID:         "1003",
		Name:       "Leia Organa",
		AppearsIn:  []int{4, 5, 6},
		HomePlanet: "Alderaa",
	}
	Tarkin = StarWarsChar{
		ID:        "1004",
		Name:      "Wilhuff Tarkin",
		AppearsIn: []int{4},
	}

	Threepio = StarWarsChar{
		ID:              "2000",
		Name:            "C-3PO",
		AppearsIn:       []int{4, 5, 6},
		PrimaryFunction: "Protocol",
	}
	Artoo = StarWarsChar{
		ID:              "2001",
		Name:            "R2-D2",
		AppearsIn:       []int{4, 5, 6},
		PrimaryFunction: "Astromech",
	}

	/* Initialize the relationships */
	Luke.Friends = append(Luke.Friends, []StarWarsChar{Han, Leia, Threepio, Artoo}...)
	Vader.Friends = append(Luke.Friends, []StarWarsChar{Tarkin}...)
	Han.Friends = append(Han.Friends, []StarWarsChar{Luke, Leia, Artoo}...)
	Leia.Friends = append(Leia.Friends, []StarWarsChar{Luke, Han, Threepio, Artoo}...)
	Tarkin.Friends = append(Tarkin.Friends, []StarWarsChar{Vader}...)
	Threepio.Friends = append(Threepio.Friends, []StarWarsChar{Luke, Han, Leia, Artoo}...)
	Artoo.Friends = append(Artoo.Friends, []StarWarsChar{Luke, Han, Leia}...)

	HumanData = map[int]StarWarsChar{
		1000: Luke,
		1001: Vader,
		1002: Han,
		1003: Leia,
		1004: Tarkin,
	}

	DroidData = map[int]StarWarsChar{
		2000: Threepio,
		2001: Artoo,
	}

	/* Initialize episode enum information */
	episodeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Episode",
		Description: "One of the films in the Star Wars Trilogy",
		Values: graphql.EnumValueConfigMap{
			"NEWHOPE": &graphql.EnumValueConfig{
				Value:       4,
				Description: "Released in 1977.",
			},
			"EMPIRE": &graphql.EnumValueConfig{
				Value:       5,
				Description: "Released in 1980.",
			},
			"JEDI": &graphql.EnumValueConfig{
				Value:       6,
				Description: "Released in 1983.",
			},
		},
	})

	/* Initialize episode information */
	characterInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Character",
		Description: "A character in the Star Wars Trilogy",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the character.",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the character.",
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {

			if character, ok := p.Value.(StarWarsChar); ok {
				id, _ := strconv.Atoi(character.ID)
				human, _ := db.GetHuman(id)
				if human.ID != "" {
					return humanType
				}
			}
			return droidType
		},
	})

	characterInterface.AddFieldConfig("friends", &graphql.Field{
		Type:        graphql.NewList(characterInterface),
		Description: "The friends of the character, or an empty list if they have none.",
	})

	/* Initialize objects and resolvers */
	humanType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Human",
		Description: "A humanoid creature in the Star Wars universe.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"friends": &graphql.Field{
				Type:        graphql.NewList(characterInterface),
				Description: "The friends of the human, or an empty list if they have none.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						id, _ := strconv.Atoi(human.ID)
						friends, _ := db.GetFriends(id)
						return friends, nil
					}
					return []interface{}{}, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.AppearsIn, nil
					}
					return nil, nil
				},
			},
			"homePlanet": &graphql.Field{
				Type:        graphql.String,
				Description: "The home planet of the human, or null if unknown.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.HomePlanet, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	droidType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Droid",
		Description: "A mechanical creature in the Star Wars universe.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(StarWarsChar); ok {
						return droid.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(StarWarsChar); ok {
						return droid.Name, nil
					}
					return nil, nil
				},
			},
			"friends": &graphql.Field{
				Type:        graphql.NewList(characterInterface),
				Description: "The friends of the droid, or an empty list if they have none.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(StarWarsChar); ok {
						friends := []map[string]interface{}{}
						for _, friend := range droid.Friends {
							friends = append(friends, map[string]interface{}{
								"name": friend.Name,
								"id":   friend.ID,
							})
						}
						return droid.Friends, nil
					}
					return []interface{}{}, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(StarWarsChar); ok {
						return droid.AppearsIn, nil
					}
					return nil, nil
				},
			},
			"primaryFunction": &graphql.Field{
				Type:        graphql.String,
				Description: "The primary function of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(StarWarsChar); ok {
						return droid.PrimaryFunction, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	/* Initialize root query */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hero": &graphql.Field{
				Type: characterInterface,
				Args: graphql.FieldConfigArgument{
					"episode": &graphql.ArgumentConfig{
						Description: "If omitted, returns the hero of the whole saga. If " +
							"provided, returns the hero of that particular episode.",
						Type: episodeEnum,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetHero(p.Args["episode"]), nil
				},
			},
			"human": &graphql.Field{
				Type: humanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//  Get the id:
					text, isOK := p.Args["id"].(string)
					if isOK {
						id, _ := strconv.Atoi(text)
						rethuman, _ := db.GetHuman(id)
						return rethuman, nil
					}
					rethuman, _ := db.GetHuman(0)
					return rethuman, nil
				},
			},
			"droid": &graphql.Field{
				Type: droidType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the droid",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := strconv.Atoi(p.Args["id"].(string))
					return GetDroid(id), nil
				},
			},
		},
	})

	/* Initialize schema based on the root query */
	StarWarsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

// GetHuman helper function
/*
func GetHuman(id int) StarWarsChar {
	if human, ok := HumanData[id]; ok {
		return human
	}
	return StarWarsChar{}
}
*/

// GetDroid helper function
func GetDroid(id int) StarWarsChar {
	if droid, ok := DroidData[id]; ok {
		return droid
	}
	return StarWarsChar{}
}

// GetHero helper function
func GetHero(episode interface{}) interface{} {
	if episode == 5 {
		return Luke
	}
	return Artoo
}
