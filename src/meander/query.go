package meander

type Place struct {
	*googleGeometry `json:"geometry"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Photos[]*googlePhoto `json:"photos"`
	Vicinity string `json:"vicinity"`
}

