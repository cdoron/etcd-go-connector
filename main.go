/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	port := flag.Int("port", 8080, "Listening port")
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/conf/conf.yaml", "")
	flag.Parse()

	fmt.Println("port:", *port)
	fmt.Println("config file:", configFile)

	// TODO - add logging level

	log.Printf("Server started")

	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	conf := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

	DefaultApiService := openapi.NewDefaultApiService(conf)
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
