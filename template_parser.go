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

func writeLineToFile(newFileName string, line string) error {
  filename := newFileName

  file, err := os.Create(filename)
  if err != nil {
    fmt.Errorf("Failed to create file: %v", err)
  }

  _, err = file.WriteString(line)
  defer file.Close()

  return nil
}

func main() {
  filePath := readFilePath()
  fmt.Print(filePath)
}
