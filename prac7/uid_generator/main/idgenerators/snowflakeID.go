package idgenerators

import (
	"fmt"
	"time"
)

const (
	on status = iota
	off
)

type status int8
type snowflakeID int64

var dataCenterId int8 = -1
var computerId int8 = -1
var sequenceNumber int16 = 0
var currentMiliseconds int64 = 0
var startMiliseconds int64 = 0
var generationStatus status = on

func InitGenerator(datacenterid int8, computerid int8) {
	if dataCenterId != -1 || computerId != -1 {
		fmt.Println("datacenter is already setted up")
	}
	dataCenterId = datacenterid
	computerId = computerid
	startMiliseconds = time.Now().UnixMilli()
	currentMiliseconds = startMiliseconds
}

func StartGenerator() {
	generationStatus = on
	go update()
}

func StopGenerator() {
	generationStatus = off
}

func update() {
	fmt.Println("asdas")
	time.Sleep(time.Millisecond)
	currentMiliseconds++
	sequenceNumber = 0
	if generationStatus == on {
		go update()
	}
}

// 0000000000000000000000000000000000000111010000000000000000000000     difference
// 0000000000000000000000000000000000000000000000100000000000000000   dataCenterID
// 0000000000000000000000000000000000000000000000000001000000000000   computerID
// 0000000000000000000000000000000000000000000000000000000000000001    sequenceNumber

func GetSnowflakeID() snowflakeID {
	var result snowflakeID = 0
	result += snowflakeID((currentMiliseconds-startMiliseconds)<<23) >> 1
	result += snowflakeID(((dataCenterId)<<3)>>3) << 17
	result += snowflakeID(((computerId)<<3)>>3) << 12
	result += snowflakeID((sequenceNumber << 4) >> 4)
	sequenceNumber++
	return result
}

func Get100SnowFlakeIDs() *[]snowflakeID {
	res := make([]snowflakeID, 100)
	for i := 0; i < 100; i++ {
		id := GetSnowflakeID()
		res[i] = id
	}

	return &res
}
