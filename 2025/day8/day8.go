package day8

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	utils "github.com/adettinger/adventOfCode-Go/utils"
)

const fileName = "2025/day8/input.txt"
const sampleFileName = "2025/day8/sampleInput.txt"

func ProcessFile() {
	fileToProcess := fileName
	fmt.Println("Processing File ", fileToProcess, "...")
	file, err := os.Open(fileToProcess)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// points := make([]point, 0)
	circuits := make([]circuit, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		// points = append(points, parseStringToPoint(i, line))
		circuits = append(circuits, circuit{[]point{parseStringToPoint(i, line)}})
		i++
	}

	// allDistances := calculateAllDistances(points)

	for i := 1; i <= 1000; i++ {
		fmt.Println()
		fmt.Printf("Round %d...\n", i)
		circuits = joinClosestPointsInCircuits(circuits)
	}
	fmt.Println()
	printCircuits(circuits)
	fmt.Println("Count of circuits: ", len(circuits))
	// fmt.Println("Length of each circuit")
	// for _, i := range circuits {
	// 	fmt.Println(len(i.points))
	// }

	fmt.Printf("Procuct of size of 3 largest circuits: %d\n", multiplythreeLargestCircuitSizes(circuits))
}

func multiplythreeLargestCircuitSizes(circuits []circuit) int {
	maxSizes := make([]int, 3)
	smallestMax := 0
	for _, i := range circuits {
		if len(i.points) > smallestMax {
			maxSizes = utils.ReplaceSmallestInt(maxSizes, len(i.points))
			_, smallestMax = utils.GetSmallestInt(maxSizes)
		}
	}
	return utils.ProductOfInts(maxSizes)
}

// func applyNextSmallestDistance(points []point, distances []pointDifference) ([]point, []pointDifference) {
// 	distance := distances[0]
// 	distances = distances[1:]

// 	points[distance.aIndex].connectedTo = append(points[distance.aIndex].connectedTo, distance.bIndex)
// 	points[distance.bIndex].connectedTo = append(points[distance.bIndex].connectedTo, distance.aIndex)

// 	return points, distances
// }

func joinClosestPointsInCircuits(circuits []circuit) []circuit {
	minDistance := math.Inf(1) //Positiv infinity

	circuitA := -1
	pointA := -1

	circuitB := -1
	pointB := -1

	for i, _ := range circuits {
		// in circuit i
		// check each points in this circuit against points in following circuits
		for j, _ := range circuits[i].points { //Each point in circuit i
			for k, _ := range circuits { //Each circuit after circuit i
				for l, _ := range circuits[k].points { //Each point in circuit k
					temp := distanceBetweenPoints(circuits[i].points[j], circuits[k].points[l])
					if temp < minDistance {
						if i != k || j != l { //dont get distance between the same points
							if !slices.Contains(circuits[i].points[j].connectedTo, circuits[k].points[l].id) {
								minDistance = temp
								circuitA = i
								pointA = j
								circuitB = k
								pointB = l
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Joining %v and %v\n", circuits[circuitA].points[pointA].String(), circuits[circuitB].points[pointB].String())

	//Create connections
	circuits[circuitA].points[pointA].connectedTo = append(circuits[circuitA].points[pointA].connectedTo, circuits[circuitB].points[pointB].id)
	circuits[circuitB].points[pointB].connectedTo = append(circuits[circuitB].points[pointB].connectedTo, circuits[circuitA].points[pointA].id)

	if circuitA == circuitB {
		return circuits
	}

	// MERGE THE CIRCUITS!!!
	// Update circuit A to merge A and B
	circuits[circuitA].points = append(circuits[circuitA].points, circuits[circuitB].points...)
	// Remove circuit B
	circuits = append(circuits[:circuitB], circuits[circuitB+1:]...)
	return circuits
}

// Example format: 162,817,812
func parseStringToPoint(id int, input string) point {
	split := strings.Split(input, ",")
	x, _ := strconv.ParseFloat(split[0], 64)
	y, _ := strconv.ParseFloat(split[1], 64)
	z, _ := strconv.ParseFloat(split[2], 64)
	return point{id, []int{}, x, y, z}
}

func calculateAllDistances(list []point) []pointDifference {
	matrix := make([]pointDifference, 0)
	for i, _ := range list {
		for j := i + 1; j < len(list); j++ {
			matrix = insertToSortedListOfDistances(
				matrix,
				pointDifference{
					list[i].id,
					list[j].id,
					distanceBetweenPoints(list[i], list[j]),
				},
			)
		}
	}
	return matrix
}

func insertToSortedListOfDistances(list []pointDifference, input pointDifference) []pointDifference {
	l := 0
	r := len(list) - 1
	for l < r {
		m := (r + l) / 2
		if list[m].distance < input.distance {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return concatCopyPreAllocate([][]pointDifference{list[:l], {input}, list[l:]})
}

func concatCopyPreAllocate(slices [][]pointDifference) []pointDifference {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]pointDifference, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}
