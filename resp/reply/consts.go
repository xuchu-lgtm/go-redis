package reply

type PongReply struct {
}

var pongbytes = []byte("+PONG\r\n")

func (p *PongReply) ToBytes() []byte {
	return pongbytes
}

func MakePongReply() *PongReply {
	return &PongReply{}
}

type OkReply struct{}

func (o *OkReply) ToBytes() []byte {
	return okBytes
}

var okBytes = []byte("+OK\r\n")

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct {
}

var nullBulkReply = []byte("$-1\r\n")

func (n *NullBulkReply) ToBytes() []byte {
	return nullBulkReply
}

var theNullBulkReply = new(NullBulkReply)

func MakeNullBulkReply() *NullBulkReply {
	return theNullBulkReply
}

type EmptyMultiBulkReply struct{}

func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

var emptyMultiBulkBytes = []byte("*0\r\n")

var theEmptyMultiBulkReply = new(EmptyMultiBulkReply)

func MakeEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return theEmptyMultiBulkReply
}

type NoReply struct {
}

var noReply = []byte("")

func (n *NoReply) ToBytes() []byte {
	return noReply
}

var theNoReply = new(NoReply)

func MakeNoReply() *NoReply {
	return theNoReply
}
