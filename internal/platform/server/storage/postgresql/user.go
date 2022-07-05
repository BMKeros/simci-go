package postgresql

const (
	tableUsers = "users"
)

type sqlUser struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
