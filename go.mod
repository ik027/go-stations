module github.com/TechBowl-japan/go-stations

go 1.16

replace ik027.train/go-stations/handler => ./handler

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.7
	ik027.train/go-stations/handler v1.0.0
)
