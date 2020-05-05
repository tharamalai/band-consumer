const cosmosjs = require("@cosmostation/cosmosjs");

const COSMOS_LCD_URL = "http://gaia-ibc-hackathon.node.bandchain.org"
const  COSMOS_CHAIN_ID = "band-cosmoshub";

// export const cosmosJs = null

export const initiateCosmosJs = () => {
  const cosmosJs = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID);
  return cosmosJs
}