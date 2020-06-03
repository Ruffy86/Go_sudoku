package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "sudoku/pkg/sudoku"

)



func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  list := &sudoku.MyStruct{}

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "sudoku",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(list)
  app.Run()
}
