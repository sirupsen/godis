package godis

import (
  "testing"
  // "fmt"
)

func TestStringStructureBasicSetAndGet(t *testing.T) {
  ss := NewStringStructure()
  ss.set("walrus", "happy")

  if ss.get("walrus") != "happy" {
    t.Errorf("Value of key was not set or received correctly")
  }
}

func TestStringStructureExists(t *testing.T) {
  ss := NewStringStructure()

  if ss.exists("narwhal") {
    t.Errorf("Exists reported existence of non-existant key")
  }

  ss.set("walrus", "whatever")

  if !ss.exists("walrus") {
    t.Errorf("Exists reported non-existance of an existant key")
  }
}

func TestStringStructureAppend(t *testing.T) {
  ss := NewStringStructure()

  ss.set("greeting", "hello")

  if ss.append("greeting", " world") != len("hello world") {
    t.Errorf("Append is not returning length of resulting string")
  }

  if ss.get("greeting") != "hello world" {
    t.Errorf("Append did not append correctly")
  }
}

func TestStringStructureAppendWhenKeyDoesntExist(t *testing.T) {
  ss := NewStringStructure()

  if ss.append("key", "value") != len("value") {
    t.Errorf("Append did not create non-existant key")
  }
}

func TestStringStructureIncrBy(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "10")

  res, err := ss.incrby("hits", "32")

  if err != nil {
    t.Errorf("INCRBY didn't interpret 10 as a number: %s", err.Error())
  }

  if res != "42" {
    t.Errorf("INCR didn't return the new incremented value")
  }

  if ss.get("hits") != "42" {
    t.Errorf("INCR didn't set the key")
  }
}

func TestStringStructureIncr(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "10")

  res, err := ss.incr("hits")

  if err != nil {
    t.Errorf("INCR didn't interpret 10 as a number")
  }

  if res != "11" {
    t.Errorf("INCR didn't return the new incremented value")
  }

  if ss.get("hits") != "11" {
    t.Errorf("INCR didn't set the key")
  }
}

func TestStringStructureIncrInvalidValue(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "can't touch this")

  _, err := ss.incr("hits")

  if err == nil {
    t.Errorf("INCR didn't yield an error on invalid value")
  }
}

func TestStringStructureDecr(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "10")

  res, err := ss.decr("hits")

  if err != nil {
    t.Errorf("DECR didn't interpret 10 as a number: %s", err.Error())
  }

  if res != "9" {
    t.Errorf("DECR didn't decrement the value by one")
  }

  if ss.get("hits") != "9" {
    t.Errorf("DECR didn't set the key")
  }
}

func TestStringStructureDecrby(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "10")

  res, err := ss.decrby("hits", "5")

  if err != nil {
    t.Errorf("DECR didn't interpret 10 as a number: %s", err.Error())
  }

  if res != "5" {
    t.Errorf("DECR didn't decrement the value by 5")
  }

  if ss.get("hits") != "5" {
    t.Errorf("DECR didn't set the key")
  }
}

func TestStringStructureDecrbyNegative(t *testing.T) {
  ss := NewStringStructure()

  ss.set("hits", "10")

  res, err := ss.decrby("hits", "-5")

  if err != nil {
    t.Errorf("DECR didn't interpret 10 as a number: %s", err.Error())
  }

  if res != "15" {
    t.Errorf("DECR didn't decrement the value by -5 (+5)")
  }

  if ss.get("hits") != "15" {
    t.Errorf("DECR didn't set the key")
  }
}
