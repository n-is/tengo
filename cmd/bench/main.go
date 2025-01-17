package main

import (
	"fmt"
	"time"

	"github.com/n-is/tengo/compiler"
	"github.com/n-is/tengo/compiler/ast"
	"github.com/n-is/tengo/compiler/parser"
	"github.com/n-is/tengo/compiler/source"
	"github.com/n-is/tengo/objects"
	"github.com/n-is/tengo/runtime"
)

func main() {
	runFib(35)
	runFibTC1(35)
	runFibTC2(35)
}

func runFib(n int) {
	start := time.Now()
	nativeResult := fib(n)
	nativeTime := time.Since(start)

	input := `
fib := func(x) {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}
` + fmt.Sprintf("out = fib(%d)", n)

	parseTime, compileTime, runTime, result, err := runBench([]byte(input))
	if err != nil {
		panic(err)
	}

	if nativeResult != int(result.(*objects.Int).Value) {
		panic(fmt.Errorf("wrong result: %d != %d", nativeResult, int(result.(*objects.Int).Value)))
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("fibonacci(%d)\n", n)
	fmt.Println("-------------------------------------")
	fmt.Printf("Result:  %d\n", nativeResult)
	fmt.Printf("Go:      %s\n", nativeTime)
	fmt.Printf("Parser:  %s\n", parseTime)
	fmt.Printf("Compile: %s\n", compileTime)
	fmt.Printf("VM:      %s\n", runTime)
}

func runFibTC1(n int) {
	start := time.Now()
	nativeResult := fibTC1(n, 0)
	nativeTime := time.Since(start)

	input := `
fib := func(x, s) {
	if x == 0 {
		return 0 + s
	} else if x == 1 {
		return 1 + s
	}

	return fib(x-1, fib(x-2, s))
}
` + fmt.Sprintf("out = fib(%d, 0)", n)

	parseTime, compileTime, runTime, result, err := runBench([]byte(input))
	if err != nil {
		panic(err)
	}

	if nativeResult != int(result.(*objects.Int).Value) {
		panic(fmt.Errorf("wrong result: %d != %d", nativeResult, int(result.(*objects.Int).Value)))
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("fibonacci(%d) (tail-call #1)\n", n)
	fmt.Println("-------------------------------------")
	fmt.Printf("Result:  %d\n", nativeResult)
	fmt.Printf("Go:      %s\n", nativeTime)
	fmt.Printf("Parser:  %s\n", parseTime)
	fmt.Printf("Compile: %s\n", compileTime)
	fmt.Printf("VM:      %s\n", runTime)
}

func runFibTC2(n int) {
	start := time.Now()
	nativeResult := fibTC2(n, 0, 1)
	nativeTime := time.Since(start)

	input := `
fib := func(x, a, b) {
	if x == 0 {
		return a
	} else if x == 1 {
		return b
	}

	return fib(x-1, b, a+b)
}
` + fmt.Sprintf("out = fib(%d, 0, 1)", n)

	parseTime, compileTime, runTime, result, err := runBench([]byte(input))
	if err != nil {
		panic(err)
	}

	if nativeResult != int(result.(*objects.Int).Value) {
		panic(fmt.Errorf("wrong result: %d != %d", nativeResult, int(result.(*objects.Int).Value)))
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("fibonacci(%d) (tail-call #2)\n", n)
	fmt.Println("-------------------------------------")
	fmt.Printf("Result:  %d\n", nativeResult)
	fmt.Printf("Go:      %s\n", nativeTime)
	fmt.Printf("Parser:  %s\n", parseTime)
	fmt.Printf("Compile: %s\n", compileTime)
	fmt.Printf("VM:      %s\n", runTime)
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fibTC1(n, s int) int {
	if n == 0 {
		return 0 + s
	} else if n == 1 {
		return 1 + s
	}
	return fibTC1(n-1, fibTC1(n-2, s))
}

func fibTC2(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		return fibTC2(n-1, b, a+b)
	}
}

func runBench(input []byte) (parseTime time.Duration, compileTime time.Duration, runTime time.Duration, result objects.Object, err error) {
	var astFile *ast.File
	parseTime, astFile, err = parse(input)
	if err != nil {
		return
	}

	var bytecode *compiler.Bytecode
	compileTime, bytecode, err = compileFile(astFile)
	if err != nil {
		return
	}

	runTime, result, err = runVM(bytecode)

	return
}

func parse(input []byte) (time.Duration, *ast.File, error) {
	fileSet := source.NewFileSet()
	inputFile := fileSet.AddFile("bench", -1, len(input))

	start := time.Now()

	p := parser.NewParser(inputFile, input, nil)
	file, err := p.ParseFile()
	if err != nil {
		return time.Since(start), nil, err
	}

	return time.Since(start), file, nil
}

func compileFile(file *ast.File) (time.Duration, *compiler.Bytecode, error) {
	symTable := compiler.NewSymbolTable()
	symTable.Define("out")

	start := time.Now()

	c := compiler.NewCompiler(file.InputFile, symTable, nil, nil, nil)
	if err := c.Compile(file); err != nil {
		return time.Since(start), nil, err
	}

	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()

	return time.Since(start), bytecode, nil
}

func runVM(bytecode *compiler.Bytecode) (time.Duration, objects.Object, error) {
	globals := make([]objects.Object, runtime.GlobalsSize)

	start := time.Now()

	v := runtime.NewVM(bytecode, globals, -1)
	if err := v.Run(); err != nil {
		return time.Since(start), nil, err
	}

	return time.Since(start), globals[0], nil
}
