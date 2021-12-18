package password

type IPassword interface {
	Check(hashed, raw string) error
	Generate(string) (string, error)
}
