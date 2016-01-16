package muls
import (
	"time"
)

type ResultOrder uint32

const (

	DESC	ResultOrder = 1
	ASC		ResultOrder = 2

)


type Storage interface {

	Append(protocol *Protocol) bool;

	BatchAppend(protocol []*Protocol) bool;

	Read(pool string,level ProtocolLevel,limit uint32,order ResultOrder) []string;

	Clean(pool string, level ProtocolLevel,time time.Time) bool;

	Delete(pool string) bool;

	CleanAll() bool;



}