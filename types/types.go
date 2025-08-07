package types

type LoginPostRequest struct {
	Domain				string		`json:"domain"`
	DuplicateSession 	bool		`json:"duplicateSession"`
	Email				string		`json:"email"`
	Password			string		`json:"password"`
	Remember			bool		`json:"remember"`
}