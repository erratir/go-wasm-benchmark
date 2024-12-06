package main

import (
	"strconv"
	"syscall/js"
	"time"
)

func runGoBenchmark(this js.Value, args []js.Value) interface{} {
	return js.Global().Call("runGoBenchmarkAsync")
}

func runGoBenchmarkAsync(this js.Value, args []js.Value) interface{} {
	return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]

		// Получаем значение из поля ввода и преобразуем его в целое число
		iterationsStr := js.Global().Get("document").Call("getElementById", "iterations").Get("value").String()
		iterations, err := strconv.Atoi(iterationsStr)
		if err != nil {
			js.Global().Get("document").Call("getElementById", "goResult").Set("textContent", "Error: Invalid number of iterations")
			resolve.Invoke()
			return nil
		}

		start := time.Now()
		sum := 0
		for i := 0; i < iterations; i++ {
			for j := 0; j < iterations; j++ {
				sum += i + j
			}
		}
		elapsed := time.Since(start)

		// Отображение результата на странице
		js.Global().Get("document").Call("getElementById", "goResult").Set("textContent", "Go WASM Benchmark: "+elapsed.String()+", Sum: "+strconv.Itoa(sum))

		resolve.Invoke()
		return nil
	}))
}

func main() {
	// Экспортируем функцию для вызова из JavaScript
	js.Global().Set("runGoBenchmark", js.FuncOf(runGoBenchmark))
	js.Global().Set("runGoBenchmarkAsync", js.FuncOf(runGoBenchmarkAsync))

	// Удерживаем main-функцию в активном состоянии
	select {}
}
