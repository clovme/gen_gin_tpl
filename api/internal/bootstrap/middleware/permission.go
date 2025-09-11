package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/enums/code"
	"net/http"
)

func PermissionMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		count, err := query.Q.User.
			Select(query.Q.User.ID.As("count")).
			Join(query.Q.UserRole, query.Q.UserRole.UserID.EqCol(query.Q.User.ID), query.Q.UserRole.Status.Eq(1)).
			Join(query.Q.Role, query.Q.Role.ID.EqCol(query.Q.UserRole.RoleID), query.Q.Role.Status.Eq(1)).
			Join(query.Q.RolePermission, query.Q.RolePermission.RoleID.EqCol(query.Q.Role.ID)).
			Join(query.Q.Permission, query.Q.Permission.ID.EqCol(query.Q.RolePermission.PermissionID),
				query.Q.Permission.Status.Eq(1),
				query.Q.Permission.Uri.Eq(c.Request.URL.Path),
				query.Q.Permission.Method.Eq(c.Request.Method),
			).
			Where(query.Q.User.ID.Eq(c.UserInfo.ID)).Count()
		if err != nil || count <= 0 {
			// 没有权限
			responseJsonOrHtml(c, code.RequestUnauthorized, http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
