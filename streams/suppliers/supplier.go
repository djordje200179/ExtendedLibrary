package suppliers

import "github.com/djordje200179/extendedlibrary/misc/optional"

type Supplier[T any] interface {
	Supply() optional.Optional[T]
}
