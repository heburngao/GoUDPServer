package TCPConnectionDeal

import (
	cgolib "BuffUtil"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protocol/msg"
	"reflect"
)

type B struct {
	a string
	b byte
	c int
}

func (b B) test() {
	fmt.Println("this is a reflect test")
}

//===================================================================================================================

func Test() {

	fmt.Println("测试反射开始...")
	b := new(B)
	m, _ := reflect.TypeOf(b).MethodByName("test")

	fmt.Println(m.PkgPath, "测试反射结束")

	//protobuf========
	dataPB := &msg.Helloworld{
		Id:  proto.Int32(1000),
		Str: proto.String("你好世界"),
		Opt: proto.Int32(999),
	}
	//encode
	pb_encodede, err := proto.Marshal(dataPB)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	//decode
	dataDecode := &msg.Helloworld{}
	err = proto.Unmarshal(pb_encodede, dataDecode)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}

	//test run
	if dataPB.GetId() != dataDecode.GetId() {
		log.Fatal("data mismatch %q != %q", dataPB.GetId(), dataDecode.GetId())
	}
	fmt.Println("ID:", dataDecode.GetId())
	fmt.Println("Str:", dataDecode.GetStr())
	fmt.Println("Opt:", dataDecode.GetOpt())
	//=============

	body := &cgolib.WriteBody{make([]byte, 1024), 0}
	body.WriteOPCode3("DB")
	body.WriteString("OOXX")
	body.WriteInt(2017)
	body.WriteSizeAtlast()
	fmt.Println("写入字节数: ", body.GetSize_write())
	fmt.Println("开始读取>>>>>>>>>")

	//body.SetBuff_read(body.GetBuff_Write())
	body2 := &cgolib.ReadBody{body.GetBuff_Write(), 0}

	fmt.Println("读取DB:", body2.ReadString(2))
	fmt.Println("读取length:", body2.ReadInt(), "检验是否与写入时一致")
	fmt.Println("读取OOXX:", body2.ReadString(4))
	fmt.Println("读取int:", body2.ReadInt())
}
