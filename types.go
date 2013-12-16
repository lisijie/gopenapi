package gopenapi

type T_RetBase struct {
	Ret     int
	Msg     string
	Is_Lost int
}

//GetInfo() 返回的用户信息结构
type T_UserInfo struct {
	T_RetBase
	Nickname           string //昵称
	Gender             string //性别
	Country            string //国家（当pf=qzone、pengyou或qplus时返回）
	Province           string //省（当pf=qzone、pengyou或qplus时返回）
	City               string //市（当pf=qzone、pengyou或qplus时返回）
	Figureurl          string //头像URL
	Openid             string //用户QQ号码转化得到的ID（当pf=qplus时返回）
	Qq_level           int    //用户QQ等级（当pf=qplus时返回）
	Qq_vip_level       int    //用户QQ会员等级（当pf=qplus时返回）
	Qplus_level        int    //用户Q+等级（当pf=qplus时返回）
	Is_yellow_vip      int    //是否为黄钻用户（0：不是； 1：是）
	Is_yellow_year_vip int    //是否为年费黄钻用户（0：不是； 1：是）
	Yellow_vip_level   int    //黄钻等级，目前最高级别为黄钻8级（如果是黄钻用户才返回此参数）
	Is_yellow_high_vip int    //是否为豪华版黄钻用户（0：不是； 1：是）
	Is_blue_vip        int    //是否为蓝钻用户（0：不是； 1：是）
	Is_blue_year_vip   int    //是否为年费蓝钻用户（0：不是； 1：是）
	Blue_vip_level     int    //蓝钻等级（如果是蓝钻用户才返回此参数）
	Is_super_blue_vip  int    //是否是豪华蓝钻
}

//TotalVipInfo() 返回的用户VIP信息
type T_UserVipInfo struct {
	T_RetBase
	//QQ会员
	Is_vip      int //是否为QQ会员（0：不是； 1：是）
	Is_Year_Vip int //是否是年费QQ会员（0：不是； 1：是）
	Vip_level   int //QQ会员等级
	Is_high_vip int //是否豪华版QQ会员（0：不是； 1：是）
	//蓝钻
	Is_Blue      int //是否为蓝钻用户（0：不是； 1：是）
	Is_year_blue int //是否是年费蓝钻用户（0：不是； 1：是）
	Blue_level   int //蓝钻等级
	Is_high_blue int //是否豪华版蓝钻（0：不是； 1：是）
	//黄钻
	Is_yellow         int //是否为黄钻用户（0：不是； 1：是）
	Is_year_yellow    int //是否是年费黄钻用户（0：不是； 1：是）
	Yellow_level      int //黄钻等级
	Is_high_yellow    int //是否豪华版黄钻（0：不是； 1：是）
	Is_gamevip_yellow int //是否为超级套餐包月（0：不是； 1：是）
	//红钻
	Is_red      int //是否为红钻用户（0：不是； 1：是）
	Red_level   int //红钻等级
	Is_high_red int //是否豪华版红钻（0：不是； 1：是）
	//绿钻
	Is_green      int //是否为绿钻用户（0：不是； 1：是）
	Is_year_green int //是否是年费绿钻用户（0：不是； 1：是）
	Green_level   int //绿钻等级
	Is_high_green int //是否豪华绿钻（0：不是； 1：是）
	//粉钻
	Is_pink         int //是否为粉钻用户（0：不是； 1：是）
	Is_year_pink    int //是否是年费粉钻用户（0：不是； 1：是）
	Pink_level      int //粉钻等级
	Is_pink_15      int //是否是15元粉钻用户（0：不是； 1：是）
	Is_year_pink_15 int //是否是年费15元粉钻用户（0：不是； 1：是）
	//超级qq
	Is_superqq    int //是否为超级QQ用户（0：不是； 1：是）
	Superqq_level int //超级QQ等级
	//3366
	Level_3366 int //3366会员等级
}

//IsVip() 返回的用户VIP信息
type T_UserIsVip struct {
	T_RetBase
	Is_yellow_vip      int //是否为黄钻用户（0：不是； 1：是）
	Is_yellow_year_vip int //是否为年费黄钻用户（0：不是； 1：是）
	Yellow_vip_level   int //黄钻等级。目前最高级别为黄钻8级（如果是黄钻用户才返回此字段）
	Is_yellow_high_vip int //是否为豪华版黄钻用户（0：不是； 1：是）
	Yellow_vip_pay_way int //用户的付费类型
}

type T_BuyGoods struct {
	T_RetBase
	Token      string
	Url_params string
}

type T_Balance struct {
	T_RetBase
	Balance      int //获取到的游戏币余额
	Comm_balance int //如果存在公共游戏币账户，则这里返回公共游戏币账户中游戏币余额，否则不返回该参数
}
