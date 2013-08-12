package godis

import (
  "strconv"
  "errors"
)

type StringStructure struct {
  data map[string]string
}

func NewStringStructure() *StringStructure {
  ss := new(StringStructure)
  ss.data = make(map[string]string)

  return ss
}

func (s *StringStructure) set(key string, value string) {
  s.data[key] = value
}

func (s *StringStructure) exists(key string) (ok bool) {
  _, ok = s.data[key]
  return
}

func (s *StringStructure) get(key string) string {
  return s.data[key]
}

func (s *StringStructure) append(key string, suffix string) int {
  s.data[key] += suffix
  return len(s.data[key])
}

func (s *StringStructure) incrby(key string, incrementer string) (string, error) {
  n, err := strconv.ParseInt(s.data[key], 10, 64)

  if err != nil {
    return "", errors.New("value is not an integer or out of range")
  }

  int_incrementer, err := strconv.ParseInt(incrementer, 10, 64)

  if err != nil {
    return "", errors.New("value is not an integer or out of range")
  }

  s.data[key] = strconv.Itoa(int(n) + int(int_incrementer))

  return s.data[key], nil
}

func (s *StringStructure) incr(key string) (string, error) {
  n, err := s.incrby(key, "1")
  return n, err
}

func (s *StringStructure) decrby(key string, decrementer string) (string, error) {
  n, err := s.reverseSign(decrementer)

  if err != nil {
    return "", err
  }

  n, err = s.incrby(key, n)

  return n, err
}

func (s *StringStructure) reverseSign(number string) (string, error) {
  n, err := strconv.ParseInt(number, 10, 64)
  return strconv.Itoa(-int(n)), err
}

func (s *StringStructure) decr(key string) (string, error) {
  n, err := s.incrby(key, "-1")
  return n, err
}

