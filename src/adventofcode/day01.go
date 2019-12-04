package adventofcode

import (
	"math"
)

func Day01Part1() float64 {
	inputs := getInputs()

	var sum float64

	for _, v := range inputs {
		sum += getFuel(v)
	}

	return sum
}

func Day01Part2() float64 {
	inputs := getInputs()

	var sum float64

	sum = 0.0
	for _, v := range inputs {
		sum += getFuelRecursive(v)
	}

	return sum
}

func getFuel(mass float64) float64 {

	fuel := math.Floor(mass/3.0) - 2.0

	if fuel > 0.0 {
		return fuel
	}

	return 0.0
}

func getFuelRecursive(mass float64) float64 {
	fuel := getFuel(mass)

	if fuel > 0.0 {
		return fuel + getFuelRecursive(fuel)
	}

	return 0.0
}

func getInputs() []float64 {

	return []float64{
		120588,
		137636,
		114877,
		118328,
		97394,
		58497,
		139343,
		80307,
		125063,
		70956,
		119676,
		76115,
		91916,
		64618,
		82881,
		57000,
		141785,
		73460,
		68992,
		125701,
		97839,
		137800,
		111051,
		104591,
		114396,
		60210,
		80238,
		112009,
		70265,
		140582,
		58765,
		96848,
		130438,
		55615,
		53903,
		109361,
		129512,
		75888,
		93231,
		54697,
		125320,
		53614,
		87173,
		71762,
		147739,
		131840,
		123979,
		54434,
		121517,
		113518,
		83544,
		124924,
		76608,
		130483,
		149285,
		134147,
		111589,
		88174,
		136392,
		94448,
		139244,
		54064,
		85110,
		102985,
		95646,
		54649,
		129755,
		135795,
		119653,
		147633,
		108386,
		143180,
		126587,
		119273,
		130579,
		56006,
		83232,
		99948,
		147711,
		83092,
		99706,
		98697,
		143231,
		94526,
		53102,
		86002,
		71413,
		111054,
		147220,
		136504,
		59308,
		61123,
		148771,
		113986,
		55483,
		94426,
		62791,
		100959,
		63604,
		112511,
	}

}
