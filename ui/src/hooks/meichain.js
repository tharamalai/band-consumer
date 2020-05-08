import useAxios from 'axios-hooks'
import { getMeichainRestServer } from 'utils'

const MEICHAIN_REST_SERVER = getMeichainRestServer()

export const useMeichainBalance = (meiAddress) => {
  return useAxios(
    `${MEICHAIN_REST_SERVER}/bank/balances/${meiAddress}`,
  )
}

export const useMeiCDP = (meiAddress) => {
  return useAxios(
    `${MEICHAIN_REST_SERVER}/meicdp/cdp/${meiAddress}`,
  )
}