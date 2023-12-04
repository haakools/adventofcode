    
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)
	
	
func readFileToRunes(path string) [][]rune {

	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    width := 10 // should be set dynamically
    output := make([][]rune, width)
    // 
    i := 0
    for scanner.Scan() {
        fmt.Println("hello")
        output[i][:] = scanner.Text()
        fmt.Println(scanner.Text())
        fmt.Println(scanner.Text())
        i++
    }

    fmt.Println(output)


    return output
}


func main() {



    readFileToRunes("example.txt")



}
