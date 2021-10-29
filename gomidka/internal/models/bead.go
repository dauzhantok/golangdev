package models

type Bread struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Gluten          bool   `json:"Gluten"`          // yes
	Sugar           int    `json:"Sugar"`           // 50
	Expiration_date int    `json:"Expiration_date"` //3 days
	Culture         string `json:"Culture"`         //Russ or French ...
	Type            string `json:"Type"`            // Baget, Kerpich ...
	Filling         string `json:"Filling"`         // Chocolate , froot
}
