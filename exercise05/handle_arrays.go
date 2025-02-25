package exercise05

import "fmt"

func Examples_arrays() {

	// Definir un slice de letras
	letras := []string{"A", "B", "C", "D", "E"}

	// Variable para almacenar la posición de la letra "E"
	posicion := -1

	// Recorrer el slice para buscar la letra "E"
	for i, letra := range letras {
		if letra == "E" {
			posicion = i
			break // Salir del bucle una vez encontrada la letra
		}
	}

	// Verificar si se encontró la letra "E"
	if posicion != -1 {
		fmt.Printf("La letra 'E' se encuentra en la posición %d del slice.\n", posicion)
	} else {
		fmt.Println("La letra 'E' no se encuentra en el slice.")
	}

	//Arreglo multidimensional
	// Definir un arreglo multidimensional
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Recorrer la matriz
	for i, fila := range matriz {
		for j, valor := range fila {
			fmt.Printf("matriz[%d][%d] = %d\n", i, j, valor)
		}
	}

}
