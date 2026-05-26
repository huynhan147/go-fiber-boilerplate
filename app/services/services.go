package services

type Services struct {
	Auth        AuthService
	User        UserService
	Transaction TransactionService
}
