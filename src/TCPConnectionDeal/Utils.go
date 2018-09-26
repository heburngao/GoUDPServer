package TCPConnectionDeal

import (
	cgolib "BuffUtil"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

var LockRead sync.Mutex
var LockWrite sync.Mutex
var LockRemove sync.Mutex
var LockCopy sync.Mutex

//====================================================
func GetCloneArry(writeBuf []byte) []byte {
	LockCopy.Lock()
	sendbf := make([]byte, len(writeBuf))
	copy(sendbf, writeBuf) //深度复制
	LockCopy.Unlock()
	return sendbf
}

func WriteClient(cmd int, ret int, pbStream []byte) []byte {
	LockWrite.Lock()
	// DB + LEN + CMD + RET + PAYLOAD
	body := &cgolib.WriteBody{make([]byte, 1024), 0}
	body.WriteOPCode3("DB")
	body.WriteShort(cmd)
	body.WriteShort(ret)
	body.WriteBytes(pbStream)
	body.WriteSizeAtlast()
	fmt.Println("######### 写入命令字:", cmd)
	//fmt.Println("sizeofwrite: " , len(cgolib.GetBuff_Write()) , "readSize: " , cgolib.GetSize_write())
	sendBuff := body.GetBuff_Write()[:body.GetSize_write()]

	sendBuff = GetCloneArry(sendBuff)
	fmt.Println("[ end ] >>>>>> 写入字本次总节数:", body.GetSize_write(), " total buff: ", sendBuff)
	LockWrite.Unlock()
	return sendBuff
}
func ReadClient(buffer []byte) (int16, []byte) {
	//LockRead.Lock()
	// DB + LEN + CMD  + PAYLOAD
	//body := &cgolib.ReadBody{make([]byte,1024),0}
	//body.SetBuff_read(buffer)                //设置读取对象
	body := &cgolib.ReadBody{buffer, 0}
	opcode := body.ReadString(2)
	fmt.Println("ReadClient::读取DB:", opcode) //读取ushort opcode
	readlen := body.ReadInt()                //读取 int 长度
	fmt.Println("ReadClient::读取length:", readlen, "检验是否与写入时一致")

	cmd := body.ReadShort() //读取命令字，short
	fmt.Println("ReadClient::########## 读取命令字:", cmd)
	//xxxfmt.Println("ReadClient::index_read: ", body.GetIndex_Read())
	pb := body.ReadBytesByLen(int(readlen) - 8) //读取[ pb字节 ]需去掉DB两字节+长度四字节+cmd两字节 一共八字节
	pb = GetCloneArry(pb)
	fmt.Println("ReadClient::读取payload:", pb, "\n index_read: ", body.GetIndex_Read())
	//xxxfmt.Println("ReadClient::read pb bytes: ", pb)
	//LockRead.Unlock()
	return cmd, pb
}

///==============
func Remove(slice []ConnStruct, elems ...ConnStruct) []ConnStruct {
	//LockRemove.Lock()
	isInElems := make(map[ConnStruct]bool)
	for _, elem := range elems {
		isInElems[elem] = true
	}
	index := 0
	for _, elem := range slice {
		if !isInElems[elem] {
			slice[index] = elem
			index++
		}
	}
	//LockRemove.Unlock()
	return slice[0:index]
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//==== get ip address =====
func GetLocalIp() string {
	var IpAddr string
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		log.Println("Get local IP addr failed!!!")
		IpAddr = "127.0.0.1" //"localhost"
		return IpAddr
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return IpAddr
			}
		}
	}
	IpAddr = "127.0.0.1" //"localhost"
	return IpAddr
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

//这样用
//
//func main() {
//	Try(func() {
//		panic("foo")
//	}, func(e interface{}) {
//		print(e)
//	})
//}
