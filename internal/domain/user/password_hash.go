package user

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword Password, password string) error
}
