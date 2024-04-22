package main

import (
		"fmt"
		"time"
		)
func or(channels ...<-chan interface{})<- chan interface{}{
	done := make(chan interface{})
	for i := range(channels){
		go func(ch <-chan interface{}){
		select{
			case <- ch:
				close(done)
			case <-done:
				return
		}
		}(channels[i])
	}
	return done

}

func main(){
	start := time.Now()
	sig := func(after time.Duration) <-chan interface{}{
		c := make(chan interface{})
		go func(){
			defer close(c)
			time.Sleep(after)
		}()
		return c
		}
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("program finished after %v\n", time.Since(start))

}