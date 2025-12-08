package day8

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x float64
	y float64
	z float64
}

func (p point) String() string {
	return fmt.Sprintf("{X: %d, Y: %d, Z: %d}", int(p.x), int(p.y), int(p.z))
}

func distanceBetweenPoints(a, b point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
}

type circuit struct {
	points []point
}

func (c circuit) String() string {
	var sb strings.Builder
	for _, i := range c.points {
		sb.WriteString(fmt.Sprintf("%v, ", i.String()))
	}
	output := sb.String()
	return output[:len(output)-2]
}

type pointDifference struct {
	a        point
	b        point
	distance float64
}

func (pd pointDifference) String() string {
	return fmt.Sprintf("a: [%v], b: [%v], distance: %f", pd.a.String(), pd.b.String(), pd.distance)
}

const fileName = "2025/day8/input.txt"
const sampleFileName = "2025/day8/sampleInput.txt"

func ProcessFile() {
	fileToProcess := sampleFileName
	fmt.Println("Processing File ", fileToProcess, "...")
	file, err := os.Open(fileToProcess)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	circuits := make([]circuit, 0)
	pointsToCircuits := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		circuits = append(circuits, circuit{[]point{parseStringToPoint(line)}})
		pointsToCircuits[len(circuits)-1] = len(circuits) - 1
	}

	// Map from index of point -> index of circuit

	// Map from point to circuit?

	for i := 1; i <= 10; i++ {
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
			maxSizes = replaceSmallestInt(maxSizes, len(i.points))
			_, smallestMax = getSmallestInt(maxSizes)
		}
	}
	return productOfInts(maxSizes)
}

func productOfInts(input []int) int {
	result := 1
	for _, i := range input {
		result = result * i
	}
	return result
}

func getSmallestInt(input []int) (minIndex int, min int) {
	min = math.MaxInt
	for index, i := range input {
		if i < min {
			min = i
			minIndex = index
		}
	}
	return
}

func replaceSmallestInt(input []int, new int) []int {
	indexToReplace, _ := getSmallestInt(input)
	input[indexToReplace] = new
	return input
}

func printCircuits(circuits []circuit) {
	fmt.Println("Circuits:")
	for _, c := range circuits {
		fmt.Printf("%v\n", c.String())
	}
}

func joinClosestPointsInCircuits(circuits []circuit) []circuit {
	minDistance := math.Inf(1) //Positiv infinity
	circuitA := -1
	circuitB := -1
	pointA := point{}
	pointB := point{}
	for i, _ := range circuits {
		// in circuit i
		// check each points in this circuit against points in following circuits
		for j, _ := range circuits[i].points { //Each point in circuit i
			for k, _ := range circuits { //Each circuit after circuit i
				for l, _ := range circuits[k].points { //Each point in circuit k
					temp := distanceBetweenPoints(circuits[i].points[j], circuits[k].points[l])
					if temp < minDistance {
						if i != k || j != l { //dont get distance between the same points
							minDistance = temp
							circuitA = i
							circuitB = k
							pointA = circuits[i].points[j]
							pointB = circuits[k].points[l]
						}
					}
				}
			}
		}
	}
	fmt.Printf("Joining %v and %v\n", pointA.String(), pointB.String())
	if circuitA == circuitB {
		return circuits
	}
	// Update circuit A to merge A and B
	circuits[circuitA].points = append(circuits[circuitA].points, circuits[circuitB].points...)
	// Remove circuit B
	circuits = append(circuits[:circuitB], circuits[circuitB+1:]...)
	return circuits
}

// Example format: 162,817,812
func parseStringToPoint(input string) point {
	split := strings.Split(input, ",")
	x, _ := strconv.ParseFloat(split[0], 64)
	y, _ := strconv.ParseFloat(split[1], 64)
	z, _ := strconv.ParseFloat(split[2], 64)
	return point{x, y, z}
}

func calculateDistances(list []point) []pointDifference {
	matrix := make([]pointDifference, 0)
	for i, _ := range list {
		for j := i + 1; j < len(list); j++ {
			matrix = insertToSortedListOfDistances(
				matrix,
				pointDifference{
					list[i],
					list[j],
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
