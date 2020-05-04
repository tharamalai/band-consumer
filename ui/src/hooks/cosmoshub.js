import useAxios from 'axios-hooks'

export const useCosmosBalance = (cosmosAddress) => {
  return useAxios(
    `http://gaia-ibc-hackathon.node.bandchain.org:1317/bank/balances/${cosmosAddress}`,
  )
}
