// Zaradb lite faset document database
package main

import (
	"embed"
	"fmt"
	"net/http"
	"zaradb/store"
)

//go:embed  static
var content embed.FS

func main() {
	// TODO close programe greatfully.

	db := store.NewDB("test.db")
	db.CreateCollection("test")
	defer db.Close()

	fmt.Printf("zaradb run on %s:%s\n", Host, Port)

	http.Handle("/", http.FileServer(http.FS(content)))

	http.HandleFunc("/shell", shell)

	// standard endpoint
	http.HandleFunc("/ws", Ws)

	// endpoints for speed network
	http.HandleFunc("/query", Resever)
	http.HandleFunc("/result", Sender)

	http.ListenAndServe(":1111", nil)
}

// render static shell.html file
func shell(w http.ResponseWriter, r *http.Request) {
	// Open the index.html file
	f, err := content.ReadFile("static/shell.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the index.html file to the response
	fmt.Fprint(w, string(f))
}
