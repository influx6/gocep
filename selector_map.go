package gocep

type EqualsMapString struct {
	Name  string
	Key   string
	Value string
}

func (f EqualsMapString) Select(e Event) bool {
	return e.MapString(f.Name, f.Key) == f.Value
}

type EqualsMapBool struct {
	Name  string
	Key   string
	Value bool
}

func (f EqualsMapBool) Select(e Event) bool {
	return e.MapBool(f.Name, f.Key) == f.Value
}

type EqualsMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f EqualsMapInt) Select(e Event) bool {
	return e.MapInt(f.Name, f.Key) == f.Value
}

type EqualsMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f EqualsMapFloat) Select(e Event) bool {
	return e.MapFloat(f.Name, f.Key) == f.Value
}

type NotEqualsMapString struct {
	Name  string
	Key   string
	Value string
}

func (f NotEqualsMapString) Select(e Event) bool {
	return e.MapString(f.Name, f.Key) != f.Value
}

type NotEqualsMapBool struct {
	Name  string
	Key   string
	Value bool
}

func (f NotEqualsMapBool) Select(e Event) bool {
	return e.MapBool(f.Name, f.Key) != f.Value
}

type NotEqualsMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f NotEqualsMapInt) Select(e Event) bool {
	return e.MapInt(f.Name, f.Key) != f.Value
}

type NotEqualsMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f NotEqualsMapFloat) Select(e Event) bool {
	return e.MapFloat(f.Name, f.Key) != f.Value
}

type LargerThanMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f LargerThanMapInt) Select(e Event) bool {
	return e.MapInt(f.Name, f.Key) > f.Value
}

type LargerThanMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f LargerThanMapFloat) Select(e Event) bool {
	return e.MapFloat(f.Name, f.Key) > f.Value
}

type LessThanMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f LessThanMapInt) Select(e Event) bool {
	return e.MapInt(f.Name, f.Key) < f.Value
}

type LessThanMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f LessThanMapFloat) Select(e Event) bool {
	return e.MapFloat(f.Name, f.Key) < f.Value
}
