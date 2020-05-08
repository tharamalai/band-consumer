import React, { useEffect, useState, createContext } from 'react'
import { Flex, Image, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import Button from 'components/Button'
import FaucetBtn from 'components/FaucetBtn'
import { usePrice } from 'hooks/price'
import { useCosmosBalance, useCosmosHubFaucet } from 'hooks/cosmoshub'
import { toAtom, toAtomUnit, convertAtomToUsd, findTokenBySymbol, ATOM_UNIT_SYMBOL, COSMOS_CHAIN_ID } from 'utils'
import refresh from 'images/refresh.svg'
import { useCosmosHubContextState } from 'contexts/CosmosHubContext'
import Big from 'big.js'

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

const LogIn = ({ cosmosAddress }) => {
  const [{ data: cosmosBalanceData, loading: cosmosBalanceLoading, error: cosmosBalanceError }, cosmosAccountBalanceRefetch] = useCosmosBalance(cosmosAddress)
  const [{ data: priceData, loading: priceLoading, error: priceError }, priceRefetch] = usePrice()
  const [{ data: faucetData, loading: faucetLoading, error: faucetError }, requestFaucet] = useCosmosHubFaucet()
  const { sendTokenToMeichain } = useCosmosHubContextState()

  return (
    <Flex flexDirection="column" width="100%">
      <Image src={refresh} width="1vw" style={{position: "absolute", top: "1.5vw", right: "1.5vw", cursor: "pointer"}}
        onClick={() => {
          cosmosAccountBalanceRefetch()
        }}
      />
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
            {cosmosBalanceData
              ? toAtom(findTokenBySymbol(cosmosBalanceData.result, ATOM_UNIT_SYMBOL).amount)
              : 'loading...'}
          </Text>
          <Text
            mt="0.4vw"
            fontSize="0.83vw"
            fontWeight={400}
            lineHeight="1vw"
            color={colors.purple.normal}
            style={{ fontStyle: 'italic' }}
          >
            {cosmosBalanceData && priceData
              ? `≈ ${convertAtomToUsd(toAtom(findTokenBySymbol(cosmosBalanceData.result, ATOM_UNIT_SYMBOL).amount), priceData.cosmos.usd)} USD`
              : 'loading...'}
          </Text>
        </Flex>
        <Flex flexDirection="column" alignItems="flex-end">
          <FaucetBtn onClick={() => {
            requestFaucet({
              data: {
                "address": cosmosAddress,
                "chain-id": COSMOS_CHAIN_ID
              }
            })
          }} />
          <Button
            mt="0.833vw"
            py="0.55vw"
            px="1vw"
            onClick={() => {
              const amount = window.prompt('Input transfer amount')
              if (!amount) {
                return
              }

              const transferAmount = Big(toAtomUnit(amount))
              const atomBalance = Big(findTokenBySymbol(cosmosBalanceData.result, ATOM_UNIT_SYMBOL).amount)
              if (transferAmount.gt(atomBalance)) {
                alert(`Max transfer amount is ${toAtom(atomBalance)}`)
                return
              }

              const recipient = window.prompt('Input Meichain recipient address')
              if (!recipient) {
                return
              }
              
              sendTokenToMeichain(toAtomUnit(amount), recipient)
            }}
          >
            <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
              Send ATOM to MeiChain
            </Text>
          </Button>
        </Flex>
      </Flex>
    </Flex>
  )
}

export default ({ cosmosAddress, setCosmosAddress }) => {
  const { getCosmosAddress, setPrivateKeyFromMnemonic } = useCosmosHubContextState()
  
  return (
    <Card>
      {cosmosAddress ? (
        <LogIn cosmosAddress={cosmosAddress}/>
      ) : (
        <Flex flexDirection="column" alignItems="center" width="100%">
          <Image src={ConnectCosmos} width="23vw" />
          <Button
            mt="2.22vw"
            py="0.55vw"
            px="1vw"
            onClick={() => {
              const mnemonic = window.prompt('Input Cosmos address mnemonic')
              if (mnemonic) {
                try {
                  const cosmosAddress = getCosmosAddress(mnemonic)
                  setPrivateKeyFromMnemonic(mnemonic)
                  setCosmosAddress(cosmosAddress)
                } catch (error) {
                  alert("Invalid mnemonic. Cannot get account from mnemonic.")
                }
              }
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
