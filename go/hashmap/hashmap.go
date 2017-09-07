package hashmap

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

const ratio = .75 // ratio sets the capacity the hashmap has to be at before it expands

// roundUp takes a uint64 greater than 0 and rounds it up to the next
// power of 2.
func roundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}

func hash(key uint64) uint64 {
	key ^= key >> 33
	key *= 0xff51afd7ed558ccd
	key ^= key >> 33
	key *= 0xc4ceb9fe1a85ec53
	key ^= key >> 33
	return key
}

func hashAny(data interface{}) int64 {
	h, _ := hashstructure.Hash(data, nil)
	return int64(hash(h))
}

type packet struct {
	key, value interface{}
}

type packets []*packet

func (packets packets) find(key interface{}) int64 {
	h := hashAny(key)
	i := h & (int64(len(packets)) - 1)
	for packets[i] != nil && packets[i].key != key {
		i = (i + 1) & (int64(len(packets)) - 1)
	}
	if packets[i] != nil && packets[i].key != key {
		return -1
	}

	return i
}

func (packets packets) set(packet *packet) (ok bool) {
	i := packets.find(packet.key)
	if i < 0 {
		return false
	}
	if packets[i] == nil {
		packets[i] = packet
		return true
	}

	packets[i].value = packet.value
	return true
}

func (packets packets) get(key interface{}) (interface{}, bool) {
	i := packets.find(key)
	if i < 0 {
		return nil, false
	}
	if packets[i] == nil {
		return nil, false
	}

	return packets[i].value, true
}

func (packets packets) delete(key interface{}) bool {
	i := packets.find(key)
	if i < 0 {
		return false
	}
	if packets[i] == nil {
		return false
	}
	packets[i] = nil

	return true
}

func (packets packets) exists(key interface{}) bool {
	i := packets.find(key)
	return packets[i] != nil
}

type HashMap struct {
	count   uint64
	packets packets
}

func (hm *HashMap) rebuild() {
	packets := make(packets, roundUp(uint64(len(hm.packets))+1))
	for _, packet := range hm.packets {
		if packet == nil {
			continue
		}

		packets.set(packet)
	}
	hm.packets = packets
}

func (hm *HashMap) Get(key interface{}) (interface{}, bool) {
	return hm.packets.get(key)
}

func (hm *HashMap) Set(key, value interface{}) {
	if float64(hm.count+1)/float64(len(hm.packets)) > ratio {
		hm.rebuild()
	}

	if ok := hm.packets.set(&packet{key: key, value: value}); !ok {
		hm.rebuild()
		hm.packets.set(&packet{key: key, value: value})
	}
	hm.count++
}

func (hm *HashMap) Exists(key interface{}) bool {
	return hm.packets.exists(key)
}

func (hm *HashMap) Delete(key interface{}) {
	if hm.packets.delete(key) {
		hm.count--
	}
}

func (hm *HashMap) Len() uint64 {
	return hm.count
}

func (hm *HashMap) Cap() uint64 {
	return uint64(len(hm.packets))
}

func (hm *HashMap) debug() {
	fmt.Printf("packets:%+v\n", hm.packets)
}

// New returns a new HashMap with a bucket size specified
// by hint.
func New(hint uint64) *HashMap {
	if hint == 0 {
		hint = 16
	}

	hint = roundUp(hint)
	return &HashMap{
		count:   0,
		packets: make(packets, hint),
	}
}
