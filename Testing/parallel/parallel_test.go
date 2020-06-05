package parallel

import (
	"testing"
)

func TestSomething(t *testing.T) {
	//	t.Parallel()
}

func TestA(t *testing.T) {
	//	t.Parallel()
}

func TestB(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {

	})

	t.Run("sub2", func(t *testing.T) {

	})
}
