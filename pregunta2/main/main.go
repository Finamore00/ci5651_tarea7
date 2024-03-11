package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// Tipo de datos punto
type Point struct {
	x int
	y int
}

/*
Representación como string del tipo de datos punto.
*/
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func sqDist(p Point, q Point) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
}

func orientation(p Point, q Point, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	} else {
		return 2
	}
}

// Tipo de datos pila y definición de métodos sobre la pila
type PointStack []Point

func (s *PointStack) pop() (Point, error) {
	if len(*s) == 0 {
		return Point{-1, -1}, errors.New("pila vacia")
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1] //Remover último elemento
	return val, nil
}

func (s *PointStack) secondToTop() Point {
	p := s.peek()
	s.pop()
	ret := s.peek()
	s.push(p)
	return ret
}

func (s PointStack) peek() Point {
	return s[len(s)-1]
}

func (s *PointStack) push(p Point) {
	(*s) = append((*s), p)
}

/*
Implementación del algoritmo GrahamScan para el cálculo de un Convex Hull
de un conjunto de puntos. Recibe como entrada una lista de puntos y retorna
un slice con los puntos que conforman el Convex Hull del conjunto.
*/
func grahamScan(pointSet []Point) []Point {

	if len(pointSet) < 3 {
		return []Point{}
	}

	s := PointStack{} //Pila vacía

	//Se busca el elemento más bajo del conjunto
	lowestPoint := Point{math.MinInt, math.MaxInt}
	minIndex := -1
	for i, e := range pointSet {
		if e.y < lowestPoint.y {
			lowestPoint = e
			minIndex = i
		} else if e.y == lowestPoint.y {
			if e.x > lowestPoint.x {
				lowestPoint = e
				minIndex = i
			}
		}
	}

	//Se mueve el elemento más bajo al inicio de la lista
	tmp := lowestPoint
	pointSet[minIndex] = pointSet[0]
	pointSet[0] = tmp

	sort.Slice(pointSet[1:], func(i, j int) bool {
		p := pointSet[i+1]
		q := pointSet[j+1]

		o := orientation(lowestPoint, p, q)
		if o == 0 {
			return sqDist(lowestPoint, q) >= sqDist(lowestPoint, p)
		}

		return o == 2
	})

	//Remover elementos colineares dentro del conjunto, manteniendo solo el más lejano a lowestPoint
	filtSet := []Point{}
	filtSet = append(filtSet, pointSet[0])

	for i := 1; i < len(pointSet)-1; i++ {
		if orientation(lowestPoint, pointSet[i], pointSet[i+1]) == 0 {
			continue
		}
		filtSet = append(filtSet, pointSet[i])
	}

	if orientation(lowestPoint, pointSet[len(pointSet)-2], pointSet[len(pointSet)-1]) != 0 {
		filtSet = append(filtSet, pointSet[len(pointSet)-1])
	}

	if len(filtSet) < 3 {
		return []Point{}
	}

	s.push(lowestPoint)
	s.push(filtSet[1])
	s.push(filtSet[2])

	for i := 3; i < len(filtSet); i++ {
		for len(s) > 1 && orientation(s.secondToTop(), s.peek(), filtSet[i]) != 2 {
			s.pop()
		}
		s.push(filtSet[i])
	}

	return s
}

func main() {
	pointSet := []Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6},
		{0, 6}, {1, 5}, {2, 4}, {4, 2}, {5, 1}, {6, 0}}

	var currentLayer []Point
	counter := 0

	for len(pointSet) > 1 {
		currentLayer = grahamScan(pointSet)
		counter += 1
		newSet := []Point{}
		for _, e1 := range pointSet {
			foundFlag := false
			for _, e2 := range currentLayer {
				if e1.x == e2.x && e1.y == e2.y {
					foundFlag = true
					break
				}
			}
			if !foundFlag {
				newSet = append(newSet, e1)
			}
		}
		pointSet = newSet
	}

	fmt.Println(counter)
}
