package storages

import (
	"time"
	"muls"
)

type Mysql struct {

}

func NewMysqlStorage() *Mysql {

	storage := new(Mysql)

	return storage
}

func (self *Mysql)Append(protocol *muls.Protocol) bool {

	return true
}

func (self *Mysql)BatchAppend(protocol []*muls.Protocol) bool {

	return true
}

func (self *Mysql)Read(pool string,level muls.ProtocolAction,limit uint32,order muls.ResultOrder) []string {


}

func (self *Mysql)Clean(pool string, level muls.ProtocolAction,time time.Time) bool {


}

func (self *Mysql)Delete(pool string) bool {


}

func (self *Mysql)CleanAll() bool {


}