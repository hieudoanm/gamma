package tasks_schema

import "github.com/graphql-go/graphql"

var Task = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Name: "title",
			Type: graphql.String,
		},
		"completed": &graphql.Field{
			Name: "completed",
			Type: graphql.Boolean,
		},
	},
})
