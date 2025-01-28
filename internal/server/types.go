package server

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAcount(firstName, lastName string) *Account {
	return &Account{
		ID:        99,
		FirstName: firstName,
		LastName:  lastName,
		Number:    100,
		Balance:   1000,
	}
}
