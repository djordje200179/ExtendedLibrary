package suppliers

import (
	"bufio"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"io"
)

type readerSupplier struct {
	*bufio.Scanner
}

func Reader(reader io.Reader, splitFunction bufio.SplitFunc) Supplier[string] {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunction)

	return readerSupplier{scanner}
}

func (supplier readerSupplier) Supply() optional.Optional[string] {
	if supplier.Scan() {
		return optional.FromValue(supplier.Text())
	} else {
		return optional.Empty[string]()
	}
}
