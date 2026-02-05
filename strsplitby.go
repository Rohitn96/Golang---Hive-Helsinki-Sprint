package main

// import "fmt"

func StrSplitBy(s, sep string) []string {
    var result []string // start as empty slice
    part := ""

    if s == "" || sep == "" {
        return []string{} //empty slice instead of nil
    }

    for i := 0; i < len(s); {
        if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
            result = append(result, part)
            part = ""
            i += len(sep)
        } else {
            part += string(s[i])
            i++
        }
    }

    result = append(result, part)
    return result
}

// func main() {
// 	fmt.Println(StrSplitBy("HowYOUhaveYOUyouYOUbeen?", "YOU"))
// }