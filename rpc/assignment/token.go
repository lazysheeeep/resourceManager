package assignment

import (
	"resourceManager/rpc/model"
	"resourceManager/rpc/types/core"
	"strconv"
)

func AssignToken(token model.Token) *core.TokenInfo {
	id := strconv.Itoa(int(token.ID))
	createdAt := token.CreatedAt.Unix()
	updatedAt := token.UpdatedAt.Unix()
	status := uint32(token.Status)

	return &core.TokenInfo{
		Id:        &id,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Status:    &status,
		Uuid:      &token.Uuid,
		Token:     &token.Token,
		ExpiredAt: &token.ExpiredAt,
	}
}

func AssignTokens(tokens []model.Token, total int64) *core.TokenListResp {

	var data []*core.TokenInfo

	for _, v := range tokens {
		token := AssignToken(v)
		data = append(data, token)
	}

	return &core.TokenListResp{Data: data, Total: uint64(total)}
}
