# Go Standard Deviation

Check out the standard deviation in Go.

It is meant to be simple and easy to read, hopefully it is.

## Usage

    import (
        "fmt"
        "github.com/antonlindstrom/gostddev"
    )

    func MyStdDev() {
        set := float64{1,2,3,4,5}
        fmt.Printf("%f\n", gostddev.StdDev(set))
        // Prints 1.41421
    }
