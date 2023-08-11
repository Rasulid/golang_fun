package main

import "fmt"

type Person struct {
	ID   int64
	Name string
}

func main() {
	var defaultMap map[int64]string

	fmt.Printf("Type: %T Value: %#v\n", defaultMap, defaultMap)
	fmt.Printf("Len: %d\n\n", len(defaultMap))

	// map by make

	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMake, mapByMake)

	// Map by make with cap

	mapByMakeWithCap := make(map[string]string, 10)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

	// Map by Literal

	mapByLiteral := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	fmt.Printf("Type: %T Value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Len: %d\n\n", len(mapByLiteral))

	// insert and update values

	mapByMake["First"] = "value1"
	fmt.Printf("Type: %T Value: %#v\n", mapByMake, mapByMake)
	fmt.Printf("Len: %d\n\n", len(mapByMake))

	fmt.Println(mapByLiteral["key1"])
	mapByLiteral["key1"] = "Rasul"
	fmt.Println(mapByLiteral["key1"])

	// check if exists

	value, ok := mapByMake["value"]
	fmt.Printf("Value: %d IsExist: %t\n\n", value, ok)

	// delete value

	delete(mapByMake, "First")
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMake, mapByMake)

	// map iteration

	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	for key, value := range mapForIteration {
		fmt.Printf("Key: %s Value: %d\n", key, value)
	}

	// unique values

	users := []Person{
		{ID: 1, Name: "Rasul"},
		{ID: 2, Name: "Rasuwl"},
		{ID: 3, Name: "Rasul"},
		{ID: 1, Name: "Rasuwl"},
	}

	uniqueUsers := make(map[int64]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.ID]; !ok {
			uniqueUsers[user.ID] = struct{}{}
		}
	}
	fmt.Println("uniqueUsers: ", uniqueUsers)
	fmt.Printf("Type: %T Value: %#v\n", uniqueUsers, uniqueUsers)

	// find values

	usermap := make(map[int64]Person, len(users))
	for _, user := range users {
		usermap[user.ID] = user
	}

	fmt.Println(FindByMap(6, usermap))

}

func FindByMap(id int64, userMap map[int64]Person) *Person {
	if user, ok := userMap[id]; ok {
		return &user
	}
	return nil
}
