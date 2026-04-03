package hyperloglog

import (
	"cmp"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"slices"
	"sync"

	"github.com/PhanNam1501/bustub-go/include/common/utils"
	"github.com/PhanNam1501/bustub-go/include/types"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type HashFunc[KeyType comparable] func(KeyType) uint64

type CountMinSketch[KeyType comparable] struct {
	width         uint32
	depth         uint32
	matrix        [][]uint32
	hashFunctions []HashFunc[KeyType]
	mu            sync.Mutex
}

/**
 * Constructor for the count-min sketch.
 *
 * @param width The width of the sketch matrix.
 * @param depth The depth of the sketch matrix.
 * @returns error if width or depth are zero (thay cho std::invalid_argument).
 */
func NewCountMinSketch[KeyType comparable](width, depth uint32) (*CountMinSketch[KeyType], error) {
	if width == 0 || depth == 0 {
		return nil, errors.New("width or depth cannot be zero")
	}

	cms := &CountMinSketch[KeyType]{
		width: width,
		depth: depth,
	}

	// @TODO(student) Implement this function!
	cms.matrix = make([][]uint32, depth)

	for i := range cms.matrix {
		cms.matrix[i] = make([]uint32, width)
	}
	// @spring2026 PLEASE DO NOT MODIFY THE FOLLOWING
	// Initialize seeded hash functions
	cms.hashFunctions = make([]HashFunc[KeyType], 0, depth)
	for i := uint32(0); i < depth; i++ {
		cms.hashFunctions = append(cms.hashFunctions, cms.getHashFunction(i))
	}

	return cms, nil
}

func (c *CountMinSketch[KeyType]) getHashFunction(seed uint32) HashFunc[KeyType] {

	seedBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(seedBytes, seed)
	seedHash := utils.HashBytes(seedBytes)

	return func(key KeyType) uint64 {
		var keyHash utils.HashT

		switch v := any(key).(type) {
		case types.Value:
			keyHash = utils.HashValue(v)
		case string:
			keyHash = utils.HashBytes([]byte(v))

		case []byte:
			keyHash = utils.HashBytes(v)

		case int64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
			keyHash = utils.HashBytes(b)

		case uint32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, v)
			keyHash = utils.HashBytes(b)

		default:
			keyHash = utils.HashBytes([]byte(fmt.Sprintf("%v", v)))
		}
		finalHash := utils.CombineHashes(seedHash, keyHash)

		return uint64(finalHash)
	}
}

func (c *CountMinSketch[KeyType]) Insert(item KeyType) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// @TODO(student) Implement this function!
	for i, h := range c.hashFunctions {
		v := h(item)
		vMod := v % uint64(c.width)
		c.matrix[i][vMod]++
	}
}

func (c *CountMinSketch[KeyType]) Merge(other *CountMinSketch[KeyType]) error {
	if c.width != other.width || c.depth != other.depth {
		return errors.New("Incompatible CountMinSketch dimensions for merge")
	}

	// Khóa cả 2 sketch để đảm bảo thread-safe khi merge
	c.mu.Lock()
	defer c.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()

	// @TODO(student) Implement this function!
	for i := 0; i < int(c.depth); i++ {
		for j := 0; j < int(c.width); i++ {
			c.matrix[i][j] += other.matrix[i][j]
		}
	}

	return nil
}

func (c *CountMinSketch[KeyType]) Count(item KeyType) uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()

	// @TODO(student) Implement this function!
	var minn uint32 = math.MaxUint32
	for i, h := range c.hashFunctions {
		v := h(item)
		vMod := v % uint64(c.width)
		minn = min(c.matrix[i][vMod], minn)
	}
	return minn
}

func (c *CountMinSketch[KeyType]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := range c.matrix {
		clear(c.matrix[i])
	}
}

func (c *CountMinSketch[KeyType]) TopK(k uint16, candidates []KeyType) []Pair[KeyType, uint32] {
	// @TODO(student) Implement this function!
	res := make([]Pair[KeyType, uint32], len(candidates))
	for i, cand := range candidates {
		var minn uint32 = math.MaxUint32
		for j, h := range c.hashFunctions {
			v := h(cand)
			vMod := v % uint64(c.width)
			minn = min(c.matrix[j][vMod], minn)
		}
		res[i] = Pair[KeyType, uint32]{
			Key:   cand,
			Value: minn,
		}
	}

	slices.SortFunc(res, func(a, b Pair[KeyType, uint32]) int {
		return cmp.Compare(b.Value, a.Value)
	})

	kInt := int(k)

	if kInt > len(res) {
		kInt = len(res)
	}

	return res[:kInt]
}
