package verify

import "testing"

type PageInfoTest struct {
	Page     int
	PageSize int
	Name     string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.Page = 0
	testInfo.PageSize = 0
	err := ParamVerify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉0值")
	}
	testInfo.Name = ""
	testInfo.Page = 1
	testInfo.PageSize = 10
	err = ParamVerify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常检测name为空")
	}
	testInfo.Name = "test"
	testInfo.Page = 1
	testInfo.PageSize = 10
	err = ParamVerify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通过检测")
	}
}
