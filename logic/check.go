package logic

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func Check(s []string) {
	col, err := getTerminalWidth()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	stringValue := s[1]
	bannerName := s[2]

	data, err := os.ReadFile("./banners/" + bannerName + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileHash := MD5(string(data))
	if !CheckHash(fileHash, bannerName) {
		fmt.Println("Error: hash code of", bannerName+".txt", "has been changed")
		return
	}

	result := Process(stringValue, data)

	size := result
	ans := 0

	for i := 0; i < len(size); i++ {
		if size[i] == '\n' || size[i] == '\t' || size[i] == '\r' {
			ans++
		}
	}

	if len(size) == ans {
		fmt.Print(result)
		return
	} else if (len(size)-ans)/8 > int(col) {
		fmt.Println((len(size) - ans) / 8)
		fmt.Println(col)
		fmt.Println("Too much characters")
		return
	}

	fmt.Print(result)
}

func getTerminalWidth() (uint16, error) {
	ws := &winsize{}

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		return 0, errno
	}

	return ws.Col, nil
}
