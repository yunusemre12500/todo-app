module github.com/yunusemre12500/todo-app/repository/user

go 1.24.1

require (
	github.com/yunusemre12500/todo-app/model v1.0.0
	github.com/yunusemre12500/todo-app/repository v1.0.0
)

replace (
	github.com/yunusemre12500/todo-app/model => ../../model
	github.com/yunusemre12500/todo-app/repository => ../../repository
)
