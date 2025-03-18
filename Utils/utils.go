package utils

func Repeater(input <-chan int, output ...chan int) {
	for msg := range input {
		for _, ch := range output {
			select {
			case ch <- msg: // Send msg to ch unless it is full
			default:	// If ch is full, do nothing
			}
		}
	}
}
