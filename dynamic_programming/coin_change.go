package dynamicprogramming

import (
	"math"
	"slices"
)

/*
322. Coin Change

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.
*/

func coinChange(coins []int, amount int) int {
	// Approach
	// This is a classic memoization solution. A key element of dynamic programming
	// is to identify a "base case" and either build up or break down from there.
	// The int memos[i] tells us the minimum number of coins it takes to get amount i.
	// To find memos[i] we consider all the coin values c and look at the amount i-c.
	// if memos[i-c] is not MaxInt that means we got to it and therefore we can get to
	// i in at most memos[i-c] + 1 steps. Checking each coin value, we will get the
	// optimal number for memos[i].
	// Our initial case is memos[0] = 0.

	slices.Sort(coins)
	memos := make([]int, amount+1)

	for i := range memos {
		if i != 0 {
			// start by assuming we can't get there
			memos[i] = math.MaxInt
		}
		for _, c := range coins {
			// the coins are sorted so if c is bigger than i then we
			// can just move on to the next value
			if i < c {
				break
			}

			// check if we reached i-c, if we have then make sure we pick the
			// c that gives us the smallest number of coins
			if memos[i-c] != math.MaxInt {
				memos[i] = min(memos[i], 1+memos[i-c])
			}
		}
	}

	if memos[amount] == math.MaxInt {
		return -1
	} else {
		return memos[amount]
	}
}

// WARNING:
// this approach doesn't work because dfs doesn't guarantee the shortest path
// and bfs is too slow
type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Length() int {
	return len(q.arr)
}

func (q *Queue[T]) Push(value T) {
	q.arr = append(q.arr, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.arr) == 0 {
		var zero T
		return zero, false
	}
	t := q.arr[0]
	q.arr = q.arr[1:]
	return t, true
}

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Length() int {
	return len(s.arr)
}

func (s *Stack[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.arr) == 0 {
		var zero T
		return zero, false
	}
	t := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return t, true
}

func coinChangedfs(coins []int, amount int) int {
	slices.Sort(coins)

	type Item struct {
		Amount   int
		NumCoins int
	}
	container := Stack[Item]{}
	container.Push(Item{0, 0})

	for container.Length() > 0 {
		item, ok := container.Pop()
		if !ok {
			panic("bad pop")
		}
		// for i := len(coins) - 1; i >= 0; i-- {
		for i := range coins {
			if item.Amount == amount {
				return item.NumCoins
			}
			if item.Amount+coins[i] == amount {
				return item.NumCoins + 1
			}
			if item.Amount+coins[i] < amount {
				container.Push(Item{Amount: item.Amount + coins[i], NumCoins: item.NumCoins + 1})
			}
		}
	}
	return -1
}
