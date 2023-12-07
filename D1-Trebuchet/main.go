package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    //puzzleContent, _ := os.ReadFile("./Day1.txt")
    //fmt.Println(string(puzzleContent))
    file, _ := os.Open("./Day1.txt")
    defer file.Close()
    lineScanner := bufio.NewScanner(file)
    sum := 0
    for lineScanner.Scan() {
        lineText := lineScanner.Text()
        // regex first digit: ^[^\d]*(\d)
        firstNumFinder, _ := regexp.Compile("(\\d{1})")
        firstNum := firstNumFinder.FindString(lineText)
        fmt.Printf("First num: %v  ;", firstNum)
        // go doesn't support negative lookahead to preserve O(n) time
        ////regex last digit: (\d)(?!.*\d)
        //lastNumFinder, err := regexp.Compile("(\\d)(?!.*\\d)")
        //if err != nil {
        //    fmt.Println(err)
        //}
        //lastNum := lastNumFinder.FindString(lineText)
        lastNum := firstNumFinder.FindString(Reverse(lineText))
        fmt.Printf("lastNum: %v\n", lastNum)
        i, _ := strconv.Atoi(firstNum + lastNum)
        sum += i
    }
    fmt.Println(sum)

}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
