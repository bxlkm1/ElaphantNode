package types

import (
	"bytes"
	"errors"
	"fmt"
	common2 "github.com/elastos/Elastos.ELA.Elephant.Node/common"
	"github.com/elastos/Elastos.ELA/common"
	. "github.com/elastos/Elastos.ELA/core/types"
	"io"
)

var Vote TxType = 0x90

var TxTypeEnum = map[TxType]string{
	CoinBase:                "CoinBase",
	RegisterAsset:           "RegisterAsset",
	TransferAsset:           "TransferAsset",
	Record:                  "Record",
	Deploy:                  "Deploy",
	SideChainPow:            "SideChainPow",
	RechargeToSideChain:     "RechargeToSideChain",
	WithdrawFromSideChain:   "WithdrawFromSideChain",
	TransferCrossChainAsset: "TransferCrossChainAsset",
	Vote:                    "Vote",
}

type TransactionHistory struct {
	Address    common.Uint168
	Txid       common.Uint256
	Type       []byte
	Value      uint64
	CreateTime uint64
	Height     uint64
	Fee        uint64
	Inputs     []common.Uint168
	Outputs    []common.Uint168
	TxType     TxType
	Memo       []byte
}

type TransactionHistoryDisplay struct {
	Address    string
	Txid       string
	Type       string
	Value      uint64
	CreateTime uint64
	Height     uint64
	Fee        uint64
	Inputs     []string
	Outputs    []string
	TxType     string
	Memo       string
}

type ThResult struct {
	History  TransactionHistorySorter
	TotalNum int
}

func (th *TransactionHistory) Serialize(w io.Writer) error {
	err := common.WriteVarBytes(w, th.Address.Bytes())
	if err != nil {
		return errors.New("[TransactionHistory], Address serialize failed.")
	}
	err = common.WriteVarBytes(w, th.Txid.Bytes())
	if err != nil {
		return errors.New("[TransactionHistory], Txid serialize failed.")
	}
	err = common.WriteVarBytes(w, th.Type)
	if err != nil {
		return errors.New("[TransactionHistory], Type serialize failed.")
	}
	err = common.WriteUint64(w, th.Value)
	if err != nil {
		return errors.New("[TransactionHistory], Value serialize failed.")
	}
	err = common.WriteUint64(w, th.CreateTime)
	if err != nil {
		return errors.New("[TransactionHistory], CreateTime serialize failed.")
	}
	err = common.WriteUint64(w, th.Height)
	if err != nil {
		return errors.New("[TransactionHistory], Height serialize failed.")
	}
	err = common.WriteUint64(w, th.Fee)
	if err != nil {
		return errors.New("[TransactionHistory], Fee serialize failed.")
	}
	err = common.WriteVarUint(w, uint64(len(th.Inputs)))
	if err != nil {
		return errors.New("[TransactionHistory], Length of inputs serialize failed.")
	}
	for i := 0; i < len(th.Inputs); i++ {
		err = common.WriteVarBytes(w, th.Inputs[i].Bytes())
		if err != nil {
			return errors.New("[TransactionHistory], input:" + string(th.Inputs[i].Bytes()) + " serialize failed.")
		}
	}
	err = common.WriteVarUint(w, uint64(len(th.Outputs)))
	if err != nil {
		return errors.New("[TransactionHistory], Length of outputs serialize failed.")
	}
	for i := 0; i < len(th.Outputs); i++ {
		err = common.WriteVarBytes(w, th.Outputs[i].Bytes())
		if err != nil {
			return errors.New("[TransactionHistory], output:" + string(th.Outputs[i].Bytes()) + " serialize failed.")
		}
	}
	err = common.WriteVarBytes(w, []byte{byte(th.TxType)})
	if err != nil {
		return errors.New("[TransactionHistory], TxType serialize failed.")
	}
	err = common.WriteVarBytes(w, th.Memo)
	if err != nil {
		return errors.New("[TransactionHistory], Memo serialize failed.")
	}
	return nil
}

