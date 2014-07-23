[![Release Version](http://img.shields.io/github/release/Klathmon/StructToMap.svg)](https://github.com/Klathmon/StructToMap)
[![Build Status](https://img.shields.io/travis/Klathmon/StructToMap.svg)](https://travis-ci.org/Klathmon/StructToMap)
[![Coverage Status](https://img.shields.io/coveralls/Klathmon/StructToMap.svg)](https://coveralls.io/r/Klathmon/StructToMap?branch=master)
[![Go Walker Documentation](https://img.shields.io/badge/Go%20Walker-API%20Documentation-yellowgreen.svg)](https://gowalker.org/github.com/Klathmon/StructToMap)
[![Licence](https://img.shields.io/badge/Licence-MIT-brightgreen.svg)](https://www.tldrlegal.com/l/mit)
# StructToMap


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
