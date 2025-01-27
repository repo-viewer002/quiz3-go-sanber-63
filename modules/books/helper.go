package books

import "fmt"

func DefineBookThickness(totalPage int) string {
	if totalPage < 100 {
		return "tipis"
	} else {
		return "tebal"
	}
}

func ValidateReleaseYear(releaseYear int) error {
	if releaseYear < 1980 || releaseYear > 2024 {
		return fmt.Errorf("invalid release year, please input release year start from 1980 - 2024, your input : \"%d\"", releaseYear)
	}

	return nil
}
