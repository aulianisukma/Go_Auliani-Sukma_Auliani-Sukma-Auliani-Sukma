package main

import "fmt"

type Car struct {
    Type    string
    FuelIn  float64
}

func (c *Car) PerkiraanJarak() float64 {
   efesiensiBahanBakar := 1.5 // bahan bakar yang digunakan per mill
    return c.FuelIn * efesiensiBahanBakar
}

func main() {
    myCar := Car{"Sedan", 20.0} // 20.0 liters bahan bakar yang ada di dalam tangki
    distance := myCar.PerkiraanJarak()
    fmt.Printf("estimasi jarak yang ditempuh adalah %.2f mill\n", distance)
}

