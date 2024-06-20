## THEME: EXPERIMENTATION DE BD AVEC GO
<h4> 
    Author: RANDRIANARISOA Mahenina
</h4>

## Description sur le projet
Je suis en train de développer une application en Go pour gérer des données utilisateur et des comptes bancaires dans une base de données MySQL. Voici une description concise de ce que vous faites :

1. **Connexion à la base de données** : Vous vous connectez à une base de données MySQL en utilisant le package `database/sql` et le pilote MySQL.

2. **Gestion des utilisateurs** :
   - Mis en place des opérations CRUD (Create, Read, Update, Delete) pour les utilisateurs (`createUser`, `getUsers`, `updateUser`, `deleteUser`).
   - Chaque opération manipule les données des utilisateurs dans la table `users` de votre base de données.

3. **Gestion des comptes bancaires** :
   - Mis en place des opérations pour gérer les comptes bancaires (`getAllAccounts`, `createAccount`, `updateAccount`, `deleteAccount`).
   - Ces opérations interagissent avec la table `accounts` de votre base de données.

4. **Fonctionnalités supplémentaires** :
   - Vous avez développé des fonctionnalités pour exporter les données des utilisateurs (`exportUsersToFile`) et des comptes bancaires (`exportAccountsToFile`) vers des fichiers texte.

5. **Structure du programme** :
   - Ce programme est organisé en plusieurs fichiers (`main.go`, `db.go`, `user.go`, `account.go`, etc.) pour mieux structurer et séparer les responsabilités des différentes parties de votre application.

En résumé, mon application utilise Go pour se connecter à une base de données MySQL, gérer les données d'utilisateurs et de comptes bancaires à travers des opérations CRUD, et offre des fonctionnalités pour exporter ces données vers des fichiers texte, fournissant ainsi une base solide pour développer une application de gestion de données utilisateur et financières.