package cache

import (
	"fmt"
	"testing"
)

func TestRedisProxy_Follows(t *testing.T) {
	fmt.Println("--------------------------TestRedisProxy_Follows--------------------------")
	initRDB()
	p := NewRedisProxy()
	err := p.SetFollow(10, 11, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.GetFollowStatus(10, 11))

	err = p.SetFollow(10, 11, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.GetFollowStatus(10, 10))
}

func TestRedisProxy_SetFavorite(t *testing.T) {
	fmt.Println("--------------------------TestRedisProxy_SetFavorite--------------------------")
	initRDB()
	p := NewRedisProxy()
	err := p.SetFavorite(10, 10, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.GetFavoriteStatus(10, 10))

	err = p.SetFavorite(10, 10, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p.GetFavoriteStatus(10, 10))
}
