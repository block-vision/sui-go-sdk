package sui

import (
	"encoding/base64"
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestOnKeyStoreSign(t *testing.T) {
	ks, err := SetAccountKeyStore("../config/sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	t.Run("test on sign with ed25519", func(t *testing.T) {
		ed25519Addr := "0x72cffd05deb71fa9b30584cb0f512d680cb08eab"
		txDataBytes, err := base64.StdEncoding.DecodeString("VHJhbnNhY3Rpb25EYXRhOjoAAgAAAAAAAAAAAAAAAAAAAAAAAAACAQAAAAAAAAAgf4wCMzbSQGAtJy5c2FShsm5eDefCLIODnSU2sC07IXMKZGV2bmV0X25mdARtaW50AAMADAtibG9ja3Zpc2lvbgAMC2Jsb2NrdmlzaW9uAAgHdGVzdHVybA8VhOvfVMkbhXJ5Oyp54IVRTqbH2I2ztReyGFA+2Q33fyXewleRiQMEAAAAAAAAACBvr7UgN38gsUyCYE9wVOFX6mj20zu+W+y5dUyfsb01EAEAAAAAAAAA6AMAAAAAAAA=")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		expected := []byte("dQ5BbqgrPtbsoHBowT1f8WVxO5trpHHXNs4g0Mj9p2/8oOnzKyd8VnEt6eUUBWLiHnfuDIYxEzwFC89nD+XZDQ==")
		actualBytes, err := ks.Sign(ed25519Addr, txDataBytes)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, string(expected), string(base64.StdEncoding.EncodeToString(actualBytes)))
	})

	t.Run("test on sign with secp256k1", func(t *testing.T) {
		secpAddr := "0x4d6f1a54e805038f44ecd3112927af147e9b9ecb"
		txDataBytes, err := base64.StdEncoding.DecodeString("VHJhbnNhY3Rpb25EYXRhOjoAAgAAAAAAAAAAAAAAAAAAAAAAAAACAQAAAAAAAAAgf4wCMzbSQGAtJy5c2FShsm5eDefCLIODnSU2sC07IXMKZGV2bmV0X25mdARtaW50AAMADAtibG9ja3Zpc2lvbgAMC2Jsb2NrdmlzaW9uAAgHdGVzdHVybA8VhOvfVMkbhXJ5Oyp54IVRTqbH2I2ztReyGFA+2Q33fyXewleRiQMEAAAAAAAAACBvr7UgN38gsUyCYE9wVOFX6mj20zu+W+y5dUyfsb01EAEAAAAAAAAA6AMAAAAAAAA=")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		expected := []byte("0mkXvSOEq57hJPdd+svJ6CNYXRV9Go9pSTuvHIbRkaZxcKWQSz1U3aKywCDtBZsFIzXP8Wf1g5zbXFcwkt4A/wA=")
		actual, err := ks.Sign(secpAddr, txDataBytes)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, string(expected), base64.StdEncoding.EncodeToString(actual))
	})
}

func TestOnKeyStoreKeys(t *testing.T) {
	ks, err := SetAccountKeyStore("../config/sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	t.Run("test on get public keys", func(t *testing.T) {
		publicKeys := ks.Keys()
		if len(publicKeys) != 6 {
			t.FailNow()
		}

		containStr := func(strs []string, str string) bool {
			for i := range strs {
				if str == strs[i] {
					return true
				}
			}
			return false
		}
		for i := range publicKeys {
			if !containStr(publicKeys, publicKeys[i]) {
				t.FailNow()
			}
		}
	})
}

func TestOnKeyStoreGetKey(t *testing.T) {
	ks, err := SetAccountKeyStore("../config/sui.keystore.fortest")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	t.Run("test on keystore GetKey()", func(t *testing.T) {
		_kp, err := ks.GetKey("0x4d6f1a54e805038f44ecd3112927af147e9b9ecb")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "A5/ex++LPumQg36qll95WuyJ2Dejf3AKiGc1LqGSYLOR", _kp.PublicKeyBase64)
		_kp, err = ks.GetKey("0x72cffd05deb71fa9b30584cb0f512d680cb08eab")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "XoFR32cGFHSHqGK96LSfBwx7U38aesHth5wy7wbjJWk=", _kp.PublicKeyBase64)

		_kp, err = ks.GetKey("0x8f3cf7d8ebb187bd655cea775802d0d9c1c5b145")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "ILFuDpjV0wfTYP9GqDN5AsuU4kHoVqdwsINIVnSh7dc=", _kp.PublicKeyBase64)

		_kp, err = ks.GetKey("0xc697e5fdd38d5f63ebeb14c2b49a864d473849db")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "RCZJ/xmuALnCxErBTAs/FxowPtOJSOnVIzGikcpj78Y=", _kp.PublicKeyBase64)

		_kp, err = ks.GetKey("0xf354bb3497c5879d68b49582d3a8887dbd26e3f0")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "M1UztctB3tvHRgWpjhHon7nzWWIoh4pBaERNfUg7loU=", _kp.PublicKeyBase64)

		_kp, err = ks.GetKey("0xfde3698d3e7da3f359e1036078da9cfbfb31f203")
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		assert.Equal(t, "owj3ywGVAs9e280oC78MDuYMkAJ24Qj/Op0hnO/2QFE=", _kp.PublicKeyBase64)
	})
}
