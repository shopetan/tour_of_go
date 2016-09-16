package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64{
    z := float64(1)
    for diff := float64(1); math.Abs(diff) > 1e-10;{
        diff = ((math.Pow(z, 2) - x) / (2.0 * z))
        z = z - diff
    }
    return z
}

func main(){
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
