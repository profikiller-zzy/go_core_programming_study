package es_hotel

// Hotel 酒店结构体，对应MySQL表和ES文档
type Hotel struct {
	ID        int64   `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Address   string  `json:"address" db:"address"`
	Price     float64 `json:"price" db:"price"`
	Score     float64 `json:"score" db:"score"`
	Brand     string  `json:"brand" db:"brand"`
	City      string  `json:"city" db:"city"`
	StarName  string  `json:"star_name" db:"star_name"`
	Business  string  `json:"business" db:"business"`
	Latitude  string  `json:"latitude" db:"latitude"`
	Longitude string  `json:"longitude" db:"longitude"`
	Pic       string  `json:"pic" db:"pic"`
}

func (h Hotel) TableName() string {
	return "tb_hotel" // MySQL表名
}
