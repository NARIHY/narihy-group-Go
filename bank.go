package main

import (
	"database/sql"
)

// createAccount crée un nouveau compte dans la base de données.
func createAccount(db *sql.DB, ownerID int, initialBalance float64, currency string) error {
	sqlStatement := "INSERT INTO accounts (owner_id, balance, currency) VALUES (?, ?, ?)"
	_, err := db.Exec(sqlStatement, ownerID, initialBalance, currency)
	if err != nil {
		return err
	}
	return nil
}

// depositMoney dépose de l'argent sur un compte existant.
func depositMoney(db *sql.DB, accountID int, amount float64) error {
	sqlStatement := "UPDATE accounts SET balance = balance + ? WHERE owener_id = ?"
	_, err := db.Exec(sqlStatement, amount, accountID)
	if err != nil {
		return err
	}
	return nil
}

// withdrawMoney retire de l'argent d'un compte existant.
func withdrawMoney(db *sql.DB, accountID int, amount float64) error {
	sqlStatement := "UPDATE accounts SET balance = balance - ? WHERE id = ? AND balance >= ?"
	_, err := db.Exec(sqlStatement, amount, accountID, amount)
	if err != nil {
		return err
	}
	return nil
}

// getAllAccounts récupère tous les comptes de la base de données.
func getAllAccounts(db *sql.DB) ([]Account, error) {
	sqlStatement := "SELECT id, owner_id, balance, currency FROM accounts"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		err = rows.Scan(&account.ID, &account.OwnerID, &account.Balance, &account.Currency)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
