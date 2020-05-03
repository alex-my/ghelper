package regexp

import (
	"testing"
)

type RegexpExpect struct {
	Context string
	Result  bool
}

func TestChinesePhone(t *testing.T) {
	expects := []RegexpExpect{
		{
			Context: "23456789012",
			Result:  false,
		},
		{
			Context: "13345678901",
			Result:  true,
		},
	}

	for _, expect := range expects {
		if IsChinesePhone(expect.Context) != expect.Result {
			t.Errorf("%s result: %t", expect.Context, expect.Result)
		}
	}
}

func TestNickName(t *testing.T) {
	names := []string{
		"name1",
		"Name2",
		"name_3",
		"name-4",
		"名字5",
		"name名字6",
		"0007",
	}
	for _, name := range names {
		if !IsNickName(name) {
			t.Errorf("name: %s is the correct nick name", name)
		}
	}
}

func TestAccount(t *testing.T) {
	expects := []RegexpExpect{
		{
			Context: "name1",
			Result:  true,
		},
		{
			Context: "1name2",
			Result:  false,
		},
		{
			Context: "name三",
			Result:  false,
		},
	}

	for _, expect := range expects {
		if IsAccount(expect.Context) != expect.Result {
			t.Errorf("%s result: %t", expect.Context, expect.Result)
		}
	}
}

func TestEmail(t *testing.T) {
	expects := []RegexpExpect{
		{
			Context: "abc@mn.com",
			Result:  true,
		},
		{
			Context: "abc@mn",
			Result:  true,
		},
		{
			Context: "123@mn.com",
			Result:  true,
		},
		{
			Context: "abc@",
			Result:  false,
		},
		{
			Context: "abc",
			Result:  false,
		},
	}

	for _, expect := range expects {
		if IsEmail(expect.Context) != expect.Result {
			t.Errorf("%s result: %t", expect.Context, expect.Result)
		}
	}
}
