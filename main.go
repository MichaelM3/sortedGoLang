package main

import (
    "fmt"
    "sort"
)

func main() {
    strArr := []string{"Oregon", "Massachussets", "Wisconsin", "Washington", "Appls", "Apple"} 

    sort.Slice(strArr, func(i, j int) bool {
        return len(strArr[i]) < len(strArr[j])
    })

    fmt.Println(strArr)
    
}
