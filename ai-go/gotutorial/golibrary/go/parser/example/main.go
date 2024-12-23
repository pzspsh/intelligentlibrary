/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:53:31
*/
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/gohade/hade/framework/cobra"
	"github.com/pkg/errors"
)

var genClassFile string

func InitMermaidCmd() *cobra.Command {
	GenClassCommand.Flags().StringVarP(&genClassFile, "file", "f", "", "file path")
	//_ = GenClassCommand.MarkFlagRequired("file")
	MermaidCommand.AddCommand(GenClassCommand)
	return MermaidCommand
}

var MermaidCommand = &cobra.Command{
	Use:   "mermaid",
	Short: "mermaid",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

// GenClassCommand generate a file's class diagrames
var GenClassCommand = &cobra.Command{
	Use:   "genClass",
	Short: "genClass",
	RunE: func(c *cobra.Command, args []string) error {
		// check file exist
		if _, err := os.Stat(genClassFile); os.IsNotExist(err) {
			return errors.Wrap(err, "file not exist")
		}

		// parse file
		fset := token.NewFileSet()
		asf, err := parser.ParseFile(fset, genClassFile, nil, parser.ParseComments)
		if err != nil {
			return errors.Wrap(err, "parse file error")
		}
		//ast.Print(fset, asf)
		//return nil

		fileContent, _ := os.ReadFile(genClassFile)

		// print class diagrames
		tab := "\t"
		fmt.Println("classDiagram")

		// class map
		classFieldMap := map[string][]string{}
		classFuncMap := map[string][]string{}
		interfaceFuncMap := map[string][]string{}

		for _, decl := range asf.Decls {
			switch decl.(type) {
			case *ast.GenDecl:
				genDecl := decl.(*ast.GenDecl)
				switch genDecl.Tok {
				case token.TYPE:
					for _, spec := range genDecl.Specs {
						typeSpec := spec.(*ast.TypeSpec)

						switch typeSpec.Type.(type) {
						case *ast.StructType:
							className := typeSpec.Name.Name
							if classFieldMap[className] == nil {
								classFieldMap[className] = []string{}
							}

							structType := typeSpec.Type.(*ast.StructType)
							for _, field := range structType.Fields.List {
								startPos := field.Type.Pos()
								endPos := field.Type.End()
								fieldType := string(fileContent[startPos-1 : endPos-1])

								for _, name := range field.Names {
									classFieldMap[className] = append(classFieldMap[className],
										name.Name+" "+fieldType)
								}
							}
						case *ast.InterfaceType:
							interfaceName := typeSpec.Name.Name
							if interfaceFuncMap[interfaceName] == nil {
								interfaceFuncMap[interfaceName] = []string{}
							}

							interfaceType := typeSpec.Type.(*ast.InterfaceType)
							for _, method := range interfaceType.Methods.List {
								startPos := method.Type.Pos()
								endPos := method.Type.End()
								methodType := string(fileContent[startPos-1 : endPos-1])

								if method.Names == nil {
									if ident, ok := method.Type.(*ast.Ident); ok {
										//spew.Dump(ident.Obj)
										if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
											if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
												for _, method := range interfaceType.Methods.List {
													startPos := method.Type.Pos()
													endPos := method.Type.End()
													methodType := string(fileContent[startPos-1 : endPos-1])
													interfaceFuncMap[interfaceName] = append(interfaceFuncMap[interfaceName],
														methodType)
												}
											}
										}
									}
								}
								for _, name := range method.Names {
									interfaceFuncMap[interfaceName] = append(interfaceFuncMap[interfaceName],
										name.Name+" "+methodType)
								}
							}
						}
					}
				}
			case *ast.FuncDecl:
				funcDecl := decl.(*ast.FuncDecl)
				if funcDecl.Recv != nil {
					// method
					recv := funcDecl.Recv.List[0]
					switch recv.Type.(type) {
					case *ast.StarExpr:
						starExpr := recv.Type.(*ast.StarExpr)
						className := string(fileContent[starExpr.X.Pos()-1 : starExpr.X.End()-1])
						if classFuncMap[className] == nil {
							classFuncMap[className] = []string{}
						}
						classFuncMap[className] = append(classFuncMap[className], funcDecl.Name.Name)
					case *ast.StructType:
						structType := recv.Type.(*ast.StructType)
						structType.Fields.Pos()
						structType.Fields.End()
						className := string(fileContent[structType.Fields.Pos()-1 : structType.Fields.End()-1])
						if classFuncMap[className] == nil {
							classFuncMap[className] = []string{}
						}
						classFuncMap[className] = append(classFuncMap[className], funcDecl.Name.Name)
					}
				}
			}
		}
		//fmt.Println(tab)

		for className, fields := range classFieldMap {
			fmt.Println(tab + "class " + className)
			for _, field := range fields {
				fmt.Println(tab + className + ": " + field)
			}

			if classFuncMap[className] != nil {
				for _, funcName := range classFuncMap[className] {
					fmt.Println(tab + className + ": " + funcName + "()")
				}
			}
			fmt.Println()
		}

		for interfaceName, funcs := range interfaceFuncMap {
			fmt.Println(tab + "class " + interfaceName)
			fmt.Println(tab + "<<interface>> " + interfaceName)

			for _, funcName := range funcs {
				fmt.Println(tab + interfaceName + ": " + funcName)
			}
			fmt.Println()
		}

		return nil
	},
}
