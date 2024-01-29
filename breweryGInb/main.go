// // // // // package main

// // // // // import (
// // // // // 	"fmt"
// // // // // 	"math"
// // // // // 	"math/rand"
// // // // // 	"time"
// // // // // )

// // // // // type Shape interface {
// // // // // 	Area() float64
// // // // // }

// // // // // type Square struct {
// // // // // 	Side int
// // // // // }

// // // // // type Rectangle struct {
// // // // // 	Width  float64
// // // // // 	Length float64
// // // // // }

// // // // // func (r Rectangle) Area() float64 {
// // // // // 	return r.Length * r.Width
// // // // // }

// // // // // type Users struct {
// // // // // 	FirstName string

// // // // // 	LastName string
// // // // // }

// // // // // func (s Square) Area() float64 {
// // // // // 	return float64(math.Pow(float64(s.Side), 2.0))
// // // // // }

// // // // // func main() {

// // // // // 	// m := map[string]int{
// // // // // 	// 	"sunday":  0,
// // // // // 	// 	"monday":  1,
// // // // // 	// 	"tuesday": 2,
// // // // // 	// }
// // // // // 	// value, ok := m["mondays"]

// // // // // 	// fmt.Println(value, ok)

// // // // // 	// if ok {
// // // // // 	// 	fmt.Println(value)
// // // // // 	// } else {
// // // // // 	// 	fmt.Println(m)
// // // // // 	// }

// // // // // 	// reader := bufio.NewReader(os.Stdin)

// // // // // 	// fmt.Println("Please enter your age")

// // // // // 	// input, _ := reader.ReadString('\n')

// // // // // 	// fmt.Println("Your age is", input)

// // // // // 	// var currentTime = time.Now()

// // // // // 	// fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

// // // // // 	// var birthDate = time.Date(2002, time.March, 18, 17, 30, 0, 0, currentTime.Local().Location())

// // // // // 	// fmt.Println(birthDate.Weekday())

// // // // // 	// var arr = make([]int, 10)

// // // // // 	// var i int

// // // // // 	// var rem int = 1

// // // // // 	// for i = 0; i < 10; i++ {

// // // // // 	// 	arr[i] = rem

// // // // // 	// 	rem += 2
// // // // // 	// }

// // // // // 	// fmt.Println(arr)

// // // // // 	// r := Rectangle{
// // // // // 	// 	Length: 12,
// // // // // 	// 	Width:  8,
// // // // // 	// }

// // // // // 	// s := Square{
// // // // // 	// 	Side: 6,
// // // // // 	// }

// // // // // 	//fmt.Println(r.Area(), s.Area())

// // // // // 	// var slice = []string{}

// // // // // 	// slice = append(slice, "beta", "gama")

// // // // // 	// slice[0] = "alpha"

// // // // // 	// fmt.Println(slice)

// // // // // 	// scores := make([]int, 4)

// // // // // 	// scores[0] = 10

// // // // // 	// scores[1] = 20

// // // // // 	// scores[2] = 30

// // // // // 	// scores[3] = 40

// // // // // 	// scores = append(scores, 50)

// // // // // 	// fmt.Println(scores)

// // // // // 	// courses := []string{"Java", "Python", "C", "Go"}

// // // // // 	// courses = append(courses[:1], courses[2:]...)

// // // // // 	// fmt.Println(courses)

// // // // // 	// user1 := Users{FirstName: "Kamesh Sarvan", LastName: "Bhavaraju"}

// // // // // 	// fmt.Printf("The details of user1 are %+v", user1)

// // // // // 	rand.Seed(time.Now().UnixNano())

// // // // // 	diceNumber := rand.Intn(6) + 1

// // // // // 	switch diceNumber {
// // // // // 	case 1:
// // // // // 		fmt.Println("welcome to game")
// // // // // 	case 6:
// // // // // 		fmt.Printf("Hey you can move forward %v times, And you got chance to roll the dice again", diceNumber)
// // // // // 	default:
// // // // // 		fmt.Printf("Hey you can move forward %v times", diceNumber)

// // // // // 	}
// // // // // }

// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // )

// // // // func display() {

// // // // 	fmt.Println("Enter You choice")

// // // // 	fmt.Println("1.Add\n2.Subtract\n3.Multiply\n4.Divide\n5.Exit")

// // // // 	res := decision()

// // // // 	if res == "entry" {
// // // // 		display()
// // // // 	}

// // // // }

// // // // func decision() string {

// // // // 	var choice int

// // // // 	fmt.Scanln(&choice)

// // // // 	if choice == 5 {
// // // // 		return "end"
// // // // 	}

// // // // 	if choice < 1 || choice > 5 {
// // // // 		fmt.Println("Invalid input")
// // // // 		return "entry"
// // // // 	} else {
// // // // 		calculator(choice)
// // // // 		return "entry"
// // // // 	}

// // // // }
// // // // func takeInput() (int, int) {

// // // // 	fmt.Println("Enter the two numbers :")
// // // // 	var n1, n2 int
// // // // 	fmt.Scanln(&n1, &n2)
// // // // 	return n1, n2

// // // // }

// // // // func add(x, y int) {
// // // // 	fmt.Printf("The addition of %d and %d is %d\n", x, y, x+y)
// // // // }

// // // // func subtract(x, y int) {
// // // // 	fmt.Printf("The subtraction of %d and %d is %d\n", x, y, x-y)
// // // // }

// // // // func multiply(x, y int) {
// // // // 	fmt.Printf("The multiplication of %d and %d is %d\n", x, y, x*y)
// // // // }

// // // // func divide(x, y int) {
// // // // 	fmt.Printf("The division of %d and %d is %d\n", x, y, x/y)
// // // // }

