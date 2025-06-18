package models



type SKUCategory struct {
	ID   string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}
