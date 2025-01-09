package domain

type UserData struct {
	Id           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Password     string `json: "password"`
	Genero       string `json:"genero"`
	Atributos    string `json:"atributos"`
	Maneja       bool   `json:"maneja"`
	Lentes       bool   `json:"lentes"`
	Diabetico    bool   `json:"diabetico"`
	Enfermedades string `json:"enfermedades"`
	Admin        bool   `json:"admin"`
	Estado       bool   `json:"estado"`
}

type LoginData struct {
	Token  string `json: "token"`
	IdU    int    `json: "idu"`
	AdminU bool   `json:"adminu"`
}
