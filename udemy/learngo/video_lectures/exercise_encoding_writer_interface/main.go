package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var _ io.Reader = strings.NewReader("Hi")

	fmt.Println("Hello, playground")
	//Notice in the fmt package of the std lib that Println is just a wrapper of Fprintln (i.e. it returns the return of Fprintln()
	// which itself accepts the empty interface that Println takes and writes to Stdout.
	// Notice in the std lib that Fprintln accepts type io.Writer and a variadic empty interface. So try using Fprintln and pass it any
	// type that satisfies the Writer interface while specifying the first parm that it accepts, as specified in the std lib.
	fmt.Fprintln(os.Stdout, "Hello, playground")
	//Notice how Write works similarly
	io.WriteString(os.Stdout, "Hello, playground")

}
