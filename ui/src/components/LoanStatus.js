import React from 'react'
import { Flex, Text } from 'rebass'
import colors from 'ui/colors'
import { useMeiBalance } from 'hooks/meichain'
import { findMeiAmount, toMei } from 'utils'

export default ({ meiAddress }) => {
  const [{ data: meiBalanceData, loading: meiBalanceLoading, error: meiBalanceError }, meiAccountBalanceRefetch] = useMeiBalance("cosmos180plwgqxyx55vvx0eucrg5lz3q6nf06e3s27jz")


  return (
    <Flex flexDirection="column" width="100%" p="1.8vw">
      <Text
        color={colors.purple.dark}
        fontSize="1.25vw"
        fontWeight={900}
        lineHeight="1.53vw"
      >
        Loan Status
      </Text>
      <Flex flexDirection="row" justifyContent="space-between" mt="0.8vw">
        <Text
          fontSize="0.83vw"
          fontWeight={500}
          lineHeight="1vw"
          color={colors.purple.normal}
        >
          MeiChain Account
        </Text>
        <Text
          fontSize="0.83vw"
          fontWeight={800}
          lineHeight="1vw"
          color={colors.purple.dark}
        >
          {meiAddress}
        </Text>
      </Flex>
      <Flex flexDirection="row" justifyContent="space-between" mt="1vw">
        <Text
          fontSize="0.83vw"
          fontWeight={500}
          lineHeight="1vw"
          color={colors.purple.normal}
        >
          Current balance
        </Text>
        <Flex flexDirection="column">
          <Flex flexDirection="row" alignItems="flex-end">
            <Text
              fontSize="1.867vw"
              fontWeight={400}
              lineHeight="2.114vw"
              color={colors.purple.dark}
            >
              {meiBalanceLoading
                ? 'loading...'
                : meiBalanceError
                ? meiBalanceError
                : toMei(findMeiAmount(meiBalanceData).amount)
              }
            </Text>
            <Text
              fontSize="0.83vw"
              fontWeight={500}
              lineHeight="1vw"
              ml="1vw"
              mb="0.3vw"
              color={colors.purple.normal}
            >
              MEI
            </Text>
          </Flex>
          <Text
            mt="0.4vw"
            fontSize="0.83vw"
            fontWeight={400}
            lineHeight="1vw"
            color={colors.purple.normal}
            style={{ fontStyle: 'italic' }}
          >
            {meiBalanceLoading
              ? 'loading...'
              : meiBalanceError
              ? meiBalanceError
              :`≈ ${toMei(findMeiAmount(meiBalanceData).amount)} USD`}
          </Text>
        </Flex>
      </Flex>
    </Flex>
  )
}
