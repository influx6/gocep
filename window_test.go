package gocep

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkLengthWindowNoFunction128(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}

}

func BenchmarkLengthWindowSumInt(b *testing.B) {
	w := NewLengthWindow(1)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt64(b *testing.B) {
	w := NewLengthWindow(64)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt128(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt256(b *testing.B) {
	w := NewLengthWindow(256)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowAverageMap(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})
	w.SetFunction(AverageMapInt{"Record", "Value", "avg(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}

}

func BenchmarkLengthWindowAverageInt(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(AverageInt{"Value", "avg(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowLargerThanMap(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})
	w.SetSelector(LargerThanMapInt{"Record", "Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowLargerThanInt(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetSelector(LargerThanInt{"Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowOrderByMap(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})
	w.SetView(OrderByMapInt{"Record", "Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowOrderByInt(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetView(OrderByInt{"Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowOrderByReverseMap(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})
	w.SetView(OrderByMapInt{"Record", "Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowOrderByReverseInt(b *testing.B) {
	w := NewLengthWindow(128)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetView(OrderByInt{"Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func TestConcurrency(t *testing.T) {
	w := NewLengthWindow(2)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetSelector(LargerThanInt{"Value", 1})
	w.SetFunction(Count{"count"})
	w.SetView(OrderByInt{"Value", true})

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			w.Input() <- IntEvent{"foo", rand.Int()}
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		<-w.Output()
	}
}

func TestLengthWindow(t *testing.T) {
	w := NewLengthWindow(2)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.SetSelector(LargerThanInt{"Value", 1})
	w.SetFunction(Count{"count"})
	w.SetView(OrderByInt{"Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if w.Capacity() != 1024 {
		t.Error(w.Capacity())
	}

	var test = []struct {
		index int
		count int
		value int
	}{
		{0, 2, 9},
		{1, 2, 8},
	}

	for _, tt := range test {
		if event[tt.index].Record["count"] != tt.count {
			t.Error(event)
		}
		if event[tt.index].Int("Value") != tt.value {
			t.Error(event)
		}
	}

	if Oldest(w.Event()).Record["count"] != 2 {
		t.Error(w.Event())
	}

}

func TestLengthWindowMap(t *testing.T) {

	w := NewLengthWindow(2)
	defer w.Close()

	w.SetSelector(EqualsType{MapEvent{}})
	w.SetSelector(LargerThanMapInt{"Record", "Value", 1})
	w.SetFunction(Count{"count"})
	w.SetFunction(AverageMapInt{"Record", "Value", "avg(Record:Value)"})
	w.SetView(OrderByMapInt{"Record", "Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		event = w.Update(MapEvent{m})
	}

	var test = []struct {
		index int
		count int
		value int
		avg   float64
	}{
		{0, 2, 9, 8.5},
		{1, 2, 8, 8.5},
	}

	for _, tt := range test {
		if event[tt.index].Record["count"] != tt.count {
			t.Error(event)
		}
		if event[tt.index].MapInt("Record", "Value") != tt.value {
			t.Error(event)
		}
		if event[tt.index].Record["avg(Record:Value)"] != tt.avg {
			t.Error(event)
		}
	}
}

func TestLengthWindowListen(t *testing.T) {

	w := NewLengthWindow(2)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	w.Listen("")

}

func TestLengthBatchWindow(t *testing.T) {

	w := NewLengthBatchWindow(2)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if event[0].Int("Value") != 8 {
		t.Error(event)
	}

	if event[1].Int("Value") != 9 {
		t.Error(event)
	}

}

func TestTimeWindow0ms(t *testing.T) {
	w := NewTimeWindow(0 * time.Millisecond)
	defer w.Close()

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if len(event) != 0 {
		t.Error(event)
	}
}

func TestTimeWindow10ms(t *testing.T) {
	w := NewTimeWindow(1 * time.Millisecond)
	defer w.Close()

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if len(event) == 0 {
		t.Error(event)
	}
}

func TestTimeBatchWindow10ms(t *testing.T) {
	w := NewTimeBatchWindow(4 * time.Millisecond)
	defer w.Close()

	for i := 0; i < 10; i++ {
		w.Update(IntEvent{"foo", i})
	}
}

func TestLengthWindowPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	w := NewLengthWindow(10)
	defer w.Close()

	w.SetSelector(EqualsType{IntEvent{}})
	// IntEvent and Map Function -> panic!!
	w.SetFunction(AverageMapInt{"Record", "Value", "avg(Record:Value)"})
	event := w.Update(IntEvent{"foobar", 10})
	if len(event) != 0 {
		t.Error(event)
	}
}
