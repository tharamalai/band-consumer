import useAxios from 'axios-hooks'

export const useMeichainBalance = (meiAddress) => {
  return useAxios(
    `http://localhost:8010/bank/balances/${meiAddress}`,
  )
}

export const useMeiCDP = (meiAddress) => {
  return useAxios(
    `http://localhost:8010/meicdp/cdp/${meiAddress}`,
  )
}