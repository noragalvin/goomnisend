package goomnisend

import (
	"log"
	"testing"
)

func TestGetList(t *testing.T) {
    apiKey := "5cda78668653ed3e50c96af9-zq91qjo93tzNn3BVY4Yr2Njl95HjCLFK2HUPFokznrrvjkwrK9"
	listID := "5cda877f8653ed591c6056ec"
	client := New(apiKey)

	list, err := client.GetList(listID, nil)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(list)
}