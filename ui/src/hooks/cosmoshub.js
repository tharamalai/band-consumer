import useAxios from 'axios-hooks'

export const useCosmosBalance = (cosmosAddress) => {
  return useAxios(
    `http://gaia-ibc-hackathon.node.bandchain.org:1317/bank/balances/${cosmosAddress}`, 
  )
}

export const useCosmosHubFaucet = () => {
  return useAxios({
      url: `http://gaia-ibc-hackathon.node.bandchain.org:8000/`,
      method: 'POST'
    },
    { manual: true }
  )
}
