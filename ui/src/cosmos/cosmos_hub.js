const cosmosjs = require("@cosmostation/cosmosjs");

const COSMOS_LCD_URL = "http://gaia-ibc-hackathon.node.bandchain.org"
const  COSMOS_CHAIN_ID = "band-cosmoshub";

// export const cosmosJs = null

export let ecpairPriv = null

export const initiateCosmosJs = () => {
  const cosmosJs = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID);
  return cosmosJs
}

export const getCosmosAddress = (mnemonic) => {
  try {
    const cosmos = initiateCosmosJs()
    cosmos.setPath("m/44'/118'/0'/0/0");
    const address = cosmos.getAddress(mnemonic);
    ecpairPriv = cosmos.getECPairPriv(mnemonic);
    return address
  } catch (error) {
    console.log(error)
    throw `Error cannot get cosmos account from mnemonic: ${error.message}`
  }
}