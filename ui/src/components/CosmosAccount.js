import React, { useEffect } from 'react'
import { Flex, Image, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import Button from 'components/Button'
import FaucetBtn from 'components/FaucetBtn'
import { usePrice } from 'hooks/price'
import { useCosmosBalance } from 'hooks/cosmoshub'

import ConnectCosmos from 'images/connect-cosmos.svg'

const Card = styled(Flex).attrs(() => ({
  p: '1.8vw',
  width: '100%',
  mt: '2.5vw',
  height: '12.639vw',
}))`
  background: rgba(249, 251, 252, 0.9);
  box-shadow: 0px 16px 32px rgba(95, 106, 128, 0.1);
  border-radius: 0.56vw;
  position: relative;
`

export default ({ cosmosAddress, setCosmosAddress }) => {
  // cosmos1mypshseyp9z30xscvg3kfcjg6v2r0rhz6plgl6
  const [{ data: cosmosBalanceData, loading: cosmosBalanceLoading, error: cosmosBalanceError }, cosmosAccountBalanceRefetch] = useCosmosBalance("cosmos1mypshseyp9z30xscvg3kfcjg6v2r0rhz6plgl6")

  const [{ data: priceData, loading: priceLoading, error: priceError }, refetch] = usePrice()

  useEffect(() => {
    console.log("new address", cosmosAddress)
  });

  return (
    <Card>
      {cosmosAddress ? (
        <Flex flexDirection="column" width="100%">
          <Text
            fontSize="0.83vw"
            fontWeight={500}
            lineHeight="1vw"
            color={colors.purple.normal}
          >
            CosmosHub Account
          </Text>
          <Text
            fontSize="0.83vw"
            fontWeight={800}
            lineHeight="1vw"
            color={colors.purple.dark}
            mt="0.347vw"
          >
            {cosmosAddress}
          </Text>
          <Flex flexDirection="row" justifyContent="space-between" mt="2.153vw">
            <Flex flexDirection="column">
              <Text
                fontSize="0.83vw"
                fontWeight={500}
                lineHeight="1vw"
                color={colors.purple.normal}
              >
                ATOM
              </Text>
              <Text
                fontSize="1.867vw"
                fontWeight={400}
                lineHeight="2.114vw"
                color={colors.purple.dark}
              >
                {cosmosBalanceLoading
                  ? 'loading...'
                  : cosmosBalanceError
                  ? cosmosBalanceError
                  : cosmosBalanceData.result[0].amount
                }
              </Text>
              <Text
                mt="0.4vw"
                fontSize="0.83vw"
                fontWeight={400}
                lineHeight="1vw"
                color={colors.purple.normal}
                style={{ fontStyle: 'italic' }}
              >
                {priceLoading
                  ? 'loading...'
                  : priceError
                  ? priceError
                  : `≈ ${priceData.cosmos.usd} USD`}
              </Text>
            </Flex>
            <Flex flexDirection="column" alignItems="flex-end">
              <FaucetBtn onClick={() => alert('faucet')} />
              <Button
                mt="0.833vw"
                py="0.55vw"
                px="1vw"
                onClick={() => alert('Send atom to meichain')}
              >
                <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
                  Send ATOM to MeiChain
                </Text>
              </Button>
            </Flex>
          </Flex>
        </Flex>
      ) : (
        <Flex flexDirection="column" alignItems="center" width="100%">
          <Image src={ConnectCosmos} width="23vw" />
          <Button
            mt="2.22vw"
            py="0.55vw"
            px="1vw"
            onClick={() => {
              const address = window.prompt('Input Cosmos Address')
              console.log("address", address)
              setCosmosAddress(address)
            }}
          >
            <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
              Connect To CosmosHub
            </Text>
          </Button>
        </Flex>
      )}
    </Card>
  )
}
