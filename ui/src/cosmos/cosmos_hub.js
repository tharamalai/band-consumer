const cosmosjs = require("@cosmostation/cosmosjs");

const COSMOS_LCD_URL = "http://gaia-ibc-hackathon.node.bandchain.org"
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
    throw `Error cannot get cosmos account from mnemonic: ${error.message}`
  }
}

export const requestCosmosHubFaucet = () => {
  if (!ecpairPriv) {
    throw `Please connect wallet before request faucet`
  }


}