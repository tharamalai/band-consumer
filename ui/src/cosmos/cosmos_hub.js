import useAxios from 'axios-hooks'
import { ATOM_UNIT_SYMBOL, MEICHAIN_GAIA_TRANSFER_CHANNEL, GAIA_MEICHAIN_TRANSFER_CHANNEL } from 'utils';

const cosmosjs = require("@cosmostation/cosmosjs");

const COSMOS_LCD_URL = "http://localhost:8011"
const  COSMOS_CHAIN_ID = "band-cosmoshub";

export let cosmos = null

export let ecpairPriv = null

export const initiateCosmosJs = () => {
  cosmos = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID);
  return cosmos
}

export const isInitiateCosmos = () => {
  if (!cosmos) {
    throw Error(`Error cosmos js is not initialized`)
  }
}

export const getCosmosAddress = (mnemonic) => {
  try {
    isInitiateCosmos()
    cosmos.setPath("m/44'/118'/0'/0/0");
    const address = cosmos.getAddress(mnemonic);
    ecpairPriv = cosmos.getECPairPriv(mnemonic);
    return address
  } catch (error) {
    console.log(error)
    throw Error(`Error cannot get cosmos account from mnemonic: ${error.message}`)
  }
}

export const sendTokenToMeichain = (cosmosAddress, amount, receiver) => {
  try {
    isInitiateCosmos()
    if (!ecpairPriv) {
      throw `Please connect wallet before request faucet`
    }
    cosmos.getAccounts(cosmosAddress).then(data => {
      console.log("data", data)
      let stdSignMsg = cosmos.newStdMsg({
        msgs: [
          {
            type: "ibc/transfer/MsgTransfer",
            value: {
              source_port: "transfer",
              source_channel: GAIA_MEICHAIN_TRANSFER_CHANNEL,
              dest_height: "10000000",
              amount: [
                {
                  amount: String(amount),
                  denom: `transfer/${MEICHAIN_GAIA_TRANSFER_CHANNEL}/${ATOM_UNIT_SYMBOL}`
                }
              ],
              sender: cosmosAddress,
              receiver: receiver
            }
          }
        ],
        chain_id: COSMOS_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      });
    
      const signedTx = cosmos.sign(stdSignMsg, ecpairPriv);
      cosmos.broadcast(signedTx).then(response => console.log(response));
    })
  } catch (error) {
    throw Error(`Error cannot send token to meichain: ${error.message}`)
  }


}