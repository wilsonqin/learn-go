package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
    // alternative loop format
    //for i := 0 ; i < len(b) ; i++ {
    //  b[i] = 'A'
    //}
    
    for i := range b {
        b[i] = 'A'
    }
    
    return len(b), nil
}

func RunMyReader() {
    reader.Validate(MyReader{})
}