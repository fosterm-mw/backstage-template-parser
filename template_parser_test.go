package main

import (
  "testing"
  "bufio"
  "os"
)

func assertStringEquals(got string, want string, t *testing.T) bool {
  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
    return false
  }
  return true
}

func checkRead(err error, t *testing.T) {
  if err != nil {
    t.Errorf("Error %e", err)
  }
}

func assertFileEquals(got string, want string, t *testing.T) bool {
  gotFile, err := os.Open(got)
  checkRead(err, t)
  wantFile, err := os.Open(want)
  checkRead(err, t)

  gotScanner := bufio.NewScanner(gotFile)
  wantScanner := bufio.NewScanner(wantFile)
  for gotScanner.Scan() && wantScanner.Scan() {
    gotCurrentLine := gotScanner.Text()
    wantCurrentLine := wantScanner.Text()
    if gotCurrentLine != wantCurrentLine {
      t.Errorf("Got %v, wanted %v", gotCurrentLine, wantCurrentLine)
    }
  }
  
  return true
}

func TestReadFilePath (t *testing.T) {
  path := "example.yaml"
  assertStringEquals(path, "example.yaml", t)
}

func TestReadFirstLine (t *testing.T) {
  firstLine, err := readFileLineAsString("example.yaml", 1)
  checkRead(err, t)

  assertStringEquals(firstLine, "apiVersion: v1", t)
}

func TestReadSecondLine (t *testing.T) {
  secondLine, err := readFileLineAsString("example.yaml", 2)
  checkRead(err, t)

  assertStringEquals(secondLine, "kind: Backstage-Template", t)
}

func TestWriteOneLineToFile (t *testing.T) {
  line, err := readFileLineAsString("example.yaml", 1)
  checkRead(err, t)

  writeLineToFile("test.yaml", line)
  want, _ := readFileLineAsString("example.yaml", 1)
  got, _ := readFileLineAsString("test.yaml", 1)
  assertStringEquals(got, want, t)
}

func TestWriteApiVersionAndKindToFile (t *testing.T) {
  testFileName := "test.yaml"
  createNewFile(testFileName)
  
  line, err := readFileLineAsString("example.yaml", 1)
  checkRead(err, t)
  writeLineToFile("test.yaml", line)

  line, err = readFileLineAsString("example.yaml", 2)
  checkRead(err, t)
  writeLineToFile("test.yaml", line)

  assertFileEquals(testFileName, "testfiles/write_api_version_and_kind_to_file.yaml", t)
}

func TestCanParseMetadata (t *testing.T) {
  want := TemplateMetadata{
    Name: "gcp-template",
    Title: "GCP Template",
    Description: "Create GCP Resources",
  }
  got, err := parseMetadata("examples/header.yaml")
  checkRead(err, t)

  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
  }
}

func TestCanNotParseNoMetadata (t *testing.T) {
  _, got := parseMetadata("examples/error_header.yaml")
  if got == nil {
    t.Errorf("Did not return error for bug.")
  }
}

func TestCreateTemplateHeader (t *testing.T) {
  testFileName := "test2.yaml"
  matchFileName := "testfiles/create_template_header.yaml"
  generatorFileName := "examples/init_file.yaml"
  initTemplateFile(testFileName, generatorFileName)
  assertFileEquals(testFileName, matchFileName, t)
  //store any annotations
  //store any labels
}

