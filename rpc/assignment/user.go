package assignment

import (
	"resourceManager/rpc/model"
	"resourceManager/rpc/types/core"
)

func AssignUser(user model.User, role model.Role) *core.UserInfo {

	loginStatus := uint64(user.LoginStatus)

	return &core.UserInfo{
		Username:    &user.Username,
		LoginStatus: &loginStatus,
		Avatar:      &user.Avatar,
		Email:       &user.Email,
		Phone:       &user.Phone,
		RoleCodes:   &role.Code,
		RoleName:    &role.Name,
	}
}

func AssignListUser(users []model.User, roles []model.Role, total uint64) *core.UserListResp {

	var info []*core.UserInfo
	for i, _ := range users {
		user := AssignUser(users[i], roles[i])
		info = append(info, user)
	}

	return &core.UserListResp{
		Total: total,
		Data:  info,
	}
}