// // // // func operation(op func(int, int), x, y int) {
// // // // 	op(x, y)
// // // // }

// // // // func calculator(choice int) {

// // // // 	var n1, n2 int

// // // // 	n1, n2 = takeInput()

// // // // 	switch choice {
// // // // 	case 1:

// // // // 		operation(add, n1, n2)

// // // // 	case 2:

// // // // 		operation(subtract, n1, n2)

// // // // 	case 3:

// // // // 		operation(multiply, n1, n2)

// // // // 	case 4:

// // // // 		operation(divide, n1, n2)

// // // // 	}
// // // // }

// // // // // CANNOT use ":=" at package scope
// // // // // app_version := "1.0"

// // // // // var app_version string = "1.0"
// // // // //var app_version = "1.0"

// // // // func main() {
// // // // 	/*S
// // // // 	var username string
// // // // 	username = "Magesh"
// // // // 	fmt.Printf("Hi %s, Have a nice day!\n", username)
// // // // 	*/

// // // // 	/*
// // // // 		var username string = "Magesh"
// // // // 		fmt.Printf("Hi %s, Have a nice day!\n", username)
// // // // 	*/

// // // // 	// type inference
// // // // 	/*
// // // // 		var username = "Magesh"
// // // // 		fmt.Printf("Hi %s, Have a nice day!\n", username)
// // // // 	*/

// // // // 	// 	username := "Magesh"
// // // // 	// 	fmt.Printf("Hi %s, Have a nice day!\n", username)

// // // // 	// 	/* unused variable */
// // // // 	// 	/*
// // // // 	// 		var x int
// // // // 	// 		x = 100
// // // // 	// 	*/
// // // // 	// 	// fmt.Println(x)

// // // // 	// 	var (
// // // // 	// 		x, y int = 20, 30

// // // // 	// 		result string = "The sum of %d and %d is : %d\n"
// // // // 	// 	)

// // // // 	// 	fmt.Printf(result, x, y, x+y)

// // // // 	// 	var a, b, res = 40, 60, "The sum of %d and %d is : %d\n"

// // // // 	// 	fmt.Printf(res, a, b, a+b)

// // // // 	// 	var comp complex128 = 56 + 128i

// // // // 	// 	fmt.Println(real(comp), imag(comp))

// // // // 	// 	const (
// // // // 	// 		sunday = (iota + 1) * 2

// // // // 	// 		monday

// // // // 	// 		tuesday
// // // // 	// 	)

// // // // 	// 	fmt.Println(sunday, monday, tuesday)

// // // // 	// 	switch number := 9; {
// // // // 	// 	case number%2 == 0:
// // // // 	// 		fmt.Printf("The Given number %d is even", number)
// // // // 	// 	default:
// // // // 	// 		fmt.Printf("The Given number %d is odd\n", number)
// // // // 	// 	}

// // // // 	// 	if number := 68; number%2 == 0 {
// // // // 	// 		if number < 50 {
// // // // 	// 			fmt.Printf("The number is even and less than 50")
// // // // 	// 		} else {
// // // // 	// 			fmt.Printf("The number is even and greater than 50")
// // // // 	// 		}

// // // // 	// 	} else {
// // // // 	// 		fmt.Printf("The number is odd")
// // // // 	// 	}
// // // // 	// OUTER:
// // // // 	// 	for i := 0; i <= 5; i++ {
// // // // 	// 		for j := 0; j <= 5; j++ {
// // // // 	// 			fmt.Println(i, "*", j)
// // // // 	// 			if i == j {
// // // // 	// 				fmt.Println("***********")
// // // // 	// 				continue OUTER
// // // // 	// 			}
// // // // 	// 		}
// // // // 	// 	}

// // // // 	// OUTER:
// // // // 	// 	for i := 2; i <= 100; i++ {
// // // // 	// 		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
// // // // 	// 			if i%j == 0 {
// // // // 	// 				continue OUTER
// // // // 	// 			}
// // // // 	// 		}
// // // // 	// 		fmt.Println(i)
// // // // 	// 	}

// // // // 	//calculator()

// // // // 	// q, r := func(x, y int) (q, r int) {

// // // // 	// 	q = x / y
// // // // 	// 	r = x % y
// // // // 	// 	return
// // // // 	// }(50, 2)

// // // // 	// fmt.Println(q, r)

// // // // 	// msg := func(firstName, lastName string) string {
// // // // 	// 	return fmt.Sprintf("Hello Welcome %s %s", firstName, lastName)
// // // // 	// }("Kamesh", "Sarvan")

// // // // 	// fmt.Println(msg)

// // // // 	// //var greetUser func(string, string) string

// // // // 	// var greetUser = func(firstName, lastName string) string {
// // // // 	// 	return fmt.Sprintf("Hello Welcome %s %s", firstName, lastName)
// // // // 	// }

// // // // 	// fmt.Println(greetUser("Venkata", "Kamesh"))

// // // // 	// someFunction := func(x, y int) int {
// // // // 	// 	return x + y
// // // // 	// }

// // // // 	// var res = func(x, y int) int {
// // // // 	// 	return x + y
// // // // 	// }(8, 10)

// // // // 	// fmt.Println(res)

// // // // 	// fmt.Println(someFunction(8, 9))

// // // // 	// logOPeration(add, 10, 8)

// // // // 	// logOPeration(subtract, 10, 8)

// // // // 	display()

// // // // }

// // // // // func logOPeration(opera func(int, int) int, x, y int) {

// // // // // 	log.Println("Operation Started...")

// // // // // 	fmt.Println(opera(x, y))

// // // // // 	log.Println("Operation ended....")

// // // // // }

// // // // // func add(x, y int) int {
// // // // // 	return x + y
// // // // // }

