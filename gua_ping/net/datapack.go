package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"go_basic/gua_ping/itf"
	"go_basic/gua_ping/utils"
)

//封包拆包实例，无需成员
type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d DataPack) GetHeadLen() uint32 {
	return 8
}

func (d DataPack) Pack(msg itf.IMessage) ([]byte, error) {
	//存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//拆包(解压)方法
func (d DataPack) Unpack(binaryData []byte) (itf.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)
	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObj.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObj.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}
	return msg, nil
}
