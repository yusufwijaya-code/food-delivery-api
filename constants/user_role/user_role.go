package user_role

type UserRole string

const (
	Customer   UserRole = "customer"
	Restaurant UserRole = "restaurant"
	Driver     UserRole = "driver"
	Admin      UserRole = "admin"
)
