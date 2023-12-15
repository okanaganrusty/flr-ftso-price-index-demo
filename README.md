# Flare Networks

## Flare Time Series Oracle: Price Index

Start at
[https://docs.flare.network/dev/getting-started/contract-addresses/](https://docs.flare.network/dev/getting-started/contract-addresses/)
to get the Flare Contract Registry address. Then from there, you can retrieve any contract that you're wanting to work with.

I have also provided the `contracts/` directory containing the current contracts as of December 9th, 2023.

To run this program run `go build` and then execute the `flr-ftso-price-index` executable. This'll report the current prices of the
registered FTSO assets known by the Flare network.

### Testing

Example output:

```bash
Network ID: 14
Current Block number: 16737626
Coin: XRP, price: 0.621670
Coin: LTC, price: 71.729110
Coin: XLM, price: 0.125510
Coin: DOGE, price: 0.094700
Coin: ADA, price: 0.621250
Coin: ALGO, price: 0.200800
Coin: BTC, price: 42017.251930
Coin: ETH, price: 2244.441060
Coin: FIL, price: 4.873670
Coin: FLR, price: 0.016780
Coin: ARB, price: 1.128750
Coin: AVAX, price: 40.020020
Coin: BNB, price: 247.041120
Coin: MATIC, price: 0.860220
Coin: SOL, price: 74.815920
Coin: USDC, price: 1.000120
Coin: USDT, price: 0.999610
Coin: XDC, price: 0.049580
```
