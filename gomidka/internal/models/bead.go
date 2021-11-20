package models

type(
	Bread struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`

		Gluten          bool   `json:"gluten" db:"gluten"`          // yes
		Price           int    `json:"price" db:"price"`           //2$
		Culture         string `json:"culture" db:"culture"`         //Rus or French ...
		CategoryId      string `json:"category_id" db:"category_id"`            // Baget, Kerpich ...
		Filling         string `json:"filling" db:"filling"`         // izum ...
	}
	BreadsFilter struct {
		Query *string `json:"query"`
	}
)

