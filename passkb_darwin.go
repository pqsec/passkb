package passkb

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Foundation/Foundation.h>

CGEventRef cgeventcreatekeyboardevent(unsigned short c, bool down) {
	// real key code does not matter as we are overriding with CGEventKeyboardSetUnicodeString
	CGEventRef cf = CGEventCreateKeyboardEvent(NULL, 0, down);
	if (cf)
		CGEventKeyboardSetUnicodeString(cf, 1, &c);
	return cf;
}

void cgeventpost(CGEventRef cf) {
	CGEventPost(kCGAnnotatedSessionEventTap, cf);
}

void cfrelease(CGEventRef cf) {
	CFRelease(cf);
}
*/
import "C"
import (
	"fmt"
	"time"
)

type macKeyboard struct{}

func New(name string) (Keyboard, error) {
	return &macKeyboard{}, nil
}

func (kb macKeyboard) Close() error {
	return nil
}

func (kb macKeyboard) Type(str string, delay time.Duration) error {
	events := make([]C.CGEventRef, 2*len(str))

	for i, r := range str {
		i := i
		events[2*i] = C.cgeventcreatekeyboardevent(C.ushort(r), true)
		if events[2*i] == 0 {
			return fmt.Errorf("Cannot create a key down event for %v", r)
		} else {
			defer C.cfrelease(events[2*i])
		}

		events[2*i+1] = C.cgeventcreatekeyboardevent(C.ushort(r), false)
		if events[2*i+1] == 0 {
			return fmt.Errorf("Cannot create a key down event for %v", r)
		} else {
			defer C.cfrelease(events[2*i+1])
		}
	}

	time.Sleep(delay)

	for i, e := range events {
		C.cgeventpost(e)
		t := time.Millisecond
		if i&0x1 == 1 {
			t = 100 * time.Millisecond
		}
		time.Sleep(t)
	}

	return nil
}
