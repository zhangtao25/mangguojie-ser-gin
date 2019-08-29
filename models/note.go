package models

type Note struct {
	Id			string		`json:"id"`
	Title     	string 		`json:"title"`
	Type     	string 		`json:"type"`
	Desc   		string 		`json:"desc"`
	Likes 		int 		`json:"likes"`
	Cover 		string 		`json:"cover"`
	UserId     	string 		`json:"user_id"`
	Collects    int 		`json:"collects"`
	Images		string		`json:"images"`
}

func GetNotes() interface{}{
	var note []Note
	res := db.Find(&note)
	return res.Value
}

func GetNote(id string) interface{}{
	note := Note{}
	res := db.Where(&Note{Id: id}).First(&note)
	return res.Value
}
//func AddNote()  interface{}{
//
//}