package types

// CreateUser create repository types for user
type CreateUser struct {
	ID           int
	FirstName    string
	LastName     string
	MobileNumber string
}

// GetUser get repository types for user
type GetUser struct {
	ID int64
}

// UpdateUser update repository types for user
type UpdateUser struct {
	ID           int
	FirstName    string
	LastName     string
	MobileNumber string
}
