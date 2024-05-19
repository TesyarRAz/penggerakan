package model

import (
	"github.com/golang-jwt/jwt/v5"
	lo "github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

type permission struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type role struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Permissions []*permission `json:"permissions"`
}

type Auth struct {
	ID   string
	Name string

	claims jwt.MapClaims

	roles          []*role
	allPermissions []*permission
}

func NewAuth(id string, name string, claims jwt.MapClaims) *Auth {
	return &Auth{
		ID:     id,
		Name:   name,
		claims: claims,
	}
}

func (a *Auth) ParseRoleAndPermission() error {
	roles := a.claims["roles"].([]interface{})
	permissions := a.claims["permissions"].([]interface{})

	a.roles = lop.Map(roles, func(rc interface{}, _ int) *role {
		r := rc.(map[string]interface{})
		return &role{
			ID:          r["id"].(string),
			Name:        r["name"].(string),
			Permissions: parseRolePermissions(r["permissions"].([]interface{})),
		}
	})

	a.allPermissions = parseRolePermissions(permissions)
	a.allPermissions = append(a.allPermissions, lo.FlatMap(a.roles, func(r *role, _ int) []*permission {
		return r.Permissions
	})...)

	return nil
}

func (a *Auth) HasPermission(permissionName ...string) bool {
	return lo.ContainsBy(a.allPermissions, func(p *permission) bool {
		return lo.Contains(permissionName, p.Name)
	})
}

func (a *Auth) HasRole(roleName ...string) bool {
	return lo.ContainsBy(a.roles, func(r *role) bool {
		return lo.Contains(roleName, r.Name)
	})
}

func parseRolePermissions(permissions []interface{}) []*permission {
	return lop.Map(permissions, func(pc interface{}, _ int) *permission {
		p := pc.(map[string]interface{})
		return &permission{
			ID:   p["id"].(string),
			Name: p["name"].(string),
		}
	})
}
