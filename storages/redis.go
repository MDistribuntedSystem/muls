package storages
import (
	"time"
	"muls"
)

type Redis struct {

}

func NewRedisStorage() *Redis {

	storage := new(Redis)

	return storage
}

func (self *Redis)Append(protocol *muls.Protocol) bool {

	return true
}

func (self *Redis)BatchAppend(protocol []*muls.Protocol) bool {

	return true
}

func (self *Redis)Read(pool string,level muls.ProtocolAction,limit uint32,order muls.ResultOrder) []string {


}

func (self *Redis)Clean(pool string, level muls.ProtocolAction,time time.Time) bool {


}

func (self *Redis)Delete(pool string) bool {


}

func (self *Redis)CleanAll() bool {


}