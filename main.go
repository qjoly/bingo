package main

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type BingoCard struct {
	Title  string
	Grid   [5][5]string
	Phrase string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate", generateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	phrases, err := loadPhrases("phrases.txt")
	if err != nil {
		http.Error(w, "Impossible de charger les phrases", http.StatusInternalServerError)
		return
	}

	if len(phrases) < 25 {
		http.Error(w, "Il faut au moins 25 phrases pour générer une grille de bingo", http.StatusBadRequest)
		return
	}

	shuffled := make([]string, len(phrases))
	perm := rand.Perm(len(phrases))
	for i, v := range perm {
		shuffled[v] = phrases[i]
	}

	var grid [5][5]string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			grid[i][j] = shuffled[i*5+j]
		}
	}

	grid[2][2] = "Case gratuite"

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	card := BingoCard{
		Title:  "Mon Bingo Personnalisé",
		Grid:   grid,
		Phrase: "Cliquez sur les cases pour les marquer",
	}

	err = tmpl.Execute(w, card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadPhrases(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var phrases []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return phrases, nil
}
