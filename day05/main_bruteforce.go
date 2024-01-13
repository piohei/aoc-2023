package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

var seeds []int
var seedToSoil = make(map[int]int)
var soilToFertilizer = make(map[int]int)
var fertilizerToWater = make(map[int]int)
var waterToLight = make(map[int]int)
var lightToTemperature = make(map[int]int)
var temperatureToHumidity = make(map[int]int)
var humidityToLocation = make(map[int]int)

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	currentMapName := ""
	for _, line := range inputAsLines {
		if strings.HasPrefix(line, "seeds: ") {
			s := strings.Replace(line, "seeds: ", "", -1)
			for _, raw := range strings.Split(s, " ") {
				val, _ := strconv.Atoi(raw)
				seeds = append(seeds, val)
			}
			continue
		}

		switch {
			case strings.HasPrefix(line, "seed-to-soil map:"):
				currentMapName = "seedToSoil"
				continue
			case strings.HasPrefix(line, "soil-to-fertilizer map:"):
				currentMapName = "soilToFertilizer"
				continue
			case strings.HasPrefix(line, "fertilizer-to-water map:"):
				currentMapName = "fertilizerToWater"
				continue
			case strings.HasPrefix(line, "water-to-light map:"):
				currentMapName = "waterToLight"
				continue
			case strings.HasPrefix(line, "light-to-temperature map:"):
				currentMapName = "lightToTemperature"
				continue
			case strings.HasPrefix(line, "temperature-to-humidity map:"):
				currentMapName = "temperatureToHumidity"
				continue
			case strings.HasPrefix(line, "humidity-to-location map:"):
				currentMapName = "humidityToLocation"
				continue
			case line == "":
				continue
		}

		ranges := strings.Split(line, " ")
		valueStart, _ := strconv.Atoi(ranges[0])
		keyStart, _ := strconv.Atoi(ranges[1])
		length, _ := strconv.Atoi(ranges[2])

		for i := 0; i < length; i++ {
			put(currentMapName, keyStart + i, valueStart + i)
		}
	}

	printAll()
	var minLocation = 2147483647
	for _, seed := range seeds {
		location := getMapping(seed)
		// fmt.Printf("seed %v -> %v\n", seed, location)
		if location < minLocation {
			minLocation = location
		}
	}
	// fmt.Printf("res = %v\n", minLocation)
}

func printAll() {
	fmt.Printf("%s: %v\n", "seeds", seeds)
	fmt.Printf("%s: %v\n", "seedToSoil", seedToSoil)
	fmt.Printf("%s: %v\n", "soilToFertilizer", soilToFertilizer)
	fmt.Printf("%s: %v\n", "fertilizerToWater", fertilizerToWater)
	fmt.Printf("%s: %v\n", "waterToLight", waterToLight)
	fmt.Printf("%s: %v\n", "lightToTemperature", lightToTemperature)
	fmt.Printf("%s: %v\n", "temperatureToHumidity", temperatureToHumidity)
	fmt.Printf("%s: %v\n", "humidityToLocation", humidityToLocation)
}

func getMapping(key int) int {
	res := key
	// fmt.Printf("seed %v\n", res)
	res = get("seedToSoil", res)
	// fmt.Printf("-> soil %v\n", res)
	res = get("soilToFertilizer", res)
	// fmt.Printf("-> fertilizer %v\n", res)
	res = get("fertilizerToWater", res)
	// fmt.Printf("-> water  %v\n", res)
	res = get("waterToLight", res)
	// fmt.Printf("-> light %v\n", res)
	res = get("lightToTemperature", res)
	// fmt.Printf("-> tempreature %v\n", res)
	res = get("temperatureToHumidity", res)
	// fmt.Printf("-> humidity %v\n", res)
	res = get("humidityToLocation", res)
	// fmt.Printf("-> location %v\n", res)
	return res
}

func put(n string, key, value int) {
	switch n {
		case "seedToSoil":
			seedToSoil[key] = value
		case "soilToFertilizer":
			soilToFertilizer[key] = value
		case "fertilizerToWater":
			fertilizerToWater[key] = value
		case "waterToLight":
			waterToLight[key] = value
		case "lightToTemperature":
			lightToTemperature[key] = value
		case "temperatureToHumidity":
			temperatureToHumidity[key] = value
		case "humidityToLocation":
			humidityToLocation[key] = value
	}
}

func get(n string, key int) int {
	var val int
	var ok bool
	switch n {
		case "seedToSoil":
			val, ok = seedToSoil[key]
		case "soilToFertilizer":
			val, ok = soilToFertilizer[key]
		case "fertilizerToWater":
			val, ok = fertilizerToWater[key]
		case "waterToLight":
			val, ok = waterToLight[key]
		case "lightToTemperature":
			val, ok = lightToTemperature[key]
		case "temperatureToHumidity":
			val, ok = temperatureToHumidity[key]
		case "humidityToLocation":
			val, ok = humidityToLocation[key]
	}
	if ok {
		return val
	} else {
		return key
	}
}