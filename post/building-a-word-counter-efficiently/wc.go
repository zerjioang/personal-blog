package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
)

const (
    // wc chunk size 16384
    chunksize = 4096
)

type chunk [chunksize]byte

func countWords(name string) (wordCount int) {
    // open the file for reading
    f, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    // create a buffered reader
    reader := bufio.NewReader(f)
    var part chunk
    wordCount = 0
    previousWord := 0
    for {
        var totalread int
        if totalread, err = reader.Read(part[:]); err != nil {
            break
        }
        for i := 1; i < totalread; i++ {
            c := part[i-1]
            isSpecial := c == '\n' || c == '\r' || c == '\t'
            if (c == ' ' || isSpecial) && part[i] != ' ' {
                wordCount++
                // fmt.Println(wordCount, string(part[previousWord:i]))
                if isSpecial {
                    i++
                }
                previousWord = i
            }
        }
        if previousWord != totalread {
            wordCount++
        }
    }
    if err != io.EOF {
        log.Fatal("error reading ", name, ": ", err)
    } else {
        err = nil
    }
    return
}

func main() {
    file := os.Args[1]
    count := countWords(file)
    fmt.Println(count)
}