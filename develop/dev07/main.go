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
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		wg := sync.WaitGroup{}
		single := make(chan interface{})

		output := func(c <-chan interface{}) {
			for n := range c {
				single <- n
			}
			wg.Done() // HL
		}

		wg.Add(len(channels)) // HL
		for _, c := range channels {
			go output(c)
		}

		go func() {
			wg.Wait() // HL
			close(single)
		}()
		return single
	}

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
