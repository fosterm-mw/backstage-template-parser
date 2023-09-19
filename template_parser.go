package main

import (
  "os"
  "bufio"
  "fmt"
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

func main() {
  filePath := readFilePath()
  fmt.Print(filePath)
}
