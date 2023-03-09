package handlers

import (
	"board/globals"
	"board/models"
	"net/http"
	"strings"

	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/types"
)

func getPosts(board string) ([]models.Post, error) {
	rows, err := globals.DB.Query("SELECT board, content, timestamp FROM posts WHERE board IS ?", board)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	err = rows.Iterate(func(d types.Document) error {
		var post models.Post
		err = document.StructScan(d, &post)
		if err != nil {
			return err
		}
		posts = append(posts, post)
		return nil
	})

	return posts, nil
}

func Board(w http.ResponseWriter, r *http.Request) {
	board := strings.Split(r.URL.Path, "/")[2]

	if board == "" {
		http.NotFound(w, r)
		return
	}

	posts, err := getPosts(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Layout(w, r, "board.html", models.Board{Title: board, Posts: posts})
}
