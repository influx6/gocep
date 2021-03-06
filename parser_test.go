package gocep

import (
	"reflect"
	"testing"
	"time"
)

func TestParserTimeWindow(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   int
		Message string
	}

	p := NewParser()
	p.Register("LogEvent", LogEvent{})

	q := "select count(*) from LogEvent.time(10 sec) where Level > 2 and Level < 10"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if stmt.time != 10*time.Second {
		t.Fail()
	}

	if reflect.TypeOf(stmt.function[0]).Name() != "Count" {
		t.Fail()
	}

	if reflect.TypeOf(stmt.selector[0]).Name() != "EqualsType" {
		t.Fail()
	}

	if reflect.TypeOf(stmt.selector[1]).Name() != "LargerThanInt" {
		t.Fail()
	}

	if reflect.TypeOf(stmt.selector[2]).Name() != "LessThanInt" {
		t.Fail()
	}
}

func TestParserError(t *testing.T) {
	p := NewParser()

	q := "select * from MapEvent.length(10)"
	_, err := p.Parse(q)
	if err == nil {
		t.Error("failed.")
	}

	if err.Error() != "parse selector: EventType [MapEvent] is not registered" {
		t.Errorf("failed: %v", err)
	}
}

func TestNewStatementLength(t *testing.T) {
	p := NewParser()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.length(10)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
		return
	}
	window := stmt.New(1024)
	defer window.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	window.Input() <- MapEvent{m}
	event := <-window.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}

func TestNewStatementTime(t *testing.T) {
	p := NewParser()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.time(10 sec)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	window := stmt.New(1024)
	defer window.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	window.Input() <- MapEvent{m}
	event := <-window.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}
