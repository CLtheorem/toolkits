package main

import (
    "fmt"
    "github.com/CLtheorem/toolkit/colorfulPrint"
)

func main() {
    colorfulPrint.Println(colorfulPrint.ColorYellow, "colorfulPrint")
    colorfulPrint.Println(colorfulPrint.ColorGreen, "colorfulPrint")
    colorfulPrint.Println(colorfulPrint.ColorRed, "colorfulPrint")
    colorfulPrint.Println(colorfulPrint.ColorCyan, "colorfulPrint")
    colorfulPrint.Println(colorfulPrint.ColorBlue, "colorfulPrint")
    colorfulPrint.Println(colorfulPrint.ColorPurple, "colorfulPrint")
    colorfulPrint.Printf(colorfulPrint.ColorGreen, "result: %s", "success")
    fmt.Println(colorfulPrint.Sprintf(colorfulPrint.ColorRed, "error info: %s","errorrr"))
    fmt.Println(colorfulPrint.Sprint(colorfulPrint.ColorYellow, "yellllllllow"))
}
