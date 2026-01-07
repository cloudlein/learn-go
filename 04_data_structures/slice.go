package main

import "fmt"

func main() {
	names := [...]string{"salahudin", "rin", "udin", "fitria"}

	slice := names[1:2]
	slice1 := names[:]
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(cap(slice))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:] // sabtu dan minggu
	daysSlice1[0] = "Sabtu baru"
	daysSlice1[1] = "Minggu baru"
	fmt.Println(daysSlice1)
	fmt.Println(days) // data array ikut berubah

	daysSlice2 := append(daysSlice1, "Libur baru")
	daysSlice2[0] = "Sabtu lama"
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]int, 2, 5)
	newSlice[0] = 1
	newSlice[1] = 2
	//newSlice[2] = 3 // errror harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, 25)
	fmt.Println(newSlice2)
	newSlice2[0] = 10
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	thisArray := [...]int{1, 2, 3}
	thisSlice := []int{1, 2, 3}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
