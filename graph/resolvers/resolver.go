package resolvers

import "github.com/bcmmbaga/scale"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	service scale.Service
}
