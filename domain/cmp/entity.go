package cmp

type Entity interface {
	ID() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}
