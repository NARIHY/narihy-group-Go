package main

import (
	"database/sql"
)

// User représente un utilisateur dans le système.
// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// }

// createUser crée un nouvel utilisateur dans la base de données.
func createUser(db *sql.DB, name string, age int) error {
	sqlStatement := "INSERT INTO users (name, age) VALUES (?, ?)"
	_, err := db.Exec(sqlStatement, name, age)
	if err != nil {
		return err
	}
	return nil
}

// getAllUsers récupère tous les utilisateurs de la base de données.
func getAllUsers(db *sql.DB) ([]User, error) {
	sqlStatement := "SELECT id, name, age FROM users"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// updateUser met à jour les informations d'un utilisateur dans la base de données.
func updateUser(db *sql.DB, id int, name string, age int) error {
	sqlStatement := "UPDATE users SET name = ?, age = ? WHERE id = ?"
	_, err := db.Exec(sqlStatement, name, age, id)
	if err != nil {
		return err
	}
	return nil
}

// deleteUser supprime un utilisateur de la base de données à partir de son ID.
func deleteUser(db *sql.DB, id int) error {
	sqlStatement := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

// userMenu est un menu interactif pour la gestion des utilisateurs.
// func userMenu(db *sql.DB) {
// 	for {
// 		fmt.Println("\nMenu Gestion des Utilisateurs:")
// 		fmt.Println("1. Créer un utilisateur")
// 		fmt.Println("2. Lire tous les utilisateurs")
// 		fmt.Println("3. Mettre à jour un utilisateur")
// 		fmt.Println("4. Supprimer un utilisateur")
// 		fmt.Println("0. Retour au menu principal")

// 		var choice int
// 		fmt.Print("Choix : ")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 0:
// 			return
// 		case 1:
// 			createUserMenu(db)
// 		case 2:
// 			readAllUsers(db)
// 		case 3:
// 			updateUserMenu(db)
// 		case 4:
// 			deleteUserMenu(db)
// 		default:
// 			fmt.Println("Choix invalide. Veuillez réessayer.")
// 		}
// 	}
// }
