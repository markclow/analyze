package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go/ast"
	"strings"
)

var cmdGetFunctionParameters = &cobra.Command{
	Use:   "getFunctionParameters",
	Short: "Retrieves parameters for function name",
	Long:  `Retrieves comma-separated list of parameters for function name.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("fileName")
		functionName, _ := cmd.Flags().GetString("functionName")
		getFunctionParameters(fileName, functionName)
	},
}

var cmdGetFunctionResults = &cobra.Command{
	Use:   "getFunctionResults",
	Short: "Retrieves returned result(s) for function name",
	Long:  `Retrieves returned result(s) for function name.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("fileName")
		functionName, _ := cmd.Flags().GetString("functionName")
		getFunctionResults(fileName, functionName)
	},
}

var cmdGetFunctionInvokedBy = &cobra.Command{
	Use:   "getFunctionInvokedBy",
	Short: "Retrieves returned result(s) for function name",
	Long:  `Retrieves returned result(s) for function name.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("fileName")
		functionName, _ := cmd.Flags().GetString("functionName")
		getFunctionInvokedBy(fileName, functionName)
	},
}

func init() {
	cmdGetFunctionParameters.Flags().String("functionName", "", "Name of function.") // flag just for isImported
	cmdGetFunctionParameters.MarkFlagRequired("functionName")
	rootCmd.AddCommand(cmdGetFunctionParameters)

	cmdGetFunctionResults.Flags().String("functionName", "", "Name of function.") // flag just for isImported
	cmdGetFunctionResults.MarkFlagRequired("functionName")
	rootCmd.AddCommand(cmdGetFunctionResults)

	cmdGetFunctionInvokedBy.Flags().String("functionName", "", "Name of function.") // flag just for isImported
	cmdGetFunctionInvokedBy.MarkFlagRequired("functionName")
	rootCmd.AddCommand(cmdGetFunctionInvokedBy)
}

func getFunctionParameters(fileName string, functionName string) {
	root, err := parseFile(fileName)
	if err != nil {
		fmt.Printf("error: cannot parse file '%s'", fileName)
		return
	}
	// calculate slice of imported package
	found := false
	sb := strings.Builder{}
	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			fd := n.(*ast.FuncDecl)
			if fd.Name.Name == functionName {
				found = true
				if fd.Type != nil && fd.Type.Params != nil && fd.Type.Params.List != nil {
					for _, p := range fd.Type.Params.List {
						if sb.Len() > 0 {
							sb.WriteString(",")
						}
						tp := fmt.Sprintf("%v", p.Type)
						sb.WriteString(tp)
					}
				}
			}
		}
		return true
	})
	if found {
		if sb.Len() == 0 {
			fmt.Println("No parameters")
		} else {
			fmt.Printf("%v\n", sb.String())
		}
	} else {
		fmt.Println("Function not found")
	}
}

func getFunctionResults(fileName string, functionName string) {
	root, err := parseFile(fileName)
	if err != nil {
		fmt.Printf("error: cannot parse file '%s'", fileName)
		return
	}
	// calculate slice of imported package
	found := false
	sb := strings.Builder{}
	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			fd := n.(*ast.FuncDecl)
			if fd.Name.Name == functionName {
				found = true
				if fd.Type != nil && fd.Type.Results != nil && fd.Type.Results.List != nil {
					for _, p := range fd.Type.Results.List {
						if sb.Len() > 0 {
							sb.WriteString(",")
						}
						tp := fmt.Sprintf("%v", p.Type)
						sb.WriteString(tp)
					}
				}
			}
		}
		return true
	})
	if found {
		if sb.Len() == 0 {
			fmt.Println("No parameters")
		} else {
			fmt.Printf("%v\n", sb.String())
		}
	} else {
		fmt.Println("Function not found")
	}
}

func getFunctionInvokedBy(fileName string, functionName string) {
	root, err := parseFile(fileName)
	if err != nil {
		fmt.Printf("error: cannot parse file '%s'", fileName)
		return
	}
	var invokers []string
	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			fd := n.(*ast.FuncDecl)
			// look at elements inside function
			ast.Inspect(fd, func(n ast.Node) bool {
				switch n.(type) {
				case *ast.CallExpr:
					ce := n.(*ast.CallExpr)
					if ce.Fun != nil {
						ident, ok := ce.Fun.(*ast.Ident)
						if ok {
							if ident.Name == functionName {
								addInvoker(fd.Name.Name, &invokers)
							}
						}
					}
				}
				return true
			})
		}
		return true
	})
	sb := strings.Builder{}
	for _, invoker := range invokers {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(invoker)
	}
	fmt.Printf("%v\n", sb.String())
}

func addInvoker(invoker string, sl *[]string) {
	found := false
	for _, item := range *sl {
		if item == invoker {
			found = true
			break
		}
	}
	if !found {
		*sl = append(*sl, invoker)
	}
}
