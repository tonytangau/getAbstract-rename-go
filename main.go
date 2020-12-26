package main

import (
    "os"
    "fmt"
    "io/ioutil"
     "log"
     "strings"
     "regexp"
)

func main() {
  // Replace path with your downloaded pdf folder
  var path = "D:/getAbstract/"

  files, err := ioutil.ReadDir(path)

  if err != nil {
      log.Fatal(err)
  }

  for _, f := range files {
      // Print initial name
      fmt.Println(f.Name())

      // Remova trailing -en-
      reg := regexp.MustCompile(`-en-.*`)
      cleanName := reg.ReplaceAllString(f.Name(), "${1}")

      var fullName = strings.Title(strings.ToLower(strings.Replace(cleanName, "-", " ", -1)))

      err = os.Rename(path + f.Name(), path + fullName + ".pdf")

      if err != nil {
          log.Fatal(err)
      }
  }
}