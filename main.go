package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/go-chi/chi"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/tarancss/ethcli"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", Handler)
	fmt.Println("Running on :8080")
	http.ListenAndServe(":8080", r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OKJ"))

	balanceOutputs, vesperPrice, totalValue := Values()
}

type TokenPair struct {
	VTokenAddress string
	TokenAddress  string
}
type Values struct {
}

func GetValues() (*Values, error) {

	tokenPairs := map[string]*TokenPair{
		"WETH": {"0x103cc17c2b1586e5cd9bad308690bcd0bbe54d5e", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"},
		"WBTC": {"0x4b2e76ebbc9f2923d83f5fbde695d8733db1a117b", "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"},
		"USDC": {"0x0c49066c0808ee8c673553b7cbd99bcc9abf113d", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
		"DAI":  {"0xca0c34a3f35520b9490c1d58b35a19ab64014d80", "0x6b175474e89094c44da98b954eedeac495271d0f"},
		"LINK": {"0x0a27e910aee974d05000e05eab8a4b8ebd93d40c", "0x514910771af9ca656af840dff83e8264ecf986ca"},
	}
	cg := gecko.NewClient(http.DefaultClient)

	paymentSplitterBuybackAddress := "0x223809E09ec28C28219769C3FF05c790c213152C"
	vVSPBuybackAddress := "0xbA4cFE5741b357FA371b506e5db0774aBFeCf8Fc"
	multisigBuybackAddress := "0x9520b477Aa81180E6DdC006Fc09Fb6d3eb4e807A"

	// Connect to an infura node in order to interact with balancer
	conn := ethcli.Init("", "")
	balanceOutputs := []float64{}
	// Get balances of all vTokens from PaymentSplitter, Multisig and VSP + Calculate prices.
	for symbol, pair := range tokenPairs {

		totalBalance := float64(0)
		// Payment splitter balance
		_, paymentSplitterBalance, err := conn.GetBalance(paymentSplitterBuybackAddress, pair.VTokenAddress)
		if err != nil {
			return nil, err
		}
		totalBalance += math.Round(float64(paymentSplitterBalance.Int64()) / float64(1000000000000000000))

		// vVSP balance
		_, vVSPBalance, err := conn.GetBalance(vVSPBuybackAddress, pair.VTokenAddress)
		if err != nil {
			return nil, err
		}
		totalBalance += math.Round(float64(vVSPBalance.Int64()) / float64(1000000000000000000))
		// MultiSig Balance
		_, multisigBalance, err := conn.GetBalance(multisigBuybackAddress, pair.VTokenAddress)
		if err != nil {
			return nil, err
		}
		totalBalance += math.Round(float64(multisigBalance.Int64()) / float64(1000000000000000000))
		balanceOutputs = append(balanceOutputs, totalBalance)

		// Calculate total value
		cg.SimpleSinglePrice()
		ethPrice, err := cg.SimpleSinglePrice(token, "usd")

	}
	// Get Geyser balance from vesper"s geyser contract + Append it to balances array.
	cg.SimplePrice()
	// Get Geyser price and add it to total value
	cg.SimplePrice()
	// Get VSP price
	cg.SimplePrice()
}
