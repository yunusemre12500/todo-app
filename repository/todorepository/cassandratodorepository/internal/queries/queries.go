package queries

import _ "embed"

var (
	//go:embed delete_todo_by_id.cql
	DeleteTodoByIDQuery string
	//go:embed get_todo_by_id.cql
	GetTodoByIDQuery string
	//go:embed get_todos_by_list_id.cql
	GetTodosByListIDQuery string
	//go:embed save_todo.cql
	SaveTodoQuery string
	//go:embed update_todo_by_id.cql
	UpdateTodoByIDQuery string
)
