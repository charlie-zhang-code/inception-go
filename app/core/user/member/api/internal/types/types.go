// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddSysUserReq struct {
	Id          int64  `json:"id"`                   //用户唯一标识符，主键
	Identify    string `json:"identify"`             //用户的唯一标识符
	Username    string `json:"username"`             //用户登录时使用的用户名
	Password    string `json:"password"`             //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname,optional"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar,optional"`      //用户头像的URL或路径
	Quote       string `json:"quote,optional"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday,optional"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender,optional"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email,optional"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone,optional"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt,optional"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp,optional"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion,optional"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs,optional"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status,optional"`      //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted,optional"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark      string `json:"remark,optional"`      //对记录的备注信息，如特殊说明等
}

type AddSysUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteSysUserReq struct {
	Ids []int64 `form:"ids"`
}

type DeleteSysUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type QueryPageSysUserListData struct {
	Id          int64  `json:"id"`          //用户唯一标识符，主键
	Identify    string `json:"identify"`    //用户的唯一标识符
	Username    string `json:"username"`    //用户登录时使用的用户名
	Password    string `json:"password"`    //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar"`      //用户头像的URL或路径
	Quote       string `json:"quote"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status"`      //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark      string `json:"remark"`      //对记录的备注信息，如特殊说明等
	CreateAt    string `json:"createAt"`    //记录创建的时间戳
	CreateBy    string `json:"createBy"`    //创建该记录的用户标识符
	UpdateAt    string `json:"updateAt"`    //记录最后更新的时间戳
	UpdateBy    string `json:"updateBy"`    //最后更新该记录的用户标识符
}

type QueryPageSysUserListReq struct {
	Current     int64  `form:"current,default=1"`    //第几页
	PageSize    int64  `form:"pageSize,default=20"`  //每页的数量
	Id          int64  `json:"id,optional"`          //用户唯一标识符，主键
	Identify    string `json:"identify,optional"`    //用户的唯一标识符
	Username    string `json:"username,optional"`    //用户登录时使用的用户名
	Password    string `json:"password,optional"`    //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname,optional"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar,optional"`      //用户头像的URL或路径
	Quote       string `json:"quote,optional"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday,optional"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender,optional"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email,optional"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone,optional"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt,optional"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp,optional"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion,optional"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs,optional"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status,default=2"`     //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted,optional"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
}

type QueryPageSysUserListResp struct {
	Code     string                      `json:"code"`
	Message  string                      `json:"message"`
	Current  int64                       `json:"current,default=1"`
	Data     []*QueryPageSysUserListData `json:"data"`
	PageSize int64                       `json:"pageSize,default=20"`
	Success  bool                        `json:"success"`
	Total    int64                       `json:"total"`
}

type QuerySysUserDetailData struct {
	Id          int64  `json:"id"`          //用户唯一标识符，主键
	Identify    string `json:"identify"`    //用户的唯一标识符
	Username    string `json:"username"`    //用户登录时使用的用户名
	Password    string `json:"password"`    //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar"`      //用户头像的URL或路径
	Quote       string `json:"quote"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status"`      //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark      string `json:"remark"`      //对记录的备注信息，如特殊说明等
	CreateAt    string `json:"createAt"`    //记录创建的时间戳
	CreateBy    string `json:"createBy"`    //创建该记录的用户标识符
	UpdateAt    string `json:"updateAt"`    //记录最后更新的时间戳
	UpdateBy    string `json:"updateBy"`    //最后更新该记录的用户标识符
}

type QuerySysUserDetailReq struct {
	Id int64 `form:"id"`
}

type QuerySysUserDetailResp struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    QuerySysUserDetailData `json:"data"`
}

type QuerySysUserListData struct {
	Id          int64  `json:"id"`          //用户唯一标识符，主键
	Identify    string `json:"identify"`    //用户的唯一标识符
	Username    string `json:"username"`    //用户登录时使用的用户名
	Password    string `json:"password"`    //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar"`      //用户头像的URL或路径
	Quote       string `json:"quote"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status"`      //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark      string `json:"remark"`      //对记录的备注信息，如特殊说明等
	CreateAt    string `json:"createAt"`    //记录创建的时间戳
	CreateBy    string `json:"createBy"`    //创建该记录的用户标识符
	UpdateAt    string `json:"updateAt"`    //记录最后更新的时间戳
	UpdateBy    string `json:"updateBy"`    //最后更新该记录的用户标识符
}

type QuerySysUserListReq struct {
	Id          int64  `json:"id,optional"`          //用户唯一标识符，主键
	Identify    string `json:"identify,optional"`    //用户的唯一标识符
	Username    string `json:"username,optional"`    //用户登录时使用的用户名
	Password    string `json:"password,optional"`    //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname,optional"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar,optional"`      //用户头像的URL或路径
	Quote       string `json:"quote,optional"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday,optional"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender,optional"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email,optional"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone,optional"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt,optional"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp,optional"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion,optional"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs,optional"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status,default=2"`     //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted,optional"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
}

type QuerySysUserListResp struct {
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Data    []*QuerySysUserListData `json:"data"`
	Success bool                    `json:"success"`
}

type UpdateSysUserReq struct {
	Id          int64  `json:"id"`                   //用户唯一标识符，主键
	Identify    string `json:"identify"`             //用户的唯一标识符
	Username    string `json:"username"`             //用户登录时使用的用户名
	Password    string `json:"password"`             //用户账户的密码，应存储加密后的值
	Nickname    string `json:"nickname,optional"`    //用户在系统中显示的名字，可选
	Avatar      string `json:"avatar,optional"`      //用户头像的URL或路径
	Quote       string `json:"quote,optional"`       //用户设置的个人座右铭或签名
	Birthday    string `json:"birthday,optional"`    //用户的出生日期，用于个性化服务或统计
	Gender      int32  `json:"gender,optional"`      //用户的性别，0表示保密，1表示男性，2表示女性
	Email       string `json:"email,optional"`       //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone   string `json:"telephone,optional"`   //用户的联系电话，可用于身份验证或联系用户
	LoginAt     string `json:"loginAt,optional"`     //用户最近一次登录的时间戳
	LoginIp     string `json:"loginIp,optional"`     //用户最近一次登录的IP地址
	LoginRegion string `json:"loginRegion,optional"` //用户最近一次登录的IP地址所属的地理位置
	LoginOs     string `json:"loginOs,optional"`     //用户最近一次登录的操作系统
	Status      int32  `json:"status,optional"`      //状态，0表示禁用，1表示正常启用
	Deleted     int32  `json:"deleted,optional"`     //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark      string `json:"remark,optional"`      //对记录的备注信息，如特殊说明等
}

type UpdateSysUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UpdateSysUserStatusReq struct {
	Status int32 `json:"status,default=2"` //状态，0表示禁用，1表示正常启用
}

type UpdateSysUserStatusResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}