package samgol

import (
    "math/rand"
    "time"
)

// Return random place
func GetPlace() string {
    rand.Seed(time.Now().Unix())
    places := []string{
        "World",
        "Singapore",
        "渋谷",
        "函館",
    }
    return places[rand.Intn(len(places))]
}
