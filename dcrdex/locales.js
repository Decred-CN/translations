export const templateKeysPT = {
  [ID_NO_PASS_ERROR_MSG]: '密码不能为空',
  [ID_NO_APP_PASS_ERROR_MSG]: '应用密码不能为空',
  [ID_PASSWORD_NOT_MATCH]: '密码不相同',
  [ID_SET_BUTTON_BUY]: '来自{{ asset }}的买入订单',
  [ID_SET_BUTTON_SELL]: '来自{{ asset }}的卖出订单',
  [ID_OFF]: '关闭',
  [ID_READY]: '准备好',
  [ID_LOCKED]: '锁定',
  [ID_NOWALLET]: '没有钱包',
  [ID_WALLET_SYNC_PROGRESS]: '钱包同步进度{{ syncProgress }}%',
  [ID_HIDE_ADDIIONAL_SETTINGS]: '隐藏其它设置',
  [ID_SHOW_ADDIIONAL_SETTINGS]: '显示其它设置',
  [ID_BUY]: '买',
  [ID_SELL]: '卖',
  [ID_NOT_SUPPORTED]: '{{ asset }}不受支持',
  [ID_CONNECTION_FAILED]: '连接到服务器 dex 失败。您可以关闭 dexc 并稍后重试或等待尝试重新连接。',
  [ID_ORDER_PREVIEW]: '总计： {{ total }} {{ asset }}',
  [ID_CALCULATING]: '计算中...',
  [ID_ESTIMATE_UNAVAILABLE]: '估计不可用',
  [ID_NO_ZERO_RATE]: '汇率不能为零',
  [ID_NO_ZERO_QUANTITY]: '数量不能为零',
  [ID_TRADE]: '交易',
  [ID_NO_ASSET_WALLET]: '没有钱包 {{ asset }}',
  [ID_EXECUTED]: '执行',
  [ID_BOOKED]: '保留',
  [ID_CANCELING]: '取消',
  [ID_ACCT_UNDEFINED]: '帐户未定义。',
  [ID_KEEP_WALLET_PASS]: '保留钱包密码',
  [ID_NEW_WALLET_PASS]: '设置新的钱包密码'
}

const localesMap = {
  'en-us': templateKeys,
  'pt-br': templateKeysPT
}

export default class Locales {
  constructor (locale) {
    console.log((document.documentElement.lang).toLowerCase())
    // lang is set programatically by backend
    this.locale = (document.documentElement.lang).toLowerCase()
  }

  // formatDetails will format the message to its locale.
  // need to add one for plurals
  formatDetails (stringKey, args = undefined) {
    return this.stringTemplateParser(localesMap[this.locale][stringKey], args)
  }

  stringTemplateParser (expression, valueObj) {
    const templateMatcher = /{{\s?([^{}\s]*)\s?}}/g
    const text = expression.replace(templateMatcher, (substring, value, index) => {
      value = valueObj[value]
      return value
    })
    return text
  }
}
