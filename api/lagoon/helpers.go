package lagoon

import "errors"

// ErrExist indicates that an attempt was made to create an object that already
// exists.
var ErrExist = errors.New("object already exists")
