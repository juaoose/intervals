package athlete

type Athlete struct {
	ID       string `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
}
