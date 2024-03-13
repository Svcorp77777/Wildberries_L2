package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

type keyWget struct {
	URL          string
	OutDirectory string
}

func flagArgumentsСonsole() keyWget {
	wgetKey := keyWget{}

	flag.StringVar(&wgetKey.URL, "url", "", "URL сайта для скачивания")
	flag.StringVar(&wgetKey.OutDirectory, "output", ".", "директория для сохранения файлов")

	flag.Parse()

	return wgetKey
}

func main() {
	wgetKey := flagArgumentsСonsole()

	err := loading(wgetKey.URL, wgetKey.OutDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Возникла ошибка скачать не удалось: %v\n", err)
		os.Exit(1)
	}
}

func loading(urlStr string, outDirectory string) error {
	parseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	err = loadingPage(urlStr, parseURL, outDirectory)
	if err != nil {
		return err
	}

	fmt.Printf("Загрузка завершена успешно, файл сохранён в директорию: %s\n", outDirectory)
	return nil
}

func loadingPage(urlStr string, basicURL *url.URL, outDirectory string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось скачать страницу, статус: %d", resp.StatusCode)
	}

	newTokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := newTokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := newTokenizer.Token()

			for _, attr := range token.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					linkURL, err := basicURL.Parse(attr.Val)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Возникла ошибка при анализе URL: %v\n", err)

						continue
					}

					err = loadingResource(linkURL, outDirectory)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Возникла ошибка при скачивании ресурса: %v\n", err)
					}
				}
			}
		}
	}
}

func loadingResource(url *url.URL, outDirectory string) error {
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось скачать ресурс %s, статус: %d", url.String(), resp.StatusCode)
	}

	filePath := path.Join(outDirectory, url.Host, url.Path)
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Сайт скачан: %s\n", filePath)

	return nil
}
