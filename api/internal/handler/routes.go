// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "resourceManager/api/internal/handler/base"
	user "resourceManager/api/internal/handler/user"
	"resourceManager/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/core/init/database",
				Handler: base.InitDatbaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/create_user",
				Handler: user.CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/delete_user",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/get_user_by_id",
				Handler: user.GetUserByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/get_user_by_username",
				Handler: user.GetUserByUsernameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/get_user_list",
				Handler: user.GetUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update_user",
				Handler: user.UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/upload_avatar",
				Handler: user.UploadAvatarHandler(serverCtx),
			},
		},
	)
}
