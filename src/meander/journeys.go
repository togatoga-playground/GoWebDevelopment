package meander

type j struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []interface{}{
	&j{Name: "ロマンチック", PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxI_stand"}},
	&j{Name: "ショッピング", PlaceTypes: []string{"department_stand", "cafe", "clothing_store", "jewlry_store", "shoe_store"}},
	&j{Name: "ナイトライフ", PlaceTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hospital"}},
	&j{Name: "カルチャー", PlaceTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	&j{Name: "リラックス", PlaceTypes: []string{"hair_care", "beaity_salon", "cafe", "spa"}},
}