// // // // // func subtract(x, y int) int {
// // // // // 	return x - y
// // // // // }

// // // /*
// // // 	Refactor the below using functions
// // // 	Ensure that each function follows SRP
// // // */

// // // // package main

// // // // import (
// // // // 	"errors"
// // // // 	"fmt"
// // // // )

// // // // var ErrF1 = errors.New("f1 error")

// // // // var ErrF2 = errors.New("f2 error")

// // // // func main() {
// // // // 	defer fmt.Println("main completeed")
// // // // 	err := f2()
// // // // 	if errors.Is(err, ErrF1) {
// // // // 		fmt.Println(ErrF1)
// // // // 	}
// // // // 	if errors.Is(err, ErrF2) {
// // // // 		fmt.Println(ErrF2)
// // // // 	}
// // // // }

// // // // func f2() error {
// // // // 	err := f1()
// // // // 	return fmt.Errorf(" %w %w ", ErrF2, err)
// // // // }

// // // // func f1() error {
// // // // 	return ErrF1
// // // // }

// // // // func main() {

// // // // 	var arr []int = []int{1, 2, 3, 4}

// // // // 	defer fmt.Println(arr[0])

// // // // 	arr[0] = 6

// // // // 	fmt.Println(arr[0], "-------")

// // // // }

// // // // func main() {
// // // // 	var (
// // // // 		userChoice, n1, n2, result int
// // // // 		operation                  func(int, int) int
// // // // 		err                        error
// // // // 	)

// // // // 	for {
// // // // 		userChoice, err = getUserChoice()
// // // // 		if err != nil {
// // // // 			fmt.Println(err)
// // // // 		}
// // // // 		n1, n2 = getOperands()
// // // // 		operation = getOperation(userChoice)
// // // // 		result = operation(n1, n2)
// // // // 		fmt.Println("Result :", result)
// // // // 	}
// // // // }

// // // // func getUserChoice() (n int, err error) {

// // // // 	fmt.Println("1. Add")
// // // // 	fmt.Println("2. Subtract")
// // // // 	fmt.Println("3. Multiply")
// // // // 	fmt.Println("4. Divide")
// // // // 	fmt.Println("5. Exit")
// // // // 	fmt.Println("Enter your choice :")
// // // // 	fmt.Scanln(&n)

// // // // 	if n == 5 {
// // // // 		err = errors.New("exit")
// // // // 		return
// // // // 	} else if n < 1 || n > 5 {
// // // // 		err = errors.New("invalid option error")
// // // // 		return
// // // // 	}
// // // // 	return
// // // // }

// // // // func getOperands() (n1, n2 int) {
// // // // 	fmt.Println("Enter the operands :")
// // // // 	fmt.Scanln(&n1, &n2)
// // // // 	return
// // // // }

// // // // func getOperation(userChoice int) func(int, int) int {
// // // // 	switch userChoice {
// // // // 	case 1:
// // // // 		return add
// // // // 	case 2:
// // // // 		return subtract
// // // // 	case 3:
// // // // 		return multiply
// // // // 	case 4:
// // // // 		return divide
// // // // 	default:
// // // // 		return func(i1, i2 int) int { return 0 }
// // // // 	}
// // // // }

// // // // func add(x, y int) int {
// // // // 	return x + y
// // // // }

// // // // func subtract(x, y int) int {
// // // // 	return x - y
// // // // }

// // // // func multiply(x, y int) int {
// // // // 	return x * y
// // // // }

// // // //	func divide(x, y int) int {
// // // //		return x / y
// // // //	}
// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strings"
// // // )

// // // func main() {
// // // 	// var no int = 100
// // // 	// fmt.Println("Before incrementing, no :", no)
// // // 	// increment(&no, 200)
// // // 	// fmt.Println("After incrementing, no :", no)

// // // 	// n1, n2 := 100, 200

// // // 	// var ptr *int = &n1

// // // 	// n3 := ptr

// // // 	// *n3 += 10

// // // 	// fmt.Println(n3, n1)

// // // 	// fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
// // // 	// swap(&n1, &n2)
// // // 	// fmt.Printf("Aftter swapping, n1 = %d and n2 = %d\n", n1, n2)

// // // 	//var x = genPrimes(2, 100)

// // // 	//fmt.Println(x)

// // // 	// for x, y := 10, 10; x < 30; x++ {

// // // 	// 	fmt.Println(x, y)

// // // 	// 	x++

// // // 	// 	if x > 100 {
// // // 	// 		break
// // // 	// 	}

// // // 	// }

// // // 	var s string = "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea tempor Consequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"

// // // 	var words []string = strings.Split(s, " ")

// // // 	var wordmap = make(map[int]int)

// // // 	for _, val := range words {

// // // 		x := len(val)

// // // 		wordmap[x]++

// // // 	}

// // // 	var no, count int = 0, 0

// // // 	for key, val := range wordmap {
// // // 		if val > count {
// // // 			no = key
// // // 			count = val
// // // 		}

// // // 	}

// // // 	fmt.Println(no, count)

// // // }

// // // // func increment(no *int, increment int) /* NO RETURN RESULT */ {
// // // // 	*no += increment
// // // // }

// // // // func swap(n1, n2 *int) /* NO RETURN RESULT */ {

// // // // 	var third = *n2

// // // // 	*n2 = *n1

// // // // 	*n1 = third

// // // // }

// // // // func genPrimes(start, end int) []int {

// // // // 	var primes []int
// // // // 	for i := start; i <= end; i++ {
// // // // 		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
// // // // 			fmt.Println(i, j)
// // // // 			if i%j == 0 {
// // // // 				break
// // // // 			}
// // // // 		}
// // // // 		primes = append(primes, i)
// // // // 	}

