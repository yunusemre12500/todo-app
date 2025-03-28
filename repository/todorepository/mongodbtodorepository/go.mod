module github.com/yunusemre12500/todo-app/repository/todorepository/mongodbtodorepository

go 1.24.1

require (
	github.com/yunusemre12500/todo-app/model v1.0.0
	github.com/yunusemre12500/todo-app/repository v1.0.0
	github.com/yunusemre12500/todo-app/repository/todorepository v1.0.0
	go.mongodb.org/mongo-driver/v2 v2.1.0
	golang.org/x/sync v0.12.0
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

replace (
	github.com/yunusemre12500/todo-app/model => ../../../model
	github.com/yunusemre12500/todo-app/repository => ../../../repository
	github.com/yunusemre12500/todo-app/repository/todorepository => ../../../repository/todorepository
)
