package sqlpassctxcheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "sqlpassctxcheck",
	Doc:  "check for sql module method call without ctx",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		sqlModuleAlias := getSqlModuleAlias(file)

		ast.Inspect(file, func(n ast.Node) bool {
			// ここから
			if call, ok := n.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					if funx, ok := fun.X.(*ast.Ident); ok {
						if funx.Obj != nil {
							if funxobjdecl, ok := funx.Obj.Decl.(*ast.Field); ok {
								if funxobjdecltype, ok := funxobjdecl.Type.(*ast.StarExpr); ok {
									if funxobjdecltypex, ok := funxobjdecltype.X.(*ast.SelectorExpr); ok {
										if funxobjdecltypexx, ok := funxobjdecltypex.X.(*ast.Ident); ok {
											var sqlModuleOfReceiver sqlPackage
											if funxobjdecltypexx.Name == sqlModuleAlias[sql] {
												sqlModuleOfReceiver = sql
											} else if funxobjdecltypexx.Name == sqlModuleAlias[sqlx] {
												sqlModuleOfReceiver = sqlx
											}

											if restricted, ok := sqlModuleRestrictedMethodMap[sqlModuleOfReceiver].get(funxobjdecltypex.Sel.Name, fun.Sel.Name); ok {
												pass.Report(analysis.Diagnostic{
													Pos: call.Pos(),
													End: call.End(),
													Message: fmt.Sprintf(
														"use (*%s.%s).%s instead of (*%s.%s).%s",
														sqlModuleAlias[sqlModuleOfReceiver],
														restricted.ReceiverType,
														restricted.AlternateMethodName,
														sqlModuleAlias[sqlModuleOfReceiver],
														restricted.ReceiverType,
														restricted.MethodName,
													),
												})
											}
										}
									}
								}
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

type sqlModuleAliasMap map[sqlPackage]string

func getSqlModuleAlias(file *ast.File) sqlModuleAliasMap {
	sqlModuleAlias := sqlModuleAliasMap{
		sql:  "sql",
		sqlx: "sqlx",
	}

	for _, decl := range file.Decls {
		if gendecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gendecl.Specs {
				if importSpec, ok := spec.(*ast.ImportSpec); ok {
					for _, sqlPackage := range allSqlPackage {
						if importSpec.Name != nil &&
							importSpec.Path.Value == sqlModuleImportPathMap[sqlPackage] {
							sqlModuleAlias[sqlPackage] = importSpec.Name.Name
						}
					}
				}
			}
		}
	}
	return sqlModuleAlias
}
