package model

type Role int

const (
	MEMBER Role = iota
	ADMIN
)

func (r Role) string() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
