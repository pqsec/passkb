package passkb

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/bendahl/uinput"
)

type keyPress struct {
	code  int
	upper bool
}

var (
	upperSymbolMap = map[rune]int{
		'~': uinput.KeyGrave,
		'{': uinput.KeyLeftbrace,
		'}': uinput.KeyRightbrace,
		'|': uinput.KeyBackslash,
		':': uinput.KeySemicolon,
		'"': uinput.KeyApostrophe,
		'<': uinput.KeyComma,
		'>': uinput.KeyDot,
		'?': uinput.KeySlash,
	}
)

func keyFromRune(r rune) (*keyPress, error) {
	if i := strings.IndexRune("1234567890-=", r); i >= 0 {
		return &keyPress{uinput.Key1 + i, false}, nil
	}

	if i := strings.IndexRune("qwertyuiop[]", r); i >= 0 {
		return &keyPress{uinput.KeyQ + i, false}, nil
	}

	if i := strings.IndexRune("asdfghjkl;'`", r); i >= 0 {
		return &keyPress{uinput.KeyA + i, false}, nil
	}

	if i := strings.IndexRune("\\zxcvbnm,./", r); i >= 0 {
		return &keyPress{uinput.KeyBackslash + i, false}, nil
	}

	if i := strings.IndexRune("!@#$%^&*()_+", r); i >= 0 {
		return &keyPress{uinput.Key1 + i, true}, nil
	}

	if i, ok := upperSymbolMap[r]; ok {
		return &keyPress{i, true}, nil
	}

	if unicode.IsLetter(r) && unicode.IsUpper(r) {
		lower, err := keyFromRune(unicode.ToLower(r))
		if err != nil {
			return nil, fmt.Errorf("Cannot convert character %v to a key", r)
		}
		lower.upper = true
		return lower, nil
	}

	return nil, fmt.Errorf("Cannot convert character %v to a key", r)
}

func stringToKeys(s string) ([]*keyPress, error) {
	var keys []*keyPress
	for _, r := range s {
		k, err := keyFromRune(r)
		if err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}

	return keys, nil
}

type linuxKeyboard struct {
	ukb uinput.Keyboard
}

func New(name string) (Keyboard, error) {
	kb, err := uinput.CreateKeyboard("/dev/uinput", []byte(name))
	if err != nil {
		return nil, err
	}

	return &linuxKeyboard{kb}, nil
}

func (kb linuxKeyboard) Close() error {
	return kb.ukb.Close()
}

func (kb linuxKeyboard) Type(str string, delay time.Duration) error {
	keys, err := stringToKeys(str)
	if err != nil {
		return err
	}

	time.Sleep(delay)

	for _, key := range keys {
		if key.upper {
			kb.ukb.KeyDown(uinput.KeyLeftshift)
		}

		time.Sleep(100 * time.Millisecond)
		kb.ukb.KeyPress(key.code)

		if key.upper {
			kb.ukb.KeyUp(uinput.KeyLeftshift)
		}
	}

	return nil
}
