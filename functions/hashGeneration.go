package functions

import (
	"hash/fnv"
	"math/rand"
)

func HashGeneration(link string) uint32 {

	hash := fnv.New32()
	hash.Write([]byte(link))

	return hash.Sum32() * uint32(rand.Intn(10))
}