// // // // 	return primes

// // // // }

// // // package main

// // // import "fmt"

// // // type Employee struct {
// // // 	Id int

// // // 	Name string

// // // 	Org Organization
// // // }

// // // type Organization struct {
// // // 	Id int

// // // 	Name string

// // // 	City string
// // // }

// // // func main() {

// // // 	organization := Organization{Id: 101, Name: "Thoughtclan", City: "Bangalore"}

// // // 	employee := Employee{Id: 1, Name: "Venkat", Org: organization}

// // // 	fmt.Printf("%#v", employee)

// // // }

// // // package main

// // // import "fmt"

// // // type Product struct {
// // // 	Id   int
// // // 	Name string
// // // 	Cost float32
// // // }

// // // type Dummy struct {
// // // 	Id int
// // // }

// // // type PerishableProduct struct {
// // // 	Dummy
// // // 	Product
// // // 	Expiry string
// // // }

// // // func main() {

// // // 	p := Product{
// // // 		Id:   100,
// // // 		Name: "Grapes",
// // // 		Cost: 60,
// // // 	}

// // // 	var grapes PerishableProduct = PerishableProduct{
// // // 		Product: p,
// // // 		Expiry:  "2 Days",
// // // 	}
// // // 	fmt.Println(grapes.Product.Id, grapes.Name, grapes.Cost, grapes.Expiry)

// // // 	//fmt.Println(Format(grapes.Product))

// // // 	p.ApplyDiscount(10)

// // // 	grapes.ApplyDiscount(10)

// // // 	fmt.Println(grapes.Format())

// // // 	//fmt.Println(grapes.Cost)
// // // 	// fmt.Println(grapes.Product.Id, grapes.Product.Name, grapes.Product.Cost, grapes.Expiry)

// // // 	// use the Format & ApplyDiscount functions for grapes (do not change the functions)
// // // }

// // // func (product Product) Format() string {
// // // 	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
// // // }

// // // func (pp PerishableProduct) Format() string {
// // // 	return fmt.Sprintf("%s Expiry = %q ", pp.Product.Format(), pp.Expiry)
// // // }

// // // func (product *Product) ApplyDiscount(discountPercentage float32) {
// // // 	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
// // // }

// // // package main

// // // import "fmt"

// // // /*
// // // Write the apis for the following

// // // IndexOf => return the index of the given product (return -1 if not exists )
// // // 	ex:  returns the index of the given product

// // // */

// // // func IndexOf(products []Product, product Product) int {

// // // 	for idx := 0; idx < len(products); idx++ {

// // // 		if products[idx].Id == product.Id {
// // // 			return idx
// // // 		}
// // // 	}

// // // 	return -1

// // // }

// // // /*

// // // Includes => return true if the given product is present in the collection else return false
// // // 	ex:  returns true if the given product is present in the collection

// // // */

// // // func Includes(products []Product, product Product) bool {

// // // 	return IndexOf(products, product) != -1

// // // }

// // // /*

// // // Filter => return a new collection of products that satisfy the given condition
// // // 	some of the use cases:
// // // 		1. filter all costly products (cost > 1000)
// // // 			OR
// // // 		2. filter all stationary products (category = "Stationary")
// // // 			OR
// // // 		3. filter all costly stationary products
// // // 		etc
// // // */

// // // func Filter(products []Product) []Product {

// // // 	filteredProducts := []Product{}

// // // 	var maxPrice float32 = 0

// // // 	for _, product := range products {

// // // 		if product.Category == "Stationary" {

// // // 			filteredProducts = append(filteredProducts, product)
// // // 			if product.Cost > maxPrice {

// // // 				maxPrice = product.Cost

// // // 			}
// // // 		} else if product.Cost > 1000 {
// // // 			filteredProducts = append(filteredProducts, product)
// // // 		}

// // // 	}

// // // 	for _, product := range products {

// // // 		if product.Category == "Stationary" && product.Cost == maxPrice {
// // // 			filteredProducts = append(filteredProducts, product)
// // // 		}
// // // 	}

// // // 	return filteredProducts

// // // }

// // // func Any(products []Product) bool {

// // // 	filteredProducts := Filter(products)

// // // 	return len(filteredProducts) >= 1

// // // }

// // // func All(products []Product) bool {

// // // 	filteredProducts := Filter(products)

// // // 	return len(filteredProducts) == len(products)

// // // }

// // // /*
// // // Any => return true if any of the product in the collections satifies the given criteria
// // // 	some of the use cases:
// // // 		1. are there any costly product (cost > 1000)?
// // // 		OR
// // // 		2. are there any stationary product (category = "Stationary")?
// // // 		OR
// // // 		3. are there any costly stationary product?
// // // 		etc

// // // All => return true if all the products in the collections satifies the given criteria
// // // 	some of use cases:
// // // 		1. are all the products costly products (cost > 1000)?
// // // 		OR
// // // 		2. are all the products stationary products (category = "Stationary")?
// // // 		OR
// // // 		3. are all the products costly stationary products?
// // // 		etc
// // // */

// // // // type Product struct {
// // // // 	Id       int
// // // // 	Name     string
// // // // 	Cost     float32
// // // // 	Units    int
// // // // 	Category string
// // // // }

// // // // func main() {
// // // // 	products := []Product{{105, "Pen", 5, 50, "Stationary"}, {107, "Pencil", 2, 100, "Stationary"},
// // // // 		{103, "Marker", 50, 20, "Utencil"},
// // // // 		{102, "Stove", 5000, 5, "Utencil"},
// // // // 		{101, "Kettle", 2500, 10, "Utencil"},
// // // // 		{104, "Scribble Pad", 20, 20, "Stationary"}, {109, "Golden Pen", 2000, 20, "Stationary"},
// // // // 	}

