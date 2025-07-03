package main

import (
	"upordown/internal/database"
	"upordown/pkg"
)

func HandleTrack(args []string) {
	if len(args) < 1 {
		println("Usage: upordown track <url>")
		return
	}

	domain, parseErr := pkg.GetHost(args[0])

	if parseErr != nil {
		println("Error parsing URL:", parseErr.Error())
		return
	}

	result, insertErr := database.DB.Exec("INSERT INTO websites (domain) VALUES (?)", domain)
	if insertErr != nil {
		println("Error inserting website into database:", insertErr.Error())
		return
	}

	_, idErr := result.LastInsertId()
	if idErr != nil {
		println("Error getting last insert ID:", idErr.Error())
		return
	}

	println("Tracking website:", domain)
}
