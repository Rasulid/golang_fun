package main

import "fmt"

type Stream interface {
	read() string
	write(string)
	close()
}

func WriteToFile(stream Stream, text string) {
	stream.write(text)
}

func CloseFile(stream Stream) {
	stream.close()
}

// file struct
type File struct {
	text string
}

type Folder struct{}

// Method for creating *File

func (f *File) read() string {
	return f.text
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("text in file row", message)
}

func (f *File) close() {
	fmt.Println("Folder is Colused ")
}

// Method for *Folder

func (f *Folder) close() {
	fmt.Println("Folder is Colused ")
}

func main() {
	file := &File{}
	folder := &Folder{}

	WriteToFile(file, "Hello BItch")
	CloseFile(file)

	folder.close()

}
