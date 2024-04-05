package model

type Protocol struct {
	Name            string   `json:"name"`
	Port            string   `json:"port"`
	OSI             string   `json:"osi"`
	ConnnectionType []string `json:"connectionType"`
	Speed           int      `json:"speed"`
	Security        string   `json:"security"`
	Payload         []string `json:"payload"`
	Headers         []string `json:"headers"`
	Standards       []string `json:"standards"`
	Description     string   `json:"description"`
	LatestVersion   float64  `json:"latestVersion"`
}
