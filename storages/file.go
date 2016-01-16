package storages
import (
	"os"
	"time"
	"muls"
)

type File struct {

	path 		string
	handler 	*os.File

}

func NewFileStorage(path string) *File {

	storage := new(File)

	storage.path = path

	return storage
}

func (self *File)Append(protocol *muls.Protocol) bool {

	return true
}

func (self *File)BatchAppend(protocol []*muls.Protocol) bool {

	return true
}

func (self *File)Read(pool string,level muls.ProtocolAction,limit uint32,order muls.ResultOrder) []string {


}

func (self *File)Clean(pool string, level muls.ProtocolAction,time time.Time) bool {


}

func (self *File)Delete(pool string) bool {


}

func (self *File)CleanAll() bool {


}

/* ============================================ Private Method ==================================== */


func (self *File) open() bool {

	return true
}