// // // // 	fmt.Println(IndexOf(products, products[0]))

// // // // 	//fmt.Println(Includes(products,products[2]))

// // // // 	fmt.Println(Filter(products))

// // // // 	fmt.Println(All(products))

// // // // 	fmt.Println(Any(products))

// // // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strconv"
// // // 	"strings"
// // // )

// // /*
// // Write the apis for the following

// // IndexOf => return the index of the given product (return -1 if not exists )
// // 	ex:  returns the index of the given product

// // Includes => return true if the given product is present in the collection else return false
// // 	ex:  returns true if the given product is present in the collection

// // Filter => return a new collection of products that satisfy the given condition
// // 	some of the use cases:
// // 		1. filter all costly products (cost > 1000)
// // 			OR
// // 		2. filter all stationary products (category = "Stationary")
// // 			OR
// // 		3. filter all costly stationary products
// // 		etc

// // Any => return true if any of the product in the collections satifies the given criteria
// // 	some of the use cases:
// // 		1. are there any costly product (cost > 1000)?
// // 		OR
// // 		2. are there any stationary product (category = "Stationary")?
// // 		OR
// // 		3. are there any costly stationary product?
// // 		etc

// // All => return true if all the products in the collections satifies the given criteria
// // 	some of use cases:
// // 		1. are all the products costly products (cost > 1000)?
// // 		OR
// // 		2. are all the products stationary products (category = "Stationary")?
// // 		OR
// // 		3. are all the products costly stationary products?
// // 		etc
// // */

// // // type Product struct {
// // // 	Id       int
// // // 	Name     string
// // // 	Cost     float32
// // // 	Units    int
// // // 	Category string
// // // }

// // // func (p Product) Format() string {
// // // 	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
// // // }

// // // type Products []Product

// // // func (products Products) Format() string {
// // // 	var sb strings.Builder
// // // 	for _, p := range products {
// // // 		sb.WriteString(fmt.Sprintf("%s\n", p.Format()))
// // // 	}
// // // 	return sb.String()
// // // }

// // // func (products Products) IndexOf(p Product) int {
// // // 	for idx, prod := range products {
// // // 		if prod == p {
// // // 			return idx
// // // 		}
// // // 	}
// // // 	return -1
// // // }

// // /*
// // func (products Products) FilterCostlyProducts() Products {
// // 	var result Products
// // 	for _, p := range products {
// // 		if p.Cost > 50 {
// // 			result = append(result, p)
// // 		}
// // 	}
// // 	return result
// // }

// // func (products Products) FilterStationaryProducts() Products {
// // 	var result Products
// // 	for _, p := range products {
// // 		if p.Category == "Stationary" {
// // 			result = append(result, p)
// // 		}
// // 	}
// // 	return result
// // }
// // */

// // // type ProductPredicate func(Product) bool

// // // func (products Products) Filter(predicate ProductPredicate) Products {
// // // 	var result Products
// // // 	for _, p := range products {
// // // 		if predicate(p) {
// // // 			result = append(result, p)
// // // 		}
// // // 	}
// // // 	return result
// // // }

// // // func negate(predicate ProductPredicate) ProductPredicate {
// // // 	return func(p Product) bool {
// // // 		return !predicate(p)
// // // 	}
// // // }

// // // func main() {

// // // 	products := Products{
// // // 		Product{105, "Pen", 5, 50, "Stationary"},
// // // 		Product{107, "Pencil", 2, 100, "Stationary"},
// // // 		Product{103, "Marker", 50, 20, "Utencil"},
// // // 		Product{102, "Stove", 5000, 5, "Utencil"},
// // // 		Product{101, "Kettle", 2500, 10, "Utencil"},
// // // 		Product{104, "Scribble Pad", 20, 20, "Stationary"},
// // // 		Product{109, "Golden Pen", 2000, 20, "Stationary"},
// // // 	}

// // // 	fmt.Println("Initial List")
// // // 	fmt.Println(products.Format())

// // // 	stove := Product{102, "Stove", 5000, 5, "Utencil"}
// // // 	fmt.Println("Index of stove :", products.IndexOf(stove))

// // // 	/*
// // // 		fmt.Println("Filter costly products")
// // // 		costlyProducts := products.FilterCostlyProducts()
// // // 		fmt.Println(costlyProducts.Format())

// // // 		fmt.Println("Filter stationary products")
// // // 		stationaryProducts := products.FilterStationaryProducts()
// // // 		fmt.Println(stationaryProducts.Format())
// // // 	*/

// // // 	fmt.Println("Filter costly products")
// // // 	var costlyProductPredicate = func(p Product) bool {
// // // 		return p.Cost > 50
// // // 	}

// // // 	var stationaryProductPredicate = func(p Product) bool {
// // // 		return p.Category == "Stationary"
// // // 	}

// // // 	var AllPredicate = func(p Product) bool {
// // // 		return p.Cost > 50 && p.Category == "Stationary"
// // // 	}

// // // 	var AnyPredicate = func(p Product) bool {
// // // 		return costlyProductPredicate(p) || stationaryProductPredicate(p)
// // // 	}
// // // 	costlyProducts := products.Filter(costlyProductPredicate)
// // // 	fmt.Println(costlyProducts.Format())

// // // 	fmt.Println("Filter affordable products")
// // // 	/*
// // // 		var affordableProductPredicate = func(p Product) bool {
// // // 			return !costlyProductPredicate(p)
// // // 		}
// // // 	*/

// // // 	var affordableProductPredicate = negate(costlyProductPredicate)
// // // 	affordableProducts := products.Filter(affordableProductPredicate)
// // // 	fmt.Println(affordableProducts.Format())

