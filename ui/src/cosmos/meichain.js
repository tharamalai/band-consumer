
const cosmosjs = require("@cosmostation/cosmosjs");

export const COSMOS_LCD_URL = "http://localhost:8010"
export const  COSMOS_CHAIN_ID = "meichain";

export let meichain = null
export let ecpairPriv = null

export const initiateMeichain = () => {
  meichain = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID);
  return meichain
}

export const isInitiateMeichain = () => {
  if (!meichain) {
    throw Error(`Error meichain js is not initialized`)
  }
}

export const getMeichainAddress = (mnemonic) => {
  try {
    isInitiateMeichain()
    meichain.setPath("m/44'/118'/0'/0/0");
    const address = meichain.getAddress(mnemonic);
    ecpairPriv = meichain.getECPairPriv(mnemonic);
    return address
  } catch (error) {
    console.log(error)
    throw Error(`Error cannot get meichain account from mnemonic: ${error.message}`)
  }
}
