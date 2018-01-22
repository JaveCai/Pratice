package main 
import "fmt"
func main() { 
        input := "The quick brown 狐 jumped over the lazy 犬" 
        // Get Unicode code points. 
	fmt.Printf("len: %d\n", len(input))
        n := 0
        rune := make([]rune, len(input))
        for _, r := range input { 
                rune[n] = r
                n++
        } 
        rune = rune[0:n]
        fmt.Printf("[]rune len: %d\n", len(rune))
        // Reverse 
        for i := 0; i < n/2; i++ { 
                rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
        } 
        // Convert back to UTF-8. 
        output := string(rune)
        fmt.Println(output)
}