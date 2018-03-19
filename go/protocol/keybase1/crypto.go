// Auto-generated by avdl-compiler v1.3.23 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/crypto.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ED25519PublicKey [32]byte

func (o ED25519PublicKey) DeepCopy() ED25519PublicKey {
	var ret ED25519PublicKey
	copy(ret[:], o[:])
	return ret
}

type ED25519Signature [64]byte

func (o ED25519Signature) DeepCopy() ED25519Signature {
	var ret ED25519Signature
	copy(ret[:], o[:])
	return ret
}

type ED25519SignatureInfo struct {
	Sig       ED25519Signature `codec:"sig" json:"sig"`
	PublicKey ED25519PublicKey `codec:"publicKey" json:"publicKey"`
}

func (o ED25519SignatureInfo) DeepCopy() ED25519SignatureInfo {
	return ED25519SignatureInfo{
		Sig:       o.Sig.DeepCopy(),
		PublicKey: o.PublicKey.DeepCopy(),
	}
}

type EncryptedBytes32 [48]byte

func (o EncryptedBytes32) DeepCopy() EncryptedBytes32 {
	var ret EncryptedBytes32
	copy(ret[:], o[:])
	return ret
}

type BoxNonce [24]byte

func (o BoxNonce) DeepCopy() BoxNonce {
	var ret BoxNonce
	copy(ret[:], o[:])
	return ret
}

type BoxPublicKey [32]byte

func (o BoxPublicKey) DeepCopy() BoxPublicKey {
	var ret BoxPublicKey
	copy(ret[:], o[:])
	return ret
}

type CiphertextBundle struct {
	Kid        KID              `codec:"kid" json:"kid"`
	Ciphertext EncryptedBytes32 `codec:"ciphertext" json:"ciphertext"`
	Nonce      BoxNonce         `codec:"nonce" json:"nonce"`
	PublicKey  BoxPublicKey     `codec:"publicKey" json:"publicKey"`
}

func (o CiphertextBundle) DeepCopy() CiphertextBundle {
	return CiphertextBundle{
		Kid:        o.Kid.DeepCopy(),
		Ciphertext: o.Ciphertext.DeepCopy(),
		Nonce:      o.Nonce.DeepCopy(),
		PublicKey:  o.PublicKey.DeepCopy(),
	}
}

type UnboxAnyRes struct {
	Kid       KID     `codec:"kid" json:"kid"`
	Plaintext Bytes32 `codec:"plaintext" json:"plaintext"`
	Index     int     `codec:"index" json:"index"`
}

func (o UnboxAnyRes) DeepCopy() UnboxAnyRes {
	return UnboxAnyRes{
		Kid:       o.Kid.DeepCopy(),
		Plaintext: o.Plaintext.DeepCopy(),
		Index:     o.Index,
	}
}

type SignED25519Arg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Msg       []byte `codec:"msg" json:"msg"`
	Reason    string `codec:"reason" json:"reason"`
}

type SignED25519ForKBFSArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Msg       []byte `codec:"msg" json:"msg"`
	Reason    string `codec:"reason" json:"reason"`
}

type SignToStringArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Msg       []byte `codec:"msg" json:"msg"`
	Reason    string `codec:"reason" json:"reason"`
}

type UnboxBytes32Arg struct {
	SessionID        int              `codec:"sessionID" json:"sessionID"`
	EncryptedBytes32 EncryptedBytes32 `codec:"encryptedBytes32" json:"encryptedBytes32"`
	Nonce            BoxNonce         `codec:"nonce" json:"nonce"`
	PeersPublicKey   BoxPublicKey     `codec:"peersPublicKey" json:"peersPublicKey"`
	Reason           string           `codec:"reason" json:"reason"`
}

type UnboxBytes32AnyArg struct {
	SessionID   int                `codec:"sessionID" json:"sessionID"`
	Bundles     []CiphertextBundle `codec:"bundles" json:"bundles"`
	Reason      string             `codec:"reason" json:"reason"`
	PromptPaper bool               `codec:"promptPaper" json:"promptPaper"`
}

