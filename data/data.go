package data

import "fmt"

var Movies = generateMovies()

func generateMovies() []string {
	movies := []string{}

	for i := 1; i <= 10000; i++ {
		movies = append(movies, "Film Ke-"+fmt.Sprint(i))
	}

	return movies
}
