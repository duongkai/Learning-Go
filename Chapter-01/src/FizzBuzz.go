package main
import "fmt"

func main() {
    // Fizz Buzz
    for gi := 0; gi < 100; gi++ {
        var tmp string
        switch {
        case gi % 15 == 0:
            tmp = fmt.Sprintf ("%d  FizzBuzz", gi)
        case gi % 5 == 0 && gi % 3 != 0:
            tmp = fmt.Sprintf ("%d  Buzz", gi)
        case gi % 3 == 0 && gi % 5 != 0:
            tmp = fmt.Sprintf ("%d  Fizz", gi)
        default: tmp = fmt.Sprintf ("%d", gi)
        }
        fmt.Println (tmp)
    }
}
