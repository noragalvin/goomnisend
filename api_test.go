package goomnisend

import (
	"log"
	"testing"
	"time"
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

func TestGetListMembers(t *testing.T) {
	apiKey := "5cda78668653ed3e50c96af9-zq91qjo93tzNn3BVY4Yr2Njl95HjCLFK2HUPFokznrrvjkwrK9"
	listID := "5cda877f8653ed591c6056ec"
	
	client := New(apiKey)

	params := InterestCategoriesQueryParams{}
	params.ListID = listID
	members, err := client.GetMembers(&params)
	if err != nil {
		log.Println("ERROR: ", err.Error())
	} else {
		log.Println(members)
	}
}

func TestCreateMember(t *testing.T) {
	apiKey := "5cda78668653ed3e50c96af9-zq91qjo93tzNn3BVY4Yr2Njl95HjCLFK2HUPFokznrrvjkwrK9"
	listID := "5cda877f8653ed591c6056ec"
	
	client := New(apiKey)

	params := MemberRequest{}
	params.Email = "clonenora02@gmail.com"
	params.FirstName = "Nora1"
	params.LastName = "Galvin1"
	params.Status = "subscribed"
	params.StatusDate = time.Now()
	list := &ListResponse{}
	list.ListID = listID
	lists := []*ListResponse{}
	lists = append(lists, list)
	params.Lists = lists

	member, err := client.CreateMember(&params)
	if err != nil {
		log.Println("ERROR: ", err.Error())
	} else {
		log.Println(member)
	}
}