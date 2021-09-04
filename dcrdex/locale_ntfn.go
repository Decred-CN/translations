var TemplateKeys = map[string]string{
	// [host]
	SubjectAccountRegistered: "您现在可以在 %s 进行交易",
	// [confs, host]
	SubjectFeePaymentInProgress: "在 %s 交易之前等待 %d 确认",
	// [confs, required confs]
	SubjectRegUpdate: "费用支付确认 %v/%v",
	// [host, error]
	SubjectFeePaymentError: "向 %s 支付费用时遇到错误: %v",
	// [host, error]
	SubjectAccountUnlockError: "解锁 %s 的帐户时出错: %v",
	// [host]
	SubjectFeeCoinError: "%s 的费用硬币为空。",
	// [host]
	SubjectWalletConnectionWarning: "检测到 %s 的注册不完整，无法连接到 Decred 钱包",
	// [host, error]
	SubjectWalletUnlockError: "已连接到 Decred 钱包以在 %s 完成注册，但无法解锁：%v",
	// [ticker, error]
	SubjectWithdrawError: "在 %s 提取过程中遇到错误: %v",
	// [ticker, coin ID]
	SubjectWithdrawSend: "%s 的提款已成功完成。硬币 ID = %s",
	// [error]
	SubjectOrderLoadFailure: "某些订单无法从数据库加载: %v",
	// [qty, ticker, token]
	SubjectYoloPlaced: "以市场价格 (%s) 出售 %.8f %s",
	// [sell string, qty, ticker, rate string, token]
	SubjectOrderPlaced: "%sing %.8f %s, rate = %s (%s)",
	// [missing count, token, host]
	SubjectMissingMatches: "%d 订单 %s 的匹配没有被 %q 报告并被视为已撤销",
	// [token, error]
	SubjectWalletMissing: "活动订单 %s 的钱包检索错误: %v",
	// [side, token, match status]
	SubjectMatchErrorCoin: "订单 %s 的匹配 %s 处于状态 %s，但没有交换硬币服务商。",
	// [side, token, match status]
	SubjectMatchErrorContract: "订单 %s 的匹配 %s 处于状态 %s，没有服务商交换合约。",
	// [ticker, contract, token, error]
	SubjectMatchRecoveryError: "在订单 %s: %v 的交易恢复期间审核对方的交易合约 (%s %v) 时出错",
	// [token]
	SubjectOrderCoinError: "没有为活动订单 %s 记录资金硬币",
	// [token, ticker, error]
	SubjectOrderCoinFetchError: "订单 %s (%s) 的源硬币检索错误: %v",
	// [token, ticker, error]
	SubjectMissedCancel: "取消订单与订单 %s 不匹配。如果取消订单与交易在同一时期提交，或者目标订单在与取消订单匹配之前完全执行，则可能发生这种情况。",
	// [capitalized sell string, base ticker, quote ticker, host, token]
	SubjectOrderCanceled: "%s 的 %s-%s 的 %s 订单已被取消 (%s)",
	// [capitalized sell string, base ticker, quote ticker, fill percent, token]
	SubjectMatchesMade: "%s 订单 %s-%s %.1f%% 已完成 (%s)",
	// [qty, ticker, token]
	SubjectSwapSendError: "在订单 %s 上发送价值 %.8f %s 的交换输出时遇到错误",
	// [match, error]
	SubjectInitError: "通知 DEX 匹配 %s 的交换时出错：%v",
	// [match, error]
	SubjectReportRedeemError: "错误通知 DEX 匹配 %s: %v",
	// [qty, ticker, token]
	SubjectSwapsInitiated: "已发送价值 %.8f %s 的交易，订单 %s",
	// [qty, ticker, token]
	SubjectRedemptionError: "在订单 %s 上发送价值 %.8f %s 的兑换时遇到错误",
	// [qty, ticker, token]
	SubjectMatchComplete: "在订单 %s 上兑换了 %.8f %s",
	// [qty, ticker, token]
	SubjectRefundFailure: "退款％.8f％s的订单％S，但出现一些错误",
	// [qty, ticker, token]
	SubjectMatchesRefunded: "已退款 %.8f %s 订单 %s",
	// [match ID token]
	SubjectMatchRevoked: "匹配 %s 已被撤销",
	// [token, market name, host]
	SubjectOrderRevoked: "%s 市场 %s 上的订单 %s 已被服务器撤销",
	// [token, market name, host]
	SubjectOrderAutoRevoked: "%s 市场 %s 上的订单 %s 由于市场暂停而被撤销",
	// [ticker, coin ID, match]
	SubjectMatchRecovered: "找到制造商的赎回 (%s: %v) 和匹配 %s 的验证秘密",
	// [token]
	SubjectCancellingOrder: "已为订单 %s 提交了取消操作",
	// [token, old status, new status]
	SubjectOrderStatusUpdate: "订单 %v 的状态从 %v 修改为 %v",
	// [count, host, token]
	SubjectMatchResolutionError: "%s 报告的 %d 个匹配项没有找到 %s。",
	// [token]
	SubjectFailedCancel: "取消订单 %s 的订单，该订单在 Epoch 状态中停留了 2 个时期，现在已被删除。",
	// [coin ID, ticker, match]
	SubjectAuditTrouble: "仍在搜索交易对手的合约交易 %v (%s) 以匹配 %s。您的互联网和钱包连接是否良好？",
	// [host, error]
	SubjectDexAuthError: "%s: %v",
	// [count, host]
	SubjectUnknownOrders: "未找到 DEX %s 报告的 %d 个活动订单。",
	// [count]
	SubjectOrdersReconciled: "%d 个订单的状态已更新。",
	// [ticker, address]
	SubjectWalletConfigurationUpdated: "%s 钱包的配置已更新。存款地址 = %s",
	//  [ticker]
	SubjectWalletPasswordUpdated: "%s 钱包的密码已更新。",
	// [market name, host, time]
	SubjectMarketSuspendScheduled: "%s 的市场 %s 现在计划在 %v 暂停",
	// [market name, host]
	SubjectMarketSuspended: "%s 市场 %s 的交易现已暂停。",
	// [market name, host]
	SubjectMarketSuspendedWithPurge: "%s 市场 %s 的交易现已暂停。所有预订的订单现在都已清除。",
	// [market name, host, time]
	SubjectMarketResumeScheduled: "%s 的市场 %s 现在计划在 %v 恢复",
	// [market name, host, epoch]
	SubjectMarketResumed: "M%s 的市场 %s 已在epoch %d 恢复交易",
	// [host]
	SubjectUpgradeNeeded: "您可能需要更新您的客户端以在 %s 进行交易。",
	// [host]
	SubjectDEXConnected: "%s 已连接",
	// [host]
	SubjectDEXDisconnected: "%s 已断开连接",
	// [host, rule, time, details]
	SubjectPenalized: "来自 DEX 的惩罚在 %s \n最后违反规则：%s \n时间：%v \n详细信息：\n \" %s \" \n ",
}

