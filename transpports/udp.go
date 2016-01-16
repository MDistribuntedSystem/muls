package transpports
import (
	"net"
	"fmt"
	"bytes"
	"muls"
)

const MAX_BUFFER_SIZE = 1//10M

type UdpServer struct {

	addr net.UDPAddr;

	conn *net.UDPConn;

	active bool;

	buffer *bytes.Buffer

}

func NewUdpServer(ip string,port int) *UdpServer {

	server := new(UdpServer)

	server.addr = net.UDPAddr{
		IP	: net.ParseIP(ip),
		Port:port,
	}

	server.active = true
	server.conn = nil

	server.buffer = bytes.NewBufferString("")

	return server

}

func (self *UdpServer)Run() bool {

	ser,err := net.ListenUDP("udp",&self.addr)

	if err != nil {

		fmt.Printf("Start Listen Server Error:%v\n",err)

		return false
	}

	self.conn = ser;


	self.Loop()

	return true
}

func (self *UdpServer)Stop() bool{

	return true
}


func (self *UdpServer)Loop(){

	buffer := make([]byte,6550);

	for {

		size,remoteAddr,err := self.conn.ReadFromUDP(buffer)

		if err != nil {

			fmt.Printf("Some error  %v", err)
			continue

		}

		 self.Read(buffer,size)

		go self.Send(remoteAddr,buffer)

	}

}

func (self *UdpServer) Read(buffer []byte,size int){


	protocol := muls.NewProtocolWithBytes(buffer[0:size])

	self.buffer.WriteString(protocol.Fmt())
	self.buffer.WriteString("\n")

	fmt.Printf("%d %d\n",self.buffer.Len(),size)

	if self.buffer.Len() + size >= MAX_BUFFER_SIZE{
		fmt.Printf("%s",self.buffer.Bytes())
		self.buffer.Reset()
	}

}


func (self *UdpServer) Send(remote *net.UDPAddr,buffer []byte) {



	self.conn.WriteToUDP([]byte("abcd"),remote)

}
