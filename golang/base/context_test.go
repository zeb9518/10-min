package base

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "zeb")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 24)
	return child
}

func step3(ctx context.Context) {
	fmt.Printf("name %s \n", ctx.Value("name"))
	fmt.Printf("age %d \n", ctx.Value("age"))
}

func f1() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Microsecond*100)
	defer cancel()
	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("f1() err: %v\n", err)
	}
}

func f2() {
	parent, cancel1 := context.WithTimeout(context.TODO(), time.Millisecond*1000)
	defer cancel1()
	t0 := time.Now()

	time.Sleep(time.Millisecond * 500)
	child, cancel2 := context.WithTimeout(parent, time.Millisecond*100)
	defer cancel2()
	t1 := time.Now()
	select {
	case <-child.Done():
		err := child.Err()
		t3 := time.Now()
		fmt.Println(t3.Sub(t0).Microseconds(), t3.Sub(t1).Microseconds())
		fmt.Printf("f2() err: %v\n", err)
	}
}

func f3() {
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()
	defer cancel()
	select {
	case <-ctx.Done():
		err := ctx.Err()
		t3 := time.Now()
		fmt.Println(t3.Sub(t0).Microseconds())
		fmt.Printf("f3() err: %v\n", err)
	}
}

func TestContext(t *testing.T) {
	grandpa := context.TODO()
	father := step1(grandpa)
	grandson := step2(father)
	step3(grandson)
	//f1()
	//f2()
	f3()
}
