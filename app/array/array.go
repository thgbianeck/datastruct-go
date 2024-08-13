package array

// Array representa uma estrutura de dados de array simples.
type Array struct {
	data []int
}

// New cria uma nova instância de Array.
func New() *Array {
	return &Array{data: []int{}}
}

// Add adiciona um valor ao Array.
func (a *Array) Add(value int) {
	a.data = append(a.data, value)
}

// GetAll retorna todos os valores armazenados no Array.
func (a *Array) GetAll() []int {
	return a.data
}

// Find verifica se um valor está presente no Array.
func (a *Array) Find(value int) bool {
	for _, v := range a.data {
		if v == value {
			return true
		}
	}
	return false
}
