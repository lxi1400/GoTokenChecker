package main
import (
	"fmt"
	"./utils"
	"sync"
	"time"
	"bufio"
	"os"
)



func main() {
	vaild := 0
	invaild := 0
	utils.Banner()
	tokens = utils.ReadTokens()
	utils.Rename(fmt.Sprintf("Checker | Checking %v tokens!", len(tokens)))
	fmt.Printf("Loaded %v tokens! Press enter to start...", len(tokens))
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	utils.Clear()
	utils.Banner()
	fmt.Printf("Attempting to check  %v tokens now\n", len(tokens))
	for _, token := range tokens {
		wg.Add(1)
		go func() {
			response := utils.MakeRequest(token, wg)
			switch response{
			case true:
				utils.WriteToFile(token)
				utils.PrintVaild(token)
				vaild += 1
			
			default:
				utils.PrintInvaild(token)
				invaild += 1
			}
		}()
		wg.Wait()

	}
	utils.Rename(fmt.Sprintf("Checker | Finshed checking %v tokens! [%v vaild, %v invaild]", len(tokens), vaild, invaild))

	utils.Finished(vaild, invaild)
	time.Sleep(1 *time.Hour)
}



var (
    wg = new(sync.WaitGroup)
	tokens []string
)
