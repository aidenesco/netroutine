package netroutine

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Status int

const (
	Success Status = iota
	Fail
	Retry
	Error
	Custom
)

func (s Status) String() string {
	return [...]string{"Success", "Fail", "Retry", "Error", "Custom"}[s]
}

var (
	//Success = errors.New("success")
	//Fail    = errors.New("fail")
	//Retry   = errors.New("retry")
	//Error   = errors.New("error")
	//Custom  = errors.New("custom")

	blocks = map[string]interface{}{
		idBlockAppId:             BlockAppId{},
		idBlockBase64Decode:      BlockBase64Decode{},
		idBlockBase64Encode:      BlockBase64Encode{},
		idBlockBasicAuth:         BlockBasicAuth{},
		idBlockFlagTimePassed:    BlockFlagTimePassed{},
		idBlockFlagVariables:     BlockFlagVariables{},
		idBlockGenerateString:    BlockGenerateString{},
		idBlockJSONBuilder:       BlockJSONBuilder{},
		idBlockMathAdd:           BlockMathAdd{},
		idBlockMathCeil:          BlockMathCeil{},
		idBlockMathCondition:     BlockMathCondition{},
		idBlockMathDivide:        BlockMathDivide{},
		idBlockMathFloor:         BlockMathFloor{},
		idBlockMathMultiply:      BlockMathMultiply{},
		idBlockMathSubtract:      BlockMathSubtract{},
		idBlockMD5Hash:           BlockMD5Hash{},
		idBlockParseCookies:      BlockParseCookies{},
		idBlockParseHeader:       BlockParseHeader{},
		idBlockParseJSON:         BlockParseJSON{},
		idBlockParseLR:           BlockParseLR{},
		idBlockParseRegex:        BlockParseRegex{},
		idBlockParseTime:         BlockParseTime{},
		idBlockParseURL:          BlockParseURL{},
		idBlockRandomChoiceList:  BlockRandomChoiceList{},
		idBlockRandomUA:          BlockRandomUA{},
		idBlockRecaptcha:         BlockRecaptcha{},
		idBlockRecaptchaV3:       BlockRecaptchaV3{},
		idBlockRequest:           BlockRequest{},
		idBlockSetCookie:         BlockSetCookie{},
		idBlockSetVariable:       BlockSetVariable{},
		idBlockSHA1Hash:          BlockSHA1Hash{},
		idBlockStringBuilder:     BlockStringBuilder{},
		idBlockToFloat:           BlockToFloat{},
		idBlockMathTotal:         BlockMathTotal{},
		idBlockUnix:              BlockUnix{},
		idBlockUnixMilli:         BlockUnixMilli{},
		idBlockUnixNano:          BlockUnixNano{},
		idBlockURLDecode:         BlockURLDecode{},
		idBlockURLEncode:         BlockURLEncode{},
		idBlockURLEncodedBuilder: BlockURLEncodedBuilder{},
		idBlockUUID:              BlockUUID{},
		idBlockVarContainsFilter: BlockVarContainsFilter{},
		idBlockVarLenFilter:      BlockVarLenFilter{},
		idBlockVarToCap:          BlockVarToCap{},
	}
)

type Routine struct {
	blocks []Runnable
}

type Runnable interface {
	Run(wce *Environment) (message string, status Status)
	kind() string
	toBytes() ([]byte, error)
	fromBytes([]byte) error
}

type intermediateData struct {
	Kind string
	Data []byte
}

func RoutineFromBytes(raw []byte) (*Routine, error) {
	var toBuild []intermediateData
	var builtBlocks []Runnable

	if err := json.Unmarshal(raw, &toBuild); err != nil {
		return nil, err
	}

	for _, b := range toBuild {
		foundType, ok := blocks[b.Kind]
		if !ok {
			return nil, errors.New("unknown block type:" + b.Kind)
		}

		block, ok := reflect.New(reflect.TypeOf(foundType)).Interface().(Runnable)
		if !ok {
			return nil, errors.New("block does not conform to the Runnable interface")
		}

		if err := block.fromBytes(b.Data); err != nil {
			return nil, err
		}

		builtBlocks = append(builtBlocks, block)
	}

	return &Routine{blocks: builtBlocks}, nil
}

func (r *Routine) ToBytes() ([]byte, error) {
	var blocks []intermediateData

	for _, b := range r.blocks {
		data, err := b.toBytes()
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, intermediateData{Kind: b.kind(), Data: data})
	}
	return json.Marshal(blocks)
}

func (r *Routine) Run(wce *Environment) {
	defer wce.Client.CloseIdleConnections()

	for _, v := range r.blocks {
		attempts := 0
		for {
			if attempts >= wce.maxRetry && wce.maxRetry != -1 {
				return
			}

			msg, status := v.Run(wce)

			wce.addLog(msg)
			wce.Status = status

			switch status {
			case Error:
				return
			case Retry:
				wce.Client.CloseIdleConnections()

				//Retries are often network failures
				time.Sleep(wce.retrySleep)

				attempts++
				continue
			case Fail:
				return
			case Custom:
				return
			case Success:
			}
			break
		}
	}
}

func (r *Routine) ToSum() (string, error) {
	bytes, err := r.ToBytes()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", sha1.Sum(bytes)), nil
}

func NewRoutine(b []Runnable) *Routine {
	return &Routine{blocks: b}
}
