package store

import (
	"github.com/ivanlebron/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
