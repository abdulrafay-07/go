package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	X, Y, Z int
}

type Obj struct {
	Name   string
	Coords Coordinates
}

var Objects []Obj

func AddObj(name string, coords Coordinates) {
	newObj := Obj{
		Name: name,
		Coords: Coordinates{
			X: coords.X,
			Y: coords.Y,
			Z: coords.Z,
		},
	}

	Objects = append(Objects, newObj)

	fmt.Println("Object added successfully")
}

func PrintObjs() {
	for _, v := range Objects {
		fmt.Printf("%s is at the coordinates: \nX: %d, Y: %d, Z: %d\n", v.Name, v.Coords.X, v.Coords.Y, v.Coords.Z)
	}
}

func SearchObj(coords Coordinates) {
	for _, v := range Objects {
		if v.Coords.X == coords.X && v.Coords.Y == coords.Y && v.Coords.Z == coords.Z {
			fmt.Printf("Object found!\n%s is at:\nX: %d, Y: %d, Z: %d\n", v.Name, v.Coords.X, v.Coords.Y, v.Coords.Z)
			return
		}
	}

	fmt.Println("There is no such object with the entered coordinates")
}

func DelObj(coords Coordinates) {
	for i, v := range Objects {
		if v.Coords.X == coords.X && v.Coords.Y == coords.Y && v.Coords.Z == coords.Z {
			Objects = append(Objects[:i], Objects[i+1:]...)
			fmt.Println("Object deleted successfully")
			return
		}
	}

	fmt.Println("There is no such object with the entered coordinates")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var input int
		fmt.Println("Press\n[1] to add an object\n[2] to print all objects\n[3] to search an object by coordinates\n[4] to delete an object\n[5] to exit")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input", err)
		}

		switch input {
		case 1:
			var coords Coordinates

			fmt.Print("Enter the object's name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Trim newline

			fmt.Print("Enter the coordinates of the object (X Y Z): ")
			_, err := fmt.Scanf("%d %d %d\n", &coords.X, &coords.Y, &coords.Z)
			if err != nil {
				fmt.Println("Error reading coordinates")
				return
			}

			AddObj(name, coords)
		case 2:
			if len(Objects) == 0 {
				fmt.Println("There are no objects in the list")
				continue
			}

			PrintObjs()
		case 3:
			if len(Objects) == 0 {
				fmt.Println("There are no objects in the list")
				continue
			}

			var coords Coordinates
			fmt.Print("Enter the coordinates of the object (X Y Z): ")
			_, err := fmt.Scanf("%d %d %d\n", &coords.X, &coords.Y, &coords.Z)
			if err != nil {
				fmt.Println("Error reading coordinates")
				return
			}

			SearchObj(coords)
		case 4:
			if len(Objects) == 0 {
				fmt.Println("There are no objects in the list")
				continue
			}

			var coords Coordinates
			fmt.Print("Enter the coordinates of the object (X Y Z): ")
			_, err := fmt.Scanf("%d %d %d\n", &coords.X, &coords.Y, &coords.Z)
			if err != nil {
				fmt.Println("Error reading coordinates")
				return
			}

			DelObj(coords)
		default:
			os.Exit(0)
		}
	}
}
