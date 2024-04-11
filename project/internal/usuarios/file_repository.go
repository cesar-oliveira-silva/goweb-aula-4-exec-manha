package usuarios

import "github.com/cesar-oliveira-silva/goweb-aula-4-exec-manha.git/project/pkg/store"

type FileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) GetAll() ([]Usuario, error) {
	var ps []Usuario
	r.db.Read(&ps)
	return ps, nil
}

func (r *FileRepository) Store(nome string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	p := Usuario{
		Id:          lastID,
		Nome:        nome,
		Sobrenome:   sobrenome,
		Email:       email,
		Idade:       idade,
		Altura:      altura,
		Ativo:       ativo,
		DataCriacao: datacriacao,
	}

	var ps []Usuario

	// primeiro lemos o arquivo
	r.db.Read(&ps)

	// calculamos qual o próximo ID
	lastIdInserted := len(ps)
	lastIdInserted++
	p.Id = uint64(lastIdInserted)

	// inserimos o produto a ser cadastrado no slice de produtos
	ps = append(ps, p)

	// gravamos no arquivo novamente com o novo produto inserido
	err := r.db.Write(ps)
	if err != nil {
		return Usuario{}, err
	}
	return p, nil
}

func (r *FileRepository) Delete(id uint64) error {
	return nil
}

func (r *FileRepository) Update(id uint64, nome string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	return Usuario{}, nil
}
func (r *FileRepository) UpdateName(id uint64, name string) (Usuario, error) {
	return Usuario{}, nil
}

func (r *FileRepository) LastID() (uint64, error) {
	var ps []Usuario
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil

}
