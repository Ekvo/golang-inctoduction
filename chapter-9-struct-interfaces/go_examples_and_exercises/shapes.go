/*
Chapter 9: structs and interfaces

in this place:
1 interfaces: Shape
5 structs: Point, Square,Rectangle,Circle and MultiShapes

Point       is member to Square,Rectangle,Circle
interface   is store a member functions
MultiShapes is container for Shapes

interface: Shape have are -

	            area() float64
		        perimeter() float64
		        center() Point
		        north() Point
		        south() Point
		        east() Point
		        west() Point

struct: Square have a unic member function - newSquare(,,,)*
struct: Rectangle have a unic member function - newRectangle(,,,,)*
struct: Circle have a unic member function - newCircle(,,,)*

func property(shapes ...Shape)(void)        - print the all about Shape
func shapes_load(ms *MultiShapes)(void)     - add shape to container (MultiShapes) if MultiShapes!=nil
func shapes_property(ms *MultiShapes)(void) - print container (MultiShapes) data
func program(path int)(void)                - for call on main
func program_with_pointers()(void)          - work with (MultiShapes)
func program_without_pointers()(void)       - work with (*MultiShapes)

func catch_block()(void)                    - for catch exceptions
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

const ( //const string variables for type check  when we create a shape
	str_square    = "square"
	str_rectangle = "rectangle"
	str_circle    = "circle"
	str_unknown   = "unknown type"
	str_error     = "incorrect type"
	str_panic     = "!!! PANIC !!!"
)

// ------------------------------------ Point -----------------------------------------
type Point struct {
	x, y float64
}

// ------------------------------------ Point - END -----------------------------------
// ------------------------------------ Shape interface -------------------------------
type Shape interface {
	area() float64
	perimeter() float64
	center() Point
	north() Point
	south() Point
	east() Point
	west() Point
}

// -------------------------------------------------------------------------------------
type MultiShapes struct {
	shapes []Shape
}

func property(shapes ...Shape) {
	for _, num_sh := range shapes {
		fmt.Println("Type:", num_sh,
			"\nArea:", num_sh.area(),
			"\nPerimetr:", num_sh.perimeter(),
			"\nCenter:", num_sh.center(),
			"\nNorth:", num_sh.north(),
			"\nSouth:", num_sh.south(),
			"\nEast:", num_sh.east(),
			"\nWest:", num_sh.west(),
		)
	}
}

// ------------------------------------ Shape interface - END -------------------------
// ------------------------------------ Square ----------------------------------------
type Square struct {
	type_name string //for check_type
	p         Point  //left up point
	side      float64
}

func newSquare(str_type string, coord Point, s_len float64) (ptr_s *Square) {
	if str_type == str_square {
		if s_len > 0 {
			ptr_s = new(Square)
			ptr_s.type_name = str_type
			ptr_s.p = coord
			ptr_s.side = s_len
			return
		}
		panic(str_panic + " : create " + str_square + " with negativ side " + fmt.Sprintf("%f", s_len))
	}
	panic(str_panic + " : create " + str_square + " with type " + str_type)
}

func (s Square) area() float64 {
	return s.side * s.side
}
func (s Square) perimeter() float64 {
	return s.side * 4.0
}
func (s Square) center() Point {
	cen := Point{s.p.x + s.side/2.0, s.p.y + s.side/2.0}
	return cen
}
func (s Square) north() Point {
	nor := Point{s.p.x + s.side/2.0, s.p.y}
	return nor
}
func (s Square) south() Point {
	sou := Point{s.p.x + s.side/2.0, s.p.y + s.side}
	return sou
}
func (s Square) east() Point {
	e := Point{s.p.x + s.side, s.p.y + s.side/2.0}
	return e
}
func (s Square) west() Point {
	w := Point{s.p.x, s.p.y + s.side/2.0}
	return w
}

// ------------------------------------ Square - END ----------------------------------
// ------------------------------------ Rectangle -------------------------------------
type Rectangle struct {
	type_name string
	p         Point   //left up point
	w         float64 //widht
	h         float64 //height
}

func newRectangle(str_type string, coord Point, widht float64, height float64) (ptr_s *Rectangle) {
	if str_type == str_rectangle {
		if widht > 0 && height > 0 {
			ptr_s = new(Rectangle)
			ptr_s.type_name = str_type
			ptr_s.p = coord
			ptr_s.w = widht
			ptr_s.h = height
			return
		}
		panic(str_panic + " : create " + str_rectangle +
			" with negativ side: widht = " + fmt.Sprintf("%f", widht) +
			"; height = " + fmt.Sprintf("%f", height))

	}
	panic(str_panic + " : create " + str_rectangle + " with type " + str_type)
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}
func (r Rectangle) perimeter() float64 {
	return (r.w + r.h) * 2.0
}
func (r Rectangle) center() Point {
	cen := Point{r.p.x + r.w/2.0, r.p.y + r.h/2.0}
	return cen
}
func (r Rectangle) north() Point {
	nor := Point{r.p.x + r.w/2.0, r.p.y}
	return nor
}
func (r Rectangle) south() Point {
	sou := Point{r.p.x + r.w/2.0, r.p.y + r.h}
	return sou
}
func (r Rectangle) east() Point {
	e := Point{r.p.x + r.w, r.p.y + r.h/2.0}
	return e
}
func (r Rectangle) west() Point {
	w := Point{r.p.x, r.p.y + r.h/2.0}
	return w
}

// ------------------------------------ Rectangle - END -------------------------------
// ------------------------------------ Circle ----------------------------------------
type Circle struct {
	type_name string
	p         Point   //left up point
	r         float64 //radius
}

func newCircle(str_type string, coord Point, radius float64) (ptr_c *Circle) {
	if str_type == str_circle { //if type good
		if radius > 0 {
			ptr_c = new(Circle)
			ptr_c.type_name = str_type
			ptr_c.p = coord
			ptr_c.r = radius
			return
		}
		panic(str_panic + " : create " + str_circle + " with negativ radius " + fmt.Sprintf("%f", radius))
	}
	panic(str_panic + " : create " + str_circle + " with type " + str_type)
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.r
}
func (c Circle) center() Point {
	cen := Point{c.p.x + c.r, c.p.y + c.r}
	return cen
}
func (c Circle) north() Point {
	nor := Point{c.p.x + c.r, c.p.y}
	return nor
}
func (c Circle) south() Point {
	sou := Point{c.p.x + c.r*2.0, c.p.y + c.r*2.0}
	return sou
}
func (c Circle) east() Point {
	e := Point{c.p.x + c.r*2.0, c.p.y + c.r}
	return e
}
func (c Circle) west() Point {
	w := Point{c.p.x, c.p.y + c.r/2.0}
	return w
}

// ------------------------------------ Circle - END ----------------------------------
// ------------------------------------ Free functions --------------------------------
const loop int = 10 //for fucntion 'shapes_load'
/*
shapes_load(ms *MultiShapes)

take pointer to MultiShapes
load shapes to ms.shapes array
*/
func shapes_load(ms *MultiShapes) {
	if ms == nil {
		panic("MultiShapes pointer is nil")
	}
	var ( //type float
		ox  = 0.0 //this varaibles for comfortable reed code
		oy  = 0.0
		d1  = 0.0
		d2  = 0.0 //need only a rectangle
		tmp = 0.0 //buffer : for solo casting to float64
	)
	for i := 1; i < loop; i++ { //var (i) start on 1, because 0 leads to panic in create shape (radius!=0 etc)
		tmp = float64(i)
		ox = tmp
		oy = tmp
		d1 = tmp
		switch i {
		case 1, 2, 3:
			c := newCircle(str_circle, Point{ox, oy}, d1)
			ms.shapes = append(ms.shapes, c) //load circle
		case 4, 5, 6:
			d2 = tmp
			r := newRectangle(str_rectangle, Point{ox, oy}, d1, d2)
			ms.shapes = append(ms.shapes, r) //load rectangle
		case 7, 8, 9:
			s := newSquare(str_square, Point{ox, oy}, d1)
			ms.shapes = append(ms.shapes, s) //load square
		default:
			panic(str_panic + " test_shapes !step" + strconv.Itoa(i))
		}
	}
}

func shapes_property(ms *MultiShapes) { //print date with help property(shape)
	if ms != nil {
		for _, num_sh := range ms.shapes {
			property(num_sh)
		}
	}
}

func program(path int) { // get result with 'path' equal 0 or 1
	switch path {
	case 0:
		program_with_pointers()
	case 1:
		program_without_pointers()
	case 2:
		panic("Bender says: I think I saw a deuce!!!(futurama)")
	default:
		panic(str_panic + " incorrect path: " + strconv.Itoa(path))
	}
}

func program_with_pointers() {
	var ms *MultiShapes //for test nil. !Comment next lane!
	ms = new(MultiShapes)
	shapes_load(ms)
	shapes_property(ms)
}

func program_without_pointers() {
	ms := MultiShapes{[]Shape{}}
	shapes_load(&ms)
	shapes_property(&ms)
}

func catch_block() { //catch errors
	e_str := recover()
	if e_str != nil {
		fmt.Println(e_str)
	}
}

//------------------------------------ Free functions - END --------------------------
