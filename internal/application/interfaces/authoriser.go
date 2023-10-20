package interfaces

type Authoriser interface {
	Authorised(token string) bool
}
