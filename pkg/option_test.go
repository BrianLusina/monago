package pkg

import "testing"

func TestOption(t *testing.T) {
	option := NewOption[string]()

	val, err := option.Take()

	if err == nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got %v", val)
	}

	option.Set("Hola!")
	_, err = option.Take()
	if err != nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got %v", err)
	}

	option.Clear()
	if option.IsSome() {
		t.Fatalf("Option[T] should have none, but has some")
	}
}

func BenchmarkOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		option := NewOption[string]()

		val, err := option.Take()

		if err == nil {
			b.Fatalf("[unexpected] wanted no value out of Option[T], got %v", val)
		}

		option.Set("Hola!")
		_, err = option.Take()
		if err != nil {
			b.Fatalf("[unexpected] wanted no value out of Option[T], got %v", err)
		}

		option.Clear()
		if option.IsSome() {
			b.Fatalf("Option[T] should have none, but has some")
		}
	}
}
