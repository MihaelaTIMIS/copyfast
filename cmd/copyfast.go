package main

import (
	"fmt"

	"github.com/MihaelaTIMIS/copyfast/internal/fssync"
	"github.com/MihaelaTIMIS/copyfast/internal/zoo"
	"github.com/spf13/cobra"
)

func main() {

	//var name string // nil
	//var name2 string="MHAELA"
	//var name3 = "MIhaela"
	// name4 := "Mihaela"
	// age := 28

	// fmt.Println("Hello " + name4 + " " + lib.CovertIntToString(age))
	// fmt.Println(lib.Division(1, 5))

	// chien := zoo.MakeChien("Rex", "wuaf!")
	// fmt.Println(chien)
	// chien.Say()

	// chat := zoo.MakeChat("Kitty", "Miau..")
	// fmt.Println(chat)
	// chat.Say()

	// oiseau := zoo.MakeOiseau("Coco", "Cracra..")
	// fmt.Println(oiseau)
	// oiseau.Say()

	// adopte(chien)
	// adopte(chat)
	// adopte(oiseau)

	// a := "a"
	// b := Hello(&a)

	// fmt.Println(a, b)
	// var versionCmd = &cobra.Command{
	// 	Use:   "version",
	// 	Short: "Print the version number of Hugo",
	// 	Long:  `All software have versions. This is Hugo's`,
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("Hugo Static Site Generator v0.9 --HEAD")
	// 	},
	// }
	// versionCmd.AddCommand()

	var rootCmd = &cobra.Command{
		Use:   "copyfast",
		Short: "Sync folder to target",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			origin := args[0]
			target := args[1]
			ftp := args[2]
			// fssync.CopyRecursive(origin, target)
			fssync.ScanRecursive(origin, target, ftp)

			// tuto.MakeChannel()
		},
	}

	rootCmd.Execute()
}

func Hello(s *string) *string {
	*s = "b"
	// ns := "b"

	return s
}
func adopte(animal zoo.IsAnimal) {
	switch a := animal.(type) {
	case *zoo.Chat:
		fmt.Println("Super, je peux adopter", a.Name)
	case *zoo.Chien:
		fmt.Println("Super, je peux adopter", a.Name)
	case *zoo.Oiseau:
		fmt.Println("Super, je peux adopter", a.Name)
	}
}
