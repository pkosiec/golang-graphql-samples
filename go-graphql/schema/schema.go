package schema

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
)

func New() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RootQuery",
			Description: "All available queries",
			Fields: graphql.Fields{
				"artist": &graphql.Field{
					Name:        "Artist Query",
					Description: "Get artist",
					Type:        ArtistType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						id := params.Args["id"]

						if id == "404" {
							return nil, nil
						}

						return Artist{
							ID:   "krzysztofKrawczyk",
							Name: "Krzysztof Krawczyk",
						}, nil
					},
				},
				"artists": &graphql.Field{
					Name:        "Artists Query",
					Description: "Get all artists",
					Type:        graphql.NewNonNull(graphql.NewList(ArtistType)),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return []Artist{
							{
								ID:   "krzysztofKrawczyk",
								Name: "Krzysztof Krawczyk",
							},
						}, nil
					},
				},
			},
		}),
		Mutation: nil,
	}

	// Field for ArtistType: songs: [Song!]!
	ArtistType.AddFieldConfig("songs", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(SongType))),
		Name:        "Songs",
		Description: "Songs field description",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			root, ok := params.Source.(Artist)
			if !ok {
				return nil, errors.New("source is nil")
			}

			log.Printf("\nArtists songs for %s\n", root.Name)

			return []Song{
				{
					ID:       root.ID,
					Title:    "Mój przyjacielu",
					Duration: 40,
				},
				{
					ID:       root.ID,
					Title:    "Parostatek",
					Duration: 50,
				},
			}, nil
		},
	})

	return graphql.NewSchema(schemaConfig)
}
