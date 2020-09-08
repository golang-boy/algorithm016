package Week_01

import (
	"log"
	"testing"
)

func TestMyCircularDeque(t *testing.T) {
	deque := Constructor(3)

	log.Println(deque.InsertLast(1))
	log.Println(deque.InsertLast(2))

	log.Println(deque.InsertFront(6))
	log.Println(deque.InsertFront(7))

	log.Println(deque.GetRear())
}
