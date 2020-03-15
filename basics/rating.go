package main

// func main() {
// 	var name string
// 	var userRating string

// 	// Front End
// 	// take uisr input
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter a name ")
// 	name, _ = reader.ReadString('\n')

// 	// take rating form user  and conver it into string
// 	reader = bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter the rating between 1 to 10")
// 	userRating, _ = reader.ReadString('\n')
// 	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 32)
// 	if myNumRating > 10 {
// 		println("You have Entered a Wrong Value")
// 	} else {
// 		// BackENd
// 		fmt.Printf("Hello %v!Thanks for rating our shop with %v Star Rating. \n\n your rating was recorded in our system at %v", name, myNumRating, time.Now().Format(time.Stamp))
// 	}
// }
