package my_handlers

import (
	"errors"
	"fmt"
	kook_CardBuild "github.com/Quinlivanner/kook-CardBuild"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/kaiheila/golang-bot/api/base"
	event2 "github.com/kaiheila/golang-bot/api/base/event"
	"github.com/kaiheila/golang-bot/api/helper"
	log "github.com/sirupsen/logrus"
	"strings"
)

type FaqEventHandler struct {
	Token   string
	BaseUrl string
}

func (gteh *FaqEventHandler) Handle(e event.Event) error {
	log.WithField("event", fmt.Sprintf("%+v", e.Data())).Info("收到频道内的文字消息.")
	err := func() error {
		if _, ok := e.Data()[base.EventDataFrameKey]; !ok {
			return errors.New("data has no frame field")
		}
		frame := e.Data()[base.EventDataFrameKey].(*event2.FrameMap)
		data, err := sonic.Marshal(frame.Data)
		if err != nil {
			return err
		}
		msgEvent := &event2.MessageKMarkdownEvent{}
		err = sonic.Unmarshal(data, msgEvent)
		log.Infof("Received json event:%+v", msgEvent)
		if err != nil {
			return err
		}
		client := helper.NewApiHelper("/v3/message/create", gteh.Token, gteh.BaseUrl, "", "")
		if msgEvent.Author.Bot {
			log.Info("bot message")
			return nil
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "账号") && ContainsPhrase(msgEvent.KMarkdown.RawContent, "注册") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你可以在我们的官方网站注册账号" + linkUrl("https://wotlk.everlook-wow.net"))
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "改密码") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("登录官网后，点击个人资料，点击更改密码" + linkUrl("https://wotlk.everlook-wow.net"))
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "忘记密码") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("打开官网，点击登录，点击忘记密码，输入你注册时的邮箱地址，点击重置密码。" + linkUrl("https://wotlk.everlook-wow.net"))
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "客户端") && ContainsPhrase(msgEvent.KMarkdown.RawContent, "下载") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你可以在kook的客户端频道下载客户端" + "(chn)" + "7922271820000424" + "(chn)")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "高清客户端") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("高清客户端存在较多Bug，不推荐使用，除非你能自己解决遇到的问题。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "启动器") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("由于网络原因，EverLook启动器需要开启VPN软件才能使用")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "服务器地址") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("服务器地址为：logon.everlook-wow.net")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "登录错误") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请逐一尝试以下方法：")
			Cards.AddKmarkdown("1. 确保你是在80级网页注册的账号，而不是60级网页，两个版本的游戏账号不通用。")
			Cards.AddKmarkdown("2. 确保输入的是注册时的用户名，而不是邮箱。")
			Cards.AddKmarkdown("3. 确保输入的账号密码正确，可以通过是否能成功登录官网判断。")
			Cards.AddKmarkdown("4. 确保你的realmlist.wtf文件中的服务器地址为set realmlist logon.everlook-wow.net。")
			Cards.AddKmarkdown("5. 确保你的客户端是从kook频道下载的。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "登录不上") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请逐一尝试以下方法：")
			Cards.AddKmarkdown("1. 确保你是在80级网页注册的账号，而不是60级网页，两个版本的游戏账号不通用。")
			Cards.AddKmarkdown("2. 确保输入的是注册时的用户名，而不是邮箱。")
			Cards.AddKmarkdown("3. 确保输入的账号密码正确，可以通过是否能成功登录官网判断。")
			Cards.AddKmarkdown("4. 确保你的realmlist.wtf文件中的服务器地址为set realmlist logon.everlook-wow.net。")
			Cards.AddKmarkdown("5. 确保你的客户端是从kook频道下载的。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "断开连接") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("先尝试开启加速器，然后打开游戏目录下WTF文件，找到Config.wtf文件，使用记事本工具打开它，将SET realmName修改为：SET realmName \"冰封王座\" 。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "加入世界频道") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请在游戏聊天框内输入/join 大脚世界频道")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "频道提示") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请在游戏聊天框内输入/script ChatFrame_RemoveMessageGroup(ChatFrame1, \"CHANNEL\")")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "无法发言") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你需要完成50个任务，等级达到30级。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "试玩账号") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你需要达到10级才能邀请玩家组队。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "开启3倍经验") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("在游戏聊天框内输入.xp view来查看当前经验倍数，请输入.xp set 3开启三倍经验。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "开启硬核模式") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请在游戏出生地找到名为“奇怪的卷轴”告示牌，点击后选择“与死神搏斗”。请注意：硬核模式无法使用邮件功能，请勿使用大号邮寄传家宝、金币等物资，我们不不会为您恢复因此丢失的物资。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "硬核角色恢复") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("硬核角色死亡后无法再次进入游戏。请在硬核模式下开启显卡的视频回放功能，通过视频证明为服务器bug导致的死亡，可以申请复活。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "证据链接") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你可以选择如下方式获得链接：")
			Cards.AddKmarkdown("1:图片链接推荐使用：https://gyazo.com/")
			Cards.AddKmarkdown("2:视频链接推荐使用：https://www.bilibili.com/")

			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "封禁申诉") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请登录官网，点击个人资料，点击联系支持，查看封禁原因，并提交申诉。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "捐赠问题") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请前往https://everlook.zendesk.com/hc/en-us/requests/new，寻求帮助。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "联系GM") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("在游戏内输入/GM命令，打开帮助面板，点击联系GM，描述问题。寻求帮助。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "角色卡死") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("在游戏内输入/GM命令，打开帮助面板，点击角色卡死，选择自动脱离卡死。如果你使用的是高清客户端，请更换为普通客户端。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "任务无法完成") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请关闭所有插件，退出游戏，打开游戏目录下Cache文件，删除WDB后再次尝试。如果你使用的是高清客户端，请更换为普通客户端。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "任务奖励无法选择") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请关闭所有插件，退出游戏，打开游戏目录下Cache文件，删除WDB后再次尝试。如果你使用的是高清客户端，请更换为普通客户端。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "物品无法退还") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请关闭所有插件，退出游戏，打开游戏目录下Cache文件，删除WDB后再次尝试。如果你使用的是高清客户端，请更换为普通客户端。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "弥补关系") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("你应该先击杀杀死弗约恩和5名雷铸铁巨人，获得熔渣包裹的金属。如果还是存在问题，请在游戏内联系GM寻求帮助。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "修复的误会") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请尝试接受任务后不要进入战斗，不要使用坐骑，到达护送目的地后与受伤的雨声神谕者重叠站位。如果还是存在问题，请在游戏内联系GM寻求帮助。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "一个也不能少") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请尝试于解救的NPC对话后不要进入战斗，不要使用坐骑。如果还是存在问题，请在游戏内联系GM寻求帮助。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "伊米隆的回响") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请查找攻略，在正确的位置使用道具，并且与NPC重叠站位。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "风暴之子瓦杜兰") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请确保你输出的伤害大于布鲁沃·斩铁和塑石者布德克拉格造成的伤害。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "委以重任") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("此任务目前存在bug，还未修复，请先跳过。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "边缘科学的益处") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("此任务目前存在bug，还未修复，请先跳过。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "宝石描述") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("珠宝的显示目前存在问题，请以实际效果为准。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "132错误") {
			Cards, err := kook_CardBuild.NewCardWithOption(kook_CardBuild.CardThemeNone, "lg", "#000000")
			if err != nil {
				return err
			}

			Cards.AddKmarkdown("(met)" + msgEvent.Author.ID + "(met)")
			Cards.AddKmarkdown("请降低游戏效果，减少插件数量，删除WDB文件。或者尝试Large Address Aware大内存工具。")
			CardsContent, err := kook_CardBuild.GenerateCardMessageContent(Cards)
			if err != nil {
				return err
			}

			SendGroupCardessage(CardsContent, msgEvent.TargetId, client)
		}

		return nil
	}()
	if err != nil {
		log.WithError(err).Error("GroupTextEventHandler err")
	}

	return nil
}

func DoesNotContainPhrase(text, phrase string) bool {
	// 将文字和词组都转换为小写，然后检查是否不包含词组
	return !strings.Contains(strings.ToLower(text), strings.ToLower(phrase))
}
