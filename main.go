package main

import (
    "fmt"
    "sort"
)

func main() {
    strArr := []string{"Oregon", "Massachussets", "Wisconsin", "Washington"} 

    sort.Strings(strArr)
    fmt.Println(strArr)
    
}
