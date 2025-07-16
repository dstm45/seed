# Documentation Technique de l'Application "Seed"

## 1. Vue d'ensemble

"Seed" est une application web complète pour la gestion et la découverte d'événements. Construite avec Go pour le backend et `templ` pour le rendu frontend, elle offre une plateforme robuste pour les utilisateurs, les organisateurs d'événements et les administrateurs.

L'architecture est monolithique, avec une séparation claire des préoccupations (contrôleurs, base de données, vues) et une forte dépendance à l'écosystème standard de Go.

## 2. Démarrage Rapide

### Prérequis

-   Go (version 1.21 ou supérieure)
-   PostgreSQL
-   Node.js et npm (pour la gestion des dépendances frontend)
-   `goose` pour les migrations de base de données (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
-   `sqlc` pour la génération de code SQL (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)
-   `templ` pour la génération de templates (`go install github.com/a-h/templ/cmd/templ@latest`)

### Installation

1.  **Cloner le dépôt :**
    ```bash
    git clone <url-du-repo>
    cd seed
    ```

2.  **Configurer la base de données :**
    -   Assurez-vous que PostgreSQL est en cours d'exécution.
    -   Créez une base de données pour le projet.
    -   Appliquez les migrations avec `goose`:
        ```bash
        goose -dir "pkg/database/migrations" postgres "user=<user> password=<password> dbname=<dbname> sslmode=disable" up
        ```

3.  **Configurer les variables d'environnement :**
    -   Créez un fichier `.env` à la racine du projet en vous basant sur `.env.example` (s'il existe) ou en ajoutant les variables suivantes :
        ```env
        DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"
        SECRET="votre_secret_jwt_super_secret"
        ```

4.  **Installer les dépendances frontend :**
    ```bash
    npm install
    ```

5.  **Générer les assets et le code :**
    -   Générez le code Go à partir des requêtes SQL :
        ```bash
        sqlc generate
        ```
    -   Générez les templates Go à partir des fichiers `.templ`:
        ```bash
        templ generate
        ```
    -   Compilez les fichiers CSS avec TailwindCSS :
        ```bash
        npm run build:css
        ```

6.  **Lancer l'application :**
    ```bash
    go run ./cmd/main.go
    ```
    Le serveur sera accessible à l'adresse `http://localhost:8888`.

## 3. Architecture et Structure du Projet

L'application suit une architecture monolithique avec une séparation claire des responsabilités.

-   `cmd/main.go`: Point d'entrée de l'application. Il charge les variables d'environnement et démarre le serveur.
-   `pkg/`: Contient toute la logique métier de l'application.
    -   `config/`: Gère la configuration du serveur, le routage (`http.ServeMux`), l'injection des dépendances et la configuration du serveur de fichiers statiques.
    -   `controllers/`: Contient les handlers HTTP qui répondent aux requêtes, interagissent avec la base de données et rendent les vues.
    -   `database/`: Gère tout ce qui concerne la base de données.
        -   `connexion.go`: Établit la connexion à la base de données.
        -   `migrations/`: Contient les fichiers de migration `goose`.
        -   `queries/`: Contient les requêtes SQL brutes pour `sqlc`.
        -   Fichiers `*.sql.go`: Code Go type-safe généré par `sqlc`.
    -   `middlewares/`: Contient les middlewares pour l'authentification (`IsAuthenticated`) et la vérification des rôles (`IsAdmin`).
    -   `utils/`: Fonctions utilitaires pour le hashage (`bcrypt`), la gestion des JWT, et la sauvegarde de fichiers.
    -   `views/`: Contient les templates `templ` et les assets statiques.
        -   `component/`: Composants de vue réutilisables (headers, layout de base).
        -   `static/`: Fichiers CSS, JavaScript et images.
-   `sqlc.yaml`: Fichier de configuration pour l'outil `sqlc`.
-   `.env`: Fichier (non versionné) pour les variables d'environnement.

## 4. Backend en Détail

### 4.1. Contrôleurs

Les contrôleurs sont responsables de la gestion des requêtes HTTP. Chaque contrôleur est une struct qui détient une dépendance vers la base de données (`*database.Queries`).

-   **`AuthController`**: Gère l'inscription (`SignUp`), la connexion (`SignIn`) et la déconnexion (`Logout`).
-   **`UserController`**: Gère les actions des utilisateurs connectés comme la modification du profil et du mot de passe.
-   **`AdminController`**: Gère le panneau d'administration (dashboard, gestion des utilisateurs et des événements).

### 4.2. Middlewares

-   **`IsAuthenticated`**: Vérifie la validité d'un cookie JWT nommé "auth". Si le cookie est absent ou invalide, l'utilisateur est redirigé vers la page de connexion. Sinon, les informations (claims) du token sont ajoutées au contexte de la requête.
-   **`IsAdmin`**: S'assure que l'utilisateur authentifié a le rôle "admin". Ce middleware est appliqué après `IsAuthenticated`.

### 4.3. Utilitaires

-   **`pkg/utils/authentication.go`**: Gère la création et la validation de tokens JWT (avec `github.com/golang-jwt/jwt/v5`).
-   **`pkg/utils/hash.go`**: Fonctions pour hacher et comparer les mots de passe avec `bcrypt`.
-   **`pkg/utils/files.go`**: Logique pour sauvegarder les fichiers téléchargés (ex: photos de profil).

## 5. Frontend en Détail

Le frontend est rendu côté serveur, ce qui le rend simple et performant.

### 5.1. Templating (`templ`)

-   Le projet utilise `templ`, un moteur de templates qui combine la syntaxe HTML avec la puissance de Go.
-   Les vues sont définies dans des fichiers `.templ` et compilées en code Go (`_templ.go`), garantissant la sécurité de type et d'excellentes performances.
-   Une approche par composants est utilisée (`pkg/views/component/`) pour les éléments réutilisables comme le layout de base et les headers.

### 5.2. Style (TailwindCSS)

-   Le style est géré par TailwindCSS pour un développement rapide d'interfaces modernes.
-   Le fichier source est `pkg/views/static/css/input.css`.
-   La commande `npm run build:css` compile le CSS final dans `pkg/views/static/css/output.css`, qui est ensuite servi statiquement.

## 6. Base de Données

### 6.1. Moteur et Connexion

-   **Moteur**: PostgreSQL.
-   **Connexion**: Le package `database` utilise `jackc/pgx` pour se connecter à la base de données via une URL fournie par la variable d'environnement `DATABASE_URL`.

### 6.2. Génération de Code (`sqlc`)

-   `sqlc` génère du code Go type-safe à partir de requêtes SQL brutes définies dans `pkg/database/queries/`.
-   Cela élimine le besoin d'écrire du code SQL répétitif et prévient les injections SQL.

### 6.3. Migrations (`goose`)

-   `goose` est utilisé pour gérer les migrations du schéma de la base de données.
-   Pour créer une nouvelle migration, utilisez la commande :
    ```bash
    goose -dir "pkg/database/migrations" create nom_de_la_migration sql
    ```

## 7. Flux d'Authentification

L'authentification est basée sur les JSON Web Tokens (JWT) stockés dans des cookies.

1.  **Connexion (`/signin`)**: L'utilisateur soumet ses identifiants. Le serveur vérifie les informations, et si elles sont correctes, génère un token JWT.
2.  **Stockage du Token**: Le JWT est encapsulé dans un cookie HTTP `auth` qui est envoyé au client.
3.  **Requêtes Protégées**: Pour chaque requête vers une route protégée, le middleware `IsAuthenticated` intercepte la requête.
4.  **Validation**: Le middleware vérifie la présence et la validité du cookie `auth`. Si le token est valide, les informations de l'utilisateur sont extraites et passées au handler suivant. Sinon, l'utilisateur est redirigé.
5.  **Déconnexion (`/logout`)**: Le cookie `auth` est supprimé, invalidant la session de l'utilisateur.

## 8. Endpoints de l'API

Voici un aperçu des routes principales de l'application.

### 8.1. Authentification

-   `GET /signup`: Affiche la page d'inscription.
-   `POST /signup`: Traite le formulaire d'inscription.
-   `GET /signin`: Affiche la page de connexion.
-   `POST /signin`: Traite le formulaire de connexion.
-   `GET /logout`: Déconnecte l'utilisateur.

### 8.2. Utilisateurs (Authentification requise)

-   `GET /` ou `GET /index`: Affiche le tableau de bord de l'utilisateur.
-   `GET /editprofile`: Affiche la page de modification du profil.
-   `POST /editprofile`: Met à jour les informations du profil.
-   `GET /editpassword`: Affiche la page de changement de mot de passe.
-   `POST /editpassword`: Met à jour le mot de passe de l'utilisateur.
-   `GET /liste`: Affiche la liste de tous les événements.
-   `GET /evenement/detail/{uuid}`: Affiche les détails d'un événement spécifique.

### 8.3. Administration (Rôle "admin" requis)

-   `GET /admin/dashboard`: Affiche le tableau de bord de l'administrateur.
-   `GET /admin/settings`: Affiche la page des paramètres.
-   `GET /admin/encours`: Affiche les événements en cours.
-   `GET /admin/manageusers`: Affiche la page de gestion des utilisateurs.
-   `GET /admin/managevent`: Affiche la page de gestion des événements.

## 9. Axes d'Amélioration Potentiels

-   **Validation des Données**: La validation des formulaires pourrait être centralisée à l'aide d'une bibliothèque dédiée pour éviter la logique répétitive dans les contrôleurs.
-   **Gestion des Erreurs**: Mettre en place une gestion des erreurs plus structurée pour retourner des pages d'erreur informatives à l'utilisateur au lieu de simples logs.
-   **Refactoring**: Unifier la logique de sauvegarde de fichiers, qui semble dupliquée entre `pkg/utils/files.go` et `pkg/controllers/userControllers.go`.
-   **Interactivité Frontend**: L'expérience utilisateur pourrait être améliorée en utilisant une bibliothèque JavaScript légère comme Alpine.js ou htmx pour ajouter de l'interactivité sans la complexité d'un framework complet.