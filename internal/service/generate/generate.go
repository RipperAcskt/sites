package generate

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Gen() error {
	file, err := os.Create("./internal/service/generate/check.go")
	if err != nil {
		return err
	}
	_, err = file.WriteString(codeStr)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("./internal/service/sites")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {

		_, err = file.WriteString(fmt.Sprintf(`wg.Add(1)
		fmt.Print("POST%v ")
			fmt.Println(sites.POST%v(email))
		`, strings.Split(f.Name(), ".")[0], strings.Split(f.Name(), ".")[0]))
		if err != nil {
			return err
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("wg.Wait()}")
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
