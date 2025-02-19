package tests

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/multi"
)

func TestTenancyUserList(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/getTenancyList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(1)
	data.Value("list").Array().Length().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "userName", "email", "phone", "nickName", "headerImg", "authorityName", "authorityType", "authorityId", "tenancyName", "defaultRouter", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

	baseLogOut(auth)
}

func TestTenancyUserProcess(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityId": source.TenancyAuthorityId, "authorityType": multi.TenancyAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("注册成功")

	user := obj.Value("data").Object()
	user.Value("id").Number().Ge(0)
	user.Value("userName").String().Equal("admin")
	user.Value("authorityId").String().Equal(source.TenancyAuthorityId)
	userId := user.Value("id").Number().Raw()

	// changePassword error
	obj = auth.POST("/v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "authorityType": multi.TenancyAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("/v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "authorityType": multi.TenancyAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("修改失败，原密码与当前账户不符")

	// setUserAuthority
	obj = auth.POST("/v1/admin/user/setUserAuthority").
		WithJSON(map[string]interface{}{"id": userId, "authorityId": source.TenancyAuthorityId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// setTenancyInfo
	obj = auth.PUT(fmt.Sprintf("/v1/admin/user/setUserInfo/%d", int(userId))).
		WithJSON(map[string]interface{}{"email": "tenancy@master.com", "phone": "13800138001"}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/user/deleteUser").
		WithJSON(map[string]interface{}{"id": userId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestTenancyUserRegisterError(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "a303176530", "password": "123456", "authorityId": source.TenancyAuthorityId, "authorityType": multi.TenancyAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("注册失败")

	baseLogOut(auth)
}

func TestTenancyUserRegisterAuthorityIdEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "chindeo_tenancy", "password": "123456", "authorityId": 0, "authorityType": multi.TenancyAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("AuthorityId值不能为空")

	baseLogOut(auth)
}

func TestTenancyUserRegisterAuthorityTypeEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "chindeo_tenancy", "password": "123456", "authorityId": source.TenancyAuthorityId, "authorityType": 0}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("AuthorityType值不能为空")

	baseLogOut(auth)
}
