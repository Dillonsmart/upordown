package main

func HandleTrack(args []string) {
	if len(args) < 1 {
		println("Usage: upordown track <url>")
		return
	}
	url := args[0]

	println("Tracking URL:", url)
}
