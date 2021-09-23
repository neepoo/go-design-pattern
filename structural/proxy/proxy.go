package proxy

import "fmt"

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

func (u *UserList) addUser(user User) {
	*u = append(*u, user)
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList // fifo
	StackSize           int
	LastSearchUserCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	// First check cache
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		u.LastSearchUserCache = true
		return user, nil
	}
	// not in cache, check db
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	// find in db
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
	u.LastSearchUserCache = false
	return user, nil
}
