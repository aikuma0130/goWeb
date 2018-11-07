package meander

type j struct {
	Name       string
	PlaceTypes []string
}

// Journeys is datasets for choices of going out
var Journeys = []interface{}{
	&j{Name: "ロマンティック", PlaceTypes: []string{"park", "bar",
		"movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name: "ショッピング", PlaceTypes: []string{"department_store",
		"cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	&j{Name: "ナイトライフ", PlaceTypes: []string{"bar", "casino",
		"food", "bar", "night_club", "bar", "bar", "hospital"}},
	&j{Name: "カルチャー ", PlaceTypes: []string{"museum", "cafe",
		"cemetery", "library", "art_gallery"}},
	&j{Name: "リラックス", PlaceTypes: []string{"hair_care",
		"beauty_salon", "cafe", "spa"}},
}
