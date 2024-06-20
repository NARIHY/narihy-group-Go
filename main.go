package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost     = "localhost"
	dbPort     = 3307
	dbUser     = "root"
	dbPassword = ""
	dbName     = "go"
)

func main() {
	// Connexion à la base de données
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Menu principal
	for {
		fmt.Println("NARIHY GROUP: \n")
		fmt.Println("\nMenu Principal:")
		fmt.Println("1. Gestion des utilisateurs")
		fmt.Println("2. Gestion des comptes bancaires")
		fmt.Println("3. Exporter les utilisateurs dans un fichier txt")
		fmt.Println("0. Quitter")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			fmt.Println("Merci d'avoir utilisé notre service.")
			return
		case 1:
			userMenu(db)
		case 2:
			bankMenu(db)
		case 3:
			exportUsersToFile(db)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func exportUsersToFile(db *sql.DB) {
	// Nom du fichier de sortie
	filename := "users.txt"

	// Ouvrir le fichier en écriture
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Impossible de créer le fichier %s: %v", filename, err)
	}
	defer file.Close()

	// Récupérer les utilisateurs depuis la base de données
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des utilisateurs depuis la base de données: %v", err)
	}

	// Écrire les utilisateurs dans le fichier
	for _, user := range users {
		line := fmt.Sprintf("%d, %s, %d\n", user.ID, user.Name, user.Age)
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatalf("Erreur lors de l'écriture dans le fichier %s: %v", filename, err)
		}
	}

	fmt.Printf("Exportation des utilisateurs terminée. Les données ont été écrites dans le fichier %s.\n", filename)
}

func userMenu(db *sql.DB) {
	for {
		fmt.Println("\nMenu Gestion des Utilisateurs:")
		fmt.Println("1. Créer un utilisateur")
		fmt.Println("2. Lire tous les utilisateurs")
		fmt.Println("3. Mettre à jour un utilisateur")
		fmt.Println("4. Supprimer un utilisateur")
		fmt.Println("0. Retour au menu principal")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			return
		case 1:
			createUserMenu(db)
		case 2:
			readAllUsers(db)
		case 3:
			updateUserMenu(db)
		case 4:
			deleteUserMenu(db)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func bankMenu(db *sql.DB) {
	for {
		fmt.Println("\nMenu Gestion des Comptes Bancaires:")
		fmt.Println("1. Créer un compte bancaire")
		fmt.Println("2. Lire tous les comptes bancaires")
		fmt.Println("3. Effectuer un dépôt")
		fmt.Println("4. Effectuer un retrait")
		fmt.Println("0. Retour au menu principal")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			return
		case 1:
			createAccountMenu(db)
		case 2:
			readAllAccounts(db)
		case 3:
			depositMenu(db)
		case 4:
			withdrawMenu(db)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

// Fonctions pour les sous-menus spécifiques à la gestion des utilisateurs

func createUserMenu(db *sql.DB) {
	var name string
	var age int

	fmt.Print("Nom de l'utilisateur : ")
	fmt.Scanln(&name)
	fmt.Print("Âge de l'utilisateur : ")
	fmt.Scanln(&age)

	err := createUser(db, name, age)
	if err != nil {
		fmt.Println("Erreur lors de la création de l'utilisateur:", err)
	} else {
		fmt.Println("Utilisateur créé avec succès.")
	}
}

func readAllUsers(db *sql.DB) {
	users, err := getAllUsers(db)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des utilisateurs:", err)
		return
	}

	fmt.Println("Liste des utilisateurs :")
	fmt.Println("ID   | Nom        | Âge")
	fmt.Println("-----|------------|-----")
	for _, user := range users {
		fmt.Printf("%-5d| %-11s| %-5d\n", user.ID, user.Name, user.Age)
	}

}

func updateUserMenu(db *sql.DB) {
	var id int
	var name string
	var age int

	fmt.Print("ID de l'utilisateur à mettre à jour : ")
	fmt.Scanln(&id)
	fmt.Print("Nouveau nom de l'utilisateur : ")
	fmt.Scanln(&name)
	fmt.Print("Nouvel âge de l'utilisateur : ")
	fmt.Scanln(&age)

	err := updateUser(db, id, name, age)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour de l'utilisateur:", err)
	} else {
		fmt.Println("Utilisateur mis à jour avec succès.")
	}
}

func deleteUserMenu(db *sql.DB) {
	var id int

	fmt.Print("ID de l'utilisateur à supprimer : ")
	fmt.Scanln(&id)

	err := deleteUser(db, id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression de l'utilisateur:", err)
	} else {
		fmt.Println("Utilisateur supprimé avec succès.")
	}
}

// Fonctions pour les sous-menus spécifiques à la gestion des comptes bancaires

func createAccountMenu(db *sql.DB) {
	var ownerID int
	var initialBalance float64
	var currency string

	fmt.Print("ID du propriétaire du compte : ")
	fmt.Scanln(&ownerID)
	fmt.Print("Solde initial du compte : ")
	fmt.Scanln(&initialBalance)
	fmt.Print("Devise du compte : ")
	fmt.Scanln(&currency)

	err := createAccount(db, ownerID, initialBalance, currency)
	if err != nil {
		fmt.Println("Erreur lors de la création du compte bancaire:", err)
	} else {
		fmt.Println("Compte bancaire créé avec succès.")
	}
}

func readAllAccounts(db *sql.DB) {
	accounts, err := getAllAccounts(db)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des comptes bancaires:", err)
		return
	}

	fmt.Println("Liste des comptes bancaires :")
	fmt.Println("ID   | Propriétaire ID | Solde   | Devise")
	fmt.Println("-----|-----------------|---------|--------")
	for _, account := range accounts {
		fmt.Printf("%-5d| %-15d| %-8.2f| %-7s\n", account.ID, account.OwnerID, account.Balance, account.Currency)
	}

}

func depositMenu(db *sql.DB) {
	var accountID int
	var amount float64

	fmt.Print("ID du compte pour le dépôt : ")
	fmt.Scanln(&accountID)
	fmt.Print("Montant à déposer : ")
	fmt.Scanln(&amount)

	err := depositMoney(db, accountID, amount)
	if err != nil {
		fmt.Println("Erreur lors du dépôt d'argent:", err)
	} else {
		fmt.Println("Dépôt effectué avec succès.")
	}
}

func withdrawMenu(db *sql.DB) {
	var accountID int
	var amount float64

	fmt.Print("ID du compte pour le retrait : ")
	fmt.Scanln(&accountID)
	fmt.Print("Montant à retirer : ")
	fmt.Scanln(&amount)

	err := withdrawMoney(db, accountID, amount)
	if err != nil {
		fmt.Println("Erreur lors du retrait d'argent:", err)
	} else {
		fmt.Println("Retrait effectué avec succès.")
	}
}
