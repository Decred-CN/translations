var TemplateKeys = map[string]string{
	// [host]
	SubjectAccountRegistered: "You may now trade at %s",
	// [confs, host]
	SubjectFeePaymentInProgress: "Waiting for %d confirmations before trading at %s",
	// [confs, required confs]
	SubjectRegUpdate: "Fee payment confirmations %v/%v",
	// [host, error]
	SubjectFeePaymentError: "Error encountered while paying fees to %s: %v",
	// [host, error]
	SubjectAccountUnlockError: "error unlocking account for %s: %v",
	// [host]
	SubjectFeeCoinError: "Empty fee coin for %s.",
	// [host]
	SubjectWalletConnectionWarning: "Incomplete registration detected for %s, but failed to connect to the Decred wallet",
	// [host, error]
	SubjectWalletUnlockError: "Connected to Decred wallet to complete registration at %s, but failed to unlock: %v",
	// [ticker, error]
	SubjectWithdrawError: "Error encountered during %s withdraw: %v",
	// [ticker, coin ID]
	SubjectWithdrawSend: "Withdraw of %s has completed successfully. Coin ID = %s",
	// [error]
	SubjectOrderLoadFailure: "Some orders failed to load from the database: %v",
	// [qty, ticker, token]
	SubjectYoloPlaced: "selling %.8f %s at market rate (%s)",
	// [sell string, qty, ticker, rate string, token]
	SubjectOrderPlaced: "%sing %.8f %s, rate = %s (%s)",
	// [missing count, token, host]
	SubjectMissingMatches: "%d matches for order %s were not reported by %q and are considered revoked",
	// [token, error]
	SubjectWalletMissing: "Wallet retrieval error for active order %s: %v",
	// [side, token, match status]
	SubjectMatchErrorCoin: "Match %s for order %s is in state %s, but has no maker swap coin.",
	// [side, token, match status]
	SubjectMatchErrorContract: "Match %s for order %s is in state %s, but has no maker swap contract.",
	// [ticker, contract, token, error]
	SubjectMatchRecoveryError: "Error auditing counter-party's swap contract (%s %v) during swap recovery on order %s: %v",
	// [token]
	SubjectOrderCoinError: "No funding coins recorded for active order %s",
	// [token, ticker, error]
	SubjectOrderCoinFetchError: "Source coins retrieval error for order %s (%s): %v",
	// [token, ticker, error]
	SubjectMissedCancel: "Cancel order did not match for order %s. This can happen if the cancel order is submitted in the same epoch as the trade or if the target order is fully executed before matching with the cancel order.",
	// [capitalized sell string, base ticker, quote ticker, host, token]
	SubjectOrderCanceled: "%s order on %s-%s at %s has been canceled (%s)",
	// [capitalized sell string, base ticker, quote ticker, fill percent, token]
	SubjectMatchesMade: "%s order on %s-%s %.1f%% filled (%s)",
	// [qty, ticker, token]
	SubjectSwapSendError: "Error encountered sending a swap output(s) worth %.8f %s on order %s",
	// [match, error]
	SubjectInitError: "Error notifying DEX of swap for match %s: %v",
	// [match, error]
	SubjectReportRedeemError: "Error notifying DEX of redemption for match %s: %v",
	// [qty, ticker, token]
	SubjectSwapsInitiated: "Sent swaps worth %.8f %s on order %s",
	// [qty, ticker, token]
	SubjectRedemptionError: "Error encountered sending redemptions worth %.8f %s on order %s",
	// [qty, ticker, token]
	SubjectMatchComplete: "Redeemed %.8f %s on order %s",
	// [qty, ticker, token]
	SubjectRefundFailure: "Refunded %.8f %s on order %s, with some errors",
	// [qty, ticker, token]
	SubjectMatchesRefunded: "Refunded %.8f %s on order %s",
	// [match ID token]
	SubjectMatchRevoked: "Match %s has been revoked",
	// [token, market name, host]
	SubjectOrderRevoked: "Order %s on market %s at %s has been revoked by the server",
	// [token, market name, host]
	SubjectOrderAutoRevoked: "Order %s on market %s at %s revoked due to market suspension",
	// [ticker, coin ID, match]
	SubjectMatchRecovered: "Found maker's redemption (%s: %v) and validated secret for match %s",
	// [token]
	SubjectCancellingOrder: "A cancel order has been submitted for order %s",
	// [token, old status, new status]
	SubjectOrderStatusUpdate: "Status of order %v revised from %v to %v",
	// [count, host, token]
	SubjectMatchResolutionError: "%d matches reported by %s were not found for %s.",
	// [token]
	SubjectFailedCancel: "Cancel order for order %s stuck in Epoch status for 2 epochs and is now deleted.",
	// [coin ID, ticker, match]
	SubjectAuditTrouble: "Still searching for counterparty's contract coin %v (%s) for match %s. Are your internet and wallet connections good?",
	// [host, error]
	SubjectDexAuthError: "%s: %v",
	// [count, host]
	SubjectUnknownOrders: "%d active orders reported by DEX %s were not found.",
	// [count]
	SubjectOrdersReconciled: "Statuses updated for %d orders.",
	// [ticker, address]
	SubjectWalletConfigurationUpdated: "Configuration for %s wallet has been updated. Deposit address = %s",
	//  [ticker]
	SubjectWalletPasswordUpdated: "Password for %s wallet has been updated.",
	// [market name, host, time]
	SubjectMarketSuspendScheduled: "Market %s at %s is now scheduled for suspension at %v",
	// [market name, host]
	SubjectMarketSuspended: "Trading for market %s at %s is now suspended.",
	// [market name, host]
	SubjectMarketSuspendedWithPurge: "Trading for market %s at %s is now suspended. All booked orders are now PURGED.",
	// [market name, host, time]
	SubjectMarketResumeScheduled: "Market %s at %s is now scheduled for resumption at %v",
	// [market name, host, epoch]
	SubjectMarketResumed: "Market %s at %s has resumed trading at epoch %d",
	// [host]
	SubjectUpgradeNeeded: "You may need to update your client to trade at %s.",
	// [host]
	SubjectDEXConnected: "%s is connected",
	// [host]
	SubjectDEXDisconnected: "%s is disconnected",
	// [host, rule, time, details]
	SubjectPenalized: "Penalty from DEX at %s\nlast broken rule: %s\ntime: %v\ndetails:\n\"%s\"\n",
}

