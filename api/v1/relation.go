package v1

import (
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []model.ReplyUser `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)
	followerID, _ := strconv.Atoi(c.Query("to_user_id"))

	var isfollow uint8 = 0
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	if uint8(action_type) == 1 {
		isfollow = 1
	} else {
		isfollow = 0
	}

	_, err = service.SetOrUpdateRelation(user.ID, uint(followerID), isfollow)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	var user model.User
	var repley []model.ReplyUser
	var followList []model.Realtion
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
	user = cUser.(model.User)

	followList, err = service.GetFollowID(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
	for i := 0; i < len(followList); i++ {
		follower := model.ReplyUser{
			ID: followList[i].FollowerId,
		}
		information, err := service.FindAuthor(followList[i].FollowerId)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.Username = information.Username
		follower.FollowCount, err = service.GetFollowNum(follower.ID)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.FollowerCount, err = service.GetFollowerNum(follower.ID)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.IsFollow = service.IsFollow(user.ID, follower.ID)
		repley = append(repley, follower)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: repley,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	var user model.User
	var repley []model.ReplyUser
	var followList []model.Realtion
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
	user = cUser.(model.User)
	followList, err = service.GetFollowerID(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
	for i := 0; i < len(followList); i++ {
		follower := model.ReplyUser{
			ID: followList[i].UserId,
		}
		information, err := service.FindAuthor(followList[i].UserId)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.Username = information.Username
		follower.FollowCount, err = service.GetFollowNum(follower.ID)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.FollowerCount, err = service.GetFollowerNum(follower.ID)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: repley,
			})
		}
		follower.IsFollow = service.IsFollow(follower.ID, user.ID)
		repley = append(repley, follower)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: repley,
	})
}
