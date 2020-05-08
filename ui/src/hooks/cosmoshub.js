import useAxios from 'axios-hooks'
import { getCosmosLcdUrl, getCosmosRestServer } from 'utils'

export const useCosmosBalance = (cosmosAddress) => {
  return useAxios(
    `${getCosmosRestServer()}bank/balances/${cosmosAddress}`, 
  )
}

export const useCosmosHubFaucet = () => {
  return useAxios({
      url: getCosmosLcdUrl(),
      method: 'POST'
    },
    { manual: true }
  )
}
