export const templateKeysPT = {
  [ID_NO_PASS_ERROR_MSG]: 'senha não pode ser vazia',
  [ID_NO_APP_PASS_ERROR_MSG]: 'senha do app não pode ser vazia',
  [ID_PASSWORD_NOT_MATCH]: 'senhas diferentes',
  [ID_SET_BUTTON_BUY]: 'Ordem de compra de {{ asset }}',
  [ID_SET_BUTTON_SELL]: 'Ordem de venda de {{ asset }}',
  [ID_OFF]: 'desligar',
  [ID_READY]: 'pronto',
  [ID_LOCKED]: 'trancado',
  [ID_NOWALLET]: 'sem carteira',
  [ID_WALLET_SYNC_PROGRESS]: 'carteira está {{ syncProgress }}% sincronizada',
  [ID_HIDE_ADDIIONAL_SETTINGS]: 'esconder configurações adicionais',
  [ID_SHOW_ADDIIONAL_SETTINGS]: 'mostrar configurações adicionais',
  [ID_BUY]: 'Comprar',
  [ID_SELL]: 'Vender',
  [ID_NOT_SUPPORTED]: '{{ asset }} não tem suporte',
  [ID_CONNECTION_FAILED]: 'Conexão ao server dex falhou. Pode fechar dexc e tentar novamente depois ou esperar para tentar se reconectar.',
  [ID_ORDER_PREVIEW]: 'Total: {{ total }} {{ asset }}',
  [ID_CALCULATING]: 'calculando...',
  [ID_ESTIMATE_UNAVAILABLE]: 'estimativa indisponível',
  [ID_NO_ZERO_RATE]: 'taxa não pode ser zero',
  [ID_NO_ZERO_QUANTITY]: 'quantidade não pode ser zero',
  [ID_TRADE]: 'trade',
  [ID_NO_ASSET_WALLET]: 'Sem carteira {{ asset }}',
  [ID_EXECUTED]: 'executado',
  [ID_BOOKED]: 'reservado',
  [ID_CANCELING]: 'cancelando',
  [ID_ACCT_UNDEFINED]: 'conta não definida.',
  [ID_KEEP_WALLET_PASS]: 'manter senha da carteira',
  [ID_NEW_WALLET_PASS]: 'definir nova senha para carteira'
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
