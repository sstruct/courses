package main

// Conditional
if day == "sunday" || day == "saturday" {
	rest()
} else if day == "monday" && isTired() {
	groan()
} else {
	work()
}

// Statements in if
if _, err := getResult(); err != nil {	// 这句啥意思
	fmt.Println("Un oh")
}

// Switch
switch day {
    case "sunday":
    	fallthrough
	case "saturday":
		rest()
	default:
		work()
}
