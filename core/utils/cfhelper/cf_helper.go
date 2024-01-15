package cfhelper

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type CfUserInfoResp struct {
	Status string        `json:"status"`
	Result []*CfUserInfo `json:"result"`
}

type CfUserInfo struct {
	LastName                string `json:"lastName"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	Rating                  int    `json:"rating"`
	FriendOfCount           int    `json:"friendOfCount"`
	TitlePhoto              string `json:"titlePhoto"`
	Handle                  string `json:"handle"`
	Avatar                  string `json:"avatar"`
	FirstName               string `json:"firstName"`
	Contribution            int    `json:"contribution"`
	Organization            string `json:"organization"`
	Rank                    string `json:"rank"`
	MaxRating               int    `json:"maxRating"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	MaxRank                 string `json:"maxRank"`
	Country                 string `json:"country,omitempty"`
	City                    string `json:"city,omitempty"`
}

type CfHelper struct {
	client *resty.Client
}

func NewCfHelper(client *resty.Client) *CfHelper {
	return &CfHelper{client: client}
}

func (c *CfHelper) GetUserInfo(handle string) (resp *CfUserInfo, err error) {
	doResp, err := c.getUserInfo(handle)
	if err != nil {
		return nil, err
	}
	return doResp.Result[0], nil
}

func (c *CfHelper) getUserInfo(handle string) (resp *CfUserInfoResp, err error) {
	resp = &CfUserInfoResp{}

	response, err := c.client.R().
		SetResult(resp).
		SetQueryParam("handles", handle).
		Get("https://codeforces.com/api/user.info")

	if err != nil {
		return nil, err
	}

	if resp.Status != "OK" {
		err = errors.New(response.String())
		return nil, err
	}
	return
}
