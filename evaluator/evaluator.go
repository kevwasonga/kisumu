package evaluator

import (
	"kisumu/ast"
	"kisumu/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.PackageStatement:
		return evalPackageStatement(node, env)
	case *ast.CoreStatement:
		return evalCoreStatement(node, env)
	case *ast.DeclareStatement:
		return evalDeclareStatement(node, env)
	case *ast.DisplayStatement:
		return evalDisplayStatement(node, env)
	default:
		return nil
	}
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, stmt := range program.Statements {
		result = Eval(stmt, env)
	}
	return result
}

func evalPackageStatement(stmt *ast.PackageStatement, env *object.Environment) object.Object {
	// No direct evaluation for PackageStatement, it's a container
	return nil
}

func evalCoreStatement(stmt *ast.CoreStatement, env *object.Environment) object.Object {
	for _, s := range stmt.Statements {
		Eval(s, env)
	}
	return nil
}

func evalDeclareStatement(stmt *ast.DeclareStatement, env *object.Environment) object.Object {
	value := Eval(stmt.Value, env)
	env.Set(stmt.Name.Value, value)
	return nil
}

func evalDisplayStatement(stmt *ast.DisplayStatement, env *object.Environment) object.Object {
	obj, _ := env.Get(stmt.Name.Value)
	if obj != nil {
		return obj
	}
	return &object.Null{}
}
