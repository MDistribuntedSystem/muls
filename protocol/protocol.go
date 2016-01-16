package protocol

/**

	| 4 bytes 		| 4 bytes | 32 bytes | 4 bytes | 4 bytes |4 bytes |  N bytes |

	| magic_code|	| ActionName| PoolName | level | time | log |
 */


import (
	"bytes"
	"fmt"
	"encoding/binary"
)

const PROTOCOL_NAME_SIZE = 32;
const PROTOCOL_LOG_SIZE  = 32000;

const PROTOCOL_DATA_MIN_LEN = 4 + 4 + 32 + 4 + 4 + 4;

const PROTOCOL_MAGIC_CODE = 0xEF000000 | 0x00A90000 | 0x0000B100 | 0x000000FF

type Protocol struct {

	magic_code uint32;

	action uint32;

	pool [32]byte;

	level uint32;

	time uint32;

	data_len uint32

	log *bytes.Buffer;

}


func NewProtocol(action uint32,pool string,level uint32 ,time uint32) *Protocol{

	protocol := new(Protocol)

	protocol.action = action

	pool_bytes := []byte(pool)

	for i:=0;i<len(pool_bytes);i++ {

		protocol.pool[i] = pool_bytes[i]

	}

	protocol.level = level

	protocol.time = time

	protocol.data_len = 0

	protocol.magic_code = PROTOCOL_MAGIC_CODE

	return protocol;

}



func NewProtocolWithBytes(raw []byte) *Protocol {

	fmt.Println(raw)

	if len(raw) < PROTOCOL_DATA_MIN_LEN {

		fmt.Printf("Protocol Error\n")

		return nil

	}

	dian := 1;

	if( raw[0] == PROTOCOL_MAGIC_CODE >> (0 * 8) & 0xFF &&
		raw[1] == PROTOCOL_MAGIC_CODE >> (1 * 8) & 0xFF &&
		raw[2] == PROTOCOL_MAGIC_CODE >> (2 * 8) & 0xFF &&
		raw[3] == PROTOCOL_MAGIC_CODE >> (3 * 8) & 0xFF ){

		dian = 0;
	}

	protocol := new(Protocol)

	buf := bytes.NewReader(raw)

	if(dian == 1){

		binary.Read(buf,binary.BigEndian,&protocol.magic_code)
		binary.Read(buf,binary.BigEndian,&protocol.action)
		binary.Read(buf,binary.BigEndian,&protocol.pool)
		binary.Read(buf,binary.BigEndian,&protocol.level)
		binary.Read(buf,binary.BigEndian,&protocol.time)
		binary.Read(buf,binary.BigEndian,&protocol.data_len)

	}else{

		binary.Read(buf,binary.LittleEndian,&protocol.magic_code)
		binary.Read(buf,binary.LittleEndian,&protocol.action)
		binary.Read(buf,binary.LittleEndian,&protocol.pool)
		binary.Read(buf,binary.LittleEndian,&protocol.level)
		binary.Read(buf,binary.LittleEndian,&protocol.time)
		binary.Read(buf,binary.LittleEndian,&protocol.data_len)

	}

	return protocol

}

func (self *Protocol)Fmt() string {

	buffer := fmt.Sprintf("%d %s %d %d",self.action,self.pool,self.level,self.time)

	return buffer

}


func (self *Protocol)Pack() []byte {

	buf := new (bytes.Buffer)

	binary.Write(buf,binary.BigEndian,uint32(self.magic_code))
	binary.Write(buf,binary.BigEndian,uint32(self.action))
	binary.Write(buf,binary.BigEndian,self.pool)
	binary.Write(buf,binary.BigEndian,uint32(self.level))
	binary.Write(buf,binary.BigEndian,uint32(self.time))
	binary.Write(buf,binary.BigEndian,uint32(self.data_len))

	return buf.Bytes()
}





