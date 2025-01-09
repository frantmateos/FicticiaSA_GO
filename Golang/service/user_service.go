package services

import (
	Domain "Golang/domain"
	Model "Golang/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type userClients interface {
	GetUserById(Id int) (Model.User, error)
	UpdateUser(ctx context.Context, User Model.User) (Model.User, error)
	InsertUser(user Model.User) (Model.User, error)
	GetUserByName(Usuario Model.User) (Model.User, error)
	GetAllUsers() ([]Model.User, error)
}

type Service struct {
	UserService userClients
}

func NewService(UserService userClients) Service {
	return Service{
		UserService: UserService,
	}
}

func (s Service) InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, error) {

	hash := md5.New()
	hash.Write([]byte(usuarioDomain.Password))
	usuarioDomain.Password = hex.EncodeToString(hash.Sum(nil))

	usuario := Model.User{
		Nombre:       usuarioDomain.Nombre,
		Password:     usuarioDomain.Password,
		Genero:       usuarioDomain.Genero,
		Atributos:    usuarioDomain.Atributos,
		Maneja:       usuarioDomain.Maneja,
		Lentes:       usuarioDomain.Lentes,
		Diabetico:    usuarioDomain.Diabetico,
		Enfermedades: usuarioDomain.Enfermedades,
		Admin:        usuarioDomain.Admin,
		Estado:       true,
	}

	//var result, er = s.UserService.GetUserByName(usuario)
	usuario2, err := s.UserService.InsertUser(usuario)

	if err != nil {
		return usuarioDomain, fmt.Errorf("Error Inserting User.")
	}

	usuarioDomain.Id = usuario2.Id

	return usuarioDomain, nil

}

func (s Service) GetUserByName(usuarioDomain Domain.UserData) (Domain.UserData, error) {

	usuario := Model.User{
		Nombre: usuarioDomain.Nombre,
	}

	user, err := s.UserService.GetUserByName(usuario)

	if err != nil {
		return Domain.UserData{}, fmt.Errorf("Error al buscar el usuario")
	}

	var userDomain Domain.UserData

	userDomain.Id = user.Id
	userDomain.Nombre = user.Nombre
	userDomain.Genero = user.Genero
	userDomain.Atributos = user.Atributos
	userDomain.Maneja = user.Maneja
	userDomain.Lentes = user.Lentes
	userDomain.Diabetico = user.Diabetico
	userDomain.Enfermedades = user.Enfermedades
	userDomain.Admin = user.Admin

	userDomain.Estado = user.Estado

	return userDomain, nil

}

func (s Service) GetUserById(userId int) (Domain.UserData, error) {
	user, err := s.UserService.GetUserById(userId)
	if err != nil {
		return Domain.UserData{}, fmt.Errorf("Error al obtener el usuario: %v", err)
	}

	userDomain := Domain.UserData{
		Id:           user.Id,
		Nombre:       user.Nombre,
		Genero:       user.Genero,
		Atributos:    user.Atributos,
		Maneja:       user.Maneja,
		Lentes:       user.Lentes,
		Diabetico:    user.Diabetico,
		Enfermedades: user.Enfermedades,
		Estado:       user.Estado,
	}

	return userDomain, nil
}

func (s Service) UpdateUser(usuarioDomain Domain.UserData) (Domain.UserData, error) {

	usuario := Model.User{
		Id:           usuarioDomain.Id,
		Nombre:       usuarioDomain.Nombre,
		Genero:       usuarioDomain.Genero,
		Atributos:    usuarioDomain.Atributos,
		Maneja:       usuarioDomain.Maneja,
		Lentes:       usuarioDomain.Lentes,
		Diabetico:    usuarioDomain.Diabetico,
		Enfermedades: usuarioDomain.Enfermedades,
		Admin:        usuarioDomain.Admin,
		Estado:       usuarioDomain.Estado,
	}
	ctx := context.Background()

	user, err := s.UserService.UpdateUser(ctx, usuario)

	if err != nil {
		return Domain.UserData{}, fmt.Errorf("Error al buscar el usuario")
	}

	var userDomain Domain.UserData

	userDomain.Id = user.Id
	userDomain.Nombre = user.Nombre
	userDomain.Genero = user.Genero
	userDomain.Atributos = user.Atributos
	userDomain.Maneja = user.Maneja
	userDomain.Lentes = user.Lentes
	userDomain.Diabetico = user.Diabetico
	userDomain.Enfermedades = user.Enfermedades
	userDomain.Admin = user.Admin
	userDomain.Estado = user.Estado

	return userDomain, nil

}

func (s Service) Login(User Domain.UserData) (Domain.LoginData, error) {
	usuario := Model.User{
		Nombre: User.Nombre,
		Admin:  User.Admin,
	}

	user, err := s.UserService.GetUserByName(usuario)
	fmt.Println("user ", user)

	var tokenDomain Domain.LoginData

	if err != nil {
		return tokenDomain, fmt.Errorf("error")
	}

	var Logpsw = md5.Sum([]byte(User.Password))
	psw := hex.EncodeToString(Logpsw[:])
	fmt.Println("contra ", psw)
	fmt.Println("contra ", user.Password)

	if psw == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"idU":    user.Id,
			"Adminu": user.Admin,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		})
		t, _ := token.SignedString([]byte("bitsion"))
		tokenDomain.Token = t
		tokenDomain.IdU = user.Id
		tokenDomain.AdminU = user.Admin
		return tokenDomain, nil
	} else {
		fmt.Println("eeror contra")
		return tokenDomain, fmt.Errorf("Contrasenia incorrecta")
	}

}

func (s Service) GetAllUsers() ([]Domain.UserData, error) {
	users, err := s.UserService.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("Error al obtener la lista de usuarios: %v", err)
	}

	var userDomainList []Domain.UserData
	for _, user := range users {
		userDomain := Domain.UserData{
			Id:           user.Id,
			Nombre:       user.Nombre,
			Genero:       user.Genero,
			Atributos:    user.Atributos,
			Maneja:       user.Maneja,
			Lentes:       user.Lentes,
			Diabetico:    user.Diabetico,
			Enfermedades: user.Enfermedades,
			Admin:        user.Admin,
			Estado:       user.Estado,
		}
		userDomainList = append(userDomainList, userDomain)
	}

	return userDomainList, nil
}
