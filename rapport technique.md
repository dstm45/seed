# Rapport Technique de l'Application "Seed"

## 1. Introduction

Ce document fournit une analyse technique de l'application "Seed". L'application est un service web de gestion et de découverte d'événements, développé en Go pour le backend et utilisant le moteur de templating `templ` avec TailwindCSS pour le frontend. Elle offre des fonctionnalités pour les utilisateurs standards, les organisateurs d'événements et les administrateurs.

L'architecture est monolithique, avec une séparation claire des préoccupations (contrôleurs, base de données, vues) et une forte dépendance à l'écosystème standard de Go.

## 2. Structure du Projet

Le projet est organisé de manière logique, séparant le code de l'application (`pkg`) du point d'entrée (`cmd`).

-   `cmd/main.go`: Point d'entrée de l'application.
-   `pkg/`: Contient toute la logique métier de l'application.
    -   `config/`: Configuration et initialisation du serveur web.
    -   `controllers/`: Gère la logique des requêtes HTTP et les interactions avec les vues.
    -   `database/`: Gère la connexion, le schéma, les migrations et les requêtes à la base de données.
    -   `middlewares/`: Contient les middlewares pour l'authentification et l'autorisation.
    -   `utils/`: Fonctions utilitaires (JWT, hash, gestion de fichiers).
    -   `views/`: Contient les templates `templ` pour le rendu HTML, organisés par fonctionnalité (admin, user, etc.).
-   `.env`: Fichier (non fourni) pour les variables d'environnement (URL de la base de données, secret JWT).
-   `sqlc.yaml`: Fichier de configuration pour l'outil `sqlc`.

## 3. Backend (Go)

### 3.1. Point d'Entrée et Serveur

-   **`cmd/main.go`**: Le programme principal charge les variables d'environnement depuis un fichier `.env` à l'aide de `joho/godotenv`. Ensuite, il initialise et démarre le serveur défini dans le package `config`.
-   **`pkg/config/server.go`**: Ce fichier est le cœur de l'application.
    -   Il utilise le `http.ServeMux` de la bibliothèque standard de Go pour le routage.
    -   Il initialise la connexion à la base de données et les contrôleurs.
    -   Les routes sont définies manuellement, associant des URL à des handlers de contrôleurs.
    -   Les middlewares d'authentification (`IsAuthenticated`) et d'autorisation (`IsAdmin`) sont appliqués de manière sélective à certaines routes.
    -   Un serveur de fichiers statiques est configuré pour servir les assets (CSS, JS, images) depuis `/pkg/views/static/` via l'URL `/static/`.
    -   Le serveur est configuré pour écouter sur le port `8888`.

### 3.2. Contrôleurs

Les contrôleurs sont responsables de la gestion des requêtes HTTP. Chaque contrôleur est une struct qui détient une dépendance vers la base de données (`*database.Queries`).

-   **`AuthController`**: Gère l'inscription (`SignUp`), la connexion (`SignIn`) et la déconnexion (`Logout`).
    -   `SignUp`: Valide le formulaire, hashe le mot de passe avec `bcrypt` et crée un nouvel utilisateur.
    -   `SignIn`: Vérifie les identifiants, compare le hash du mot de passe et, en cas de succès, génère un cookie JWT.
-   **`UserController`**: Gère les actions des utilisateurs connectés.
    -   `Index`: Affiche le tableau de bord de l'utilisateur avec ses informations et les événements qu'il a créés.
    -   `EditProfile`: Gère la mise à jour du profil, y compris le téléchargement d'un avatar (via `multipart/form-data`).
    -   `EditPassword`: Permet de changer le mot de passe.
-   **`AdminController`**: Gère le panneau d'administration, avec des pages pour le tableau de bord, la gestion des utilisateurs et des événements.

### 3.3. Middlewares

-   **`pkg/middlewares/authentication.go`**:
    -   `IsAuthenticated`: Vérifie la validité d'un cookie JWT nommé "auth". Si le cookie est absent ou invalide, l'utilisateur est redirigé vers la page de connexion. Si valide, les informations (claims) du token sont ajoutées au contexte de la requête pour être utilisées par les handlers suivants.
    -   `IsAdmin`: S'assure que l'utilisateur authentifié a le rôle "admin". Ce middleware est appliqué après `IsAuthenticated`.

### 3.4. Base de Données

