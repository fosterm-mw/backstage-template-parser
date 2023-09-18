package main

import (
  "testing"
)

func assertStringEquals(got string, want string, t *testing.T) bool {
  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
    return false
  }
  return true
}

func TestReadFilePath (t *testing.T) {
  path := "example.yaml"
  assertStringEquals(path, "example.yaml", t)
}

func TestReadFirstLine (t *testing.T) {
  firstLine, err := readFileLineAsString("example.yaml", 1)
  if err != nil {
    t.Errorf("Error %e", err)
  }

  assertStringEquals(firstLine, "apiVersion: v1", t)
}

func TestReadSecondLine (t *testing.T) {
  secondLine, err := readFileLineAsString("example.yaml", 2)
  if err != nil {
    t.Errorf("Error %e", err)
  }

  assertStringEquals(secondLine, "kind: Backstage-Template", t)
}

func TestWriteOneLineToFile (t *testing.T) {
  line, err := readFileLineAsString("example.yaml", 1)
  if err != nil {
    t.Errorf("Error %e", err)
  }
  writeLineToFile("test.yaml", line)
  want, _ := readFileLineAsString("example.yaml", 1)
  got, _ := readFileLineAsString("test.yaml", 1)
  assertStringEquals(got, want, t)
}

