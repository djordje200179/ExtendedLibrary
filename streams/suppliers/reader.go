package suppliers

import (
	"bufio"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
	"io"
)

func Reader(reader io.Reader, splitFunction bufio.SplitFunc) streams.Supplier[string] {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunction)

	return func() optional.Optional[string] {
		if scanner.Scan() {
			return optional.FromValue(scanner.Text())
		} else {
			return optional.Empty[string]()
		}
	}
}
