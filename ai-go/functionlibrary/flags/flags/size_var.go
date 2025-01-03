package flags

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Size int

func (s *Size) Set(size string) error {
	sizeInBytes, err := FileSizeToByteLen(size)
	if err != nil {
		return err
	}
	*s = Size(sizeInBytes)
	return nil
}

func FileSizeToByteLen(fileSize string) (int, error) {
	fileSize = strings.ToLower(fileSize)
	// default to mb
	if size, err := strconv.Atoi(fileSize); err == nil {
		return size * 1024 * 1024, nil
	}
	if len(fileSize) < 3 {
		return 0, errors.New("invalid size value")
	}
	sizeUnit := fileSize[len(fileSize)-2:]
	size, err := strconv.Atoi(fileSize[:len(fileSize)-2])
	if err != nil {
		return 0, errors.New("parse error: " + err.Error())
	}
	if size < 0 {
		return 0, errors.New("size cannot be negative")
	}
	if strings.EqualFold(sizeUnit, "kb") {
		return size * 1024, nil
	} else if strings.EqualFold(sizeUnit, "mb") {
		return size * 1024 * 1024, nil
	} else if strings.EqualFold(sizeUnit, "gb") {
		return size * 1024 * 1024 * 1024, nil
	} else if strings.EqualFold(sizeUnit, "tb") {
		return size * 1024 * 1024 * 1024 * 1024, nil
	}
	return 0, errors.New("unsupported size unit")
}

func (s *Size) String() string {
	return strconv.Itoa(int(*s))
}

// SizeVar converts the given fileSize with a unit (kb, mb, gb, or tb) to bytes.
// For example, '2kb' will be converted to 2048.
// If no unit is provided, it will fallback to mb. e.g: '2' will be converted to 2097152.
func (flagSet *FlagSet) SizeVar(field *Size, long string, defaultValue string, usage string) *FlagData {
	return flagSet.SizeVarP(field, long, "", defaultValue, usage)
}

// SizeVarP converts the given fileSize with a unit (kb, mb, gb, or tb) to bytes.
// For example, '2kb' will be converted to 2048.
// If no unit is provided, it will fallback to mb. e.g: '2' will be converted to 2097152.
func (flagSet *FlagSet) SizeVarP(field *Size, long, short string, defaultValue string, usage string) *FlagData {
	if field == nil {
		panic(fmt.Errorf("field cannot be nil for flag -%v", long))
	}
	if defaultValue != "" {
		if err := field.Set(defaultValue); err != nil {
			panic(fmt.Errorf("failed to set default value for flag -%v: %v", long, err))
		}
	}
	flagData := &FlagData{
		usage:        usage,
		long:         long,
		defaultValue: defaultValue,
	}
	if short != "" {
		flagData.short = short
		flagSet.CommandLine.Var(field, short, usage)
		flagSet.flagKeys.Set(short, flagData)
	}
	flagSet.CommandLine.Var(field, long, usage)
	flagSet.flagKeys.Set(long, flagData)
	return flagData
}
