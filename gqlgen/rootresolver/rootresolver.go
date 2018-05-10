package rootresolver

import (
	"context"
	"errors"
	"log"

	"github.com/pkosiec/golang-graphql-samples/gqlgen/gqlschema"
)

type RootResolver struct{}

func New() *RootResolver {
	return &RootResolver{}
}

func (r *RootResolver) Artist_songs(ctx context.Context, obj *gqlschema.Artist) ([]gqlschema.Song, error) {
	if obj == nil {
		return []gqlschema.Song{}, errors.New("obj is nil")
	}
	log.Printf("\nArtists songs for %s\n", obj.Name)

	return []gqlschema.Song{
		{
			ID:       obj.ID,
			Title:    "Mój przyjacielu",
			Duration: 40,
		},
		{
			ID:       obj.ID,
			Title:    "Parostatek",
			Duration: 50,
		},
	}, nil
}

func (r *RootResolver) Query_artist(ctx context.Context, id string) (*gqlschema.Artist, error) {
	if id == "404" {
		return nil, nil
	}

	return &gqlschema.Artist{
		ID:   "krzysztofKrawczyk",
		Name: "Krzysztof Krawczyk",
	}, nil
}

func (r *RootResolver) Query_artists(ctx context.Context) ([]gqlschema.Artist, error) {
	return []gqlschema.Artist{
		{
			ID:   "krzysztofKrawczyk",
			Name: "Krzysztof Krawczyk",
		},
	}, nil
}
