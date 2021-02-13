package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints arraY
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var X, Y float64

	for idX, val := range vals {

		if idX%2 == 0 {
			X, _ = strconv.ParseFloat(val, 64)
		} else {
			Y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{X, Y})
		}
	}
	return points, nil
}

func onSegment(p, q, r Point) bool {
	if q.X <= math.Max(p.X, r.X) && q.X >= math.Min(p.X, r.X) && q.Y <= math.Max(p.Y, r.Y) && q.Y >= math.Min(p.Y, r.Y) {
		return true
	}

	return false
}

func orientation(p, q, r Point) int {

	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	} else {
		return 2
	}

}

func doIntersect(points []Point) bool {
	if len(points) == 4 {
		o1 := orientation(points[0], points[1], points[2])
		o2 := orientation(points[0], points[1], points[3])
		o3 := orientation(points[2], points[3], points[0])
		o4 := orientation(points[2], points[3], points[1])

		if o1 != o2 && o3 != o4 {
			return true
		}
		if o1 == 0 && onSegment(points[0], points[2], points[1]) {
			return true
		}

		if o2 == 0 && onSegment(points[0], points[3], points[1]) {
			return true
		}

		if o3 == 0 && onSegment(points[2], points[0], points[3]) {
			return true
		}

		if o4 == 0 && onSegment(points[2], points[1], points[3]) {
			return true
		}

		return false
	}
	if len(points) == 5 {
		o1 := orientation(points[0], points[1], points[2])
		o2 := orientation(points[0], points[1], points[3])
		o3 := orientation(points[2], points[3], points[0])
		o4 := orientation(points[2], points[3], points[1])
		o5 := orientation(points[0], points[1], points[4])
		o6 := orientation(points[2], points[3], points[4])

		if o1 != o2 && o3 != o4 && o5 != o6 {
			return true
		}
		if o1 == 0 && onSegment(points[0], points[2], points[1]) {
			return true
		}

		if o2 == 0 && onSegment(points[0], points[3], points[1]) {
			return true
		}

		if o3 == 0 && onSegment(points[2], points[0], points[3]) {
			return true
		}

		if o4 == 0 && onSegment(points[2], points[1], points[3]) {
			return true
		}
		if o5 == 0 && onSegment(points[0], points[4], points[1]) {
			return true
		}
		if o6 == 0 && onSegment(points[2], points[4], points[3]) {
			return true
		}

		return false
	}
	return false
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	// Muchos ifs de 2,3,4,5 vertices para sacar area de cada uno

	if len(points) == 4 {
		return -1 * (((points[0].X*points[1].Y - points[0].Y*points[1].X) + (points[1].X*points[2].Y - points[1].Y*points[2].X) + (points[2].X*points[3].Y - points[2].Y*points[3].X) + (points[3].X*points[0].Y - points[3].Y*points[0].X)) / 2)
	} else if len(points) == 2 {
		return 0.0
	} else if len(points) == 3 {
		return -1 * (((points[0].X*points[1].Y - points[0].Y*points[1].X) + (points[1].X*points[2].Y - points[1].Y*points[2].X) + (points[2].X*points[0].Y - points[2].Y*points[0].X)) / 2)
	} else if len(points) == 5 {
		return -1 * (((points[0].X*points[1].Y - points[0].Y*points[1].X) + (points[1].X*points[2].Y - points[1].Y*points[2].X) + (points[2].X*points[3].Y - points[2].Y*points[3].X) + (points[3].X*points[4].Y - points[3].Y*points[4].X) + points[4].X*points[0].Y - points[4].Y*points[0].X) / 2)
	}
	return 0.0
}

// getPerimeter gets the perimeter from a given arraY of connected points
func getPerimeter(points []Point) float64 {
	if len(points) == 4 {
		return math.Sqrt(math.Pow(points[1].X-points[0].X, 2)+math.Pow(points[1].Y-points[0].Y, 2)) + math.Sqrt(math.Pow(points[2].X-points[1].X, 2)+math.Pow(points[2].Y-points[1].Y, 2)) + math.Sqrt(math.Pow(points[3].X-points[2].X, 2)+math.Pow(points[3].Y-points[2].Y, 2)) + math.Sqrt(math.Pow(points[3].X-points[0].X, 2)+math.Pow(points[3].Y-points[0].Y, 2))
	} else if len(points) == 2 {
		return 0.0
	} else if len(points) == 3 {
		return (math.Sqrt(math.Pow(points[1].X-points[0].X, 2)+math.Pow(points[1].Y-points[0].Y, 2)) + (math.Sqrt(math.Pow(points[2].X-points[1].X, 2) + math.Pow(points[2].Y-points[1].Y, 2))) + (math.Sqrt(math.Pow(points[2].X-points[0].X, 2) + math.Pow(points[2].Y-points[0].Y, 2))))
	} else if len(points) == 5 {
		return math.Sqrt(math.Pow(points[1].X-points[0].X, 2)+math.Pow(points[1].Y-points[0].Y, 2)) + math.Sqrt(math.Pow(points[2].X-points[1].X, 2)+math.Pow(points[2].Y-points[1].Y, 2)) + math.Sqrt(math.Pow(points[3].X-points[2].X, 2)+math.Pow(points[3].Y-points[2].Y, 2)) + math.Sqrt(math.Pow(points[4].X-points[3].X, 2)+math.Pow(points[4].Y-points[3].Y, 2)) + math.Sqrt(math.Pow(points[4].X-points[0].X, 2)+math.Pow(points[4].Y-points[0].Y, 2))
	}
	return 0.0
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	if doIntersect(vertices) == false {
		area := getArea(vertices)
		perimeter := getPerimeter(vertices)

		// Logging in the server side
		log.Printf("Received vertices arraY: %v", vertices)

		// Response construction
		if area == 0 {
			response := fmt.Sprintf("Welcome to the Remote Shapes AnalYzer\n")
			response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
			response += fmt.Sprintf("ERROR - Your shape is not compliYing with the minimum number of vertices.")
			fmt.Fprintf(w, response)
		} else {
			response := fmt.Sprintf("Welcome to the Remote Shapes AnalYzer\n")
			response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
			response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
			response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
			response += fmt.Sprintf(" - Area            : %v\n", area)
			fmt.Fprintf(w, response)
		}

	}

	// Send response to client

}
