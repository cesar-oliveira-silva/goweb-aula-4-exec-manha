package usuarios

type Usuario struct {
	Id          uint64 `json:"id"`
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"dataCriacao"`
}
