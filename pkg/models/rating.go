package models

type Rating struct {
	Thing
	Author      *Person `json:"author"`
	BestRating  int32   `json:"bestRating"`
	WorstRating int32   `json:"worstRating"`
	RatingValue int32   `json:"ratingValue"`
}
