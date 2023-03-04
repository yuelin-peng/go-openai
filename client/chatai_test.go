package goopenai_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	goopenai "github.com/yuelin-pengk/go-openai/client"
)

func getDefaultToken() string {
	return "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..BtF6KRe6XCgXg2zB.yW2DolNFAXC6_MfdIFI1ivLnQD9KLidawxs58zIRa9T9LU2kH6nl9Lmf3jy4jCayKABkIAWyQY6DDZOaV47zcW0P2cQlHS78dxoK1zpwdi7nYrAtpGdyPShqQRZ_c_RqFuk1-fbSvi3j6z_b7TG_65--lMelJuSKa5GTIPRCkegah3o-LitHBmTZPJg8IGrQpOeT2_Fv5yJ1XSPqQdetHFYJXtEjWxXM2uKuCcNXmVzHcRmc18b1wJktw8W1Ba8lJtcMHRbMDCNtbTTXuX2d-l3ptjXDUscNhndcTYsfFbOvUAsYP7ifMHL8KqyQt6JeRmSSd6nCV0j1GZhK1Zas98zn2koS0GE2P0LXE-nT6wOqYqZshPnUIYgzd1FO66uddqEmOKVI1Zct2sR3I9SSk_e8N8i8z3rJI0N3bwlFbIVy6Eu7z94goOsXtFfksCRlb-znWQkAvd4L6N0TX8lp7_DwrZtzMlQ5Fy9KxKJh0gamjO-8DTnQsXRCWA8pNYiklliR3JlsgJBe3AyT11a-tTqHXSDucCJ4jJFUGAmUrLCn42ZZtOnhL1G_hRFhquB6nHp0ZKlk2gthpbwEA1bEcJftImhhkyZQFAr1jgDYZanldenYCu5_JI_RxYhz3-DpuymkfL3LJPKDFOCK2172G_1CDHBtwKsM8f-wjjjAhKGs7En2q0Z1LJTzb2mxBQi2HWtNviZfunkhBNVRr_y7Vneeb58sMr7FNXFnapTtffxlaU4Ky-IKJTXLiD7K5kujDfA3JEfJvsyuQ7IwvxFE8ElvedPXG2WBtBQzp0iC4daseIUvYEvH2QlbIwTn1Rez-MPEnxqIKnL2C2jfsCFWBvFDP17E45F7VETo5DGZgxNvowJ3TeJUnfxl7U0oci1giwmsoSUXKbXLVQoFpiAleRFtNJ1Zb658WTjFoiBwYWAplXN6IJX-dz54ASgLc8OcymCdjFFF3Vg4G2uUFebz2FzDoBtavofQ8UjTvmEglArsotQTCJEdx_gmDrXzwIXthLJFliIlcBGZPqBsW6IZUPcIMGzxt4TCHx3De8Hu7l8R1Qu-dSjPyP9qUj9_V6ACL1C-ZS5BGyh2LeoukOH97VV6rJ4TBtIQi-AakhlIJSNmNEDPR9BgTXy9Lt_JYZD257uysy-Jfovsu-HW181MVF2KCgoXpYeG4SfrcwVj1cpm8LmWNuYS0933VcnHqHq14VnQa0aqc8gRdmxELsPqal7WAtD5UA93C5Wm80qiU0-uIGruMxM4eK8FxLXApfrZdpKP57RuxXFae-WWLa3pNw_Qv1RRBjTWQ12TqcMilAI8sWlFoiNMGkjkEMoEmJxIRHom9XJli8QMs6v0P07-m2ZVg2C8xsImkK4cz3H_AGVnU0qQQaN9gLTdcndHes6Yyg2uOjt9OX-z7q_PRUjtkCEuyO2_GXvj6gAK_bx6MRCiO3ukczY__rB3tth2lnndr4RrfUzPpi88p3LUUdXGRLAMWLGiXds1k7N6O9KjAb3sbR1WQFCYvf4YG2V2ponpnl90-gNVrfkP1Mj5iABcAgrUtEy595dgJWCleTLoXyTz5n6exP2ZG5qgUYIn3KFiC1Jtq5UrkYQVBw8-VX9XjPZe3vVSUHXAgzCScp0uvsMWA5vKtgGgjBxS58WHEX4r04JvCZ-mgkKgbGkTZce36q2HsdVIg7q6gGe3j_lUFexpAb7MGeHeGNPOH5Q30_tOOpVo79XOueBzht9In6wh_bDVwmqlw26-Uh8I1VPhyyq_TOrNf6sM7K-oD3HcXm0ziehxb-JBFON4oP5sVBNtI5ucgOm8DiaP1jq1iS1L6TRUepPsyitAnsuzqyk27LCJdOtkj9l3rci7-S0_NEyRuJ9C6hExv416e_QDKdkluu-2E5fBW8s965D1mnEMp5xbwdxXgwc3ZXxdLkEmnvHRCOFuCNd5jdzv2fo4s5Jazz6Q1B_qaWOfJFAO6NwhgPEJDHAIcTqqJ_rktu6HvhPvghoqwM6LkUm4DlO3mjnzh0pUbB_gUrdwGPyUiElDlfu7hOqYbgyB28jztt8lNF7JY_Ja-xp0kPntbqtAv3qkEybEwDAVN1OWSD5_N6L35NFewmz8UgdHXf45fCAhebHvRgNhD5QAvN0Si0brFr15YriJ_3kkqRTTAzdMMnzm4qbUj6-qd6DrKMIoN0bm-r3tGFBClNlpn2Uf7XZpk3rUlDZkwSzKbLT1k4xoCI46NQfx0HsCAz0nq5NFSbgUXPTkNz0d49R1Dyv8p4OeE1hWZbJEA3n7ghfaJoREFyzeGagGKUaNXaze7uGhh1qNOqz3qkeV7x1-CrC1P7u2BqHKxqJjTi6SQea9LutqPbgoFbo-BNw_QeiGzbJk7nQqsGMRy1FEL7IhwHvvfATaNSwqT-JkRONdGyVqkWqWOH3D0l8ppM91J6W0pTnUH4RGEAl1UHs1RC4h6aZ5MxciTEwhomicLdZMQ_CKvTDCk5XjK99ecD2z4rpq_yRTzNrHwJpNJb1Ypev4lNvP_wMVtyo1ciW5z7jfO2g2dkk2vfDHOboIA9MkLh90trVlP8eEJg97A-3yUcD5DtX1cPgIOKV3Tw1DOIAscHmYCzFGWDJdumVX8oZfPdXxB1Cozfjt.TbGCydFoOSya7hjTR8P7sQ"
}

func createDefaultOpenAI(t *testing.T) *goopenai.ChatAI {
	chatAI, err := goopenai.NewChatAI(getDefaultToken())
	assert.Nil(t, err)
	assert.NotNil(t, chatAI)

	return chatAI
}

func TestNewOpenAIWithSession(t *testing.T) {
	chatAI, err := goopenai.NewChatAI(getDefaultToken())
	assert.Nil(t, err)
	assert.NotNil(t, chatAI)

	chatAI, err = goopenai.NewChatAI("")
	assert.NotNil(t, err)
	assert.Nil(t, chatAI)
}

func TestNewConversation(t *testing.T) {
	chatAI := createDefaultOpenAI(t)
	c, err := chatAI.NewConversation()
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestChatAIGetAccessToken(t *testing.T) {
	t.Skip()

	chatAI := createDefaultOpenAI(t)
	accessToken, err := chatAI.GetAccessToken()
	assert.Nil(t, err)
	assert.Greater(t, len(accessToken), 0)
}

func TestAskQueryForConversation(t *testing.T) {
	chatAI := createDefaultOpenAI(t)

	c, err := chatAI.NewConversation()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	result, err := c.AskAndAnswer(context.Background(), "hello")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	fmt.Println(result)
}