// // // 	fmt.Println("Filter stationary products")

// // // 	stationaryProducts := products.Filter(stationaryProductPredicate)
// // // 	fmt.Println(stationaryProducts.Format())

// // // 	allFilteredProducts := products.Filter(AllPredicate)
// // // 	fmt.Println(allFilteredProducts.Format())

// // // 	anyFilteredProducts := products.Filter(AnyPredicate)
// // // 	fmt.Println(anyFilteredProducts.Format())
// // // }

// // // package main

// // // /* Use strconv.atoi() to convert from string to int */

// // // import (
// // // 	"fmt"
// // // 	"math"
// // // )

// // // func main() {
// // // 	fmt.Println(sum(10, 20))                                 //=> 30
// // // 	fmt.Println(sum(10, 20, 30, 40))                         //=> 100
// // // 	fmt.Println(sum(10))                                     //=> 10
// // // 	fmt.Println(sum())                                       //=> 0
// // // 	fmt.Println(sum(10, "20", 30, "40"))                     //=> 100 (use strconv.Atoi())
// // // 	fmt.Println(sum(10, "20", 30, "40", "abc"))              //=> 100
// // // 	fmt.Println(sum(10, 20, []int{30, 40}))                  //=> 100
// // // 	fmt.Println(sum(10, 20, []any{30, 40, []int{10, 20}}))   //=> 130
// // // 	fmt.Println(sum(10, 20, []any{30, 40, []any{10, "20"}})) //=> 130
// // // }

// // // func sum(arr ...any) int {
// // // 	var res int

// // // 	for _, val := range arr {

// // // 		switch value := val.(type) {
// // // 		case int:
// // // 			res += value
// // // 		case string:
// // // 			x, _ := strconv.Atoi(value)
// // // 			res += x
// // // 		case []any:
// // // 			res += sum(value...)
// // // 		}
// // // 	}
// // // 	return res
// // // }

// // // sprint:1

// // // type Shape interface {
// // // 	Area() float64
// // // 	Perimeter() float64
// // // }
// // // type Circle struct {
// // // 	Radius float64
// // // }

// // // func (c Circle) Area() float64 {
// // // 	return math.Pi * c.Radius * c.Radius
// // // }

// // // func (c Circle) Perimeter() float64 {
// // // 	return 2.0 * math.Pi * c.Radius
// // // }

// // // // sprint:2
// // // type Rectangle struct {
// // // 	Length  float64
// // // 	Breadth float64
// // // }

// // // func (r Rectangle) Area() float64 {
// // // 	return r.Length * r.Breadth
// // // }

// // // func (r Rectangle) Perimeter() float64 {
// // // 	return 2 * (r.Length + r.Breadth)
// // // }

// // // func PrintArea(s Shape) {
// // // 	fmt.Println("Area :", s.Area())
// // // }
// // // func PrintPerimeter(s Shape) {
// // // 	fmt.Println("Perimeter :", s.Perimeter())
// // // }

// // // func main() {
// // // 	c := Circle{Radius: 15}
// // // 	// fmt.Println("Area :", c.Area())
// // // 	PrintArea(c)

// // // 	r := Rectangle{Length: 10, Breadth: 14}
// // // 	// fmt.Println("Area :", r.Area())
// // // 	PrintArea(r)

// // // 	PrintPerimeter(c)

// // // 	PrintPerimeter(r)
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"sort"
// // 	"strings"
// // )

// // /*
// // Write an api for sorting the products by any attribute
// // IMPORTANT: MUST use sort.Sort()
// // */

// // type Product struct {
// // 	Id       int
// // 	Name     string
// // 	Cost     float32
// // 	Units    int
// // 	Category string
// // }

// // // fmt.Stringer interface implementation
// // func (p Product) String() string {
// // 	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
// // }

// // type Products []Product

// // func (p Products) Len() int           { return len(p) }
// // func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// // func (p Products) Less(i, j int) bool { return p[i].Cost < p[j].Cost }

// // // fmt.Stringer interface implementation
// // func (products Products) String() string {
// // 	var sb strings.Builder
// // 	for _, p := range products {
// // 		sb.WriteString(fmt.Sprintf("%s\n", p))
// // 	}
// // 	return sb.String()
// // }

// // func (products Products) IndexOf(p Product) int {
// // 	for idx, prod := range products {
// // 		if prod == p {
// // 			return idx
// // 		}
// // 	}
// // 	return -1
// // }

// // /*
// // func (products Products) FilterCostlyProducts() Products {
// // 	var result Products
// // 	for _, p := range products {
// // 		if p.Cost > 50 {
// // 			result = append(result, p)
// // 		}
// // 	}
// // 	return result
// // }

// // func (products Products) FilterStationaryProducts() Products {
// // 	var result Products
// // 	for _, p := range products {
// // 		if p.Category == "Stationary" {
// // 			result = append(result, p)
// // 		}
// // 	}
// // 	return result
// // }
// // */

// // type ProductPredicate func(Product) bool

// // func (products Products) Filter(predicate ProductPredicate) Products {
// // 	var result Products
// // 	for _, p := range products {
// // 		if predicate(p) {
// // 			result = append(result, p)
// // 		}
// // 	}
// // 	return result
// // }

// // func negate(predicate ProductPredicate) ProductPredicate {
// // 	return func(p Product) bool {
// // 		return !predicate(p)
// // 	}
// // }

// // func main() {

// // 	products := Products{
// // 		Product{105, "Pen", 5, 50, "Stationary"},
// // 		Product{107, "Pencil", 2, 100, "Stationary"},
// // 		Product{103, "Marker", 50, 20, "Utencil"},
// // 		Product{102, "Stove", 5000, 5, "Utencil"},
// // 		Product{101, "Kettle", 2500, 10, "Utencil"},
// // 		Product{104, "Scribble Pad", 20, 20, "Stationary"},
// // 		Product{109, "Golden Pen", 2000, 20, "Stationary"},
// // 	}

