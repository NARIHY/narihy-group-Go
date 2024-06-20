package main

// Définition des structures de données utilisées dans l'application
type (
	// User représente un utilisateur dans le système.
	User struct {
		ID   int
		Name string
		Age  int
	}

	// Account représente un compte bancaire dans le système.
	Account struct {
		ID       int
		OwnerID  int
		Balance  float64
		Currency string
	}
)
