package main

import "fmt"
import "io"
import "io/ioutil"

type ByteReader byte

// implement io.Reader interface
func (b ByteReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = byte(b)
		fmt.Println("reading...")
	}
	return len(buf), nil
}

type LogReader struct {
	io.Reader
}

// 合成したio.ReaderではすでにReadが定義されているが
// 別の処理を付け加えるためにwrapする
func (r LogReader) Read(b []byte) (int, error) {
	n, err := r.Reader.Read(b)
	fmt.Printf("read %d bytes, error: %v", n, err)

	return n, err
}

func main() {
	// wrapping interfaces
	{
		// ByteReaderをwrapして利用することで
		// ByteReader.Readの特性を利用しつつ(r.Reader.Read(b))
		// LogReader.Readのaddonした特性も利用している
		r := LogReader{ByteReader('A')}

		b := make([]byte, 10)
		r.Read(b)

		fmt.Printf("b: %q", b)
	}

	// chaining interfaces
	{
		// ByteReaderはio.Readerのinterfaceを満たしているため代入可能
		var r io.Reader = ByteReader('a')
		//// LimitReader returns a Reader that reads from r
		//// but stops with EOF after n bytes.
		//// The underlying implementation is a *LimitedReader.
		//func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }
		r = io.LimitReader(r, 1e6)

		// ByteReaderに対してLimitReaderの特性(?)を持たせた上で
		// LogReaderのio.Readerとして合成する
		r = LogReader{r}
		// NOTE:ここまだわかってない
		io.Copy(ioutil.Discard, r)

		// bb := make([]byte, 10)
		// r.Read(bb)
		//
		// fmt.Printf("bb: %q", bb)
	}
}
