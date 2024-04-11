package usuarios

import "fmt"

var Usuarios []Usuario
var lastID uint64 = 0

type MemoryRepository struct {
}

// Update implements Repository.
// func (m *MemoryRepository) Update(Id uint, name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
// 	panic("unimplemented")
// }

func (m *MemoryRepository) GetAll() ([]Usuario, error) {
	return Usuarios, nil
}

func (m *MemoryRepository) GetId() (Usuario, error) {
	user := Usuario{}
	return user, nil
}

func (m *MemoryRepository) Store(nome string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	lastID++
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
	Usuarios = append(Usuarios, p)
	return p, nil
}

func (m *MemoryRepository) LastID() (uint64, error) {
	return lastID, nil
}

func (m *MemoryRepository) Update(id uint64, nome string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	usuario := Usuario{id, nome, sobrenome, email, idade, altura, ativo, datacriacao}
	fmt.Println("usuario: ", usuario)
	updated := false
	for i := range Usuarios {
		if Usuarios[i].Id == usuario.Id {
			Usuarios[i] = usuario
			updated = true

		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d nao encontrado", id)
	}
	return usuario, nil

}

func (m *MemoryRepository) UpdateName(id uint64, name string) (Usuario, error) {
	var user Usuario
	updated := false
	for i := range Usuarios {
		if Usuarios[i].Id == id {
			Usuarios[i].Nome = name
			updated = true
			user = Usuarios[i]
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d nao encontrado", id)
	}

	return user, nil
}

func (m *MemoryRepository) Delete(id uint64) error {
	deleted := false
	var index int
	for i := range Usuarios {
		if Usuarios[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d nao encontrado", id)
	}
	Usuarios = append(Usuarios[:index], Usuarios[index+1:]...)
	return nil
}
