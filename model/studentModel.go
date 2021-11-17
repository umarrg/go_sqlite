package studentModel

type Student struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}

var CreateTableQuery = `create table if not exists students(
	id integer not null primary key autoincrement,
	firstName text,
	lastName text,
	email text,
	gender text,
	age integer
)`
