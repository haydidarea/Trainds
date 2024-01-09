package models

type Product struct {
	id 			int64  'gorm:"primarykey" json:"id"'
	NamaProduct string 'gorm:"type:varchar(300)" json:"NamaProduct"'
	Deskripsi 	string 'gorm:"type:text" json:"Deskripsi"'
}