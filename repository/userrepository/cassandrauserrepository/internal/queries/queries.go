package queries

import _ "embed"

var (
	//go:embed delete_user_by_id.cql
	DeleteUserByIDQuery string
	//go:embed get_user_by_id.cql
	GetUserByIDQuery string
	//go:embed get_user_by_name.cql
	GetUserByNameQuery string
	SaveUserQuery      string = "INSERT INTO ? (created_at, display_name, email_address, id, name, password_hash) VALUES (?, ?, ?, ?, ?, ?);"
)
