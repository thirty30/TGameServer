package rof

import (
	"TCore"
	"encoding/binary"
	"os"
)

type iRofRow interface {
	readBody(aBuffer []byte) int32
}

type iRofTable interface {
	newTypeObj() iRofRow
	setRowNum(aNum int32)
	setColNum(aNum int32)
	setIDMap(aKey int32, aValue iRofRow)
	setRowMap(aKey int32, aValue int32)
}

func analysisRof(aPath string, aTable iRofTable) bool {
	pFile, err := os.Open(aPath)
	if err != nil {
		return false
	}
	defer pFile.Close()
	pBuffer := make([]byte, 1024*1024*4)
	nLen, err := pFile.Read(pBuffer)
	if err != nil {
		return false
	}
	if nLen == 1024*1024*4 {
		return false
	}

	nOffset := 64
	nRowNum := int32(binary.BigEndian.Uint32(pBuffer[nOffset:]))
	aTable.setRowNum(nRowNum)
	nOffset += 4
	nColNum := int32(binary.BigEndian.Uint32(pBuffer[nOffset:]))
	aTable.setColNum(nColNum)
	nOffset += 4
	//解析头
	for i := 0; i < int(nColNum); i++ {
		nNameLen := int8(pBuffer[nOffset])
		nOffset += 1 + int(nNameLen) + 2
	}
	//解析行
	for i := 0; i < int(nRowNum); i++ {
		nID := int32(binary.BigEndian.Uint32(pBuffer[nOffset:]))
		pData := aTable.newTypeObj()
		nSize := pData.readBody(pBuffer[nOffset:])
		nOffset += int(nSize)
		aTable.setIDMap(nID, pData)
		aTable.setRowMap(int32(i), nID)
	}

	return true
}

//RofManager export
type RofManager struct {
	mTableExample sRofExampleTable
}

//Init exported
func (pOwn *RofManager) Init() error {
	if pOwn.mTableExample.init("./rofs/RofExample.bytes") == false {
		return tcore.TNewError("init RofExample fail...")
	}

	return nil
}

func (pOwn *RofManager) GetExampleTable() *sRofExampleTable {
	return &pOwn.mTableExample
}
