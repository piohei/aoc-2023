package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "sort"
)

//go:embed input
var input string

var seeds []Range
var seedToSoil []Range
var soilToFertilizer []Range
var fertilizerToWater []Range
var waterToLight []Range
var lightToTemperature []Range
var temperatureToHumidity []Range
var humidityToLocation []Range

type Range struct {
	start, end, delta int64
}

func (r Range) Contains(x int64) bool {
	return r.start <= x && x <= r.end
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

const MAX_INT int64 = 9223372036854775807

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	currentMapName := ""
	for _, line := range inputAsLines {
		if strings.HasPrefix(line, "seeds: ") {
			s := strings.Replace(line, "seeds: ", "", -1)
			rawNumbers := strings.Split(s, " ")
			for i := 0; i < len(rawNumbers); i = i + 2 {
				start, _ := strconv.ParseInt(rawNumbers[i], 10, 64)
				length, _ := strconv.ParseInt(rawNumbers[i + 1], 10, 64)
				seeds = append(seeds, Range{start: start, end: start + length - 1})
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
		valueStart, _ := strconv.ParseInt(ranges[0], 10, 64)
		keyStart, _ := strconv.ParseInt(ranges[1], 10, 64)
		length, _ := strconv.ParseInt(ranges[2], 10, 64)

		put(currentMapName, Range{start: keyStart, end: keyStart + length - 1, delta: valueStart - keyStart})
	}

	sort.Sort(ByStart(seeds))
	sort.Sort(ByStart(seedToSoil))
	sort.Sort(ByStart(soilToFertilizer))
	sort.Sort(ByStart(fertilizerToWater))
	sort.Sort(ByStart(waterToLight))
	sort.Sort(ByStart(temperatureToHumidity))
	sort.Sort(ByStart(humidityToLocation))

	var currentRanges []Range
	// currentRanges = append(seeds, Range{start: 0, end: seeds[0].start - 1}, Range{start: seeds[len(seeds)-1].end + 1, end: MAX_INT})
	currentRanges = seeds
	sort.Sort(ByStart(currentRanges))
	fmt.Printf("currentRanges = %v\n", currentRanges)

	// {
	// 	var resRanges []Range
	// 	for _, r := range seedToSoil {
	// 		for _, cr := range currentRanges {
	// 			var s1, e2, s2, e2, s3, e3 int64
	// 			if cr.Contains(r.start) {
	// 				if cr.Contains(r.end) {
	// 					s1 = cr.start
	// 					e1 = s.start - 1
	// 					if s1 <= e1 {
	// 						resRanges = append(resRanges, Range{start: s1, end: e1})
	// 					}
	// 					s2 = s.start
	// 					e2 = s.end
	// 					if s2 <= e2 {
	// 						resRanges = append(resRanges, Range{start: s2 + r.delta, end: e2 + r.delta})
	// 					}
	// 					s3 = s.end + 1
	// 					e3 = cr.end
	// 					if s3 <= e3 {
	// 						resRanges = append(resRanges, Range{start: s3, end: e3})
	// 					}
	// 				} else {
	// 					s1 = cr.start
	// 					e1 = s.start - 1
	// 					if s1 <= e1 {
	// 						resRanges = append(resRanges, Range{start: s1, end: e1})
	// 					}
	// 					s2 = s.start
	// 					e2 = cr.end
	// 					if s2 <= e2 {
	// 						resRanges = append(resRanges, Range{start: s2 + r.delta, end: e2 + r.delta})
	// 					}
	// 				}
	// 			} else {
	// 				if cr.Contains(r.end) {
	// 					s1 = cr.start
	// 					e1 = r.end
	// 					if s1 <= e1 {
	// 						resRanges = append(resRanges, Range{start: s1 + r.delta, end: e1 + r.delta})
	// 					}
	// 					s2 = r.end + 1
	// 					e2 = cr.end
	// 					if s1 <= e1 {
	// 						resRanges = append(resRanges, Range{start: s2, end: e2})
	// 					}
	// 				} else {
	// 					continue	
	// 				}
	// 			}
	// 		}
	// 	}
	// 	currentRanges = resRanges
	// }
	fmt.Printf("seedToSoil = %v\n", seedToSoil)
	currentRanges = modifyRanges(currentRanges, seedToSoil)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("soilToFertilizer = %v\n", soilToFertilizer)
	currentRanges = modifyRanges(currentRanges, soilToFertilizer)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("fertilizerToWater = %v\n", fertilizerToWater)
	currentRanges = modifyRanges(currentRanges, fertilizerToWater)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("waterToLight = %v\n", waterToLight)
	currentRanges = modifyRanges(currentRanges, waterToLight)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("lightToTemperature = %v\n", lightToTemperature)
	currentRanges = modifyRanges(currentRanges, lightToTemperature)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("temperatureToHumidity = %v\n", temperatureToHumidity)
	currentRanges = modifyRanges(currentRanges, temperatureToHumidity)
	fmt.Printf("currentRanges = %v\n", currentRanges)
	fmt.Printf("humidityToLocation = %v\n", humidityToLocation)
	currentRanges = modifyRanges(currentRanges, humidityToLocation)
	fmt.Printf("currentRanges = %v\n", currentRanges)

	var minLocation int64 = MAX_INT
	for _, r := range currentRanges {
		if r.start < minLocation {
			minLocation = r.start
		}
	}

	fmt.Printf("minLocation = %v\n", minLocation)


	// for _, r := range seeds {

	// }

	// var allRanges []Range

	// var allRanges []Range
	// allRanges = append(allRanges, seeds...)
	// allRanges = append(allRanges, seedToSoil...)
	// allRanges = append(allRanges, soilToFertilizer...)
	// allRanges = append(allRanges, fertilizerToWater...)
	// allRanges = append(allRanges, waterToLight...)
	// allRanges = append(allRanges, temperatureToHumidity...)
	// allRanges = append(allRanges, humidityToLocation...)

	// var points []int64
	// for _, r := range allRanges {
	// 	points = append(points, r.start, r.end)
	// }

	// // printAll()
	// var minLocation int64 = MAX_INT
	// var minIndex int64 = -1
	// for _, i := range points {
	// 	location := getMapping(i, false)
	// 	// fmt.Printf("seed %v -> %v\n", i, location)
	// 	if location < minLocation {
	// 		minLocation = location
	// 		minIndex = i
	// 	}
	// }
	// getMapping(minIndex, true)
	// fmt.Printf("res = %v (index %v)\n", minLocation, minIndex)
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

func getMapping(key int64, log bool) int64 {
	res := key
	if log {
		fmt.Printf("seed %v\n", res)
	}
	res = get("seedToSoil", res)
	if log {
		fmt.Printf("-> soil %v\n", res)
	}
	res = get("soilToFertilizer", res)
	if log {
		fmt.Printf("-> fertilizer %v\n", res)
	}
	res = get("fertilizerToWater", res)
	if log {
		fmt.Printf("-> water  %v\n", res)
	}
	res = get("waterToLight", res)
	if log {
		fmt.Printf("-> light %v\n", res)
	}
	res = get("lightToTemperature", res)
	if log {
		fmt.Printf("-> tempreature %v\n", res)
	}
	res = get("temperatureToHumidity", res)
	if log {
		fmt.Printf("-> humidity %v\n", res)
	}
	res = get("humidityToLocation", res)
	if log {
		fmt.Printf("-> location %v\n", res)
	}
	return res
}

func put(n string, r Range) {
	switch n {
		case "seedToSoil":
			seedToSoil = append(seedToSoil, r)
		case "soilToFertilizer":
			soilToFertilizer = append(soilToFertilizer, r)
		case "fertilizerToWater":
			fertilizerToWater = append(fertilizerToWater, r)
		case "waterToLight":
			waterToLight = append(waterToLight, r)
		case "lightToTemperature":
			lightToTemperature = append(lightToTemperature, r)
		case "temperatureToHumidity":
			temperatureToHumidity = append(temperatureToHumidity, r)
		case "humidityToLocation":
			humidityToLocation = append(humidityToLocation, r)
	}
}

func get(n string, key int64) int64 {
	switch n {
		case "seedToSoil":
			for _, r := range seedToSoil {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "soilToFertilizer":
			for _, r := range soilToFertilizer {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "fertilizerToWater":
			for _, r := range fertilizerToWater {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "waterToLight":
			for _, r := range waterToLight {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "lightToTemperature":
			for _, r := range lightToTemperature {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "temperatureToHumidity":
			for _, r := range temperatureToHumidity {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
		case "humidityToLocation":
			for _, r := range humidityToLocation {
				if r.start <= key && key <= r.end {
					return key + r.delta
				}
			}
	}
	return key
}

func modifyRanges(currentRanges, modifyRanges []Range) []Range {
	var resRanges []Range
	for _, cr := range currentRanges {
		// fmt.Printf("cr = %v\n", cr)
		subranges := filterRanges(modifyRanges, cr)
		// fmt.Printf("filtered = %v\n", subranges)
		if len(subranges) > 0 {
			if cr.start < subranges[0].start {
				subranges = append(subranges, Range{start: cr.start, end: subranges[0].start - 1})
			}
			if subranges[len(subranges) - 1].end < cr.end {
				subranges = append(subranges, Range{start: subranges[len(subranges) - 1].end + 1, end: cr.end})
			}
		} else {
			subranges = append(subranges, Range{start: cr.start, end: cr.end})
		}
		// fmt.Printf("subranges = %v\n", subranges)
		for _, r := range subranges {
			resRanges = append(resRanges, Range{start: r.start + r.delta, end: r.end + r.delta})
		}
	}
	return resRanges
}

func filterRanges(ranges []Range, overlappingRange Range) []Range {
	var res []Range
	for _, r := range ranges {
		// fmt.Printf("checking overlappingRange = %v, r = %v\n", overlappingRange, r)
		if overlappingRange.Contains(r.start) {
			if overlappingRange.Contains(r.end) {
				res = append(res, Range{start: r.start, end: r.end, delta: r.delta})
			} else {
				res = append(res, Range{start: r.start, end: overlappingRange.end, delta: r.delta})
			}
		} else if overlappingRange.Contains(r.end) {
			if !overlappingRange.Contains(r.start) {
				res = append(res, Range{start: overlappingRange.start, end: r.end, delta: r.delta})
			}
		} else if r.start <= overlappingRange.start && overlappingRange.end <= r.end {
			res = append(res, Range{start: overlappingRange.start, end: overlappingRange.end, delta: r.delta})
		}
	}
	return res
}