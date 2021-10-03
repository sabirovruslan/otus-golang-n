package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestGetBack(t *testing.T) {
	l := NewList()
	value := "back"

	l.PushBack(value)
	l.PushBack(0)
	require.Equal(t, 2, l.Len())

	l.MoveToFront(l.Back())
	require.Equal(t, value, l.Back().Value.(string))
}

func TestRemoveByFrontList(t *testing.T) {
	l := NewList()
	for _, v := range []string{"", " ", "1", "test"} {
		l.PushFront(v)
	}
	require.Equal(t, 4, l.Len())

	l.Remove(l.Front())
	l.Remove(l.Front())
	l.Remove(l.Front())
	l.Remove(l.Front())

	require.Equal(t, 0, l.Len())
}

func TestRemoveByBackList(t *testing.T) {
	l := NewList()
	for _, v := range []int{1, 2, 2, 4} {
		l.PushBack(v)
	}
	require.Equal(t, 4, l.Len())

	l.Remove(l.Back())
	l.Remove(l.Back())
	l.Remove(l.Back())
	l.Remove(l.Back())

	require.Equal(t, 0, l.Len())
}

func TestGetFront(t *testing.T) {
	l := NewList()
	value := "front"

	l.PushBack(0)
	l.PushBack(value)
	require.Equal(t, 2, l.Len())

	l.MoveToFront(l.Back())
	require.Equal(t, value, l.Front().Value.(string))
}

func TestPushFront(t *testing.T) {
	l := NewList()
	v := l.PushFront(100)
	require.Equal(t, 100, v.Value)
	require.Equal(t, 100, l.Front().Value)

	v = l.PushFront(200)
	require.Equal(t, 200, v.Value)
	require.Equal(t, 200, l.Front().Value)
}
