module github.com/yunusemre12500/todo-app/repository/todorepository/cassandratodorepository

go 1.24.1

require (
	github.com/gocql/gocql v1.7.0
	github.com/yunusemre12500/todo-app/model v1.0.0
	github.com/yunusemre12500/todo-app/repository v1.0.0
	github.com/yunusemre12500/todo-app/repository/todorepository v1.0.0
	golang.org/x/sync v0.12.0
)

require (
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace (
	github.com/yunusemre12500/todo-app/model => ../../../model
	github.com/yunusemre12500/todo-app/repository => ../../../repository
	github.com/yunusemre12500/todo-app/repository/todorepository => ../../../repository/todorepository
)
