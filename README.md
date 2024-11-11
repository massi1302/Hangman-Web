🎮 Jeu du Pendu (Hangman Game)
Un jeu du pendu en ligne de commande développé en Go, avec plusieurs niveaux de difficulté et une interface utilisateur interactive.

📋 Table des matières
Fonctionnalités
Prérequis
Installation
Utilisation
Structure du projet
Règles du jeu
Contribution
✨ Fonctionnalités
🎯 Trois niveaux de difficulté (Facile, Moyen, Difficile)
🎨 Interface colorée et animations en ASCII art
💖 Système de vies avec affichage visuel
🔤 Option de révélation initiale de lettres
⌨️ Navigation intuitive dans les menus
🛠️ Prérequis
Go 1.16 ou supérieur
Les dépendances suivantes :
github.com/fatih/color
github.com/eiannone/keyboard
📥 Installation
Clonez le répertoire :

git clone https://github.com/votre-username/jeu-de-pendu.git
cd jeu-de-pendu
Installez les dépendances :

go mod download
🎮 Utilisation
Il existe deux façons de lancer le jeu :

Mode menu interactif :

go run main.go
Mode direct avec fichier de mots et lettres révélées :

go run main.go [fichier-mots] [nombre-lettres]
Exemples :

go run main.go easy-words.txt 2    # Mode facile avec 2 lettres révélées
go run main.go medium-words.txt 1   # Mode moyen avec 1 lettre révélée
go run main.go hard-words.txt 0     # Mode difficile sans lettre révélée
📁 Structure du projet
jeu-de-pendu/
├── main.go            # Point d'entrée du programme
├── game/
│   ├── hangman.go     # Logique principale du jeu
│   ├── menu.go        # Gestion du menu
│   ├── affichage.go   # Fonctions d'affichage et couleurs
│   ├── wordsutil.go   # Utilitaires de gestion des mots
│   ├── asciiart.go    # Art ASCII pour le pendu
│   └── clearconsole.go # Utilitaire console
├── data/
│   ├── easy-words.txt  # Liste de mots faciles
│   ├── medium-words.txt # Liste de mots moyens
│   └── hard-words.txt  # Liste de mots difficiles
└── README.md
📌 Règles du jeu
Un mot est choisi aléatoirement selon le niveau de difficulté
Le joueur commence avec 10 vies (❤️)
À chaque tour, le joueur peut :
Proposer une lettre
Deviner le mot complet
Le joueur perd une vie (💔) pour chaque :
Lettre incorrecte
Mot incorrect (2 vies)
La partie est gagnée si le mot est trouvé avant de perdre toutes les vies
🎯 Niveaux de difficulté
Facile : Mots courts (4-5 lettres)
Moyen : Mots de longueur moyenne
Difficile : Mots longs et complexes
🤝 Contribution
Les contributions sont les bienvenues ! Pour contribuer :

Forkez le projet
Créez une branche pour votre fonctionnalité
Committez vos changements
Poussez vers la branche
Ouvrez une Pull Request
Développé avec ❤️ et Go
