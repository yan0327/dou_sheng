package v1

import (
	"net/http"
	"simple-demo/middleware"
	"simple-demo/model"
	"simple-demo/model/response"
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
	token := c.Query("token")
	followerID, _ := strconv.Atoi(c.Query("to_user_id"))

	var isfollow uint8 = 0
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	var err error
	var user model.User
	if uint8(action_type) == 1 {
		isfollow = 1
	} else {
		isfollow = 0
	}
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
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
	token := c.Query("token")
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	var user model.User
	var repley []model.ReplyUser
	var followList []model.Realtion

	//user.Username = claims.Username
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
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
	token := c.Query("token")
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	var user model.User
	var repley []model.ReplyUser
	var followList []model.Realtion

	//user.Username = claims.Username
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: repley,
		})
	}
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
