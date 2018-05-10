# GraphQL API Samples for Go
Sample apps written in Go using two popular solutions:
- [GQLGen](https://github.com/vektah/gqlgen)
- [GraphQL Go](https://github.com/graphql-go/graphql)

## Sample app
Those apps use following GraphQL schema:

```
type Artist {
  id: ID!
  name: String!
  songs: [Song!]!
}

type Song {
  id: ID!
  title: String!
  duration: Float!
}

type Query {
  artist(id: ID!): Artist
  artists: [Artist!]!
}
```
