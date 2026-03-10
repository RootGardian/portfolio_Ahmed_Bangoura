package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 1. Définir le répertoire de travail (racine du projet)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Dossier racine du projet :", dir)

	// 2. Servir les fichiers statiques
	// Dossier "scripts" (JS, CSS) -> accessible via /scripts/
	fsScripts := http.FileServer(http.Dir(filepath.Join(dir, "scripts")))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fsScripts))

	// Dossier "templates" (HTML fragments) -> accessible via /templates/
	fsTemplates := http.FileServer(http.Dir(filepath.Join(dir, "templates")))
	http.Handle("/templates/", http.StripPrefix("/templates/", fsTemplates))

	// Dossier "styles" (PDF, images) -> accessible via /styles/
	fsStyles := http.FileServer(http.Dir(filepath.Join(dir, "styles")))
	http.Handle("/styles/", http.StripPrefix("/styles/", fsStyles))

	// 3. Route principale "/"
	// Sert le fichier "templates/home.html" qui est la coquille de l'application (SPA)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Pour la racine, on envoie home.html
		if r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join(dir, "templates", "home.html"))
			return
		}
		// Pour les autres URL non gérées, on renvoie 404
		http.NotFound(w, r)
	})

	// 4. Démarrer le serveur
	port := "8080"
	log.Printf("Serveur démarré ! Ouvrez http://localhost:%s dans votre navigateur", port)

	// Lancement du serveur
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur : ", err)
	}
}
