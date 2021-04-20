package common_func

const BatchDealCount = 100

func SliceGroutinue() {
	tempSlice := []string{}
	goroutineCount := 0
	tempChannel := make(chan []string)

	result := []string{}

	left := 0
	right := 0
	for right < len(tempSlice) {
		left = right
		right = left + BatchDealCount
		if right >= len(tempSlice) {
			right = len(tempSlice)
		}
		go DealFunc(tempSlice[left:right], tempChannel)
		goroutineCount++

	}

	for i := 0; i < goroutineCount; i++ {
		result = append(result, <-tempChannel...)
	}
	close(tempChannel)
}

func DealFunc(tempSlice []string, tempChannel chan []string) {
	//deal logic
}
