package sdk

import (
	"errors"
	"strconv"
)

type ID int64

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(id), 10) + `"`), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*id = ID(i)
	return nil
}

var (
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidLoginToken = errors.New("invalid login token")
	ErrInvalidComboToken = errors.New("invalid combo token")
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Retcode int32  `json:"retcode"`
}

var ResponseMessage = map[int32]string{
	0:    "OK",
	-101: "系统错误",
	-102: "密码格式错误，密码格式为8-30位，并且由数字、大小写字母、英文特殊符号两种以上组合",
	-103: "参数错误",
	-104: "缺少配置",
	-106: "协议加载失败",
	-107: "渠道错误",
	-111: "???",
	-115: "请前往官网/商店下载最新版本",
	-202: "账号或密码错误",
	-210: "为了您的账号安全，请重新登录。",
}

func NewResponse(retcode int32, data any) *Response {
	return &Response{
		Data:    data,
		Message: ResponseMessage[retcode],
		Retcode: retcode,
	}
}

type CompareProtocolVersionRequestData struct {
	ID        string `json:"id"`
	AppID     string `json:"app_id"`
	ChannelID string `json:"channel_id"`
	Language  string `json:"language"`
	Major     string `json:"major"`
	Minimum   string `json:"minimum"`
}

type ProtocolVersion struct {
	ID            int64  `json:"id"`
	AppID         int64  `json:"app_id"`
	Language      string `json:"language"`
	UserProto     string `json:"user_proto"`
	PrivProto     string `json:"priv_proto"`
	Major         int32  `json:"major"`
	Minimum       int32  `json:"minimum"`
	CreateTime    string `json:"create_time"`
	TeenagerProto string `json:"teenager_proto"`
	ThirdProto    string `json:"third_proto"`
}

type CompareProtocolVersionResponseData struct {
	Modified bool             `json:"modified"`
	Protocol *ProtocolVersion `json:"protocol"`
}

type TokenCheckRequestData struct {
	AppID      int32  `json:"app_id"`
	ChannelID  int32  `json:"channel_id"`
	OpenID     ID     `json:"open_id"`
	ComboToken string `json:"combo_token"`
	Sign       string `json:"sign"`
	Region     string `json:"region"`
}

type TokenCheckResponseData struct {
	AccountType int32 `json:"account_type"`
	IPInfo      struct {
		CountryCode string `json:"country_code"`
	} `json:"ip_info"`
}

type ShieldLoginRequestData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

type Account struct {
	UID               ID     `json:"uid"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	IsEmailVerify     string `json:"is_email_verify"`
	RealName          string `json:"realname"`
	IdentityCard      string `json:"identity_card"`
	Token             string `json:"token"`
	SafeMobile        string `json:"safe_mobile"`
	FacebookName      string `json:"facebook_name"`
	GoogleName        string `json:"google_name"`
	TwitterName       string `json:"twitter_name"`
	GameCenterName    string `json:"game_center_name"`
	AppleName         string `json:"apple_name"`
	SonyName          string `json:"sony_name"`
	TapName           string `json:"tap_name"`
	Country           string `json:"country"`
	ReactivateTicket  string `json:"reactivate_ticket"`
	AreaCode          string `json:"area_code"`
	DeviceGrantTicket string `json:"device_grant_ticket"`
	SteamName         string `json:"steam_name"`
}

type ShieldLoginResponseData struct {
	Account             *Account `json:"account"`
	DeviceGrantRequired bool     `json:"device_grant_required"`
	SafeMobileRequired  bool     `json:"safe_mobile_required"`
	RealPersonRequired  bool     `json:"realperson_required"`
	ReactivateRequired  bool     `json:"reactivate_required"`
	RealNameOperation   string   `json:"realname_operation"`
}

type ShieldVerifyRequestData struct {
	UID   ID     `json:"uid"`
	Token string `json:"token"`
}

type ShieldVerifyResponseData struct {
	Account             *Account `json:"account"`
	DeviceGrantRequired bool     `json:"device_grant_required"`
	SafeMobileRequired  bool     `json:"safe_mobile_required"`
	RealPersonRequired  bool     `json:"realperson_required"`
	RealNameOperation   string   `json:"realname_operation"`
}

type ComboLoginRequestData struct {
	AppID     int32  `json:"app_id"`
	ChannelID int32  `json:"channel_id"`
	Data      string `json:"data"`
	Device    string `json:"device"`
	Sign      string `json:"sign"`
}

type ComboLoginData struct {
	UID   ID     `json:"uid"`
	Guest bool   `json:"guest"`
	Token string `json:"token"`
}

type ComboLoginResponseData struct {
	ComboID       string `json:"combo_id"`
	OpenID        ID     `json:"open_id"`
	AccountType   int32  `json:"account_type"`
	ComboToken    string `json:"combo_token"`
	Data          string `json:"data"`
	FatigueRemind any    `json:"fatigue_remind"`
	Heartbeat     bool   `json:"heartbeat"`
}
