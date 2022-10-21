package main

import (
	"embed"
	"real-estate/cmd"
)

//go:embed migrations/estate/01_estate_schema.sql
var embedMigrations embed.FS

// Call the entry point
func main() {
	cmd.Start(embedMigrations)
}
