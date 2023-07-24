package main

import "github.com/urfave/cli"

func main() {
	// cmd := exec.Command("firefox")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd := exec.Command("storages", "hello-world")
	// cmd.Stdin = strings.NewReader("and old falcon")
	// fmt.Println(cmd)

	// var out bytes.Buffer
	// cmd.Stdout = &out

	// err := cmd.Run()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("translate phrase: %q\n", out.String())

	// cmd := exec.Command("ls", "./")
	// out, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("could not run command: ", err)
	// }
	// fmt.Println("Output: ", string(out))

	cli.NewApp()
}
