package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции,
которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

*/

// fan-in chan

func main() {
	// or - функция, которая принимает несколько каналов типа `<-chan interface{}` в качестве входных данных и возвращает канал типа `<-chan interface{}`.
	// Эта функция объединяет все входные каналы в один канал и возвращает его.
	// Функция использует WaitGroup для того, чтобы убедиться, что все входные каналы были опустошены, прежде чем закрыть выходной канал.
	//
	// Параметры:
	// - channels: вариативный параметр типа `<-chan interface{}`, представляющий входные каналы.
	//
	// Возвращает:
	// - `<-chan interface{}`: канал, который объединяет все входные каналы в один канал.
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		wg := sync.WaitGroup{}
		single := make(chan interface{})

		// функция для опустошения входного канала и отправки его значений в одиночный канал
		output := func(c <-chan interface{}) {
			for n := range c {
				single <- n
			}
			wg.Done()
		}

		wg.Add(len(channels))
		// запускаем горутину для опустошения каждого входного канала
		for _, c := range channels {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(single)
		}()
		return single
	}

	// sig - функция, которая принимает время `after` в качестве входных данных и возвращает канал типа `<-chan interface{}`.
	// Эта функция создает канал и запускает горутину, которая задерживает выполнение на указанное время и затем закрывает канал.
	//
	// Параметры:
	// - after: время в формате `time.Duration`, представляющее задержку выполнения.
	//
	// Возвращает:
	// - `<-chan interface{}`: канал, который будет закрыт после выполнения задержки.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		// sig(1*time.Hour),
		// sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
