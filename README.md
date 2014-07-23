[![Build Status](https://img.shields.io/travis/Klathmon/Go-StructToMap.svg)](https://travis-ci.org/Klathmon/Go-StructToMap)
[![Coverage Status](https://img.shields.io/coveralls/Klathmon/Go-StructToMap.svg)](https://coveralls.io/r/Klathmon/Go-StructToMap?branch=master)
[![Go Walker Documentation](https://img.shields.io/badge/Go%20Walker-API%20Documentation-yellowgreen.svg)](https://gowalker.org/github.com/Klathmon/Go-StructToMap)
[![Licence](https://img.shields.io/badge/Licence-MIT-brightgreen.svg)](https://www.tldrlegal.com/l/mit)
# Go-StructToMap


A simple library for converting a `struct` to a `map` in Go.

    package main

    import "fmt"

    type AwesomeStruct struct {
      Stuff     string
      MoreStuff string
      IntStuff  int
    }

    func main() {
      newStruct := AwesomeStruct{"testing", "testing more", 7}

      newMap := structtomap.convert(newStruct)

      fmt.Println(newMap["Stuff"]) //Prints "testing"
    }
