package main

import ("fmt"
		"math"
	)

func main() {
	var a, s, t, v float64


	fmt.Println("Enter values for acceleration, initial displacement, time, and initial velocity")
	_, err := fmt.Scanf("%f", &a, &s, &t, &v)
	if err != nil {
		fmt.Println(err)
	}
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t));

}

func GenDisplaceFn(a, v, s float64) func(float64) float64{
	return func(t float64) float64 {
		return (0.5 * a) * math.Pow(t, 2) * ( v * t) * s;
	}
}
