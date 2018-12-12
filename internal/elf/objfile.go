package elf

import (
	"debug/elf"
	"io"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

type elfFile struct {
	elf *elf.File
}

func openElf(r io.ReaderAt) (*elfFile, error) {
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}
	return &elfFile{f}, nil
}

// Open loads elf format binary file.
func Open(name string) (*elfFile, error) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return openElf(f)
}

// GetFuncBytes loads function assembler instructions.
func GetFuncBytes(elfFile *elfFile, funcName string) []byte {
	symbols, _ := elfFile.elf.Symbols()
	var fSymbol elf.Symbol
	for _, s := range symbols {
		if strings.Contains(s.Name, funcName) {
			fSymbol = s
		}
	}
	textSection := elfFile.elf.Section(".text")
	fAddr := fSymbol.Value - textSection.Addr
	textb := getSectionBytes(textSection)
	pp.Println(len(textb), fAddr, fAddr+fSymbol.Size)
	return textb[fAddr : fAddr+fSymbol.Size]
}

func getSectionBytes(s *elf.Section) []byte {
	b, _ := s.Data()
	return b
}
