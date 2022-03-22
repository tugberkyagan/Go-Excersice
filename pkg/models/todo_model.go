package models

type Todo struct {
	ID       int
	Title    string
	UserID   int
	Progress int
	IsDone   bool
}
