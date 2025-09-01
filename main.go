package goarea

import "math"

// P1 é uma proporção númerica definida pela relação entre
// o perímetro de uma circunfêrencia e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi

}

// Rect é responsável por cacular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
