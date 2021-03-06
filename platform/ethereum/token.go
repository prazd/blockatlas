package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, *p), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Token, coinIndex uint) (t blockatlas.Token, ok bool) {
	t = blockatlas.Token{
		Name:     srcToken.Contract.Name,
		Symbol:   srcToken.Contract.Symbol,
		TokenID:  srcToken.Contract.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Contract.Decimals,
		Type:     blockatlas.TokenTypeERC20,
	}

	return t, true
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Token, p Platform) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token, ok := NormalizeToken(&srcToken, p.CoinIndex)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}
