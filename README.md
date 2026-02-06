# TheStreetPost
The main repository for the Street Post website.

## Go + Gin + HTMX + Tailwind + DuckDB

### Requirements
- Go 1.22+

### Run
1. Install Go dependencies:
	- `go mod tidy`
2. Start the server:
	- `go run .`

The app will be available at http://localhost:3000.

### Notes
- Pages use HTMX for boosted navigation and partial swaps.
- Tailwind is loaded via CDN.
- DuckDB is stored at data/thestreetpost.duckdb and managed through GORM.
