package KCPUtil

import (
	//	//util "ConnectionDeal"
	"bytes"
	"fmt"
)

const BUF_MAX int32 = 1024

type CombineBody struct {
	Index_getCombineSize int
	//================= 粘包相关 =================
	/// <summary>
	/// 读取的缓存带
	/// </summary>
	Buffer []byte //= make([]byte,BUF_MAX) //new byte[BUF_MAX] //缓存包

	/// <summary>
	/// 残包 ，缓存带
	/// </summary>
	TooShortPackage []byte //= nil//　短包
	/// <summary>
	/// 读取缓存 整包
	/// </summary>
	CombinedBuff []byte //= make([]byte,BUF_MAX) //　最终的包

	//		private static int _i

	/// <summary>
	/// 残包， 粘包 下标
	/// </summary>
	BufIndex int
}

/// <summary>
/// 读取整包，且返回长度
/// </summary>
/// <returns>The real read buffer.</returns>
/// <param name="byteSize">Byte size.</param>
func (p *CombineBody) GetRealReadBuf(byteSize int) (int, []byte) {

	p.BufIndex = 0
	if p.TooShortPackage != nil { //合并上次收到的残余包
		combineSize := len(p.TooShortPackage) + byteSize
		p.CombinedBuff = make([]byte, combineSize)
		//DebugTool.LogBlue(string.Format("合包后，长: {0}", realReadBuf.Length))  TODO 注释掉
		//Array.Copy ( TooShortPackage, 0, realReadBuf, 0,  TooShortPackage.Length);//把上次收到的太短的包粘上
		//Array.Copy (buffer, 0, realReadBuf,  TooShortPackage.Length, byteSize);//在太短的包后面，追加新收到的buffer ->

		copy(p.CombinedBuff, p.TooShortPackage) //把上次收到的太短的包粘上 <-
		//copy(RealReadBuf, buffer)          //在太短的包后面，追加新收到的buffer <-
		cpbytes := p.GetCloneArry(p.Buffer) //对新收到的buffer作一次copy
		//RealReadBuf = append(RealReadBuf,Buffer)//在太短的包后面，追加新收到的buffer
		p.CombinedBuff = p.BytesCombine(p.CombinedBuff, cpbytes) //在太短的包后面，追加新收到的buffer
		p.TooShortPackage = nil
	} else {
		p.CombinedBuff = make([]byte, byteSize)
		//Array.Copy (buffer, 0, realReadBuf, 0, byteSize)  ->
		copy(p.CombinedBuff, p.Buffer) // <-
	}
	return len(p.CombinedBuff), p.CombinedBuff
}
func (p *CombineBody) BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
func (p *CombineBody) GetCloneArry(writeBuf []byte) []byte {
	sendbf := make([]byte, len(writeBuf))
	copy(sendbf, writeBuf) //深度复制
	return sendbf
}

/// <summary>
/// 按fullSize指定尺寸截取一段buff , bufIndex += fullSize 向后顺推
/// </summary>
/// <returns>The real read buffer item.</returns>
/// <param name="bytesRead">Bytes read.</param>
func (p *CombineBody) GetUsefulBuff(fullSize int) []byte {
	//xxxvar fullSizeBuff []byte = make([]byte, bytesRead)
	//Array.Copy (realReadBuf, bufIndex, fullSizeBuff, 0, bytesRead)  ->  c# version
	endindex := p.BufIndex + fullSize
	//xxxcopy(fullSizeBuff, CombinedBuff[BufIndex:endindex]) // <-

	fmt.Println("endindex:", endindex, "len(CombinedBuff): ", len(p.CombinedBuff), "BuffIndex: ", p.BufIndex, " fullSize:", fullSize)
	fullUsefulBuff := p.GetCloneArry(p.CombinedBuff[p.BufIndex:endindex])
	p.BufIndex += fullSize
	return fullUsefulBuff
}

/// <summary>
/// 残包,缓存起来
/// </summary>
/// <param name="bytesRead"></param>
func (p *CombineBody) TooShortReceive(bytesRead int) {
	p.TooShortPackage = make([]byte, bytesRead) //new byte[bytesRead]
	//Array.Copy(realReadBuf, bufIndex,  TooShortPackage, 0, bytesRead) ->
	copy(p.TooShortPackage, p.CombinedBuff[p.BufIndex:]) //<-

}

//		public static short GetOpcode ()//"WO"
//		{
//			// 要预留2个字节出来记录Size的
//			return (short)((realReadBuf [0] << 8) + realReadBuf [1])
////			return (short)((realReadBuf [bufIndex] << 8) + realReadBuf [1])
//		}
//----------------读取长度--------------------------
//		private static int _msgsize

/// <summary>
/// 读取当前字节序 长度
/// </summary>
/// <returns>The send buffer length.</returns>
func (p *CombineBody) GetMsgSize() int {

	//return (int)(realReadBuf[2] << 24) + (realReadBuf[3] << 16) + (realReadBuf[4] << 8) + realReadBuf[5] //超长包粘包
	u0 := int(p.CombinedBuff[2]) << 24
	u1 := int(p.CombinedBuff[3]) << 16
	u2 := int(p.CombinedBuff[4]) << 8
	u3 := int(p.CombinedBuff[5])
	return u3 + u2 + u1 + u0
}
func (p *CombineBody) GetMsgSize2(startIndex int32) int32 {
	fmt.Println("GeMsgSize2::### 1 ### startIndex: ", startIndex, "len(p.CombineBuff): ", len(p.CombinedBuff))
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err,"GetMsgSize2 fault")
	//	}
	//}()
	//panic("GetMsgSize2 error")
	startIndex += 2 //绕过opcode
	//fmt.Println("GeMsgSize2::### 2 ### startIndex: ",startIndex, "len(p.CombineBuff): " , len(p.CombinedBuff))

	//return (int)((realReadBuf[startIndex++] << 24) | (realReadBuf[startIndex++] << 16) | (realReadBuf[startIndex++] << 8) | realReadBuf[startIndex++]) //如果是粘包(两包合一)，则继续bufIndex++
	u0 := int32(p.CombinedBuff[startIndex]) << 24
	startIndex++
	//fmt.Println("GeMsgSize2::### 3 ### startIndex: ",startIndex, "len(p.CombineBuff): " , len(p.CombinedBuff))
	u1 := int32(p.CombinedBuff[startIndex]) << 16
	startIndex++
	//fmt.Println("GeMsgSize2::### 4 ### startIndex: ",startIndex, "len(p.CombineBuff): " , len(p.CombinedBuff))
	u2 := int32(p.CombinedBuff[startIndex]) << 8
	startIndex++
	//fmt.Println("GeMsgSize2::### 5 ### startIndex: ",startIndex, "len(p.CombineBuff): " , len(p.CombinedBuff))
	u3 := int32(p.CombinedBuff[startIndex])
	startIndex++
	//fmt.Println("GeMsgSize2::### 6 ### startIndex: ",startIndex, "len(p.CombineBuff): " , len(p.CombinedBuff))
	return u3 | u2 | u1 | u0
	//return  u3+u2+u1+u0
}
