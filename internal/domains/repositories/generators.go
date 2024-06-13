package repositories

type IDGen interface {
	GenerateID() string
}
type HashGen interface {
	GenerateFromString(s string) string
}
