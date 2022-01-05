//Como pueden ver se puede usar el lenguaje Go como si tuviera reglas de C
package main;

import "fmt";
import "os";

func main() {
	programa := os.Args[0];

	fmt.Println("\nCantidad de Argumentos: ", len(os.Args));
	fmt.Println("Nombre del Programa: ", programa);
	
	if len(os.Args) == 1 {
		fmt.Println("\nArgumentos Extras: ",0);
	} else {
		fmt.Println("\nArgumentos Extras: ");
		for i := 1; i < len(os.Args); i++ {
			fmt.Printf("\tArgumento[%d]: %s\n", i, os.Args[i]);
		}
		fmt.Printf("\n");
	}
}
