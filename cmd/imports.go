package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go/ast"
)

var cmdIsImported = &cobra.Command{
	Use:   "isImported",
	Short: "Was the package imported in the file?",
	Long:  `Was the specified package imported in the specified file?`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fn, _ := cmd.Flags().GetString("fileName")
		p, _ := cmd.Flags().GetString("package")
		isImported(fn, p)
	},
}

func init() {
	cmdIsImported.Flags().String("package", "", "Name of package.") // flag just for isImported
	cmdIsImported.MarkFlagRequired("package")
	rootCmd.AddCommand(cmdIsImported)
}

func isImported(fileName string, packageName string) {
	root, err := parseFile(fileName)
	if err != nil {
		fmt.Printf("error: cannot parse file '%s'", fileName)
		return
	}
	// calculate slice of imported packages
	importedPackages := make([]string, 0)
	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.GenDecl:
			g := n.(*ast.GenDecl)
			for _, s := range g.Specs {
				switch s.(type) {
				case *ast.ImportSpec:
					inportSpec := s.(*ast.ImportSpec)
					importedPackages = append(importedPackages, inportSpec.Path.Value)
				}
			}
		}
		return true
	})
	// check slice to see if package was imported
	imported := false
	packageNameInQuotes := fmt.Sprintf(`"%s"`, packageName)
	for _, ip := range importedPackages {
		if ip == packageNameInQuotes {
			imported = true
		}
	}
	fmt.Printf("%t\n", imported)
}
