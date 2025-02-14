/*
Sample of Binary Summator, Little Indian
*/
package main

import (
    "fmt"
    "strconv"
    "strings"
    "log"
)

const (
    sumMaxBits = 3 // Maximum bits capacity of summed numbers
)

// Decimal to binary form
func dec2Bin(i int) string {
    return fmt.Sprintf("%0" + strconv.Itoa(sumMaxBits) + "b", i)
}

// Bytes to ints
func bytes2Ints(bb []byte) []int {
    ii := make([]int, 0, len(bb))
    for _, b := range bb {
        i, _ := strconv.Atoi(string(b))
        ii = append(ii, i)
    }
    return ii
}

// Ints[bin] to int
func ints2Int(s []int) int {
    iStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ""), "[]")
    iInt, _ := strconv.ParseInt(iStr, 2, 64)
    return int(iInt)
}

// Revert ints[bin]
func revert(s []int) { 
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    
    aMax := 1 << (sumMaxBits - 1)
    bMax := (1 << (sumMaxBits - 1)) - 1
    
    log.Printf("\tNum A max: %v, num B max: %v", aMax, bMax)
    
    for a := 0; a <= aMax; a++ {
            
        for b := 0; b <= bMax; b++ {
            
            log.Printf("\t%v (%v) + %v (%v) -> %v (%v)\n", 
                       dec2Bin(a), a, dec2Bin(b), b, dec2Bin(a + b), a + b)
            
            aBin := bytes2Ints([]byte(dec2Bin(a)) )
            revert(aBin)
            
            bBin := bytes2Ints([]byte(dec2Bin(b)) )
            revert(bBin)
            
            sum := make([]int, 0, sumMaxBits)
            mov := 0 // Carry bit
            for i := 0; i < sumMaxBits; i++ {
                xor1 := (aBin[i] ^ bBin[i])
                xor2 := (xor1 ^ mov)
                sum = append(sum, xor2)
                
                if (aBin[i] != 0 && bBin[i] != 0) ||
                    (xor1 != 0 && mov != 0) {
                    mov = 1
                } else {
                    mov = 0
                }
            }
            revert(sum)
            
            // Check
            if s := ints2Int(sum); s != a + b {
                log.Fatalf("sum: %v != %v\n", s, a + b)
            }
        }
    }
    
    log.Printf("\tOK\n")
}
