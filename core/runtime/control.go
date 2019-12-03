package runtime

import "errors"

const BLOCK_CMDOK byte = 0x20
const BLOCK_CMDDUMP byte = 0x40

var NoCommandVisitorError error = errors.New("no remote command visitor")
