package resolvers

import (
	 "github.com/bcmmbaga/scale/service"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	service service.Service
}
