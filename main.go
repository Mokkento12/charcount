package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
    counts := make(map[rune]int)        // Счетчик символов
    var utflen [utf8.UTFMax + 1]int     // Массив [0,1,2,3,4] для длин 1-4 байта
    invalid := 0                        // Счетчик ошибок

    in := bufio.NewReader(os.Stdin)

    fmt.Println("Введите текст (Ctrl+C для завершения):")
    
    for {
        r, n, err := in.ReadRune()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        
        // Проверяем ошибки UTF-8
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        
        counts[r]++     // Считаем символы
        utflen[n]++     // Считаем длины кодировок
    }
    
    // Выводим символы
    fmt.Printf("\n=== Символы ===\n")
    fmt.Printf("Символ\tКоличество\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    
    // Выводим длины кодировок
    fmt.Printf("\n=== Длины кодировок ===\n")
    fmt.Printf("Байты\tКоличество\n")
    for i, n := range utflen {
        if i > 0 {  // Пропускаем индекс 0
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    
    // Выводим ошибки
    if invalid > 0 {
        fmt.Printf("\n=== Ошибки ===\n")
        fmt.Printf("%d неверных символов UTF-8\n", invalid)
    }
}