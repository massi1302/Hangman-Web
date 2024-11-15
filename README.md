# Hangman Web 🎮

Un jeu du pendu interactif avec une interface web, développé en Go.

## 📝 Description

Le projet **Hangman Web** est une version web du célèbre jeu du pendu,précédemment développé en CLI (ligne de commande). Cette version intègre une interface graphique accessible via un navigateur, permettant de jouer au pendu avec différentes fonctionnalités interactives.
L'objectif principal est de réutiliser le module **Hangman CLI** implémentée en utilisant Go pour le backend et une interface utilisateur HTML/CSS pure pour le frontend. Le projet utilise un serveur HTTP natif et des templates Go pour créer une expérience de jeu interactive et engageante.

## ✨ Fonctionnalités

### Page de lancement du jeu
- Sélection d'une nouvelle partie
- Saisie du pseudo joeur
- Sélection de la page de scores et historique

### Page d'Accueil 
- Sélection du niveau de difficulté
- Redirection automatique vers la page de jeu si une partie est en cours

### Page de Jeu
- Affichage du mot masqué
- Les lettres déjà essayées
- Visualisation du pendu évolutive
- Compteur de vies restantes (coeurs)

### Page de Fin de Partie
- Messages  de victoire/défaite
- Option pour rejouer
- Option pour aller a l'acceuil 

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
git clone https://github.com/massi1302/Hangman-Web.git
cd HANGMAN-WEB
go mod tidy
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
HANGMAN-WEB/
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

## 📄 Documentation des Routes  
Routes de Vues (Frontend)
- GET / : Page d'accueil pour démarrer une nouvelle partie.
- GET /game : Page de jeu pour jouer au pendu.
- GET /end : Page de fin de partie affichant le résultat.
- GET /scores : Page affichant le tableau des scores.
Routes API (Backend)
- POST /start : Démarrer une nouvelle partie avec le pseudo et le niveau de difficulté.
- POST /guess : Envoyer une lettre ou un mot pour deviner.
- GET /leaderboard : Récupérer les scores.

## 🎮 Comment Jouer

1. Accédez à la page de lancement du jeu
2. Entrez votre pseudo et choisissez la difficulté
3. Devinez le mot en proposant des lettres 
4. Surveillez vos points de vie restants
5. Consultez le tableau des scores pour voir votre classement

## ⚙️ Configuration Requise

- Go 1.23 ou supérieur
- Navigateur web moderne

## 👥 Équipe

- [Massinissa AHFIR]  Frontend &  Backend
- [Antony FONTAINE]   Frontend

## 📊 Synthèse du Projet
Le rapport détaillant le déroulement du projet, la répartition des tâches, et la gestion du temps est disponible dans le fichier rapport.pdf.

## 📚 Ressources Utilisées
Documentation officielle Go : https://golang.org/doc/
Tutoriels et exemples sur la création d'un serveur HTTP en Go
Stack Overflow pour résoudre les problématiques liées à l'intégration des templates HTML

## 📝 Licence

Ce projet est sous licence [Type de licence]

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou proposer une pull request.
