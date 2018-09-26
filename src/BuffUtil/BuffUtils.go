package BuffUtil //被引用包 ，放在 src的 BuffUtil 目录中
// ========= 读取写入值
import (
	"fmt"
)
// fmt.Println("读取OPCODE:",ReadString(2))
// fmt.Println("读取字符串总长度:",ReadInt())
// fmt.Println("读取字符串:",ReadString(3))

// fmt.Println("读取整型：",ReadInt())
// fmt.Println("读取长整型：",ReadLong())
// fmt.Println("写入时的总长度:" ,GetSize_write())
// // fmt.Println(len(buffer_write) ,"|", len(buffer_read))

// fmt.Println(">>>>>>>>>>> init Socket server")
// netListen,err:=net.Listen("tcp","127.0.0.1:2020")
// checkErr(err)
// defer netListen.Close()
// log.Println(">>>>>>>>>> waiting for clients >>>>>>")
// for{
//     conn,err :=netListen.Accept()
// 	if err != nil{
// 		continue
// 	}

// 	log.Println(conn.RemoteAddr().String(),"tcp connect success!")
// 	go handlerConnection(conn)
// }//end for
//翻转byte数组
func ExchangeBytes(byteA []byte) {
	var temp byte
	ll := len(byteA)
	for i := 0; i < ll/2; i++ {
		temp = byteA[ll-i-1]
		byteA[ll-1-i] = byteA[i]
		byteA[i] = temp
	}
}

//从字节序读取
type ReadBody struct {
	 Buffer_read []byte//   := make([]byte, 1024)
	 Index_read int
}

func (p *ReadBody)GetBuff_Read() []byte {
	return p.Buffer_read
}
func (p *ReadBody)GetBuff_Read_HasData() []byte {
	return p.Buffer_read[0:p.Index_read]
}
func (p *ReadBody)ResetIndex_Read() {
	p.Index_read = 0
}
func (p *ReadBody)SetIndex_Read(v int) {
	p.Index_read = v
}
func (p * ReadBody)GetIndex_Read() int {
	return p.Index_read
}
func (p *ReadBody)SetBuff_read(buff []byte) {
	p.Buffer_read = buff
	p.Index_read = 0
}

func (p *ReadBody)ReadBool() bool {
	boo := p.Buffer_read[p.Index_read] != 0
	p.Index_read++
	return boo
}
func (p *ReadBody)ReadBytes() []byte {
	length := len(p.Buffer_read) - p.Index_read
	buff := p.Buffer_read[p.Index_read : p.Index_read+length]
	p.Index_read += length
	return buff
}
func (p *ReadBody)ReadBytesByLen(length int) []byte {
	buff := p.Buffer_read[p.Index_read : p.Index_read+length]
	p.Index_read += length
	return buff
}
func (p *ReadBody)ReadByte() byte {
	byt := p.Buffer_read[p.Index_read]
	p.Index_read++
	return byt
}
func (p *ReadBody)ReadShort() int16 {
	u0 := int16(p.Buffer_read[p.Index_read]) << 8
	// u0 := int(buffer_read[index_read]) << 8
	p.Index_read++
	u1 := int16(p.Buffer_read[p.Index_read])
	// u1 := int(buffer_read[index_read])
	p.Index_read++
	return u1 | u0
}
func (p *ReadBody)ReadInt() int32 {
	//defer checkError(p)
	u0 := int32(p.Buffer_read[p.Index_read]) << 24
	p.Index_read++
	u1 := int32(p.Buffer_read[p.Index_read]) << 16
	p.Index_read++
	u2 := int32(p.Buffer_read[p.Index_read]) << 8
	p.Index_read++
	u3 := int32(p.Buffer_read[p.Index_read])
	p.Index_read++
	return u3 | u2 | u1 | u0
	// u:=binary.BigEndian.int32(buffer_read[index_read:index_read+4])
	// index_read += 4
	// return u
}
func checkError(p * ReadBody){
	if err:= recover(); err != nil{
		fmt.Println("err:",err , " ...buf len:  " , len(p.Buffer_read) , " index: " , p.Index_read)
	}
}
func (p *ReadBody)ReadLong() int64 {
	u0 := int64(p.Buffer_read[p.Index_read]) << 56
	p.Index_read++
	u1 := int64(p.Buffer_read[p.Index_read]) << 48
	p.Index_read++
	u2 := int64(p.Buffer_read[p.Index_read]) << 40
	p.Index_read++
	u3 := int64(p.Buffer_read[p.Index_read]) << 32
	p.Index_read++
	u4 := int64(p.Buffer_read[p.Index_read]) << 24
	p.Index_read++
	u5 := int64(p.Buffer_read[p.Index_read]) << 16
	p.Index_read++
	u6 := int64(p.Buffer_read[p.Index_read]) << 8
	p.Index_read++
	u7 := int64(p.Buffer_read[p.Index_read])
	p.Index_read++
	return u7 | u6 | u5 | u4 | u3 | u2 | u1 | u0
}
func (p *ReadBody)ReadString(length int) string {
	str := string(p.Buffer_read[p.Index_read : p.Index_read+length])
	p.Index_read += length //len(buffer_read) - index_read
	return str
}

//================ 写入到字节序 =================

