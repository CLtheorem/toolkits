package colorfulPrint

import "fmt"

func Println(color int, a interface{}) {
    fmt.Printf("\033[%d;%d;1m%v\033[0m\n", color, color, a)
}

func Printf(color int, format string, a interface{}) {
    fmt.Printf("\033[%d;%d;1m%s\033[0m\n", color, color, fmt.Sprintf(format, a))
}

func Sprint(color int, a interface{}) string {
    return fmt.Sprintf("\033[%d;%d;1m%s\033[0m", color, color, a)
}

func Sprintf(color int, format string, a interface{}) string {
    return fmt.Sprintf("\033[%d;%d;1m%s\033[0m", color, color, fmt.Sprintf(format, a))
}

// front color
const (
    ColorRed    = 31
    ColorGreen  = 32
    ColorYellow = 33
    ColorBlue   = 34
    ColorPurple = 35
    ColorCyan   = 36
)