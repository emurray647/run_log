package db

type User struct {
	ID uint
}

type Activity struct {
	ID       uint
	User     User
	DataGlob []byte
}