-   **Moteur**: PostgreSQL.
-   **Connexion**: Le package `database` utilise `jackc/pgx` pour se connecter à la base de données via une URL fournie par la variable d'environnement `DATABASE_URL`.
-   **Génération de Code (sqlc)**: Le projet utilise `sqlc` pour générer du code Go type-safe à partir de requêtes SQL brutes.
    -   Les requêtes SQL sont définies dans `pkg/database/queries/`.
    -   `sqlc` génère les fichiers `*.sql.go` et `models.go`, qui fournissent des fonctions Go pour exécuter ces requêtes et des structs pour modéliser les tables de la base de données.
-   **Migrations**: Les migrations de schéma sont gérées avec `goose`. Les fichiers de migration se trouvent dans `pkg/database/migrations/` et décrivent l'évolution de la structure de la base de données.
-   **Schéma**: Les tables principales incluent `users`, `evenements`, `abonnements`, `commentaires`, et `activites`. Les relations clés sont établies via des clés étrangères (par exemple, `evenements.organisateur` référence `users.id`).

### 3.5. Utilitaires

-   **`pkg/utils/authentication.go`**: Gère la création et la validation de tokens JWT (avec `github.com/golang-jwt/jwt/v5`). Le secret pour signer les tokens est récupéré depuis la variable d'environnement `SECRET`.
-   **`pkg/utils/hash.go`**: Fonctions pour hacher les mots de passe avec `bcrypt`.
-   **`pkg/utils/files.go`**: Logique pour sauvegarder les fichiers téléchargés.

## 4. Frontend

Le frontend est rendu côté serveur à l'aide du moteur de templating `templ`.

### 4.1. Templating (`templ`)

-   Les vues sont définies dans des fichiers `.templ`. Ces fichiers contiennent un mélange de syntaxe HTML et de code Go.
-   Le compilateur `templ` traite ces fichiers pour générer du code Go (`_templ.go`) qui peut être appelé directement depuis les contrôleurs pour rendre le HTML.
-   **Composants**: Le projet utilise une approche par composants. `pkg/views/component/` contient des composants réutilisables comme `Base` (la structure de page principale), `Header`, et `Button`.
-   **Vues**: Les vues spécifiques à chaque section (admin, utilisateur, etc.) sont organisées dans leurs propres dossiers.

### 4.2. Style (TailwindCSS)

-   Le style est géré par TailwindCSS.
-   Le fichier de configuration de base est `tailwind.config.js`.
-   Le fichier source CSS est `pkg/views/static/css/input.css`, qui utilise les directives `@tailwind`.
-   Le fichier compilé `output.css` est servi comme un asset statique et est lié dans le composant de base (`Base.templ`).

## 5. Fonctionnalités Clés

-   **Authentification**: Système complet basé sur JWT avec des cookies HTTP.
-   **Rôles Utilisateurs**: Distinction entre utilisateurs standards et administrateurs, protégée par des middlewares.
-   **Gestion de Profil**: Les utilisateurs peuvent modifier leurs informations personnelles et leur photo de profil.
-   **Gestion d'Événements**: Les utilisateurs peuvent créer et voir des événements. Les administrateurs ont des vues dédiées pour la gestion.
-   **Interface d'Administration**: Un panneau d'administration simple pour superviser l'application.

## 6. Axes d'Amélioration Potentiels

-   **Validation des Données**: La validation des données des formulaires pourrait être renforcée et centralisée, potentiellement à l'aide d'une bibliothèque dédiée pour éviter la logique de validation dispersée dans les contrôleurs.
-   **Gestion des Erreurs**: La gestion des erreurs est principalement gérée par des logs (`utils.AfficherErreur`). Une approche plus structurée pourrait être mise en place pour retourner des pages d'erreur plus informatives à l'utilisateur.
-   **Duplication de Code**: La logique de sauvegarde de fichiers semble être présente à la fois dans `pkg/utils/files.go` et `pkg/controllers/userControllers.go`. Ce code pourrait être unifié.
-   **Sécurité**: Le décodage de token dans `utils.DecodeToken` ne semble pas toujours vérifier la validité du token, ce qui pourrait être un risque si utilisé pour des décisions de sécurité. Il est préférable de toujours utiliser `ParseToken` qui inclut la validation.
-   **Frontend**: L'interactivité du frontend est limitée. L'utilisation d'une bibliothèque JavaScript légère comme Alpine.js (qui est déjà importée mais non utilisée dans `index.js`) ou htmx pourrait améliorer l'expérience utilisateur sans nécessiter un framework lourd.