//var buffer_write []byte = make([]byte, 1024)
//var p.Index_write int
type WriteBody struct {
	Buffer_write []byte//   := make([]byte, 1024)
	Index_write int
}
func (p *WriteBody)GetBuff_Write() []byte {
	return p.Buffer_write
}
func (p *WriteBody)GetBuff_Write_HasData() []byte {
	return p.Buffer_write[0:p.Index_write]
}
func (p *WriteBody)ResetIndex_Write() {
	p.Index_write = 0
}
func (p *WriteBody)SetIndex_Write(index int) {
	p.Index_write = index
}
func (p *WriteBody)GetIndex_Write() int {
	return p.Index_write
}
func (p *WriteBody)WriteOPCode3(opcode string) {
	// WriteBytes(2)
	byts := []byte(opcode)
	p.Buffer_write[0] = byts[0]
	p.Buffer_write[1] = byts[1]
	p.Index_write = 6
}
func (p *WriteBody) WriteOPCode2(opcode []byte) {
	// WriteBytes(2)
	p.Buffer_write[0] = opcode[0]
	p.Buffer_write[1] = opcode[1]
	p.Index_write = 6
}
func (p *WriteBody) WriteOPCode(opcode int) {
	// WriteBytes(2)
	s := int16(opcode)
	// s := opcode
	p.Buffer_write[0] = byte(s) >> 8
	p.Buffer_write[1] = byte(s) & 0xff
	p.Index_write = 6
}
func (p *WriteBody) WriteSizeAtlast() {
	s := int32(p.Index_write)
	p.Buffer_write[2] = byte(s) >> 24
	p.Buffer_write[3] = byte(s) >> 16 & 0xff
	p.Buffer_write[4] = byte(s) >> 8 & 0xff
	p.Buffer_write[5] = byte(s) & 0xff
	p.Index_write = 0 //在每次 buffer_write的结束时，写入长度，然后重置p.Index_write方便以后写入
}

//只能在 WriteSizeAtlast() 之后
func (p *WriteBody) GetSize_write() int32 {
	p.Index_write = 2 //跳过opcode的下标
	u0 := int32(p.Buffer_write[p.Index_write]) << 24
	p.Index_write++
	u1 := int32(p.Buffer_write[p.Index_write]) << 16
	p.Index_write++
	u2 := int32(p.Buffer_write[p.Index_write]) << 8
	p.Index_write++
	u3 := int32(p.Buffer_write[p.Index_write]) & 0xff
	p.Index_write = 0 //重置
	return u3 | u2 | u1 | u0
}

func (p *WriteBody) WriteBool(b bool) {
	var bb byte
	if b == true {
		bb = 1
	} else {
		bb = 0
	}
	p.Buffer_write[p.Index_write] = byte(bb) //byte(b?1:0)
	p.Index_write++
}
func (p *WriteBody) WriteByte(b byte) {
	p.Buffer_write[p.Index_write] = b
	p.Index_write++
}
func (p *WriteBody) WriteBytes(byts []byte) {
	// copy(buffer_write, byts)
	// p.Index_write += len(byts)
	for i := 0; i < len(byts); i++ {
		p.Buffer_write[p.Index_write] = byts[i]
		p.Index_write++
	}

}
func (p *WriteBody) WriteShort(d int) {
	s := int16(d)
	// s := d
	s0 := byte(s >> 8)
	p.Buffer_write[p.Index_write] = s0
	p.Index_write++
	s1 := byte(s & 0xff)
	p.Buffer_write[p.Index_write] = s1
	p.Index_write++

}

func (p *WriteBody) WriteInt(d int) {
	s := int32(d)
	s0 := byte(s >> 24)
	p.Buffer_write[p.Index_write] = s0
	p.Index_write++
	s1 := byte(s >> 16 & 0xff)
	p.Buffer_write[p.Index_write] = s1
	p.Index_write++
	s2 := byte(s >> 8 & 0xff)
	p.Buffer_write[p.Index_write] = s2
	p.Index_write++
	s3 := byte(s & 0xff)
	p.Buffer_write[p.Index_write] = s3
	p.Index_write++
	// binary.BigEndian.PutUint32(buffer_write[p.Index_write:p.Index_write+4],s)
	// p.Index_write += 4

}
func (p *WriteBody) WriteLong(d int) {
	s := int64(d)
	s0 := byte(s >> 56)
	p.Buffer_write[p.Index_write] = s0
	p.Index_write++
	s1 := byte(s >> 48 & 0xff)
	p.Buffer_write[p.Index_write] = s1
	p.Index_write++
	s2 := byte(s >> 40 & 0xff)
	p.Buffer_write[p.Index_write] = s2
	p.Index_write++
	s3 := byte(s >> 32 & 0xff)
	p.Buffer_write[p.Index_write] = s3
	p.Index_write++
	s4 := byte(s >> 24 & 0xff)
	p.Buffer_write[p.Index_write] = s4
	p.Index_write++
	s5 := byte(s >> 16 & 0xff)
	p.Buffer_write[p.Index_write] = s5
	p.Index_write++
	s6 := byte(s >> 8 & 0xff)
	p.Buffer_write[p.Index_write] = s6
	p.Index_write++
	s7 := byte(s & 0xff)
	p.Buffer_write[p.Index_write] = s7
	p.Index_write++

}
func (p *WriteBody) WriteString(str string) {
	byts := []byte(str)
	p.WriteBytes(byts)
}
