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
  gotFileLen, err := os.Stat(got)
  checkRead(err, t)
  wantFileLen, err := os.Stat(want)
  checkRead(err, t)
  if gotFileLen.Size() != wantFileLen.Size() {
    t.Errorf("Files are not the same size.")
  }

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

func TestCanParseMetadata (t *testing.T) {
  wantMetadata := TemplateMetadata{
    Name: "gcp-template",
    Title: "GCP Template",
    Description: "Create GCP Resources",
  }
  wantSpec := TemplateSpec{
    Owner: "team1",
  }
  metadata := TemplateMetadata{}
  spec := TemplateSpec{}
  err := parseMetadata("../examples/header.yaml", &metadata, &spec)
  checkRead(err, t)
  gotMetadata := metadata
  gotSpec := spec

  if gotMetadata != wantMetadata {
    t.Errorf("Got %v, wanted %v", gotMetadata, wantMetadata)
  }
  if gotSpec != wantSpec {
    t.Errorf("Got %v, wanted %v", gotSpec, wantSpec)
  }
}

func TestCanNotParseNoMetadata (t *testing.T) {
  metadata := TemplateMetadata{}
  got := parseMetadata("../examples/error_header.yaml", &metadata, &TemplateSpec{})
  if got == nil {
    t.Errorf("Did not return error for bug.")
  }
}

func TestCreateTemplateHeader (t *testing.T) {
  testFileName := "../testgeneratefiles/create_template_header.yaml"
  matchFileName := "../testfiles/create_template_header.yaml"
  generatorFileName := "../examples/init_file.yaml"

  initTemplateFile(testFileName, generatorFileName)
  assertFileEquals(testFileName, matchFileName, t)
}

func TestCanCreateSpec (t *testing.T) {
  spec := TemplateSpec{}
  if &spec == nil {
    t.Errorf("Did not create spec struct")
  }
}

func TestCreateTemplateWithNoObjects (t *testing.T) {
  testFileName := "../testgeneratefiles/create_template_with_no_objects.yaml"
  matchFileName := "../testfiles/create_template_with_no_objects.yaml"
  generatorFileName := "../examples/no_objects.yaml"

  initTemplateFile(testFileName, generatorFileName)
  assertFileEquals(testFileName, matchFileName, t)
}


