////COURSE 1

// package main

// import "fmt"

//	func main() {
//		x := 42
//		y := "James Bond"
//		z := true
//		fmt.Println(x)
//		fmt.Println(y)
//		fmt.Println(z)
//	}
// package main

// import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

// func main() {

// 	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
// 	fmt.Println(s)
// }

// package main

// import "fmt"

// type dhvani int

// var x dhvani
// var y int

// func main() {
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	x = 42
// 	fmt.Println(x)
// 	y = int(x)
// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)
// }

// package main

// import "fmt"

// func main() {
// 	x := 42
// 	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := (4 == 3)
// 	b := (5 <= 10)
// 	c := (2 >= 4)
// 	d := (0 != 2)
// 	e := (4 < 2)
// 	f := (10 > 8)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Println(f)
// }

// package main

// import "fmt"

// func main() {
// 	a := 2
// 	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
// 	b := a << 1
// 	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
// }

// package main

// import "fmt"

// func main() {
// 	a := "Hello my name is Dhvani Shah. I am a Software Engineer at Searce. I am learning `GOLANG`"
// 	b := `Hello my name is Dhvani Shah. I am a Software Engineer at Searce. I am learning "GOLANG"`
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import "fmt"

// const (
// 	a = 2019 + iota
// 	b = 2019 + iota
// 	c = 2019 + iota
// 	d = 2019 + iota
// )

// func main() {
// 	fmt.Println(a, b, c, d)
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10000; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for i := 65; i <= 90; i++ {
// 		fmt.Println(i)
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("%#U\n", i)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	by := 2001
// 	for by <= 2023 {
// 		fmt.Println(by)
// 		by++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	by := 2001
// 	for {
// 		if by <= 2023 {
// 			fmt.Println(by)
// 			by++
// 		}
// 	}
// }

// package main

// import "fmt"

//	func main() {
//		for i := 10; i <= 100; i++ {
//			fmt.Printf("%v modulus 4 %v\n", i, i%4)
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var x [5]int
// 	x[0] = 2
// 	x[1] = 4
// 	x[2] = 6
// 	x[3] = 8
// 	x[4] = 10

//		for _, xval := range x {
//			fmt.Println(xval)
//		}
//		fmt.Printf("%T\n", x)
//	}
// package main

// import "fmt"

// func main() {
// 	x := [5]int{2, 4, 6, 8, 10}
// 	for _, xval := range x {
// 		fmt.Println(xval)
// 	}
// 	fmt.Printf("%T\n", x)
// }

// package main

// import "fmt"

// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	fmt.Println(x[:5])
// 	fmt.Println(x[5:])
// 	fmt.Println(x[2:7])
// 	fmt.Println(x[1:6])
// }

// package main

// import "fmt"

// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	x = append(x, 52)
// 	fmt.Println(x)
// 	x = append(x, 53, 54, 55)
// 	fmt.Println(x)
// 	y := []int{56, 57, 58, 59, 60}
// 	x = append(x, y...)
// 	fmt.Println(x)
// }

// package main

// import "fmt"

// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	x = append(x[:3], x[6:]...)
// 	fmt.Println(x)
// }

// package main

// import "fmt"

// func main() {
// 	x := []string{"James", "Bond", "Shaken, not stirred"}
// 	y := []string{"Miss", "Moneypenny", "Helloooooo, James"}
// 	z := [][]string{x, y}
// 	fmt.Println(z)
// 	for _, zval := range z {
// 		fmt.Println(zval)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	m := map[string][]string{
// 		"bond_james":      []string{"Shaken,not stirred", "Martinis", "Women"},
// 		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
// 		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
// 	}
// 	for _, val := range m {
// 		fmt.Println(val)
// 		for _, mval := range val {
// 			fmt.Println(mval)
// 		}
// 	}
// 	fmt.Println(m)
// }

// package main

// import "fmt"

// func main() {
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken,not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}
// 	m[`Dhvani`] = []string{`chips`, `soda`, `chocolates`}
// 	for _, val := range m {
// 		fmt.Println(val)
// 		for _, mval := range val {
// 			fmt.Println(mval)
// 		}
// 	}
// 	fmt.Println(m)
// }

// package main

// import "fmt"

// func main() {
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken,not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}
// 	m[`Dhvani`] = []string{`chips`, `soda`, `chocolates`}

