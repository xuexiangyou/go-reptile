package persist

import (
	"avatar/lib"
	"log"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 1
		for {
			items := <- out
			log.Printf("Get itemCount#%d item %v", itemCount, items)
			lib.Upload(items.(string), itemCount)
			itemCount++
		}
	}()

	return out
}
