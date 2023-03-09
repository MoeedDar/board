package handlers

import (
	"board/globals"
	"net/http"

	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/types"
)

func getBoards() ([]string, error) {
	rows, err := globals.DB.Query("SELECT title FROM boards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boards []string
	err = rows.Iterate(func(d types.Document) error {
		var title string
		err = document.Scan(d, &title)
		if err != nil {
			return err
		}
		boards = append(boards, title)
		return nil
	})

	return boards, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	boards, err := getBoards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Layout(w, r, "index.html", boards)
}
