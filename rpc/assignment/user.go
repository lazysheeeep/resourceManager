package assignment

import (
	"resourceManager/rpc/model"
	"resourceManager/rpc/types/core"
)

func AssignUser(user model.User) *core.UserInfo {

	loginStatus := uint64(user.LoginStatus)

	return &core.UserInfo{
		Username:    &user.Username,
		LoginStatus: &loginStatus,
		Avatar:      &user.Avatar,
		Email:       &user.Email,
		Phone:       &user.Phone,
	}
}

func AssignListUser(users []model.User, total uint64) *core.UserListResp {

	var info []*core.UserInfo
	for i, _ := range users {
		user := AssignUser(users[i])
		info = append(info, user)
	}

	return &core.UserListResp{
		Total: total,
		Data:  info,
	}
}
