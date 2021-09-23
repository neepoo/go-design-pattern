package proxy

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var mockedDatabase UserList

func TestMain(m *testing.M) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{
			ID: n,
		})
	}
	os.Exit(m.Run())
}

func TestUserListProxy(t *testing.T) {
	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackCache:     UserList{},
		StackSize:      2,
	}

	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID,
		mockedDatabase[5].ID}

	t.Run("FIndUser empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		require.NoError(t, err)
		//  As the description implies, the cache is empty at this point, and the user will have
		//	to be retrieved from mockedDatabase.
		require.Equal(t, knownIDs[0], user.ID, "Returned user name doesn't match with expected")

		require.Equal(t, 1, len(proxy.StackCache), "After one successful search in an empty cache, the size of it must e one")

		require.False(t, proxy.LastSearchUserCache, "No user can be returned from an empty cache")
	})

	t.Run("FindUser One User From Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		require.NoError(t, err)
		require.Equal(t, knownIDs[0], user.ID, "Returned user name doesn't match with expected")
		require.Equal(t, 1, len(proxy.StackCache), "Cache must not grow if we asked for an object that is stored on it")
		require.True(t, proxy.LastSearchUserCache, "The user should have been returned from the cache")
	})

	user1, err := proxy.FindUser(knownIDs[0])
	require.NoError(t, err)

	user2, _ := proxy.FindUser(knownIDs[1])
	require.False(t, proxy.LastSearchUserCache, "The user wasn't stored on the proxy cache yet")

	user3, _ := proxy.FindUser(knownIDs[2])
	require.False(t, proxy.LastSearchUserCache, "The user wasn't stored on the proxy cache yet")

	for i := 0; i < len(proxy.StackCache); i++ {
		require.NotEqual(t, proxy.StackCache[i], user1.ID, "User that should be gone was found")
	}

	require.Equal(t, 2, len(proxy.StackCache), "After inserting 3 users the cache should not grow more than to two")

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Fatal("user2 or user3 must be in cache")
		}
	}
}