// 	delete(m, `no_dr`)
// 	for _, val := range m {
// 		fmt.Println(val)
// 		for _, mval := range val {
// 			fmt.Println(mval)
// 		}
// 	}
// 	fmt.Println(m)
// }

// package main

// import "fmt"
// 	type person struct {
// 		first   string
// 		last    string
// 		flavors []string
// 	}
// func main() {

// 	p1 := person{
// 		first:   "Dhvani",
// 		last:    "Shah",
// 		flavors: []string{"Chocolate", "BubleGum"},
// 	}
// 	p2 := person{
// 		first:   "Abhishek",
// 		last:    "Pranis",
// 		flavors: []string{"Chocolate", "Strawberry"},
// 	}
// 	m := map[string]person{
// 		p1.last: p1,
// 		p2.last: p2,
// 	}
// 	fmt.Println(m)

// 	for _, mval := range m {

// 		fmt.Println(mval)

// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p1.first)
// 	fmt.Println(p1.last)
// 	for _, val := range p1.flavors {
// 		fmt.Println(val)
// 	}
// 	fmt.Println(p2)
// 	fmt.Println(p2.first)
// 	fmt.Println(p2.last)
// 	for _, val := range p2.flavors {
// 		fmt.Println(val)
// 	}
// }

// package main

// import "fmt"

// type vehicle struct {
// 	doors int
// 	color string
// }
// type sedan struct {
// 	vehicle
// 	luxury bool
// }
// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// func main() {
// 	t := truck{
// 		vehicle: vehicle{
// 			doors: 2,
// 			color: "red",
// 		},
// 		fourWheel: true,
// 	}
// 	s := sedan{
// 		vehicle: vehicle{
// 			doors: 4,
// 			color: "black",
// 		},
// 		luxury: true,
// 	}
// 	fmt.Println(t)
// 	fmt.Println(t.doors)
// 	fmt.Println(s)
// 	fmt.Println(s.color)
// }

// package main

// import "fmt"

// func main() {
// 	a := struct {
// 		first   string
// 		age     int
// 		flavors []string
// 	}{
// 		first:   "Dhvani",
// 		age:     21,
// 		flavors: []string{"choco", "vanilla"},
// 	}
// 	fmt.Println(a)
// 	for _, val := range a.flavors {
// 		fmt.Println(val)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	x := foo()
// 	y, z := bar()
// 	fmt.Println(x, y, z)
// }
// func foo() int {
// 	return 5
// }
// func bar() (int, string) {
// 	return 25, "Dhvani"
// }

// package main

// import "fmt"

// func main() {
// 	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	x := foo(ii...)
// 	iii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	y := bar(iii)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }
// func foo(xi ...int) int {
// 	sum := 0
// 	for _, val := range xi {
// 		sum = sum + val
// 	}
// 	return sum
// }
// func bar(xi []int) int {
// 	sum := 0
// 	for _, val := range xi {
// 		sum = sum + val
// 	}
// 	return sum
// }

// package main

// import "fmt"

// func main() {
// 	defer foo()
// 	fmt.Println("Dhvani Shah")

// }
// func foo() {
// 	fmt.Println("foo ran")
// }

// package main

// import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// func (p person) speak() {
// 	fmt.Println("My name is: ", p.first, " and my age is: ", p.age)
// }
// func main() {
// 	p := person{
// 		first: "Dhvani",
// 		last:  "Shah",
// 		age:   21,
// 	}
// 	p.speak()
// }

// package main

// import "fmt"

// func main() {
// 	x := 25
// 	fmt.Println(x)
// 	fmt.Println(&x)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// }

// func main() {
// 	p1 := person{
// 		name: "Dhvani Shah",
// 	}
// 	fmt.Println(p1)
// 	changeMe(&p1)
// 	fmt.Println(p1)
// }

