package main

import (
    "io"
    "os"
    "strings"
    "unicode"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
    readAmt, err := r13.r.Read(b)
    
    for i := 0; i < readAmt; i++ {
        s := rune(b[i])
        if unicode.IsLetter(s){
            switch {
                case b[i] < 'n':
                    b[i] += 13
                case b[i] < 'A':
                    b[i] -= 13
                case b[i] < 'N':
                    b[i] += 13
                default:
                    b[i] -= 13
            }
        }
    }
    
    return len(b), err
}

func RunRot13Reader() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}