package main

/*
Задание

Создать простое веб-приложение.

Нужно добавить в поисковик простое веб-приложение для отображения содержимого индекса и хранилища.
Задача №1

Создать пакет «webapp», который бы предоставлял возможность посмотреть с помощью браузера содержимое индекса и
хранилища проиндексированных документов.

Для определённости привязать эти возможности к адресам “/index” и “/docs”.

Формат представления данных выбирайте на своё усмотрение (json, xml, html-список и т.д.)
Задача №2

Написать модульные тесты для обработчиков http-запросов.
*/

import (
	"fmt"
	"go-core-4/homework12/pkg/crawler"
	"go-core-4/homework12/pkg/crawler/spider"
	"go-core-4/homework12/pkg/file"
	"go-core-4/homework12/pkg/index"
	"go-core-4/homework12/pkg/webapp"
	"log"
	"net/http"
)

func main() {
	log.Println("Собираем информацию и индексируем ...")
	api, err := createWebApi()
	if err != nil {
		log.Fatalf("Error in createWebApi", err)
	}
	log.Println("Завершено.")

	err = startMux(api)
	if err != nil {
		log.Fatal(err)
	}

}

func startMux(api *webapp.API) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Home)
	mux.HandleFunc("/index", api.Index)
	mux.HandleFunc("/docs", api.Docs)
	log.Println("Запуск веб-сервера")
	err := http.ListenAndServe(":8000", mux)
	return err
}

func createWebApi() (*webapp.API, error) {
	var id int
	var doc []crawler.Document
	var array = [...]string{"https://go.dev", "https://golang.org"}
	for _, link := range array {
		s := spider.New()
		d, err := s.Scan(link, 2, &id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		doc = append(doc, d...)

	}
	idb := index.Index(doc)
	api := webapp.Fill(doc, idb)
	err := createFiles(doc, idb)
	if err != nil {
		log.Printf("Err in createFile:", err)
		return nil, err
	}

	return api, nil
}
func createFiles(c []crawler.Document, m map[string][]int) error {
	var ind []*webapp.IndexerData
	for key, list := range m {
		ind = append(ind, &webapp.IndexerData{Word: key, Indexes: list})
	}

	var cra []*webapp.CrawlerData
	for _, val := range c {
		cra = append(cra, &webapp.CrawlerData{Id: val.ID, Title: val.Title, Body: val.Body, URL: val.URL})
	}

	cdoc, err := file.CreateFile("./crawlerDoc.json")
	if err != nil {
		return err
	}
	err = file.WriteToFile(cra, cdoc)
	if err != nil {
		return err
	}

	idoc, err := file.CreateFile("./indexDoc.json")
	if err != nil {
		return err
	}
	err = file.WriteToFile(ind, idoc)
	if err != nil {
		return err
	}
	return nil
}
