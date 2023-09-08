package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

type Coordenada struct {
	Lat  float64
	Long float64
}

// Formula para convertir grados centígrados a Farentheid
func centToFaren(cent float64) float64 {
	return (cent * 9 / 5) + 32
}

func esTrianguloRectangulo(d1, d2, d3 int) bool {
	// Coloca las longitudes en un slice y ordénalas de menor a mayor.
	lado := []int{d1, d2, d3}
	sort.Ints(lado)

	// Comprueba si es un triángulo válido.
	if lado[0] <= 0 {
		return false
	}
	// Aplica el teorema de Pitágoras.
	return lado[0]*lado[0]+lado[1]*lado[1] == lado[2]*lado[2]
}

func valorCarta(carta string) int {
	switch carta {
	case "As":
		return 1
	case "Dos":
		return 2
	case "Tres":
		return 3
	case "Cuatro":
		return 4
	case "Cinco":
		return 5
	case "Seis":
		return 6
	case "Siete":
		return 7
	case "Ocho":
		return 8
	case "Nueve":
		return 9
	case "Diez", "Jota", "Reina", "Rey":
		return 10
	default:
		return 0 // Carta no válida
	}
}

func eliminarCaracteresLargos(input string) string {
	var resultado string

	for _, runeValue := range input {
		if utf8.RuneLen(runeValue) <= 1 {
			resultado += string(runeValue)
		}
	}

	return resultado
}

func main() {
	//Problema 1: Dada una variable que registra la temperatura en grados centígrados,
	//calcula el valor de la misma temperatura pero en grados Farenheit
	cent := 37.3

	//faren := (cent * 9 / 5) + 32
	faren := centToFaren(cent)
	fmt.Printf("Solución al Problema 1\n")
	fmt.Printf("%v grados centígrados son %v grados Farenheit\n", cent, faren)
	fmt.Printf("\n")

	// Problema 2: Determinar si tres longitudes enteras dadas en las variable
	// d1, d2 y d3 forman un triángulo rectángulo, o no
	d1, d2, d3 := 4, 3, 5
	fmt.Printf("Solución al Problema 2\n")
	if esTrianguloRectangulo(d1, d2, d3) {
		fmt.Printf("Las longitudes: %v, %v y %v forman un triángulo rectángulo.\n", d1, d2, d3)
		fmt.Printf("\n")
	} else {
		fmt.Printf("Las longitudes: %v, %v y %v no forman un triángulo rectángulo.\n", d1, d2, d3)
		fmt.Printf("\n")
	}

	// Problema 3: Dados dos valores de longitud y latitud, determinar si está al este, oeste, sur, norte,
	// sureste, suroeste, noreste, noroeste o en el centro de la tierra

	// Coordenadas proporcionadas
	lat, long := -34.172, -17.039

	// Inicializa la variable loc con un valor vacío
	var loc string

	// Determina la dirección cardinal en función de las coordenadas
	if lat == 0 && long == 0 {
		loc = "En el centro de la Tierra"
	} else {
		if lat > 0 {
			loc += "Norte"
		} else if lat < 0 {
			loc += "Sur"
		}

		if long > 0 {
			if loc != "" {
				loc += "-"
			}
			loc += "Este"
		} else if long < 0 {
			if loc != "" {
				loc += "-"
			}
			loc += "Oeste"
		}
	}

	// Imprime la dirección en la que se encuentra el punto
	fmt.Printf("Solución al Problema 3\n")
	fmt.Printf("La dirección está al:\n")
	fmt.Println(loc)
	fmt.Printf("\n")

	//Problema 4: En el juego del blackjack los valores que toman las cartas son los siguientes:
	// de la carta dos(2) hasta la carta diez(10), el valor de la carta es su propio número;
	//las figuras valen diez(10); los Ases pueden valer uno(1) u once(11).
	//Dadas dos cartas del blackjack que debes representar con su nombre, determinar el valor en puntos de la mano.
	// Si el uso del As con valor 11 mejora la puntuación (sin exceder 21) entonces usar preferibleemente esa opción.
	carta1 := "Dos"
	carta2 := "As"

	// Calcula el valor de las cartas
	valor1 := valorCarta(carta1)
	valor2 := valorCarta(carta2)

	// Calcula la puntuación inicial
	suma := valor1 + valor2

	// Verifica si un As como 11 mejora la puntuación sin exceder 21
	if carta1 == "As" || carta2 == "As" && suma <= 21 {
		suma += 10 // Suma 10 para usar el As como 11
	}

	fmt.Printf("Solución al Problema 4\n")
	fmt.Printf("El valor final es %v\n", suma)
	fmt.Printf("\n")

	//Problema 5: Calcula la suma de todos los números desde 1 hasta 1,000,000 que sean múltiplos de 3, o de 5

	limite := 1000000
	sum := 0

	for i := 1; i <= limite; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Solución al Problema 5\n")
	fmt.Printf("La suma es %v\n", sum)
	fmt.Printf("\n")

	//Problema 6: Dada una cadena de caracteres que puede contener códigos UTF-8, devuelve otra que elimine todos
	//aquellos cuya codificación consuma más de un byte

	saludo := "Hola 🗺 Que seas muy 🙋 para siempre"
	fmt.Println(saludo)

	fmt.Printf("Solución al Problema 6\n")
	resultado := eliminarCaracteresLargos(saludo)
	fmt.Println(resultado)

}
