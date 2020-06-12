

package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("arreglo vacio:", a)

	a[4] = 100
	fmt.Println("set en la ultima posicion:", a)
	fmt.Println("obtener la ultima posision:", a[4])

	fmt.Println("largo del arreglo:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arreglo inicializado:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var threeD [2][3][3]int
	for k:=0; k < 3; k++ {
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d: ", threeD)


}
