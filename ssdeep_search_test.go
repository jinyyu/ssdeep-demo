package ssdeep_search

import "testing"

func TestGenerateKeys(t *testing.T) {
	hash := "12345678"
	keys := GenerateKeys(hash)

	if len(keys) != 2 || keys[0] != "1234567" || keys[1] != "2345678" {
		t.Error("error")
	}

	ssdeep := "196608:V84H0CBigCD2lwAwh9kAzYou1KgervWxbGoH9Vv:V84H+BYwUAzYVRU1U"
	result, err := ParseSsdeep(ssdeep)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if result.BlockSize != 196608 || result.HashBlockSize != "V84H0CBigCD2lwAwh9kAzYou1KgervWxbGoH9Vv" || result.Hash2BlockSIze != "V84H+BYwUAzYVRU1U" {
		t.Error("ERROR")
	}
}
