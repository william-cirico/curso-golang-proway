package interfaces

type Veiculo interface {
	Ligar() string
	Acelerar() (string, error)
	Desligar() (string, error)
}