// // 	fmt.Println("Initial List")
// // 	fmt.Println(products)

// // 	stove := Product{102, "Stove", 5000, 5, "Utencil"}
// // 	fmt.Println("Index of stove :", products.IndexOf(stove))

// // 	/*
// // 		fmt.Println("Filter costly products")
// // 		costlyProducts := products.FilterCostlyProducts()
// // 		fmt.Println(costlyProducts)

// // 		fmt.Println("Filter stationary products")
// // 		stationaryProducts := products.FilterStationaryProducts()
// // 		fmt.Println(stationaryProducts)
// // 	*/

// // 	fmt.Println("Filter costly products")
// // 	var costlyProductPredicate = func(p Product) bool {
// // 		return p.Cost > 50
// // 	}
// // 	costlyProducts := products.Filter(costlyProductPredicate)
// // 	fmt.Println(costlyProducts)

// // 	fmt.Println("Filter affordable products")
// // 	/*
// // 		var affordableProductPredicate = func(p Product) bool {
// // 			return !costlyProductPredicate(p)
// // 		}
// // 	*/

// // 	var affordableProductPredicate = negate(costlyProductPredicate)
// // 	affordableProducts := products.Filter(affordableProductPredicate)
// // 	fmt.Println(affordableProducts)

// // 	fmt.Println("Filter stationary products")
// // 	var stationaryProductPredicate = func(p Product) bool {
// // 		return p.Category == "Stationary"
// // 	}
// // 	stationaryProducts := products.Filter(stationaryProductPredicate)
// // 	fmt.Println(stationaryProducts)

// // 	fmt.Println("---------------------------------------------------------------------------")

// // 	sort.Sort(products)

// // 	fmt.Println(products)

// // }

// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	if x := 992; x%2 == 0 {
// // 		fmt.Printf("Given number %d is a even number", x)
// // 	} else {
// // 		fmt.Printf("Given number %d is a odd number", x)
// // 	}
// // }

// /* Handling multiple channels */

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // 	"time"
// // )

// // func main() {

// // 	stopCh := make(chan int)

// // 	fmt.Println("Hit Enter to Stop...")

// // 	go func() {
// // 		fmt.Scanln()
// // 		close(stopCh)
// // 	}()

// // 	ch := genPrime(stopCh)

// // 	for data := range ch {
// // 		fmt.Println(data)
// // 	}

// // }

// // func genPrime(chann <-chan int) chan int {
// // 	x := make(chan int)

// // 	go func() {
// // 	LOOP:
// // 		for i := 3; ; i++ {
// // 			select {
// // 			case <-chann:
// // 				break LOOP
// // 			default:
// // 				if isPrime(i) {
// // 					time.Sleep(100 * time.Millisecond)
// // 					x <- i
// // 				}

// // 			}

// // 		}
// // 		close(x)
// // 	}()
// // 	return x
// // }

// // func isPrime(no int) bool {
// // 	for i := 2; i <= int(math.Sqrt(float64(no))); i++ {
// // 		if no%i == 0 {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // 	"sync/atomic"
// // )

// // var count int64

// // func main() {
// // 	wg := &sync.WaitGroup{}
// // 	for i := 0; i < 200; i++ {
// // 		wg.Add(1)
// // 		go increment(wg)
// // 	}
// // 	wg.Wait()
// // 	fmt.Println(count)
// // }

// // func increment(wg *sync.WaitGroup) {
// // 	defer wg.Done()
// // 	atomic.AddInt64(&count, 1)
// // }

// // package main

// // import "fmt"

// // func main() {
// // 	//non-buffered channel
// // 	//ch := make(chan int)

// // 	//non-buffered channel (size 0)
// // 	//ch := make(chan int, 0)

// // 	//buffered channel (size 1)
// // 	// ch := make(chan int, 1)

// // 	//buffered channel (size 2)
// // 	ch := make(chan int, 2)

// // 	//buffered channel (size 3)
// // 	//ch := make(chan int, 3)

// // 	go writeData(ch)

// // 	fmt.Println("[@main] attempting to receiv 10")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 10")

// // 	fmt.Println("[@main] attempting to receiv 20")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 20")

// // 	fmt.Println("[@main] attempting to receiv 30")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 30")

// // 	fmt.Println("[@main] attempting to receiv 40")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 40")

// // 	fmt.Println("[@main] attempting to receiv 50")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 50")

// // 	fmt.Println("[@main] attempting to receiv 60")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 60")

// // 	fmt.Println("[@main] attempting to receiv 70")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 70")

// // 	fmt.Println("[@main] attempting to receiv 80")
// // 	fmt.Println(<-ch)
// // 	fmt.Println("[@main] completed receiving 80")

// // }

// // func writeData(ch chan int) {
// // 	fmt.Println("	[@writeData] attempting to send 10")
// // 	ch <- 10
// // 	fmt.Println("	[@writeData] completed sending 10")

// // 	fmt.Println("	[@writeData] attempting to send 20")
// // 	ch <- 20
// // 	fmt.Println("	[@writeData] completed sending 20")

// // 	fmt.Println("	[@writeData] attempting to send 30")
// // 	ch <- 30
// // 	fmt.Println("	[@writeData] completed sending 30")

// // 	fmt.Println("	[@writeData] attempting to send 40")
// // 	ch <- 40
// // 	fmt.Println("	[@writeData] completed sending 40")

// // 	fmt.Println("	[@writeData] attempting to send 50")
// // 	ch <- 50
// // 	fmt.Println("	[@writeData] completed sending 50")

