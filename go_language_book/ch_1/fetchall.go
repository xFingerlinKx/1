// Fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go-подпрограмма представляет собой параллельное выполнение функции. Канал
		// является механизмом связи, который позволяет одной go-подпрограмме передавать
		// значения определенного типа другой go-подпрограмме. Функция main выполняется
		// в go-подпрограмме, а инструкция go создает дополнительные go-подпрограммы.
		go fetch(url, ch) // Запуск go-подпрограммы
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // получение из канала ch
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // отправка в канал ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