// func changeMe(p *person) {
// 	p.name = "Ved Nimavat"s
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type user struct {
// 	First string
// 	Age   int
// }

// func main() {
// 	u1 := user{
// 		First: "Dhvani",
// 		Age:   21,
// 	}

// 	u2 := user{
// 		First: "Ved",
// 		Age:   22,
// 	}
// 	users := []user{u1, u2}
// 	fmt.Println(users)
// 	bs, err := json.Marshal(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bs))
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {

// 	fmt.Println("begin cpu: ", runtime.NumCPU())
// 	fmt.Println("begin gs: ", runtime.NumGoroutine())
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		fmt.Println("Hello routine 1")
// 		wg.Done()
// 	}()
// 	go func() {
// 		fmt.Println("Hello routine 2")
// 		wg.Done()
// 	}()
// 	fmt.Println("mid cpu: ", runtime.NumCPU())
// 	fmt.Println("mid gs: ", runtime.NumGoroutine())
// 	wg.Wait()
// 	fmt.Println("about to exit")
// 	fmt.Println("end cpu: ", runtime.NumCPU())
// 	fmt.Println("end gs: ", runtime.NumGoroutine())
// }

// package main

// import "fmt"

// type person struct {
// 	first string
// }
// type human interface {
// 	speak()
// }

// func (p *person) speak() {
// 	fmt.Println("Hello, I am speaking")
// }
// func saySomething(h human) {
// 	h.speak()
// }
// func main() {
// 	p1 := person{
// 		first: "Dhvani",
// 	}
// 	saySomething(&p1)
// 	p1.speak()
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	inc := 0
// 	gs := 100
// 	wg.Add(gs)
// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			v := inc
// 			runtime.Gosched()
// 			v++
// 			inc = v
// 			fmt.Println(inc)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", inc)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	var m sync.Mutex
// 	inc := 0
// 	gs := 100
// 	wg.Add(gs)
// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			m.Lock()
// 			v := inc
// 			v++
// 			inc = v
// 			m.Unlock()
// 			fmt.Println(inc)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", inc)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	var inc int64
// 	gs := 100
// 	wg.Add(gs)
// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			atomic.AddInt64(&inc, 1)
// 			fmt.Println(atomic.LoadInt64(&inc))
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", inc)
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

//	func main() {
//		fmt.Println(runtime.GOOS, runtime.GOARCH)
//	}
// package main

// import "fmt"

// func main() {
// 	c := make(chan int)
// 	go func() {
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)

// }
// package main

// import "fmt"

//	func main() {
//		cs := make(chan int)
//		go func() {
//			cs <- 42
//		}()
//		fmt.Println(<-cs)
//		fmt.Printf("cs\t%T\n", cs)
//	}
// package main

// import "fmt"

// func main() {
// 	c := gen()
// 	receive(c)
// 	fmt.Println("about to exit")
// }
// func receive(c <-chan int) {
// 	for v := range c {
// 		fmt.Println(v)
// 	}

// }
// func gen() <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 	}()

//		return c
//	}
// package main

// import "fmt"

//	func main() {
//		q := make(chan int)
//		c := gen(q)
//		receive(c, q)
//		fmt.Println("about to exit")
//	}
//
//	func receive(c, q <-chan int) {
//		for {
//			select {
//			case v := <-c:
//				fmt.Println(v)
//			case <-q:
//				return
//			}
//		}
//	}
//
//	func gen(q <-chan int) <-chan int {
//		c := make(chan int)
//		go func() {
//			for i := 0; i < 100; i++ {
//				c <- i
//			}
//			close(c)
//		}()
//		return c
//	}
// package main

// import "fmt"

// func main() {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("about to exit")

// }
// package main

// import "fmt"

// func main() {
// 	c := make(chan int)
// 	for j := 0; j < 10; j++ {
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				c <- i
// 			}

// 		}()
// 	}
// 	for k := 0; k < 100; k++ {
// 		fmt.Println(<-c)
// 	}
// 	fmt.Println("about to exit")

// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

//	func main() {
//		p1 := person{
//			First:   "James",
//			Last:    "Bond",
//			Sayings: []string{"Shaken,not stirred", "Never say Never"},
//		}
//		bs, err := json.Marshal(p1)
//		if err != nil {
//			log.Fatal("Json did not marshal", err)
//		}
//		fmt.Println(string(bs))
//	}
// package main

// import "fmt"

// type customErr struct {
// 	info string
// }

// func (ce customErr) Error() string {
// 	return fmt.Sprintf("here is the error: %v", ce.info)
// }
// func main() {
// 	c1 := customErr{
// 		info: "need more coffee",
// 	}
// 	foo(c1)
// }
// func foo(e error) {
// 	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
// }
