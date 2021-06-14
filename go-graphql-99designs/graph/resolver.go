package graph

import "go-graphql-99designs/graph/model"

func NewResolver() *Resolver {
	users := make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", Name: "Gringo"})
	users = append(users, &model.User{ID: "2", Name: "Fabi"})
	return &Resolver{
		todos: make([]*model.Todo, 0),
		users: users,
		lastUserId: 3,
		usersChan: make(chan *model.User),
	}
}

type Resolver struct{
	todos [] *model.Todo
	users [] *model.User
	lastTodoId int
	lastUserId int
	usersChan chan *model.User
}
