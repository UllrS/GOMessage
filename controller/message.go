package controller

import (
	"MessageGO/models"
	"MessageGO/repository"
	"encoding/json"
	"fmt"
	"strconv"
)

func GetMessages(user, companion, startTime string) ([]byte, error) {
	if startTime == "" {
		startTime = "1"
	}
	list, err := repository.GetMessages(user, companion, startTime)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(&list)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func PostMessages(body, sender, recipient, fileLink, fileName string) ([]byte, error) {
	message := models.Message{
		Body:         body,
		Sender:       sender,
		Recipient:    recipient,
		AttachedPath: fileLink,
		AttachedName: fileName,
	}
	lastID, err := repository.PostMessage(&message)
	if err != nil {
		return nil, err
	}
	fmt.Println(lastID)
	res, err := json.Marshal(&lastID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteMessage(idString, user string) ([]byte, error) {
	id, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return nil, err
	}
	ok := repository.DeleteMessage(id, user)
	res, err := json.Marshal(ok)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteChat(user, companion string) ([]byte, error) {
	ress, err := repository.DeleteChat(user, companion)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(ress)
	if err != nil {
		return nil, err
	}
	return res, nil
}
