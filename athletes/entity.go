package athlete

type Athlete struct {
	ID       int    `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
}
