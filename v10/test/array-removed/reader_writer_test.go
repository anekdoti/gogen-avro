package avro

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/test/array-removed/reader"
	"github.com/actgardner/gogen-avro/v10/test/array-removed/writer"
	"github.com/actgardner/gogen-avro/v10/vm"
	"os"
	"testing"
)

func TestReaderWriter(t *testing.T) {

	writerSchema, err := os.ReadFile("writer.avsc")
	if err != nil {
		panic(err)
	}

	readerSchema, err := os.ReadFile("reader.avsc")
	if err != nil {
		panic(err)
	}

	prog, err := compiler.CompileSchemaBytes([]byte(writerSchema), []byte(readerSchema))
	if err != nil {
		panic(err)
	}

	writtenRecord := writer.Foo{
		F0: []writer.Bar{
			{16},
			{32},
			{64},
		},
		F1: []int32{1, 2, 3},
		F2: []string{"a", "b"},
	}

	b := make([]byte, 0)
	output := bytes.NewBuffer(b)
	if err := writtenRecord.Serialize(output); err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(output.Bytes()))

	input := bytes.NewReader(output.Bytes())

	readRecord := reader.NewFoo()

	if err := vm.Eval(input, prog, &readRecord); err != nil {
		panic(err)
	}

	bytes, err := readRecord.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
