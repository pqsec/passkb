package passkb

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

type winKeyboard struct{}

func New(name string) (Keyboard, error) {
	return &winKeyboard{}, nil
}

func (kb winKeyboard) Close() error {
	return nil
}

func typeKey(r rune) error {
	var inputs = make([]win.KEYBD_INPUT, 2, 2)
	inputs[0].Type = win.INPUT_KEYBOARD
	inputs[0].Ki.WScan = uint16(r)
	inputs[0].Ki.DwFlags = win.KEYEVENTF_UNICODE

	inputs[1] = inputs[0]
	inputs[0].Ki.DwFlags |= win.KEYEVENTF_KEYUP

	res := win.SendInput(uint32(len(inputs)), unsafe.Pointer(&inputs[0]), int32(unsafe.Sizeof(inputs[0])))
	if res != uint32(len(inputs)) {
		return fmt.Errorf("Failed to type character %v", r)
	}

	return nil
}

func (kb winKeyboard) Type(str string, delay time.Duration) error {
	time.Sleep(delay)

	for _, r := range str {
		err := typeKey(r)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
