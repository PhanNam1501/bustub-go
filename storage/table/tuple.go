// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// tuple.go
//
// ===----------------------------------------------------------------------===//

package table

// Tuple represents one row of data.
// TODO: implement Serialize/Deserialize based on slotted page format.
type Tuple struct {
	Data []byte
}

// Serialize serializes the tuple.
//
// TODO: Serialize
func (t *Tuple) Serialize() []byte { return t.Data }

// DeserializeTuple deserializes raw bytes into a Tuple.
//
// TODO: Deserialize
func DeserializeTuple(data []byte) *Tuple { return &Tuple{Data: data} }

