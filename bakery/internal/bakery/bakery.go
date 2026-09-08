package bakery

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Start(ctx context.Context, k, n, m, t1, t2 int) {

	outBake := Bake(ctx, k, n, t1)
	outPack := Pack(ctx, m, t2, outBake)
	cakes := Inspect(outPack)
	fmt.Printf("Bakes %d/%d cakes\n", len(cakes), k)
}

func Bake(ctx context.Context, k, n, t1 int) <-chan Cake {
	out := make(chan Cake)
	var counter atomic.Int32
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				currentCake := counter.Add(1)
				if int(currentCake) > k {
					return
				}

				var cake Cake
				baketime := time.Duration(i+t1) * time.Millisecond

				select {
				case <-ctx.Done():
					return
				case <-time.After(baketime):
				}

				cake.BakedBy = i
				cake.BakeTime = t1

				select {
				case <-ctx.Done():
					return
				case out <- cake:
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Pack(ctx context.Context, m, t2 int, in <-chan Cake) <-chan Cake {
	out := make(chan Cake)
	var wg sync.WaitGroup

	for i := 0; i < m; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			packtime := time.Duration(i+t2) * time.Millisecond

			for cake := range in {
				select {
				case <-ctx.Done():
					return
				case <-time.After(packtime):
				}

				cake.PackedBy = i
				cake.PackedTime = t2

				select {
				case <-ctx.Done():
					return
				case out <- cake:
				}
			}
		}()

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Inspect(in <-chan Cake) []Cake {
	var cakes []Cake
	for cake := range in {
		cakes = append(cakes, cake)
	}
	return cakes
}
