package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
  "errors"
)

func readFilePath() string {

  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Please enter the file path: ")

  filePath, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("An error occured: ", err)
  }

  return filePath[:len(filePath)-1]
}

func openFile(filePath string) *os.File {
  f, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  return f
}

func readFileLineAsString(filePath string, lineNumber int) (string, error) {
  f := openFile(filePath)
  defer f.Close()

  scanner := bufio.NewScanner(f)
  currentLine := 0
  for scanner.Scan() {
    currentLine++
    if currentLine == lineNumber {
      return scanner.Text(), nil
    }
  }
  return "", fmt.Errorf("File %s does not have line number: %d", filePath, lineNumber)
}

func createNewFile(newFileName string) error {
  filename := newFileName

  _, err := os.Create(filename)
  if err != nil {
    fmt.Errorf("Failed to create file: %v", err)
  }

  return nil
}

func writeLineToFile(filePath string, line string) error {
  file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    fmt.Errorf("Failed to open file to append: %v", err)
  }
  line = line + "\n"

  _, err = file.WriteString(line)
  if err != nil {
    fmt.Errorf("Failed to write to file: %v", err)
  }
  defer file.Close()

  return nil
}

func initTemplateFile(filePath string, generatorFileName string) {
  apiVersion := "apiVersion: scaffolder.backstage.io/v1beta3"
  kind := "kind: Template"
  createNewFile(filePath)
  writeLineToFile(filePath, apiVersion)
  writeLineToFile(filePath, kind)

  metadata := TemplateMetadata{}
  spec := TemplateSpec{}
  err := parseMetadata(generatorFileName, &metadata, &spec)
  if err != nil {
    fmt.Errorf("Failed to pase Metadata: %v", err)
  }
  writeMetadata(filePath, metadata)
  writeSpec(filePath, spec)
}

func writeMetadata(filePath string, metadata TemplateMetadata) {
  writeLineToFile(filePath, "metadata:")
  writeLineToFile(filePath, fmt.Sprintf("  name: %v", metadata.Name))
  if metadata.Title != "" {
    writeLineToFile(filePath, fmt.Sprintf("  title: %v", metadata.Title))
  }
  if metadata.Description != "" {
    writeLineToFile(filePath, fmt.Sprintf("  description: %v", metadata.Description))
  }
}

func writeSpec(filePath string, spec TemplateSpec) {
  writeLineToFile(filePath, "spec:")
  if spec.Owner != "" {
    writeLineToFile(filePath, fmt.Sprintf("  owner: %v", spec.Owner))
  }
}

func getObjectLineNumber(scanner *bufio.Scanner, objectName string) int {
  lineNumber := 0
  for scanner.Scan() {
    lineNumber++
    currentLine := scanner.Text()
    if strings.Contains(currentLine, objectName){
      return lineNumber-1
    }
  }

  return -1
}

func parseMetadata(filePath string, metadataPointer *TemplateMetadata, specPointer *TemplateSpec) error {
  metadata := TemplateMetadata{}
  spec := TemplateSpec{}
  file := openFile(filePath)
  scanner := bufio.NewScanner(file)
  metadataLineNumber := getObjectLineNumber(scanner, "metadata")
  if metadataLineNumber == -1 {
    error := errors.New("Cannot parse metadata, object not supplied.")
    return error
  }
  lineNumber := 0
  for scanner.Scan() {
    lineNumber++
    if lineNumber >= metadataLineNumber{
      currentLine := scanner.Text()
      if strings.Contains(currentLine, "spec:"){
        break
      }
      if strings.Contains(currentLine, "name:"){
        metadata.Name = strings.Split(currentLine, ": ")[1]
      }
      if strings.Contains(currentLine, "title:"){
        metadata.Title = strings.Split(currentLine, ": ")[1]
      }
      if strings.Contains(currentLine, "description:"){
        metadata.Description = strings.Split(currentLine, ": ")[1]
      }
      if strings.Contains(currentLine, "owner:"){
        spec.Owner = strings.Split(currentLine, ": ")[1]
      }
    }
  }
  if metadata.Name == "" {
    error := errors.New("Cannot parse metadata, name required.")
    return error
  }
  *metadataPointer = metadata
  *specPointer = spec

  defer file.Close()
  return nil
}

func main() {
  filePath := readFilePath()
  fmt.Print(filePath)
}
