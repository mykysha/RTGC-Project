package domain

import (
	"errors"
	"fmt"
	"testing"
)

func TestRoom_Connecter(t *testing.T) {
	t.Parallel()

	data := []struct {
		testName       string
		userID         string
		userName       string
		room           string
		userIDInRoom   string
		userNameInRoom string
		expected       error
	}{
		{
			testName:       "id is in room",
			userID:         "user1",
			userName:       "name1",
			room:           "room1",
			userIDInRoom:   "user1",
			userNameInRoom: "someName",
			expected:       fmt.Errorf("%w : '%v', '%v', '%v'", errID, "user1", "room1", "someName"),
		},
		{
			testName:       "username is in room",
			userID:         "user1",
			userName:       "name1",
			room:           "room1",
			userIDInRoom:   "someID",
			userNameInRoom: "name1",
			expected:       fmt.Errorf("%w : '%v', '%v'", errUname, "name1", "room1"),
		},
		{
			testName:       "no conflicts in room",
			userID:         "user1",
			userName:       "name1",
			room:           "room1",
			userIDInRoom:   "someID",
			userNameInRoom: "someName",
			expected:       nil,
		},
	}

	for _, tt := range data {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			r := Room{
				Name:          tt.room,
				UserList:      make(map[string]string),
				UserIDToRowID: make(map[string]int),
			}

			r.UserList[tt.userNameInRoom] = tt.userIDInRoom

			got := r.Connecter(tt.userID, tt.userName)

			if got != nil && tt.expected != nil {
				if errors.Is(got, tt.expected) {
					t.Errorf("Connecter(%s, %s) in room %s with userName %s and userId %s: expected\n'%v'\n\tgot\n'%v'",
						tt.userID, tt.userName, tt.room, tt.userNameInRoom, tt.userIDInRoom, tt.expected, got)
				}
			}
		})
	}
}

func TestRoom_Leaver(t *testing.T) {
	t.Parallel()

	data := []struct {
		testName       string
		userID         string
		room           string
		userIDInRoom   string
		userNameInRoom string
		expected       error
	}{
		{
			testName:       "id is not in room",
			userID:         "user1",
			room:           "room1",
			userIDInRoom:   "someID",
			userNameInRoom: "someName",
			expected:       fmt.Errorf("%w : '%s', '%s'", errNoUser, "user1", "room1"),
		},
		{
			testName:       "id is in room",
			userID:         "user1",
			room:           "room1",
			userIDInRoom:   "user1",
			userNameInRoom: "someName",
			expected:       nil,
		},
	}

	for _, tt := range data {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			r := Room{
				Name:          tt.room,
				UserList:      make(map[string]string),
				UserIDToRowID: make(map[string]int),
			}

			r.UserList[tt.userNameInRoom] = tt.userIDInRoom

			_, got := r.Leaver(tt.userID)

			if got != nil && tt.expected != nil {
				if errors.Is(got, tt.expected) {
					t.Errorf("Leaver(%s) in room %s with userId %s: expected\n'%v'\n\tgot\n'%v'",
						tt.userID, tt.room, tt.userIDInRoom, tt.expected, got)
				}
			}
		})
	}
}

func TestRoom_Messenger(t *testing.T) {
	t.Parallel()

	data := []struct {
		testName       string
		userID         string
		room           string
		userIDInRoom   string
		userNameInRoom string
		text           string
		expected       error
	}{
		{
			testName:       "id is not in room",
			userID:         "user1",
			room:           "room1",
			userIDInRoom:   "someID",
			userNameInRoom: "someName",
			text:           "sample",
			expected:       fmt.Errorf("%w : '%s', '%s'", errNoUser, "user1", "room1"),
		},
		{
			testName:       "id is in room",
			userID:         "user1",
			room:           "room1",
			userIDInRoom:   "user1",
			userNameInRoom: "someName",
			text:           "sample",
			expected:       nil,
		},
	}

	for _, tt := range data {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			r := Room{
				Name:          tt.room,
				UserList:      make(map[string]string),
				UserIDToRowID: make(map[string]int),
			}

			r.UserList[tt.userNameInRoom] = tt.userIDInRoom
			r.UserList["SERVER"] = "SERVER"

			_, got := r.Messenger(tt.userID, tt.room, tt.text)

			if got != nil && tt.expected != nil {
				if errors.Is(got, tt.expected) {
					t.Errorf("Messenger(%s, %s, %s) in room %s with userId %s: expected\n'%v'\n\tgot\n'%v'",
						tt.userID, tt.room, tt.text, tt.room, tt.userIDInRoom, tt.expected, got)
				}
			}
		})
	}
}