func (th *TransactionHistory) Deserialize(r io.Reader) (*TransactionHistoryDisplay, error) {
	var err error
	txhd := new(TransactionHistoryDisplay)
	buf, err := common.ReadVarBytes(r, 1024, "address")
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Address deserialize failed.")
	}
	th.Address.Deserialize(bytes.NewBuffer(buf))
	txhd.Address, _ = th.Address.ToAddress()

	//err = th.Txid.Deserialize(r)
	buf, err = common.ReadVarBytes(r, 1024, "txid")
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Txid deserialize failed.")
	}
	th.Txid.Deserialize(bytes.NewBuffer(buf))
	txhd.Txid, _ = common2.ReverseHexString(th.Txid.String())

	th.Type, err = common.ReadVarBytes(r, 1024, "transfer type")
	txhd.Type = string(th.Type)
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Type deserialize failed.")
	}
	th.Value, err = common.ReadUint64(r)
	txhd.Value = th.Value
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Value deserialize failed.")
	}
	th.CreateTime, err = common.ReadUint64(r)
	txhd.CreateTime = th.CreateTime
	if err != nil {
		return txhd, errors.New("[TransactionHistory], CreateTime deserialize failed.")
	}
	th.Height, err = common.ReadUint64(r)
	txhd.Height = th.Height
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Height deserialize failed.")
	}
	th.Fee, err = common.ReadUint64(r)
	txhd.Fee = th.Fee
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Fee deserialize failed.")
	}
	n, err := common.ReadVarUint(r, 0)
	if err != nil {
		return txhd, errors.New("[TransactionHistory], length of inputs deserialize failed.")
	}
	for i := uint64(0); i < n; i++ {
		programHash := common.Uint168{}
		buf, err = common.ReadVarBytes(r, 1024, "address")
		if err != nil {
			return txhd, errors.New("[TransactionHistory], input deserialize failed.")
		}
		programHash.Deserialize(bytes.NewBuffer(buf))
		th.Inputs = append(th.Inputs, programHash)
		addr, _ := programHash.ToAddress()
		txhd.Inputs = append(txhd.Inputs, addr)
	}
	n, err = common.ReadVarUint(r, 0)
	if err != nil {
		return txhd, errors.New("[TransactionHistory], length of outputs deserialize failed.")
	}
	for i := uint64(0); i < n; i++ {
		programHash := common.Uint168{}
		buf, err = common.ReadVarBytes(r, 1024, "address")
		if err != nil {
			return txhd, errors.New("[TransactionHistory], output deserialize failed.")
		}
		programHash.Deserialize(bytes.NewBuffer(buf))
		th.Outputs = append(th.Outputs, programHash)
		addr, _ := programHash.ToAddress()
		txhd.Outputs = append(txhd.Outputs, addr)
	}
	txt, err := common.ReadVarBytes(r, 1, "TxType")
	if err != nil {
		return txhd, errors.New("[TransactionHistory], TxType serialize failed.")
	}
	th.TxType = TxType(txt[0])
	txhd.TxType = TxTypeEnum[th.TxType]
	th.Memo, err = common.ReadVarBytes(r, common.MaxVarStringLength, "memo")
	txhd.Memo = string(th.Memo)
	if err != nil {
		return txhd, errors.New("[TransactionHistory], Memo serialize failed.")
	}
	return txhd, nil
}

func (th TransactionHistory) String() string {
	return fmt.Sprintf("addr: %s,txid: %s,value: %d,height: %d", th.Address, th.Txid, th.Value, th.Height)
}

// TransactionHistorySorter implements sort.Interface for []TransactionHistory based on
// the Height field.
type TransactionHistorySorter []TransactionHistoryDisplay

func (a TransactionHistorySorter) Len() int           { return len(a) }
func (a TransactionHistorySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TransactionHistorySorter) Less(i, j int) bool { return a[i].Height < a[j].Height }

func (a TransactionHistorySorter) Filter(from, size uint32) TransactionHistorySorter {
	rst := TransactionHistorySorter{}
	for i, v := range a {
		if uint32(i) < from {
			continue
		}
		rst = append(rst, v)
		if uint32(len(rst)) == size {
			break
		}
	}
	return rst
}
