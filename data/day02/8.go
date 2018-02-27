package main

//文件读写
import "fmt"
import "os"
import "log"

func readfile(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)

	for {
		num, _ := f.Read(buf)
		if num == 0 {
			break
		}

		fmt.Println(string(buf))
		//	fmt.Println("%s \n", buf)

	}

}

func writefile(filename string) {

	fout, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {

		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		OutString := fmt.Sprintf(" %s: %d \n", "李雪楠", i)
		fout.WriteString(OutString)
	}

}

func main() {

	//writefile("./test.txt")
	readfile("./test.txt")
}
