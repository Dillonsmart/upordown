package main

import "upordown/internal/storage"

func HandleTrack(args []string) {
	if len(args) < 1 {
		println("Usage: upordown track <url>")
		return
	}
	url := args[0]

	result, insertErr := storage.DB.Exec("INSERT INTO websites (url) VALUES (?)", url)
	if insertErr != nil {
		println("Error inserting website into database:", insertErr.Error())
		return
	}

	_, idErr := result.LastInsertId()
	if idErr != nil {
		println("Error getting last insert ID:", idErr.Error())
		return
	}

	println("Tracking website:", url)
}
