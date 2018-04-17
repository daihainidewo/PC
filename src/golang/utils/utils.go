// file create by daihao
package utils

import (
	"sync"
)

type TokenBucket struct {
	tong chan interface{}
	wg   sync.WaitGroup
}

func NewTokenBucket(size int) *TokenBucket {
	t := new(TokenBucket)
	t.tong = make(chan interface{}, size)
	for i := 0; i < size; i++ {
		t.tong <- new(interface{})
	}
	return t
}

func (this *TokenBucket) Get() {
	this.wg.Add(1)
	<-this.tong
}

func (this *TokenBucket) Put() {
	this.wg.Done()
	this.tong <- new(interface{})
}

func (this *TokenBucket) Close() {
	this.wg.Wait()
}

func (this *TokenBucket) Len() int {
	return len(this.tong)
}