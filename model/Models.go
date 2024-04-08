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

type Device struct {
	Name  string  `json:"name"`
	Event []Event `json:"event"`
}

type Event struct {
	Name            string   `json:"name"`
	AverageMbps     int      `json:"averageMbps"`
	ConnnectionType string   `json:"connectionType"`
	Protocol        Protocol `json:"protocol"`
	QoSPriority     string   `json:"qosPriority"`
}