type CryptoInterface interface {
	// Sign the given message (which should be small) using the device's private
	// signing ED25519 key, and return the signature as well as the corresponding
	// public key that can be used to verify the signature. The 'reason' parameter
	// is used as part of the SecretEntryArg object passed into
	// secretUi.getSecret().
	SignED25519(context.Context, SignED25519Arg) (ED25519SignatureInfo, error)
	// Same as the above except a KBFS-specific prefix is added to the payload to be signed.
	SignED25519ForKBFS(context.Context, SignED25519ForKBFSArg) (ED25519SignatureInfo, error)
	// Same as the above except the full marsheled and encoded NaclSigInfo.
	SignToString(context.Context, SignToStringArg) (string, error)
	// Decrypt exactly 32 bytes using nacl/box with the given nonce, the given
	// peer's public key, and the device's private encryption key, and return the
	// decrypted data. The 'reason' parameter is used as part of the
	// SecretEntryArg object passed into secretUi.getSecret().
	UnboxBytes32(context.Context, UnboxBytes32Arg) (Bytes32, error)
	UnboxBytes32Any(context.Context, UnboxBytes32AnyArg) (UnboxAnyRes, error)
}

func CryptoProtocol(i CryptoInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.crypto",
		Methods: map[string]rpc.ServeHandlerDescription{
			"signED25519": {
				MakeArg: func() interface{} {
					ret := make([]SignED25519Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SignED25519Arg)
					if !ok {
						err = rpc.NewTypeError((*[]SignED25519Arg)(nil), args)
						return
					}
					ret, err = i.SignED25519(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"signED25519ForKBFS": {
				MakeArg: func() interface{} {
					ret := make([]SignED25519ForKBFSArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SignED25519ForKBFSArg)
					if !ok {
						err = rpc.NewTypeError((*[]SignED25519ForKBFSArg)(nil), args)
						return
					}
					ret, err = i.SignED25519ForKBFS(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"signToString": {
				MakeArg: func() interface{} {
					ret := make([]SignToStringArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SignToStringArg)
					if !ok {
						err = rpc.NewTypeError((*[]SignToStringArg)(nil), args)
						return
					}
					ret, err = i.SignToString(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"unboxBytes32": {
				MakeArg: func() interface{} {
					ret := make([]UnboxBytes32Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UnboxBytes32Arg)
					if !ok {
						err = rpc.NewTypeError((*[]UnboxBytes32Arg)(nil), args)
						return
					}
					ret, err = i.UnboxBytes32(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"unboxBytes32Any": {
				MakeArg: func() interface{} {
					ret := make([]UnboxBytes32AnyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UnboxBytes32AnyArg)
					if !ok {
						err = rpc.NewTypeError((*[]UnboxBytes32AnyArg)(nil), args)
						return
					}
					ret, err = i.UnboxBytes32Any(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type CryptoClient struct {
	Cli rpc.GenericClient
}

// Sign the given message (which should be small) using the device's private
// signing ED25519 key, and return the signature as well as the corresponding
// public key that can be used to verify the signature. The 'reason' parameter
// is used as part of the SecretEntryArg object passed into
// secretUi.getSecret().
func (c CryptoClient) SignED25519(ctx context.Context, __arg SignED25519Arg) (res ED25519SignatureInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.crypto.signED25519", []interface{}{__arg}, &res)
	return
}

// Same as the above except a KBFS-specific prefix is added to the payload to be signed.
func (c CryptoClient) SignED25519ForKBFS(ctx context.Context, __arg SignED25519ForKBFSArg) (res ED25519SignatureInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.crypto.signED25519ForKBFS", []interface{}{__arg}, &res)
	return
}

// Same as the above except the full marsheled and encoded NaclSigInfo.
func (c CryptoClient) SignToString(ctx context.Context, __arg SignToStringArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.crypto.signToString", []interface{}{__arg}, &res)
	return
}

// Decrypt exactly 32 bytes using nacl/box with the given nonce, the given
// peer's public key, and the device's private encryption key, and return the
// decrypted data. The 'reason' parameter is used as part of the
// SecretEntryArg object passed into secretUi.getSecret().
func (c CryptoClient) UnboxBytes32(ctx context.Context, __arg UnboxBytes32Arg) (res Bytes32, err error) {
	err = c.Cli.Call(ctx, "keybase.1.crypto.unboxBytes32", []interface{}{__arg}, &res)
	return
}

func (c CryptoClient) UnboxBytes32Any(ctx context.Context, __arg UnboxBytes32AnyArg) (res UnboxAnyRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.crypto.unboxBytes32Any", []interface{}{__arg}, &res)
	return
}
