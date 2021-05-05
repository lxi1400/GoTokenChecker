package utils

import (
  "log"
  "os"
	"bufio"
	"fmt"
	"sync"
	"github.com/valyala/fasthttp"
	"github.com/fatih/color"
	"syscall"
	"unsafe"
	"time"
	"github.com/valyala/fastjson"
)

func Rename(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}

func Banner() {
	color.Red(`
	████████╗ ██████╗ ██╗  ██╗███████╗███╗   ██╗     ██████╗██╗  ██╗███████╗ ██████╗██╗  ██╗███████╗██████╗
	╚══██╔══╝██╔═══██╗██║ ██╔╝██╔════╝████╗  ██║    ██╔════╝██║  ██║██╔════╝██╔════╝██║ ██╔╝██╔════╝██╔══██╗
	   ██║   ██║   ██║█████╔╝ █████╗  ██╔██╗ ██║    ██║     ███████║█████╗  ██║     █████╔╝ █████╗  ██████╔╝
	   ██║   ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╗██║    ██║     ██╔══██║██╔══╝  ██║     ██╔═██╗ ██╔══╝  ██╔══██╗
	   ██║   ╚██████╔╝██║  ██╗███████╗██║ ╚████║    ╚██████╗██║  ██║███████╗╚██████╗██║  ██╗███████╗██║  ██║
	   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝     ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝

	`)
}


func Finished(vaild, invaild int) {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("%s%s%s - %s", white("[", ), red("CHECKER"), white("]"), green(fmt.Sprintf("Finished checking all tokens! Vaild: %v Invaild: %v\n", vaild, invaild)))
}

func PrintVaild(token string) {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("%s%s%s - %s", white("[", ), red("CHECKER"), white("]"), green(fmt.Sprintf("%s is vaild\n", token)))
}

func PrintInvaild(token string) {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Printf("%s%s%s - %s", white("[", ), red("CHECKER"), white("]"), red(fmt.Sprintf("%s is invaild\n", token)))
}

func WriteToFile(text string) {

    f, err := os.OpenFile("vaild.txt",  os.O_WRONLY|os.O_APPEND, 0755)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(fmt.Sprintf("%s\n", text))

    if err2 != nil {
        log.Fatal(err2)
    }

}

func ReadTokens() []string {
    file, err := os.Open("tokens.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var tokens []string
    for scanner.Scan() {
        tokens = append(tokens, scanner.Text())
    }
    file.Close()
	  return tokens
}

func Clear() {
	fmt.Print("\033[H\033[2J")

}


func MakeRequest(token string, wg *sync.WaitGroup) bool {
	// thanks to cy <3
	defer wg.Done()

    req = fasthttp.AcquireRequest()
    defer fasthttp.ReleaseRequest(req)

    req.Header.SetMethod("GET")
    req.Header.SetRequestURI("https://discordapp.com/api/v6/users/@me")
    req.Header.Set("Authorization", token)

    res = fasthttp.AcquireResponse()
    err := client.DoTimeout(req, res, 10*time.Second)
    if err != nil {
        fasthttp.ReleaseResponse(res)
        return false
    }

    switch res.StatusCode() {
    case 429:
        time.Sleep(1 *time.Second)
        MakeRequest(token, wg)
    case 200:
        return true
    default:
        return false

    }
	return false
}




var (
  client = &fasthttp.Client{}
  sparser = fastjson.Parser{}
	req *fasthttp.Request
	res *fasthttp.Response
)
