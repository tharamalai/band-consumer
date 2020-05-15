# MeiChain

![Image of Mei Cover](https://github.com/tharamalai/meichain/blob/mei-master/image/mei_cover.png)

MeiChain, built on Cosmos SDK, takes ATOM from Cosmos Hub over the IBC to collateral for the new stablecoin MEI, which its value is pegged to 1 US Dollar. Users can open their CDP (Collateralized Debt Position) and also liquidate another CDP. It uses ATOM price feed from Band Protocol.

## MeiChain Stablecoin System Overview
![Image of Mei System overview](https://github.com/tharamalai/meichain/blob/mei-master/image/mei_overview.png)

We have 3 blockchains working together. First is Cosmos Hub where its users hold their ATOMs. ATOM will be used as collateral for minting MEI. Second is Meichain which we use to store the CDP and the MEI token itself. And the last one is Bandchain which serves the ATOM/USD price feed from several exchanges to Meichian. All chains communicate to MeiChain by sending the packet across the chain using the Cosmos-IBC over  Relayers.

## MeiChain has 5 messages:
1. Lock Collateral
2. Borrow Debt
3. Return Debt
4. Unlock Collateral
5. Liquidate

## Lock Collateral Message
![Image of Lock Collateral](https://github.com/tharamalai/meichain/blob/mei-master/image/lock_collateral.png)
When a user sends a lock collateral message, MeiChain will move ATOM from your account to the CDP.

## Borrow Debt Message
![Image of Borrow Debt](https://github.com/tharamalai/meichain/blob/mei-master/image/borrow_debt.png)
When a user sends a borrow debt message. MeiChain will request ATOM price data from Band Protocol oracle over the relayer. After MeiChain receives the price data from BandChain, the CDP module will check the collateral ratio and if collateral ratio is more than 150% then CDP module will mint new MEI tokens and send to your account

## Return Debt Message
![Image of Return Debt](https://github.com/tharamalai/meichain/blob/mei-master/image/return_debt.png)
When a user sends a return debt message. MeiChain will move MEI from your account back to the CDP. This essentially free up you collateral and allows you to mint more MEI later or withdraw your ATOM collateral.

## Unlock Collateral Message
![Image of Unlock Collateral](https://github.com/tharamalai/meichain/blob/mei-master/image/unlock_collateral.png)
When a user sends an unlock collateral message. MeiChain will request ATOM price data from Band Protocol oracle over the IBC relayer. When the IBC packet is sent to MeiChain, the CDP will check the collateral ratio and if collateral ratio is more than 150% then CDP unlock ATOM tokens and send them to your account

## Liquidate Message
![Image of Liquidate](https://github.com/tharamalai/meichain/blob/mei-master/image/liquidate.png)
When a user sends a liquidate message. MeiChain will request ATOM price data from Band Protocol oracle. When an IBC packet is sent to MeiChain, the CDP will check the collateral ratio. Again, If this CDP has the collateral ratio less than 150% then you can liquidate this CDP. The liquidator has to send their MEI to CDP and CDP clears the debt on that CDP. After that CDP will unlock ATOM and send it to the liquidator account.

MeiChain Demo: https://www.youtube.com/watch?v=0g2o4JJbzgg. 
Twitter: https://twitter.com/mei_stablecoin
