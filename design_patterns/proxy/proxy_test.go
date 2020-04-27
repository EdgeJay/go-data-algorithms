package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	mockedDatabase := userList{}

	rand.Seed(2342342)
	for idx := 0; idx < 1000000; idx++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, user{ID: n})
	}

	proxy := userListProxy{
		mockedDatabase: &mockedDatabase,
		stackCache:     userList{},
		stackSize:      2,
	}

	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID, mockedDatabase[5].ID}

	t.Run("findUser with empty cache", func(t *testing.T) {
		user, err := proxy.findUser(knownIDs[0])

		if err != nil {
			t.Fatalf("Expected error to be nil, got \"%s\"", err.Error())
		}

		if user.ID != knownIDs[0] {
			t.Error("User ID mismatched")
		}

		if len(proxy.stackCache) != 1 {
			t.Error("After one successful search, size of stackCache should be 1")
		}

		if proxy.lastSearchUsedCache {
			t.Error("No user should be returned from empty cache")
		}
	})

	t.Run("findUser with same user stored in cache", func(t *testing.T) {
		user, err := proxy.findUser(knownIDs[0])

		if err != nil {
			t.Fatalf("Expected error to be nil, got \"%s\"", err.Error())
		}

		if user.ID != knownIDs[0] {
			t.Error("User ID mismatched")
		}

		if len(proxy.stackCache) != 1 {
			t.Error("Size of stackCache should be 1")
		}

		if !proxy.lastSearchUsedCache {
			t.Error("User should be returned from cache")
		}
	})

	t.Run("findUser with multiple users - overflowing the stack", func(t *testing.T) {
		user1, err := proxy.findUser(knownIDs[0])

		if err != nil {
			t.Fatalf("Expected error to be nil, got \"%s\"", err.Error())
		}

		user2, _ := proxy.findUser(knownIDs[1])
		if proxy.lastSearchUsedCache {
			t.Error("Expected proxy.lastSearchUsedCache to be false, since user not stored in cache yet")
		}

		user3, _ := proxy.findUser(knownIDs[2])
		if proxy.lastSearchUsedCache {
			t.Error("Expected proxy.lastSearchUsedCache to be false, since user not stored in cache yet")
		}

		for idx := 0; idx < len(proxy.stackCache); idx++ {
			if proxy.stackCache[idx].ID == user1.ID {
				t.Error("Expected user1 to be moved out of stackCache")
			}
		}

		if len(proxy.stackCache) > 2 {
			t.Error("Size of stackCache should not grow more than 2")
		}

		for _, v := range proxy.stackCache {
			if v != user2 && v != user3 {
				t.Error("Non expected user found in stackCache")
			}
		}
	})
}
