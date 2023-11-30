/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:03:01
*/
package main

import (
	"debug/dwarf"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"log"
)

func findSymbol(d *dwarf.Data, name string) error {
	dr := d.Reader()
	var lr *dwarf.LineReader = nil
	var files []*dwarf.LineFile = nil
	for {
		e, err := dr.Next()
		if e == nil || err != nil {
			break
		}
		if e.Tag == dwarf.TagCompileUnit {
			lr, err = d.LineReader(e)
			if err != nil {
				log.Println(err)
				lr = nil
				break
			}
			files = lr.Files()
			for i := 0; i < len(files); i++ {
				//log.Println(files[i].Name)
			}
		}
	}
	dr.Seek(0)
	for {
		e, err := dr.Next()
		if e == nil || err != nil {
			return err
		}

		aname, ok := e.Val(dwarf.AttrName).(string)
		//log.Println(aname, e.Tag)
		if !ok || aname != name {
			continue
		}
		switch e.Tag {
		case dwarf.TagVariable: // 变量
			loc, ok := e.Val(dwarf.AttrLocation).([]uint8)
			if !ok {
				continue
			}
			if loc[0] != 3 {
				return fmt.Errorf("can't determine variable addr")
			}
			lineNo, ok := e.Val(dwarf.AttrDeclLine).(int64)
			if !ok {
				lineNo = -1
			}
			fileNo, _ := e.Val(dwarf.AttrDeclFile).(int64)
			addr := uint64(0)
			switch len(loc) {
			case 5:
				addr = uint64(binary.LittleEndian.Uint32(loc[1:]))
			case 9:
				addr = uint64(binary.LittleEndian.Uint64(loc[1:]))
			default:
				return fmt.Errorf("unknown addr size")
			}

			off, ok := e.Val(dwarf.AttrType).(dwarf.Offset)
			if !ok {
				continue
			}
			typ, err := d.Type(off)
			if err != nil {
				return err
			}
			log.Printf("%s <%x> %s %s:%d\n", aname, addr, typ.String(), files[fileNo].Name, lineNo)
			return nil
		case dwarf.TagSubprogram:
			pc, ok := e.Val(dwarf.AttrLowpc).(uint64)
			if ok {
				lineNo, ok := e.Val(dwarf.AttrDeclLine).(int64)
				if !ok {
					lineNo = -1
				}
				fileNo, _ := e.Val(dwarf.AttrDeclFile).(int64)
				log.Printf("%s <%x> %s:%d\n", aname, pc, files[fileNo].Name, lineNo)
			} else {
				lineNo, ok := e.Val(dwarf.AttrDeclLine).(int64)
				if !ok {
					lineNo = -1
				}
				fileNo, _ := e.Val(dwarf.AttrDeclFile).(int64)
				log.Printf("%s %s:%d\n", aname, files[fileNo].Name, lineNo)
			}
		}
	}
}

func FindSymbol(elf_path string, vari string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	file, e1 := elf.Open(elf_path)
	if e1 == nil {
		defer file.Close()
		dwlf, e2 := file.DWARF()
		if e2 == nil {
			findSymbol(dwlf, vari)
		}
	}
}

func main() {
	FindSymbol("./c_out/a.out", "g_a")
	FindSymbol("./c_out/a.out", "add")
	FindSymbol("./c_out/a.out", "g_ar")
}
