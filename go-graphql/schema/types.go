package schema

import "github.com/graphql-go/graphql"

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Duration float64 `json:"duration"`
}

var ArtistType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Artist",
		Description: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var SongType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Song",
		Description: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"duration": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	},
)
