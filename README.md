# MeiChain

![Image of Mei Cover](https://github.com/tharamalai/meichain/blob/mei-master/image/mei_cover.png)

Meichain, built on Cosmos SDK, takes ATOM from Cosmos Hub over the IBC to collateral for the new stablecoin MEI, which its value is pegged to 1 US Dollar. Users can open their CDP (Collateralized Debt Position) and also liquidate another CDP. It uses ATOM price feed from Band Protocol.

## MeiChain Stablecoin System Overview
![Image of Mei System overview](https://github.com/tharamalai/meichain/blob/mei-master/image/mei_overview.png)

We have 3 blockchains working together. First is Cosmos Hub where its users hold their ATOMs. ATOM will be used as collateral for minting MEI. Second is Meichain which we use to store the CDP and the MEI token itself. And the last one is Bandchain which serves the ATOM/USD price feed from several exchanges to Meichian. All chains communicate to Meichain by sending the packet across the chain using the Cosmos-IBC over  Relayers.

