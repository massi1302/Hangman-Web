# Hangman Web 🎮

Un jeu du pendu interactif avec une interface web, développé en Go.

## 📝 Description

Hangman Web est une version web du célèbre jeu du pendu, implémentée en utilisant Go pour le backend et une interface utilisateur HTML/CSS pure pour le frontend. Le projet utilise un serveur HTTP natif et des templates Go pour créer une expérience de jeu interactive et engageante.

## ✨ Fonctionnalités

### Page d'Accueil
- Saisie du pseudo joueur
- Sélection du niveau de difficulté
- Redirection automatique vers la page de jeu si une partie est en cours

### Page de Jeu
- Affichage du mot masqué
- Liste des lettres déjà essayées
- Visualisation du pendu évolutive
- Compteur de vies restantes
- Messages de feedback sur les tentatives
- Validation des entrées (lettres uniquement)

### Page de Fin de Partie
- Messages aléatoires de victoire/défaite
- Option pour rejouer
- Redirection automatique vers le jeu si une partie est en cours

### Tableau des Scores
- Historique persistant des parties
- Stockage dans un fichier texte
- Classement des meilleurs scores

## 🛠️ Technologies Utilisées

- **Go** - Backend et logique de jeu
- **HTML/CSS** - Interface utilisateur
- **Templates Go** - Rendu dynamique des pages
- **Package `os`** - Gestion des fichiers pour les scores

## 🚀 Installation

1. Clonez le repository
```bash
git clone [URL_DU_REPO]
cd hangman-web
```

2. Lancez le serveur
```bash
go run main.go
```

3. Accédez au jeu via votre navigateur
```
http://localhost:8080
```

## 📂 Structure du Projet

```
hangman-web/
│
├── main.go                 # Point d'entrée de l'application
├── game/                   # Logic du jeu
├── templates/             # Templates HTML
│   ├── home.html
│   ├── game.html
│   ├── end.html
│   └── scores.html
├── static/               # Fichiers statiques
│   ├── css/
│   └── images/
└── data/                # Stockage des scores
```

## 🎮 Comment Jouer

1. Accédez à la page d'accueil
2. Entrez votre pseudo et choisissez la difficulté
3. Devinez le mot en proposant des lettres ou des mots complets
4. Surveillez vos points de vie restants
5. Consultez le tableau des scores pour voir votre classement

## ⚙️ Configuration Requise

- Go 1.16 ou supérieur
- Navigateur web moderne

## 👥 Contributeurs

- [Nom des contributeurs]

## 📝 Licence

Ce projet est sous licence [Type de licence]

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou proposer une pull request.
