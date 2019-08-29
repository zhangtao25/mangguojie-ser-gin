package _type

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