// // 	fmt.Println("	[@writeData] attempting to send 60")
// // 	ch <- 60
// // 	fmt.Println("	[@writeData] completed sending 60")

// // 	fmt.Println("	[@writeData] attempting to send 70")
// // 	ch <- 70
// // 	fmt.Println("	[@writeData] completed sending 70")

// // 	fmt.Println("	[@writeData] attempting to send 80")
// // 	ch <- 80
// // 	fmt.Println("	[@writeData] completed sending 80")
// //

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // 	"sync"
// // )

// // type PrimeResult struct {
// // 	No      int
// // 	IsPrime bool
// // }

// // func main() {
// // 	ch := genPrimes(3, 100)
// // 	for primeNo := range ch {
// // 		fmt.Println(primeNo)
// // 	}
// // }

// // func genPrimes(start, end int) <-chan int {
// // 	ch := make(chan int)
// // 	primeCh := make(chan PrimeResult)
// // 	wg := &sync.WaitGroup{}
// // 	go func() {

// // 		for i := start; i <= end; i++ {
// // 			wg.Add(1)
// // 			go isPrime(i, primeCh, wg)
// // 		}

// // 		go func() {
// // 			wg.Wait()

// // 			close(primeCh)
// // 		}()

// // 		for result := range primeCh {
// // 			if result.IsPrime {
// // 				ch <- result.No
// // 			}
// // 		}
// // 		close(ch)
// // 	}()
// // 	return ch
// // }

// // func isPrime(no int, primeCh chan<- PrimeResult, wg *sync.WaitGroup) {
// // 	defer wg.Done()
// // 	for i := 2; i <= int(math.Sqrt(float64(no))); i++ {
// // 		if no%i == 0 {
// // 			primeCh <- PrimeResult{No: no, IsPrime: false}
// // 			return
// // 		}
// // 	}
// // 	primeCh <- PrimeResult{No: no, IsPrime: true}
// // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"time"
// // // )

// // // func main() {
// // // 	ch1 := make(chan int)
// // // 	ch2 := make(chan int)

// // // 	go func() {
// // // 		time.Sleep(2 * time.Second)
// // // 		ch1 <- 100
// // // 	}()

// // // 	go func() {
// // // 		time.Sleep(4 * time.Second)
// // // 		ch2 <- 200
// // // 	}()

// // // 	ch3 := make(chan int)
// // // 	go func() {
// // // 		time.Sleep(3 * time.Second)
// // // 		d3 := <-ch3
// // // 		fmt.Println("Data from ch3 :", d3)
// // // 	}()

// // // 	for i := 0; i < 3; i++ {
// // // 		select {
// // // 		case d1 := <-ch1:
// // // 			fmt.Println(d1)
// // // 		case ch3 <- 300:
// // // 			fmt.Println("data sent to ch3")
// // // 		case d2 := <-ch2:
// // // 			fmt.Println(d2)
// // // 		}
// // // 	}

// // // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"time"
// // // )

// // // func main() {
// // // 	ch := genPrimes(3, 100)
// // // 	for primeNo := range ch {
// // // 		fmt.Println(primeNo)
// // // 	}
// // // }

// // // func genPrimes(start, end int) <-chan int {
// // // 	ch := make(chan int)
// // // 	go func() {
// // // 		for i := start; i <= end; i++ {
// // // 			if isPrime(i) {
// // // 				time.Sleep(500 * time.Millisecond)
// // // 				ch <- i
// // // 			}
// // // 		}
// // // 		close(ch)
// // // 	}()
// // // 	return ch
// // // }

// // // func isPrime(no int) bool {
// // // 	for i := 2; i <= (no / 2); i++ {
// // // 		if no%i == 0 {
// // // 			return false
// // // 		}
// // // 	}
// // // 	return true
// // // }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type PrimeResult struct {
// 	No      int
// 	IsPrime bool
// }

// func main() {

// 	rootCntx := context.Background()

// 	timeCntx, cancel := context.WithTimeout(rootCntx, 3*time.Second)

// 	defer func() {
// 		fmt.Println(time.Now())
// 	}()
// 	go func() {
// 		fmt.Scanln()
// 		cancel()
// 	}()

// 	fmt.Println(time.Now())

// 	ch := genPrimes(3, 1300, timeCntx)
// 	for primeNo := range ch {
// 		fmt.Println(primeNo)
// 	}

// }

// func genPrimes(start, end int, timeCntx context.Context) <-chan int {
// 	ch := make(chan int)
// 	primeCh := make(chan PrimeResult)
// 	go func() {
// 		wg := &sync.WaitGroup{}
// 	LOOP:
// 		for i := start; i <= end; i++ {
// 			select {
// 			case <-timeCntx.Done():
// 				fmt.Println("[printTime] Cancel signal received.. stopping")
// 				break LOOP
// 			default:
// 				wg.Add(1)
// 				go isPrime(i, primeCh, wg)

// 			}

// 		}

// 		go func() {
// 			wg.Wait()
// 			close(primeCh)
// 		}()

// 		for result := range primeCh {
// 			if result.IsPrime {
// 				ch <- result.No
// 			}
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func isPrime(no int, primeCh chan<- PrimeResult, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(500 * time.Millisecond)
// 	for i := 2; i <= (no / 2); i++ {

// 		if no%i == 0 {
// 			primeCh <- PrimeResult{No: no, IsPrime: false}
// 			return
// 		}
// 	}
// 	primeCh <- PrimeResult{No: no, IsPrime: true}
// }

package main

import (
	"brewery/router"
	"net/http"
)

func main() {

	route := router.Router()

	http.ListenAndServe(":8080", route)

}
