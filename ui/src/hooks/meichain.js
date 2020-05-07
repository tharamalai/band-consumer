import useAxios from 'axios-hooks'

export const useMeichainBalance = (meiAddress) => {
  return useAxios(
    `http://13.250.187.211:1317/bank/balances/${meiAddress}`,
  )
}

export const useMeiCDP = (meiAddress) => {
  return useAxios(
    `http://13.250.187.211:1317/meicdp/cdp/${meiAddress}`,
  )
}