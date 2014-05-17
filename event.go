package userstream

import (
	"encoding/json"
	"github.com/k0kubun/twitter"
	"strings"
)

type FriendList struct {
	Friends []int64
}

type Tweet twitter.Tweet

type Delete struct {
	Id     int64
	UserId int64 `json:"user_id"`
}

type Favorite struct {
	Source       *twitter.User
	Target       *twitter.User
	TargetObject *twitter.Tweet `json:"target_object"`
}
type Unfavorite Favorite

type Follow struct {
	Source *twitter.User
	Target *twitter.User
}
type Unfollow Follow

type ListMemberAdded struct {
	Source       *twitter.User
	Target       *twitter.User
	TargetObject *twitter.List `json:"target_object"`
}
type ListMemberRemoved ListMemberAdded

func ParseJson(jsonText string) interface{} {
	hash := map[string]string{}
	decoder := json.NewDecoder(strings.NewReader(jsonText))
	decoder.Decode(&hash)

	decoder = json.NewDecoder(strings.NewReader(jsonText))
	if _, hasKey := hash["friends"]; hasKey {
		friendList := FriendList{}
		decoder.Decode(&friendList)
		return &friendList
	} else if _, hasKey := hash["event"]; hasKey {
		return parseEvent(decoder, hash["event"])
	} else if _, hasKey := hash["delete"]; hasKey {
		deleteHash := map[string]map[string]*Delete{}
		decoder.Decode(&deleteHash)
		return deleteHash["delete"]["status"]
	} else if _, hasKey := hash["created_at"]; hasKey {
		tweet := Tweet{}
		decoder.Decode(&tweet)
		return &tweet
	}
	return nil
}

func parseEvent(decoder *json.Decoder, eventName string) interface{} {
	switch eventName {
	case "favorite":
		favorite := Favorite{}
		decoder.Decode(&favorite)
		return &favorite
	case "unfavorite":
		unfavorite := Unfavorite{}
		decoder.Decode(&unfavorite)
		return &unfavorite
	case "follow":
		follow := Follow{}
		decoder.Decode(&follow)
		return &follow
	case "unfollow":
		unfollow := Unfollow{}
		decoder.Decode(&unfollow)
		return &unfollow
	case "list_member_added":
		listMemberAdded := ListMemberAdded{}
		decoder.Decode(&listMemberAdded)
		return &listMemberAdded
	case "list_member_removed":
		listMemberRemoved := ListMemberRemoved{}
		decoder.Decode(&listMemberRemoved)
		return &listMemberRemoved
	}
	return nil
}
