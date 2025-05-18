package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	tombStoneLive = true
	tombStoneDead = false
)

type KVStore struct {
	File   *os.File
	Index  map[string]int64
	Writer bufio.Writer
}

func NewKVStore(fileName string) (*KVStore, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", fileName, err)
		return nil, err
	}

	return &KVStore{
		File:   file,
		Index:  make(map[string]int64),
		Writer: *bufio.NewWriter(file),
	}, nil
}

func (kv *KVStore) Put(key string, value []byte) error {
	offset, err := kv.File.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Printf("seek file failed, err:%v\n", err)
		return err
	}
	// Format : [keyLen][valueLen][tombstone][key][value]

	if err := binary.Write(&kv.Writer, binary.LittleEndian, int32(len(key))); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return err
	}
	if err := binary.Write(&kv.Writer, binary.LittleEndian, int32(len(value))); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return err
	}

	if err := binary.Write(&kv.Writer, binary.LittleEndian, tombStoneLive); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return err
	}

	if _, err := kv.Writer.WriteString(key); err != nil {
		fmt.Printf("write key failed, err:%v\n", err)
		return err
	}
	if _, err := kv.Writer.Write(value); err != nil {
		fmt.Printf("write key failed, err:%v\n", err)
		return err
	}
	kv.Index[key] = offset
	return kv.Writer.Flush()
}

func (kv *KVStore) Delete(key string) (bool, error) {
	_, err := kv.File.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Printf("seek file failed, err:%v\n", err)
		return false, err
	}

	if err = binary.Write(&kv.Writer, binary.LittleEndian, int32(len(key))); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return false, err
	}
	// writing 0 as value length
	if err = binary.Write(&kv.Writer, binary.LittleEndian, int32(0)); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return false, err
	}

	if err = binary.Write(&kv.Writer, binary.LittleEndian, tombStoneDead); err != nil {
		fmt.Printf("write key length failed, err:%v\n", err)
		return false, err
	}

	if _, err = kv.Writer.WriteString(key); err != nil {
		fmt.Printf("write key failed, err:%v\n", err)
		return false, err
	}

	delete(kv.Index, key)
	return true, nil
}

func (kv *KVStore) Get(key string) ([]byte, error) {
	offSet, ok := kv.Index[key]
	if !ok {
		fmt.Printf("key %s not found\n", key)
		return []byte{}, nil
	}

	if _, err := kv.File.Seek(offSet, io.SeekStart); err != nil {
		fmt.Printf("seek file failed, err:%v\n", err)
		return nil, err
	}
	var keyLen, valueLen int32
	var tombStone bool
	if err := binary.Read(kv.File, binary.LittleEndian, &keyLen); err != nil {
		return nil, err
	}
	if err := binary.Read(kv.File, binary.LittleEndian, &valueLen); err != nil {
		return nil, err
	}

	if err := binary.Read(kv.File, binary.LittleEndian, &tombStone); err != nil {
		return nil, err
	}

	keyBytes := make([]byte, keyLen)
	if _, err := kv.File.Read(keyBytes); err != nil {
		fmt.Printf("read key failed, err:%v\n", err)
		return nil, err
	}

	valueBytes := make([]byte, valueLen)
	if _, err := kv.File.Read(valueBytes); err != nil {
		fmt.Printf("read value failed, err:%v\n", err)
		return nil, err
	}

	return valueBytes, nil
}

func main() {
	store, err := NewKVStore("data.log")
	if err != nil {
		fmt.Printf("create kv store failed, err:%v\n", err)
		return
	}
	defer func(File *os.File) {
		err := File.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
		}
	}(store.File)
	store.Put("divya", []byte("thakur"))
	store.Put("neelu", []byte("kanwar"))
	store.Put("divya", []byte("THAKUR"))

	value, err := store.Get("divya")
	if err != nil {
		fmt.Printf("get value failed, err:%v\n", err)
		return
	}
	fmt.Printf("value for key 'divya': %s\n", string(value))

	value, err = store.Get("neelu")
	if err != nil {
		fmt.Printf("get value failed, err:%v\n", err)
		return
	}
	fmt.Printf("value for key 'neelu': %s\n", string(value))

	_, err = store.Delete("neelu")
	if err != nil {
		fmt.Printf("delete value failed, err:%v\n", err)
		return
	}

	store.Get("neelu")

}
