package usuarios

type Service interface {
	GetAll() ([]Usuario, error)
	Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
	Update(Id uint64, name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
	UpdateName(id uint64, name string) (Usuario, error)
	Delete(id uint64) error
}

type service struct {
	repository Repository
}

// Update implements Service.
func (s *service) Update(Id uint64, name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	return s.repository.Update(Id, name, sobrenome, email, idade, altura, ativo, datacriacao)
}

func (s *service) GetAll() ([]Usuario, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (s *service) Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {

	usuarios, err := s.repository.Store(name, sobrenome, email, idade, altura, ativo, datacriacao)
	if err != nil {
		return Usuario{}, err
	}

	return usuarios, nil
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) UpdateName(id uint64, name string) (Usuario, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id uint64) error {
	return s.repository.Delete(id)
}
