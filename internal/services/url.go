package services

import (
"bytes"
"regexp"
"fmt"
"log"
"github.com/alissonbrunosa/shortner/internal/stores"
)

type URL struct {
	Store stores.URLStore
}

func (this *URL) Find(key string) (string, error) {
  return this.Store.Find(key)
}

func (this *URL) Create(link string) (string, error) {
  regex := regexp.MustCompile(`^(http|https)://`)
  
  hex, err := Hex(4)
  if err != nil {
    return "", err
  }

  if !regex.MatchString(link) {
    var buffer bytes.Buffer
    buffer.WriteString("http://")
    buffer.WriteString(link)
    link = buffer.String()
  }

  fail := this.Store.Create(hex, link)
  if fail != nil {
    return "", fail
  }

  newLink := fmt.Sprintf("http://localhost:8080/%s", hex)
	return newLink, nil
}

func NewURLService() *URL {
	return &URL{
		Store: stores.NewRedisURLStore(),
	}
}
