package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"time"
	"watchAlert/internal/cache"
	"watchAlert/internal/global"
	"watchAlert/pkg/utils/cmd"
	jwtUtils "watchAlert/pkg/utils/jwt"

	"watchAlert/internal/models"
	"watchAlert/pkg/ctx"
)

type ldapService struct {
	ctx *ctx.Context
}

type InterLdapService interface {
	Create(req interface{}) (interface{}, interface{})
	Update(req interface{}) (interface{}, interface{})
	Get() (interface{}, interface{})
	Ping(req interface{}) (interface{}, interface{})
	Login(req interface{}) (interface{}, interface{})
}

func newInterLdapService(ctx *ctx.Context) InterLdapService {
	return &ldapService{
		ctx: ctx,
	}
}

func (ls *ldapService) Create(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Ldap)
	err := ls.ctx.DB.Ldap().Create(*r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (ls *ldapService) Update(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Ldap)
	err := ls.ctx.DB.Ldap().Update(*r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 仅有第一条数据为准，后续数据会被忽略

func (ls *ldapService) Get() (interface{}, interface{}) {
	data, err := ls.ctx.DB.Ldap().Get()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ls *ldapService) Ping(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Ldap)
	var (
		ld  *ldap.Conn
		err error
	)
	if r.SSL == 1 {
		ld, err = ldap.DialURL("ldaps://"+r.Address, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	} else {
		ld, err = ldap.DialURL("ldap://" + r.Address)
	}
	if err != nil {
		return nil, err
	}
	defer ld.Close()
	if ld != nil {
		if err = ld.Bind(r.AdminUser, r.Password); err != nil {
			return nil, err
		}

	}
	return nil, nil
}

func (ls *ldapService) Login(req interface{}) (interface{}, interface{}) {
	loginUser := req.(*models.Member)
	// ldap 登录
	var ld *ldap.Conn
	data, err := ls.Get()
	if err != nil {
		return nil, err
	}
	r := data.(*models.Ldap)
	if r.SSL == 1 {
		ld, err = ldap.DialURL("ldaps://"+r.Address, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	} else {
		ld, err = ldap.DialURL("ldap://" + r.Address)
	}
	if err != nil {
		return nil, err
	}
	defer ld.Close()
	searchRequest := ldap.NewSearchRequest(r.DN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(r.Filter, loginUser.UserName), []string{}, nil)
	sr, err := ld.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if len(sr.Entries) != 1 {
		return nil, fmt.Errorf("user does not exist or too many entries returned")
	}
	userDN := sr.Entries[0].DN
	if err = ld.Bind(userDN, loginUser.Password); err != nil {
		return nil, fmt.Errorf("登录失败")
	}
	var userMap map[string]interface{}
	if err = json.Unmarshal(r.Mapping, &userMap); err != nil {
		return nil, fmt.Errorf("ldap字段映射失败")
	}
	for k, v := range userMap {
		userMap[k] = sr.Entries[0].GetAttributeValue(v.(string))
	}
	jsonData, _ := json.Marshal(userMap)
	_ = json.Unmarshal(jsonData, &loginUser)
	loginUser.CreateBy = "ldap"
	// 判断用户是否存在
	_, err = UserService.Register(loginUser)
	if fmt.Sprintf("%s", err) != "用户已存在" && err != nil {
		fmt.Println("register err: ", err)
		return nil, err
	}
	// 搜索用户
	var filter = models.MemberQuery{
		UserName: loginUser.UserName,
	}
	Userdata, err := UserService.Get(&filter)
	fmt.Println("userdata: ", Userdata)
	if err != nil {
		return nil, err
	}
	UserData := Userdata.(models.Member)
	tokenData, err := jwtUtils.GenerateToken(*loginUser)
	if err != nil {
		return nil, err
	}
	duration := time.Duration(global.Config.Jwt.Expire) * time.Second
	cache.NewEntryCache().Redis().Set("uid-"+UserData.UserId, cmd.JsonMarshal(r), duration)

	return tokenData, nil

}
