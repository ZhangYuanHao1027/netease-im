package netease

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	createTeam = neteaseBaseURL + "/team/create.action" //建群
	removeTeam = neteaseBaseURL + "/team/remove.action" //解散群

	updateTeam      = neteaseBaseURL + "/team/update.action" //更新群信息
	queryTeamDetail = neteaseBaseURL + "/team/update.action" //获取群信息

	addMember  = neteaseBaseURL + "/team/add.action"   //拉人
	kickMember = neteaseBaseURL + "/team/kick.action"  //踢人
	leaveTeam  = neteaseBaseURL + "/team/leave.action" //退群

	changeOwner   = neteaseBaseURL + "/team/changeOwner.action"   //移交群主
	addManager    = neteaseBaseURL + "/team/addManager.action"    //任命管理员
	removeManager = neteaseBaseURL + "/team/removeManager.action" //移除管理员

	muteTlist    = neteaseBaseURL + "/team/muteTlist.action"    //禁言群成员
	muteTlistAll = neteaseBaseURL + "/team/muteTlistAll.action" //禁言全部成员
)

//CreateTeam 建群
/**
 * @param tname 群名称
 * @param owner 创建者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param msg 0 邀请发送的文字，最大长度150字符
 * @param members ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 * @param magree 0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群
 * @param joinmode 0不用验证，1需要验证,2不允许任何人加入
 */
func (c *ImClient) CreateTeam(tname, owner, msg string, members []string, magree, joinmode int, icon string) (string, error) {
	param := map[string]string{"tname": tname}
	param["owner"] = owner
	param["msg"] = msg

	m, err := jsonTool.MarshalToString(members)
	if err != nil {
		return "", err
	}
	param["members"] = m
	param["magree"] = strconv.Itoa(magree)
	param["joinmode"] = strconv.Itoa(joinmode)
	param["icon"] = icon

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(createTeam)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	var tid string
	var faccid map[string]interface{}
	err = json.Unmarshal(*jsonRes["code"], &code)
	err = json.Unmarshal(*jsonRes["tid"], &tid)
	err = json.Unmarshal(*jsonRes["faccid"], &faccid)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//RemoveTeam 建群
/**
 * @param tid 群ID
 * @param owner 创建者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 */
func (c *ImClient) RemoveTeam(tid, owner string) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(removeTeam)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//UpdateTeam 更新群信息
/**
 * @param tid 群ID
 * @param tname 群名称
 * @param owner 创建者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param announcement 群公告
 * @param intro 群描述
 * @param icon 群头像
 * @param joinmode 0不用验证，1需要验证,2不允许任何人加入

 */
func (c *ImClient) UpdateTeam(tid, tname, owner, announcement, intro, icon string, joinmode int) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner
	param["tname"] = tname
	param["announcement"] = announcement
	param["intro"] = intro
	param["icon"] = icon
	param["icon"] = icon

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateTeam)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//QueryTeamDetail 获取群组详细信息
/**
 * @param tid 群ID
 */
func (c *ImClient) QueryTeamDetail(tid string) (string, error) {
	param := map[string]string{"tid": tid}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryTeamDetail)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//AddMember 拉人入群
/**
 * @param tid 群ID
 * @param owner 创建者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param members ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 * @param magree 0不需要被邀请人同意加入群，1需要被邀请人同意才可以加入群
 * @param msg 0 邀请发送的文字，最大长度150字符
 */
func (c *ImClient) AddMember(tid, owner string, members []string, magree int, msg string) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	m, err := jsonTool.MarshalToString(members)
	if err != nil {
		return "", err
	}
	param["members"] = m
	param["magree"] = strconv.Itoa(magree)
	param["msg"] = msg

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(addMember)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//KickMember 踢人出群
/**
 * @param tid 群ID
 * @param owner 创建者accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param members ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
 */
func (c *ImClient) KickMember(tid, owner string, members []string) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	m, err := jsonTool.MarshalToString(members)
	if err != nil {
		return "", err
	}
	param["members"] = m

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(kickMember)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//LeaveTeam 主动退群
/**
 * @param tid 群ID
 * @param accid accid，用户帐号，最大32字符，必须保证一个APP内唯一
 */
func (c *ImClient) LeaveTeam(tid, accid string) (string, error) {
	param := map[string]string{"tid": tid}
	param["accid"] = accid

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(leaveTeam)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//ChangeOwner 移交群主
/**
 * @param tid 群ID
 * @param owner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param newowner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param leave 1:群主解除群主后离开群，2：群主解除群主后成为普通成员。其它414
 */
func (c *ImClient) ChangeOwner(tid, owner, newowner string, leave int) (string, error) {
	param := map[string]string{"tid": tid}
	param["accid"] = accid

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(changeOwner)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//AddManager 任命管理员
/**
 * @param tid 群ID
 * @param owner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param members ["aaa","bbb"]（JSONArray对应的accid，长度最大1024字符（一次添加最多10个管理员）
 */
func (c *ImClient) AddManager(tid, owner string, members []string) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	m, err := jsonTool.MarshalToString(members)
	if err != nil {
		return "", err
	}
	param["members"] = m

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(addManager)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//RemoveManager 移除管理员
/**
 * @param tid 群ID
 * @param owner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param members ["aaa","bbb"]（JSONArray对应的accid，长度最大1024字符（一次添加最多10个管理员）
 */
func (c *ImClient) RemoveManager(tid, owner string, members []string) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	m, err := jsonTool.MarshalToString(members)
	if err != nil {
		return "", err
	}
	param["members"] = m

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(removeManager)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//MuteTlist 禁言群成员
/**
 * @param tid 群ID
 * @param owner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param accid accid，禁言对象的accid
 * @param mute 1-禁言，0-解禁
 */
func (c *ImClient) MuteTlist(tid, owner, accid string, mute int) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner
	param["accid"] = accid

	param["mute"] = strconv.Itoa(mute)

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteTlist)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}

//MuteTlistAll 禁言群成员
/**
 * @param tid 群ID
 * @param owner accid，用户帐号，最大32字符，必须保证一个APP内唯一
 * @param accid accid，禁言对象的accid
 * @param mute 禁言类型 0:解除禁言，1:禁言普通成员 3:禁言整个群(包括群主)
 */
func (c *ImClient) muteTlistAll(tid, owner string, mute int) (string, error) {
	param := map[string]string{"tid": tid}
	param["owner"] = owner

	param["muteType"] = strconv.Itoa(mute)

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteTlistAll)

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)

	if err != nil {
		return string(resp.Body()), err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)

	if err != nil {
		return string(resp.Body()), err
	}

	if code != 200 {
		return string(resp.Body()), errors.New("云信接口返回错误")
	}

	return string(resp.Body()), nil
}
