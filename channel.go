package main
import "fmt"

func main() {
	// Create a channel of integers
	ch := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i * 2 // Send even numbers to the channel
		}
		close(ch) // Close the channel after sending all values
	}()

	// Receive values from the channel
	for value := range ch {
		fmt.Println("Received:", value) // Print received values
	}

}