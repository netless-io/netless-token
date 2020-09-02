package main

import (
	"strings"
	"testing"
)

func Test_bufferToBase64(t *testing.T) {
	b := bufferToBase64([]byte("netless"))

	if b != "bmV0bGVzcw" {
		t.Errorf("bufferToBase64([]byte(\"netless\")) expected be \"bmV0bGVzcw\", but %s got", b)
	}
}

func Test_getMapKeys(t *testing.T) {
	m := map[string]string{
		"b": "1",
		"a": "2",
	}
	b := getMapKeys(&m)

	for i, k := range []string{"a", "b"} {
		if k != b[i] {
			t.Errorf("getMapKeys result expected be [a b], but %s got", b)
		}
	}
}

func Test_encodeURIComponent(t *testing.T) {
	s := encodeURIComponent("net less")

	if s != "net%20less" {
		t.Errorf("encodeURIComponent(\"net less\") expected be \"net%%20ess\", but %s got", s)
	}
}

func Test_stringify(t *testing.T) {
	m := map[string]string{
		"b": "1",
		"a": "2",
	}
	s := stringify(&m)

	//if s != "{\"b\":\"1\",\"a\":\"2\"}" {
	if s != "a=2&b=1" {
		t.Errorf("encodeURIComponent(\"net less\") expected be \"net%%20ess\", but %s got", s)
	}
}

func Test_mergeMap(t *testing.T) {
	m1 := map[string]string{
		"b": "3",
		"c": "0",
	}

	m2 := map[string]string{
		"a": "2",
		"b": "1",
	}
	s := mergeMap(&m1, &m2)

	if s["a"] != "2" || s["b"] != "1" || s["c"] != "0" {
		t.Errorf("mergeMap result expected be map[a:2 b:1 c:0], but %s got", s)
	}
}

func Test_jsonStringify(t *testing.T) {
	m := map[string]string{
		"b": "1",
		"a": "2",
	}

	s := jsonStringify(&m)

	if s != "{\"a\":\"2\",\"b\":\"1\"}" {
		t.Errorf("jsonStringify result expected be \"{\"a\":\"2\",\"b\":\"1\"}\", but %s got", s)
	}
}

func Test_hmac256(t *testing.T) {
	s := hmac256("netless", "key")

	if s != "4b007294b9ebeccb96fb1ef30c843b9894d6323375e042b1ef2775978225ca7f" {
		t.Errorf("jsonStringify result expected be \"4b007294b9ebeccb96fb1ef30c843b9894d6323375e042b1ef2775978225ca7f\", but %s got", s)
	}
}

func Test_SDKToken(t *testing.T) {
	c := SDKContent{
		role: AdminRole,
	}
	s := SDKToken("netless", "x", 1, &c)

	if !strings.HasPrefix(s, sdkPrefix) {
		t.Errorf("SDKToken result expected prefix is %s, but result is: %s", sdkPrefix, s)
	}
}

func Test_RoomToken(t *testing.T) {
	c := RoomContent{
		role: ReaderRole,
		uuid: "this is uuid",
	}
	s := RoomToken("netless", "x", 1, &c)

	if !strings.HasPrefix(s, roomPrefix) {
		t.Errorf("RoomToken result expected prefix is %s, but result is: %s", sdkPrefix, s)
	}
}

func Test_TaskToken(t *testing.T) {
	c := TaskContent{
		role: WriterRole,
		uuid: "this is uuid",
	}
	s := TaskToken("netless", "x", 1, &c)

	if !strings.HasPrefix(s, taskPrefix) {
		t.Errorf("TaskToken result expected prefix is %s, but result is: %s", sdkPrefix, s)
	}
}