// traduzir os subject tudo
var BrPtLocales = map[string]*Translation{
	SubjectDEXConnected: &Translation{
		Subject:  "DEX conectado",
		Template: "%s está conectado",
	},
	SubjectAccountRegistered: &Translation{
		Subject:  "Conta Registrada",
		Template: "Você agora pode trocar em %s",
	},
	SubjectFeePaymentInProgress: &Translation{
		Subject:  "Pagamento da Taxa em andamento",
		Template: "Esperando por %d confirmações antes de trocar em %s",
	},
	SubjectRegUpdate: &Translation{
		Subject:  "Atualização de registro",
		Template: "Confirmações da taxa %v/%v",
	},
	SubjectFeePaymentError: &Translation{
		Subject:  "Erro no Pagamento da Taxa",
		Template: "Erro enquanto pagando taxa para %s: %v",
	},
	SubjectAccountUnlockError: &Translation{
		Subject:  "Erro ao Destrancar carteira",
		Template: "erro destrancando conta %s: %v",
	},
	SubjectFeeCoinError: &Translation{
		Subject:  "Erro na Taxa",
		Template: "Taxa vazia para %s.",
	},
	SubjectWalletConnectionWarning: &Translation{
		Subject:  "Aviso de Conexão com a Carteira",
		Template: "Registro incompleto detectado para %s, mas falhou ao conectar com carteira decred",
	},
	SubjectWalletUnlockError: &Translation{
		Subject:  "Erro ao Destravar Carteira",
		Template: "Conectado com carteira decred para completar o registro em %s, mas falha ao destrancar: %v",
	},
	SubjectWithdrawError: &Translation{
		Subject:  "Erro Retirada",
		Template: "Erro encontrado durante retirada de %s: %v",
	},
	SubjectWithdrawSend: &Translation{
		Template: "Retirada de %s foi completada com sucesso. ID da moeda = %s",
		Subject:  "Retirada Enviada",
	},
	SubjectOrderLoadFailure: &Translation{
		Template: "Alguns pedidos falharam ao carregar da base de dados: %v",
		Subject:  "Carregamendo de Pedidos Falhou",
	},
	SubjectYoloPlaced: &Translation{
		Template: "vendendo %.8f %s a taxa de mercado (%s)",
		Subject:  "Ordem de Mercado Colocada",
	},
	SubjectOrderPlaced: &Translation{
		Subject:  "Ordem Colocada",
		Template: "%sing %.8f %s, valor = %s (%s)",
	},
	SubjectMissingMatches: &Translation{
		Template: "%d combinações para pedidos %s não foram reportados por %q e foram considerados revocados",
		Subject:  "Pedidos Faltando Combinações",
	},
	SubjectWalletMissing: &Translation{
		Template: "Erro ao recuperar pedidos ativos por carteira %s: %v",
		Subject:  "Carteira Faltando",
	},
	SubjectMatchErrorCoin: &Translation{
		Subject:  "Erro combinação de Moedas",
		Template: "Combinação %s para pedido %s está no estado %s, mas não há um executador para trocar moedas.",
	},
	SubjectMatchErrorContract: &Translation{
		Template: "Combinação %s para pedido %s está no estado %s, mas não há um executador para trocar moedas.",
		Subject:  "Erro na Combinação de Contrato",
	},
	SubjectMatchRecoveryError: &Translation{
		Template: "Erro auditando contrato de troca da contraparte (%s %v) durante troca recuperado no pedido %s: %v",
		Subject:  "Erro Recuperando Combinações",
	},
	SubjectOrderCoinError: &Translation{
		Template: "Não há Moedas de financiamento registradas para pedidos ativos %s",
		Subject:  "Erro no Pedido da Moeda",
	},
	SubjectOrderCoinFetchError: &Translation{
		Template: "Erro ao recuperar moedas de origem para pedido %s (%s): %v",
		Subject:  "Erro na Recuperação do Pedido de Moedas",
	},
	SubjectMissedCancel: &Translation{
		Template: "Pedido de cancelamento não combinou para pedido %s. Isto pode acontecer se o pedido de cancelamento foi enviado no mesmo epoque do que a troca ou se o pedido foi completamente executado antes da ordem de cancelamento ser executada.",
		Subject:  "Cancelamento Perdido",
	},
	SubjectOrderCanceled: &Translation{
		Template: "%s pedido sobre %s-%s em %s foi cancelado (%s)",
		Subject:  "Cancelamento de Pedido",
	},
	SubjectMatchesMade: &Translation{
		Template: "%s pedido sobre %s-%s %.1f%% preenchido (%s)",
		Subject:  "Combinações Feitas",
	},
	SubjectSwapSendError: &Translation{
		Template: "Erro encontrado ao enviar a troca com output(s) no valor de %.8f %s no pedido %s",
		Subject:  "Erro ao Enviar Troca",
	},
	SubjectInitError: &Translation{
		Template: "Erro notificando DEX da troca %s por combinação: %v",
		Subject:  "Erro na Troca",
	},
	SubjectReportRedeemError: &Translation{
		Template: "Erro notificando DEX da redenção %s por combinação: %v",
		Subject:  "Reportando Erro na redenção",
	},
	SubjectSwapsInitiated: &Translation{
		Template: "Enviar trocas no valor de %.8f %s no pedido %s",
		Subject:  "Trocas Iniciadas",
	},
	SubjectRedemptionError: &Translation{
		Template: "Erro encontrado enviado redenção no valor de %.8f %s no pedido %s",
		Subject:  "Erro na Redenção",
	},
	SubjectMatchComplete: &Translation{
		Template: "Resgatado %.8f %s no pedido %s",
		Subject:  "Combinação Completa",
	},
	SubjectRefundFailure: &Translation{
		Template: "Devolvidos %.8f %s no pedido %s, com algum erro",
		Subject:  "Erro no Reembolso",
	},
	SubjectMatchesRefunded: &Translation{
		Template: "Devolvidos %.8f %s no pedido %s",
		Subject:  "Reembolso Sucedido",
	},
	SubjectMatchRevoked: &Translation{
		Template: "Combinação %s foi revocada",
		Subject:  "Combinação Revocada",
	},
	SubjectOrderRevoked: &Translation{
		Template: "Pedido %s no mercado %s em %s foi revocado pelo servidor",
		Subject:  "Pedido Revocado",
	},
	SubjectOrderAutoRevoked: &Translation{
		Template: "Pedido %s no mercado %s em %s revocado por suspenção do mercado",
		Subject:  "Pedido Revocado Automatiamente",
	},
	SubjectMatchRecovered: &Translation{
		Template: "Encontrado redenção do executador (%s: %v) e validado segredo para pedido %s",
		Subject:  "Pedido Recuperado",
	},
	SubjectCancellingOrder: &Translation{
		Template: "Uma ordem de cancelamento foi submetida para o pedido %s",
		Subject:  "Cancelando Pedido",
	},
	SubjectOrderStatusUpdate: &Translation{
		Template: "Status do pedido %v revisado de %v para %v",
		Subject:  "Status do Pedido Atualizado",
	},
	SubjectMatchResolutionError: &Translation{
		Template: "%d combinações reportada para %s não foram encontradas para %s.",
		Subject:  "Erro na Resolução do Pedido",
	},
	SubjectFailedCancel: &Translation{
		Template: "Ordem de cancelamento para pedido %s presa em estado de Epoque por 2 epoques e foi agora deletado.",
		Subject:  "Falhou Cancelamento",
	},
	SubjectAuditTrouble: &Translation{
		Template: "Continua procurando por contrato de contrapartes para moeda %v (%s) para combinação %s. Sua internet e conexão com a carteira estão ok?",
		Subject:  "Problemas ao Auditar",
	},
	SubjectDexAuthError: &Translation{
		Template: "%s: %v",
		Subject:  "Erro na Autenticação",
	},
	SubjectUnknownOrders: &Translation{
		Template: "%d pedidos ativos reportados pela DEX %s não foram encontrados.",
		Subject:  "DEX Reportou Pedidos Desconhecidos",
	},
	SubjectOrdersReconciled: &Translation{
		Template: "Estados atualizados para %d pedidos.",
		Subject:  "Pedidos Reconciliados com DEX",
	},
	SubjectWalletConfigurationUpdated: &Translation{
		Template: "configuração para carteira %s foi atualizada. Endereço de depósito = %s",
		Subject:  "Configurações da Carteira Atualizada",
	},
	SubjectWalletPasswordUpdated: &Translation{
		Template: "Senha para carteira %s foi atualizada.",
		Subject:  "Senha da Carteira Atualizada",
	},
	SubjectMarketSuspendScheduled: &Translation{
		Template: "Mercado %s em %s está agora agendado para suspensão em %v",
		Subject:  "Suspensão de Mercado Agendada",
	},
	SubjectMarketSuspended: &Translation{
		Template: "Trocas no mercado %s em %s está agora suspenso.",
		Subject:  "Mercado Suspenso",
	},
	SubjectMarketSuspendedWithPurge: &Translation{
		Template: "Trocas no mercado %s em %s está agora suspenso. Todos pedidos no livro de ofertas foram agora EXPURGADOS.",
		Subject:  "Mercado Suspenso, Pedidos Expurgados",
	},
	SubjectMarketResumeScheduled: &Translation{
		Template: "Mercado %s em %s está agora agendado para resumir em %v",
		Subject:  "Resumo do Mercado Agendado",
	},
	SubjectMarketResumed: &Translation{
		Template: "Mercado %s em %s foi resumido para trocas no epoque %d",
		Subject:  "Mercado Resumido",
	},
	SubjectUpgradeNeeded: &Translation{
		Template: "Você pode precisar atualizar seu cliente para trocas em %s.",
		Subject:  "Atualização Necessária",
	},
	SubjectDEXDisconnected: &Translation{
		Template: "%s está desconectado",
		Subject:  "Server Disconectado",
	},
	SubjectPenalized: &Translation{
		Template: "Penalidade de DEX em %s\núltima regra quebrada: %s\nhorário: %v\ndetalhes:\n\"%s\"\n",
		Subject:  "Server Penalizou Você",
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