// traduzir os subject tudo
var BrPtLocales = map[string]*Translation{
	SubjectDEXConnected: &Translation{
		Subject:  "DEX 连接",
		Template: "%s 已连接",
	},
	SubjectAccountRegistered: &Translation{
		Subject:  "注册账户",
		Template: "您现在可以切换到 %s",
	},
	SubjectFeePaymentInProgress: &Translation{
		Subject:  "费用支付中",
		Template: "在切换到 %s 之前等待 %d 次确认",
	},
	SubjectRegUpdate: &Translation{
		Subject:  "记录更新",
		Template: "%v/%v 费率确认",
	},
	SubjectFeePaymentError: &Translation{
		Subject:  "费用支付错误",
		Template: "为 %s 支付费率时出错：%v",
	},
	SubjectAccountUnlockError: &Translation{
		Subject:  "解锁钱包时出错",
		Template: "解锁帐户 %s 时出错：%v",
	},
	SubjectFeeCoinError: &Translation{
		Subject:  "汇率错误",
		Template: "%s 的空置率。",
	},
	SubjectWalletConnectionWarning: &Translation{
		Subject:  "钱包连接通知",
		Template: "检测到 %s 的注册不完整，无法连接 decred 钱包",
	},
	SubjectWalletUnlockError: &Translation{
		Subject:  "解锁钱包时出错",
		Template: "与 decred 钱包连接以在 %s 上完成注册，但无法解锁：%v",
	},
	SubjectWithdrawError: &Translation{
		Subject:  "提款错误",
		Template: "删除 %s 时遇到错误：%v",
	},
	SubjectWithdrawSend: &Translation{
		Template: "成功完成 %s 的提款。货币 ID = %s",
		Subject:  "提款已发送",
	},
	SubjectOrderLoadFailure: &Translation{
		Template: "某些请求无法从数据库加载：%v",
		Subject:  "请求加载失败",
	},
	SubjectYoloPlaced: &Translation{
		Template: "以市场价格 (%s) 出售 %.8f %s",
		Subject:  "下达市价单",
	},
	SubjectOrderPlaced: &Translation{
		Subject:  "已下订单",
		Template: "%sing %.8f %s，值 = %s (%s)",
	},
	SubjectMissingMatches: &Translation{
		Template: "%s 订单的 %d 匹配项未被 %q 报告并被视为已撤销",
		Subject:  "订单缺失匹配",
	},
	SubjectWalletMissing: &Translation{
		Template: "通过钱包 %s 检索活动订单时出错：%v",
		Subject:  "丢失的钱包",
	},
	SubjectMatchErrorCoin: &Translation{
		Subject:  "货币不匹配错误",
		Template: "订单 %s 的组合 %s 处于状态 %s，但没有用于交换货币的运行程序。",
	},
	SubjectMatchErrorContract: &Translation{
		Template: "订单 %s 的组合 %s 处于状态 %s，但没有用于交换货币的运行程序。",
		Subject:  "合约组合错误",
	},
	SubjectMatchRecoveryError: &Translation{
		Template: "在检索订单 %s:%v 的交易期间审核交易对手交易合约 (%s %v) 时出错",
		Subject:  "检索匹配时出错",
	},
	SubjectOrderCoinError: &Translation{
		Template: "没有为活动订单 %s 注册资金货币",
		Subject:  "硬币订单错误",
	},
	SubjectOrderCoinFetchError: &Translation{
		Template: "检索订单 %s (%s) 的源硬币时出错：%v",
		Subject:  "硬币订单恢复错误",
	},
	SubjectMissedCancel: &Translation{
		Template: "取消订单与订单 %s 不匹配。如果取消订单与交易所同时发送，或者订单在取消订单执行之前已完全执行，则可能发生这种情况。",
		Subject:  "丢失取消",
	},
	SubjectOrderCanceled: &Translation{
		Template: "%s 上 %s-%s 上的 %s 请求已被取消 (%s)",
		Subject:  "订单取消",
	},
	SubjectMatchesMade: &Translation{
		Template: "%s 请求超过 %s-%s %.1f%% 已填充（%s）",
		Subject:  "匹配完成",
	},
	SubjectSwapSendError: &Translation{
		Template: "在以 %s 的顺序发送价值 %.8f %s 的输出的交换时遇到错误",
		Subject:  "发送交换时出错",
	},
	SubjectInitError: &Translation{
		Template: "错误通知 DEX %s 交换组合：%v",
		Subject:  "交换错误",
	},
	SubjectReportRedeemError: &Translation{
		Template: "通知 DEX %s 赎回时出错：%v",
		Subject:  "报销错误",
	},
	SubjectSwapsInitiated: &Translation{
		Template: "在订单 %s 上发送价值 %.8f %s 的交易",
		Subject:  "发起交流",
	},
	SubjectRedemptionError: &Translation{
		Template: "在订单 %s 上发现发送价值 %.8f %s 的赎回错误",
		Subject:  "赎回错误",
	},
	SubjectMatchComplete: &Translation{
		Template: "在订单 %s 上兑换了 %.8f %s",
		Subject:  "完全匹配",
	},
	SubjectRefundFailure: &Translation{
		Template: "按顺序 %s 返回 %.8f %s，有一些错误",
		Subject:  "退款错误",
	},
	SubjectMatchesRefunded: &Translation{
		Template: "在订单 %s 上返回了 %.8f %s",
		Subject:  "退款成功",
	},
	SubjectMatchRevoked: &Translation{
		Template: "组合 %s 已被撤销",
		Subject:  "撤销组合",
	},
	SubjectOrderRevoked: &Translation{
		Template: "%s 市场 %s 的订单 %s 已被服务器撤销",
		Subject:  "撤销订单",
	},
	SubjectOrderAutoRevoked: &Translation{
		Template: "%s 市场 %s 中的订单 %s 被市场暂停撤销",
		Subject:  "订单自动撤销",
	},
	SubjectMatchRecovered: &Translation{
		Template: "找到赎回 (%s: %v) 并验证了请求 %s 的秘密",
		Subject:  "恢复订单",
	},
	SubjectCancellingOrder: &Translation{
		Template: "已为订单 %s 提交取消订单",
		Subject:  "取消订单",
	},
	SubjectOrderStatusUpdate: &Translation{
		Template: "订单状态 %v 从 %v 修改为 %v",
		Subject:  "订单状态更新",
	},
	SubjectMatchResolutionError: &Translation{
		Template: "没有为 %s 找到为 %s 报告的 %d 个匹配项。",
		Subject:  "订单解析错误",
	},
	SubjectFailedCancel: &Translation{
		Template: "取消订单 %s 的订单 %s 处于 Epoque 状态 2 个 epoques，现在已被删除。",
		Subject:  "取消失败",
	},
	SubjectAuditTrouble: &Translation{
		Template: "继续寻找组合 %s 的货币 %v (%s) 的交易对手合约。您的互联网和钱包连接是否正常？",
		Subject:  "审计时的问题",
	},
	SubjectDexAuthError: &Translation{
		Template: "%s: %v",
		Subject:  "身份验证错误",
	},
	SubjectUnknownOrders: &Translation{
		Template: "未找到 DEX %s 报告的 %d 个活动订单。",
		Subject:  "DEX 报告的未知请求",
	},
	SubjectOrdersReconciled: &Translation{
		Template: "%d 个订单的更新状态。",
		Subject:  "与 DEX 协调的订单",
	},
	SubjectWalletConfigurationUpdated: &Translation{
		Template: "钱包 %s 的配置已更新。存款地址 = %s",
		Subject:  "更新的钱包设置a",
	},
	SubjectWalletPasswordUpdated: &Translation{
		Template: "钱包 %s 的密码已更新。",
		Subject:  "钱包密码更新",
	},
	SubjectMarketSuspendScheduled: &Translation{
		Template: "%s 上的市场 %s 现在计划在 %v 暂停",
		Subject:  "市场暂停预定",
	},
	SubjectMarketSuspended: &Translation{
		Template: "%s 的 %s 市场交易现已暂停。",
		Subject:  "暂停市场",
	},
	SubjectMarketSuspendedWithPurge: &Translation{
		Template: "%s 的市场交易 %s 现已暂停。订单簿中的所有订单现已被删除。",
		Subject:  "暂停市场，清除订单",
	},
	SubjectMarketResumeScheduled: &Translation{
		Template: "%s 上的市场 %s 现在计划在 %v 恢",
		Subject:  "预定市场摘要",
	},
	SubjectMarketResumed: &Translation{
		Template: "%s 上的市场 %s 已汇总用于时代 %d 中的交易",
		Subject:  "总结市场",
	},
	SubjectUpgradeNeeded: &Translation{
		Template: "您可能需要更新您的帐户以进行 %s 的交易。",
		Subject:  "需要更新",
	},
	SubjectDEXDisconnected: &Translation{
		Template: "%s 离线",
		Subject:  "服务器断开连接",
	},
	SubjectPenalized: &Translation{
		Template: "%s 上的 DEX 惩罚\n最后一条规则被破坏：%s \n时间：%v \n详细信息：\n \" %s \" \n",
		Subject:  "服务器惩罚了你",
	},
}

// EnLocale is the english translations. We can construct the EnLocale in an
// init function, since the subjects and templates are untranslated from the
// TemplateKeys. Other translators will define each entry in a struct literal.
var EnLocale = map[string]*Translation{}

var Subjects map[string]map[string]string

func registerTranslations(lang language.Tag, translations map[string]*Translation) {
	for subject, t := range translations {
		tmplKey := TemplateKeys[subject]
		message.SetString(lang, tmplKey, t.Template)
		message.SetString(lang, subject, t.Subject)
	}
}

func inintializeEnLocale() {
	for subject, t := range TemplateKeys {
		EnLocale[subject] = &Translation{
			Subject:  subject,
			Template: t,
		}
	}
	registerTranslations(language.AmericanEnglish, EnLocale)
}

func init() {
	inintializeEnLocale()
	registerTranslations(language.BrazilianPortuguese, BrPtLocales)
}
