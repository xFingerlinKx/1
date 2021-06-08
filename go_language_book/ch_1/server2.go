// Server2 - минимальный "есhо"-сервер со счетчиком запросов.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	/*
		У сервера имеется два обработчика, и запрашиваемый URL определяет, какой из
		них будет вызван: запрос /count вызывает counter, а все прочие — handler2.
		Сервер запускает обработчик для каждого входящего запроса в отдельной go-подпрограмме,
		так что несколько запросов могут обрабатываться одновременно. Однако если два параллельных
		запроса попытаются обновить счетчик count в один и тот же момент времени, он может быть
		увеличен не согласованно; в такой программе может возникнуть серьезная ошибка под названием
		состояние гонки (race condition; см. раздел 9.1). Чтобы избежать этой проблемы, нужно
		гарантировать, что доступ к переменной получает не более одной go-подпрограммы одновременно.
		Для этого каждый доступ к переменной должен быть окружен вызовами mu.Lock() и mu.Unlock().
	*/
	http.HandleFunc("/", handler2) // каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL.
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Счетчик, возвращающий количество сделанных запросов
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
