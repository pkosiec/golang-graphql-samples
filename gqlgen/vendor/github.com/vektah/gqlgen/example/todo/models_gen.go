// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package todo

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
type TodoInput struct {
	Text string `json:"text"`
	Done *bool  `json:"done"`
}
