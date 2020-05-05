const cosmosjs = require("@cosmostation/cosmosjs");

const COSMOS_LCD_URL = "http://gaia-ibc-hackathon.node.bandchain.org"
const  COSMOS_CHAIN_ID = "band-cosmoshub";

// export const cosmosJs = null

export const initiateCosmosJs = () => {
  const cosmosJs = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID);
  return cosmosJs
}

export const getCosmosAddress = (mnemonic) => {
  try {
    const cosmos = initiateCosmosJs()
    const address = cosmos.getAddress(mnemonic);
    return address
  } catch (error) {
    throw `Error cannot get cosmos account from mnemonic: ${error.message}`
  }
}