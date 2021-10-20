package types

// CreateUser create service type for user
type CreateUser struct {
	FirstName    string
	LastName     string
	MobileNumber string
	Password     string
}

// GetUser get service type for user
type GetUser struct {
	ID int64
}

// UpdateUser update service type for user
type UpdateUser struct {
	ID           int
	FirstName    string
	LastName     string
	MobileNumber string
	Password     string
}
