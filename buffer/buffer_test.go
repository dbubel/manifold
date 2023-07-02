package buffer

import (
	"testing"
	"time"
)

// chat gippity test
func TestCircularBuffer(t *testing.T) {
	// Create input and output channels
	inputChannel := make(chan []byte)
	outputChannel := make(chan []byte)

	// Create a new circular buffer
	cb := NewBuffer(inputChannel, outputChannel)

	// Input some data into the circular buffer
	data := []byte{1, 2, 3}
	inputChannel <- data

	// Wait for a short time to allow the circular buffer to process the input
	time.Sleep(100 * time.Millisecond)

	// Check if the data is correctly stored in the circular buffer
	if cb.head == nil || cb.tail == nil || cb.len != 1 {
		t.Errorf("Expected circular buffer length to be 1, but got %d", cb.len)
	}

	// Output the data from the circular buffer
	outputData := <-outputChannel

	// Check if the output data matches the input data
	if len(outputData) != len(data) {
		t.Errorf("Expected output data length to be %d, but got %d", len(data), len(outputData))
	}

	for i := 0; i < len(data); i++ {
		if outputData[i] != data[i] {
			t.Errorf("Expected output data at index %d to be %d, but got %d", i, data[i], outputData[i])
		}
	}
}

// chat gippity test
func TestCircularBufferMultipleInputs(t *testing.T) {
	// Create input and output channels
	inputChannel := make(chan []byte)
	outputChannel := make(chan []byte)

	// Create a new circular buffer
	cb := NewBuffer(inputChannel, outputChannel)

	// Input multiple data into the circular buffer
	data1 := []byte{1, 2, 3}
	data2 := []byte{4, 5, 6}
	data3 := []byte{7, 8, 9}
	inputChannel <- data1
	inputChannel <- data2
	inputChannel <- data3

	// Wait for a short time to allow the circular buffer to process the inputs
	time.Sleep(100 * time.Millisecond)

	// Check if the data is correctly stored in the circular buffer
	if cb.len != 3 {
		t.Errorf("Expected circular buffer length to be 3, but got %d", cb.len)
	}

	// Output the data from the circular buffer
	outputData1 := <-outputChannel
	outputData2 := <-outputChannel
	outputData3 := <-outputChannel

	// Check if the output data matches the input data
	checkOutputData(t, data1, outputData1, 1)
	checkOutputData(t, data2, outputData2, 2)
	checkOutputData(t, data3, outputData3, 3)
}

func checkOutputData(t *testing.T, expected, actual []byte, index int) {
	if len(actual) != len(expected) {
		t.Errorf("Expected output data length to be %d at index %d, but got %d",
			len(expected), index, len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected output data at index %d to be %d, but got %d",
				i, expected[i], actual[i])
		}
	}
}

func BenchmarkCircularBufferInputOutput(b *testing.B) {
	// Create input and output channels
	inputChannel := make(chan []byte)
	outputChannel := make(chan []byte)

	// Create a new circular buffer
	cb := NewBuffer(inputChannel, outputChannel)

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := []byte{byte(i % 256)} // Generate random data
		cb.inputChannel <- data
		<-cb.outputChannel
	}
	b.StopTimer()

	// Close the input channel
	close(cb.inputChannel)
}
