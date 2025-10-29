// !!! Linux only (TCGETS/TCSETS)

package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// terminal attributes struct (c.f. man pages)
type termios struct {
	Iflag  uint32    // input modes
	Oflag  uint32    // output modes
	Cflag  uint32    // control modes
	Lflag  uint32    // local modes
	Cc     [32]uint8 // special modes
	Ispeed uint32    // baud input speed
	Ospeed uint32    // baud output speed
}

// enableRawMode switches the terminal to raw mode and returns the original state
func enableRawMode() (*termios, error) {
	fd := int(syscall.Stdin)
	var oldState termios

	// get current terminal attributes
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd),
		uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&oldState)))
	if errno != 0 {
		return nil, errno
	}

	// modify attributes to enable raw mode
	newState := oldState
	// disble canonical mode and echo
	newState.Lflag &^= syscall.ICANON | syscall.ECHO

	// set the new terminal attricutes
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd),
		uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&newState)))
	if errno != 0 {
		return nil, errno
	}

	return &oldState, nil
}

// disableRawMode recover the terminal to it's original state
func disableRawMode(oldState *termios) error {
	fd := int(syscall.Stdin)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd),
		uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(oldState)))
	if errno != 0 {
		return errno
	}

	return nil
}

// readInput reads a keypress in raw mode
func readInput() ([]byte, int) {
	buf := make([]byte, 3)

	n, err := os.Stdin.Read(buf[:])
	if err != nil {
		fmt.Printf("Error reading the key: %v\n", err)
		return nil, -1
	}

	return buf, n
}

func main() {
	fmt.Println("Press any key (CTRL+D to exit):")

	oldState, err := enableRawMode()
	if err != nil {
		fmt.Printf("Failed to enable raw mode: %v\n", err)
		return
	}
	defer disableRawMode(oldState)

	for {
		buf, n := readInput()
		if n > 0 {
			fmt.Printf("Captured: %v (bytes: %v)\n", string(buf[:n]), buf[:n])

			// extiting with CTRL+D (4)
			if buf[0] == 4 {
				fmt.Println("Exiting...")
				return
			}
		}
	}
}
