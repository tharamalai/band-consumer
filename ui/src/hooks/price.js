import useAxios from 'axios-hooks'

export const usePrice = () => {
  return useAxios(
    'https://api.coingecko.com/api/v3/simple/price?ids=cosmos&vs_currencies=usd',
  )
}
