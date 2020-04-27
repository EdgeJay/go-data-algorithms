package proxy

import (
	"fmt"
)

type user struct {
	ID int32
}

type userFinder interface {
	findUser(id int32) (user, error)
}

type userList []user

func (list *userList) findUser(id int32) (usr user, err error) {
	for _, item := range *list {
		if item.ID == id {
			usr = item
			return
		}
	}
	err = fmt.Errorf("Unable to find user with ID: %d", id)
	return
}

type userListProxy struct {
	mockedDatabase      *userList
	stackCache          userList
	stackSize           int
	lastSearchUsedCache bool
}

func (proxy *userListProxy) findUser(id int32) (user, error) {
	usr, err := proxy.stackCache.findUser(id)

	if err == nil {
		proxy.lastSearchUsedCache = true
		return usr, nil
	}

	// not in cache, go find in database
	usr, err = proxy.mockedDatabase.findUser(id)

	if err != nil {
		return user{}, err
	}

	proxy.lastSearchUsedCache = false

	if len(proxy.stackCache) >= proxy.stackSize {
		proxy.stackCache = proxy.stackCache[1:]
	}

	proxy.stackCache = append(proxy.stackCache, usr)

	return usr, nil
}
