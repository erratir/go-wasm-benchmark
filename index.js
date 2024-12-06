async function runJavaScriptBenchmark() {
    const iterations = parseInt(document.getElementById('iterations').value, 10);
    const startTime = performance.now();
    let sum = 0;
    for (let i = 0; i < iterations; i++) {
        for (let j = 0; j < iterations; j++) {
            sum += i + j;
        }
    }
    const endTime = performance.now();
    const elapsedTime = endTime - startTime;

    // Отображение результата на странице
    document.getElementById('jsResult').textContent = `JavaScript Benchmark: ${elapsedTime.toFixed(2)} ms, Sum: ${sum}`;
}

// Экспортируем функцию для вызова из HTML
window.runJavaScriptBenchmark = runJavaScriptBenchmark;