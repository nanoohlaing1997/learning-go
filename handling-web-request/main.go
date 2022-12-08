package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const urll = "http://loc.dev"
const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=gh23442"

func main() {
	// request to some url and get response
	fmt.Println("Welcome to handling web request")

	response, err := http.Get(urll)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response.Body)

	// need to close the connection when received response
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// this is databutes
	// fmt.Println(databytes)

	content := string(databytes)
	fmt.Println(content)

	fmt.Println("This is URL in golang", myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	// query params
	qParams := result.Query()
	fmt.Printf("The type of query params are %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, val := range qParams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev:3000",
		Path:    "/tutcss",
		RawPath: "user=noh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
	fmt.Println("Raw path is", partsOfUrl.RawPath)
}
