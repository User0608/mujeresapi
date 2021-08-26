package authorization

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/user0608/mujeresapi/models/authentication"
)

var (
	//ErrInvalidToken Error devuelto cuando el token no es valido
	ErrInvalidToken error = errors.New("token no válido")
	//ErrInvalidClams Error generado cuando los datos contenidos en el clams son incorrectos o están mal estructurados
	ErrInvalidClams error = errors.New("No se pudo obtener la data")
)

//GenerageToken Genera un toquen para un usuario o cliente, el cual se empleara para futuras peticiones
// func GenerageToken(customer model.Customer) (string, error) {
func GenerageToken(usuario authentication.Usuario) (string, error) {
	claim := Clain{
		UsuarioID: usuario.ID,
		UserName:  usuario.Username,
		Roles:     usuario.GetRoles(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidateToken verifica la firma de un token para comprobar su valides
func ValidateToken(t string) (Clain, error) {

	token, err := jwt.ParseWithClaims(t, &Clain{}, verifyFunction)
	if err != nil {
		return Clain{}, err
	}
	if !token.Valid {
		return Clain{}, ErrInvalidToken
	}
	claim, ok := token.Claims.(*Clain)
	if !ok {
		return Clain{}, ErrInvalidClams
	}
	return *claim, nil
}
func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
