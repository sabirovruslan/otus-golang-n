package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("excess capacity", func(t *testing.T) {
		c := NewCache(2)

		c.Set("first", 1)
		c.Set("second", "test")
		c.Set("excess", 3)

		_, ok := c.Get("first")
		require.False(t, ok)

		_, ok = c.Get("second")
		require.True(t, ok)

		_, ok = c.Get("excess")
		require.True(t, ok)
	})

	t.Run("frequency crowding", func(t *testing.T) {
		c := NewCache(4)
		for _, i := range []string{"1", "2", "3", "4"} {
			c.Set(Key(i), i)
		}

		_, ok := c.Get("1")
		require.True(t, ok)

		c.Set("5", 5)

		_, ok = c.Get("2")
		require.False(t, ok)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(4)
		for _, i := range []string{"1", "2", "3", "4"} {
			c.Set(Key(i), i)
		}

		_, ok := c.Get("1")
		require.True(t, ok)

		c.Clear()
		_, ok = c.Get("1")
		require.False(t, ok)
	})
}
