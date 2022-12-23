package models

type Morador struct {
	ID            int64  `json:"id_morador"`
	Nome          string `json:"morador"`
	CPF           string `json:"cpf"`
	EmCasa        bool   `json:"morador em casa"`
	IDApartamento int64  `json:"apartamento"`
}

type Visitante struct {
	ID            int64  `json:"id_visitante"`
	Nome          string `json:"visitante"`
	CPF           string `json:"cpf"`
	IDApartamento int64  `json:"apartamento"`
}

type Apartamento struct {
	Numero   int64 `json:"apartamento"`
	IDPredio int64 `json:"predio"`
}

type Predio struct {
	Numero  int64  `json:"predio"`
	IDBloco string `json:"bloco"`
}

type Bloco struct {
	Letra string `json:"bloco"`
}
