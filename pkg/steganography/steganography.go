package steganography

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"

	sg "github.com/auyer/steganography"
)

type Steganography interface {
}

type steganography struct {
}

func Init() Steganography {
	return &steganography{}
}

func (s *steganography) Encode() error {
	inFile, _ := os.Open("input_file.png") // opening file
	reader := bufio.NewReader(inFile)      // buffer reader
	img, _ := png.Decode(reader)           // decoding to golang's image.Image

	w := new(bytes.Buffer) // buffer that will recieve the results

	err := sg.Encode(w, img, []byte("message")) // Encode the message into the image
	if err != nil {
		log.Printf("Error Encoding file %v", err)
		return fmt.Errorf("error encoding file: %d", err)
	}

	outFile, _ := os.Create("out_file.png") // create file
	_, err = w.WriteTo(outFile)
	if err != nil {
		return fmt.Errorf("error generating output: %d", err)
	}

	// write buffer to it
	err = outFile.Close()
	if err != nil {
		return fmt.Errorf("error writing buffer: %d", err)
	}

	return nil
}

func (s *steganography) Decode(encodedInputFile string) error {
	inFile, _ := os.Open(encodedInputFile) // opening file
	defer inFile.Close()

	reader := bufio.NewReader(inFile) // buffer reader
	img, _ := png.Decode(reader)      // decoding to golang's image.Image

	sizeOfMessage := sg.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

	msg := sg.Decode(sizeOfMessage, img) // decoding the message from the file
	fmt.Println(string(msg))

	return nil
}
