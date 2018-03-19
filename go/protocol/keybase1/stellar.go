// Auto-generated by avdl-compiler v1.3.23 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/stellar.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type StellarAccountID string

func (o StellarAccountID) DeepCopy() StellarAccountID {
	return o
}

type StellarSecretKey string

func (o StellarSecretKey) DeepCopy() StellarSecretKey {
	return o
}

type StellarRevision uint64

func (o StellarRevision) DeepCopy() StellarRevision {
	return o
}

type EncryptedStellarSecretBundle struct {
	V   int                  `codec:"v" json:"v"`
	E   []byte               `codec:"e" json:"e"`
	N   BoxNonce             `codec:"n" json:"n"`
	Gen PerUserKeyGeneration `codec:"gen" json:"gen"`
}

func (o EncryptedStellarSecretBundle) DeepCopy() EncryptedStellarSecretBundle {
	return EncryptedStellarSecretBundle{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N:   o.N.DeepCopy(),
		Gen: o.Gen.DeepCopy(),
	}
}

type StellarSecretBundleVersion int

const (
	StellarSecretBundleVersion_V1 StellarSecretBundleVersion = 1
)

func (o StellarSecretBundleVersion) DeepCopy() StellarSecretBundleVersion { return o }

var StellarSecretBundleVersionMap = map[string]StellarSecretBundleVersion{
	"V1": 1,
}

var StellarSecretBundleVersionRevMap = map[StellarSecretBundleVersion]string{
	1: "V1",
}

func (e StellarSecretBundleVersion) String() string {
	if v, ok := StellarSecretBundleVersionRevMap[e]; ok {
		return v
	}
	return ""
}

type StellarSecretBundleVersioned struct {
	Version__ StellarSecretBundleVersion `codec:"version" json:"version"`
	V1__      *StellarSecretBundleV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *StellarSecretBundleVersioned) Version() (ret StellarSecretBundleVersion, err error) {
	switch o.Version__ {
	case StellarSecretBundleVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o StellarSecretBundleVersioned) V1() (res StellarSecretBundleV1) {
	if o.Version__ != StellarSecretBundleVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return
	}
	return *o.V1__
}

func NewStellarSecretBundleVersionedWithV1(v StellarSecretBundleV1) StellarSecretBundleVersioned {
	return StellarSecretBundleVersioned{
		Version__: StellarSecretBundleVersion_V1,
		V1__:      &v,
	}
}

func (o StellarSecretBundleVersioned) DeepCopy() StellarSecretBundleVersioned {
	return StellarSecretBundleVersioned{
		Version__: o.Version__.DeepCopy(),
		V1__: (func(x *StellarSecretBundleV1) *StellarSecretBundleV1 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V1__),
	}
}

type StellarSecretBundleV1 struct {
	Revision StellarRevision      `codec:"revision" json:"revision"`
	Accounts []StellarSecretEntry `codec:"accounts" json:"accounts"`
}

func (o StellarSecretBundleV1) DeepCopy() StellarSecretBundleV1 {
	return StellarSecretBundleV1{
		Revision: o.Revision.DeepCopy(),
		Accounts: (func(x []StellarSecretEntry) []StellarSecretEntry {
			if x == nil {
				return nil
			}
			var ret []StellarSecretEntry
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Accounts),
	}
}

type StellarSecretBundle struct {
	Revision StellarRevision      `codec:"revision" json:"revision"`
	Accounts []StellarSecretEntry `codec:"accounts" json:"accounts"`
}

func (o StellarSecretBundle) DeepCopy() StellarSecretBundle {
	return StellarSecretBundle{
		Revision: o.Revision.DeepCopy(),
		Accounts: (func(x []StellarSecretEntry) []StellarSecretEntry {
			if x == nil {
				return nil
			}
			var ret []StellarSecretEntry
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Accounts),
	}
}

type StellarAccountMode int

const (
	StellarAccountMode_USER StellarAccountMode = 0
)

func (o StellarAccountMode) DeepCopy() StellarAccountMode { return o }

var StellarAccountModeMap = map[string]StellarAccountMode{
	"USER": 0,
}

var StellarAccountModeRevMap = map[StellarAccountMode]string{
	0: "USER",
}

func (e StellarAccountMode) String() string {
	if v, ok := StellarAccountModeRevMap[e]; ok {
		return v
	}
	return ""
}

type StellarSecretEntry struct {
	AccountID StellarAccountID   `codec:"accountID" json:"accountID"`
	Mode      StellarAccountMode `codec:"mode" json:"mode"`
	Signers   []StellarSecretKey `codec:"signers" json:"signers"`
	IsPrimary bool               `codec:"isPrimary" json:"isPrimary"`
	Name      string             `codec:"name" json:"name"`
}

func (o StellarSecretEntry) DeepCopy() StellarSecretEntry {
	return StellarSecretEntry{
		AccountID: o.AccountID.DeepCopy(),
		Mode:      o.Mode.DeepCopy(),
		Signers: (func(x []StellarSecretKey) []StellarSecretKey {
			if x == nil {
				return nil
			}
			var ret []StellarSecretKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Signers),
		IsPrimary: o.IsPrimary,
		Name:      o.Name,
	}
}

type StellarInterface interface {
}

func StellarProtocol(i StellarInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.stellar",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type StellarClient struct {
	Cli rpc.GenericClient
}
