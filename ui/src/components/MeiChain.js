import React from 'react'
import { Flex, Image, Text } from 'rebass'
import styled from 'styled-components'
import Button from 'components/Button'
import LoanStatus from 'components/LoanStatus'
import DebtMenu from 'components/DebtMenu'
import LockMenu from 'components/LockMenu'
import { useMeiCDP, useMeichainBalance } from 'hooks/meichain'
import { usePrice } from 'hooks/price'

import ConnectCosmos from 'images/connect-meichain.svg'

const Card = styled(Flex).attrs(() => ({
  width: '100%',
  height: '100%',
  flexDirection: 'column',
}))`
  background: rgba(249, 251, 252, 0.9);
  box-shadow: 0px 16px 32px rgba(95, 106, 128, 0.1);
  border-radius: 0.56vw;
  position: relative;
`

const LoggedInToMeiChain = ({ meiAddress }) => {
  const [{ data: meichainBalanceData, loading: meichainBalanceLoading, error: meichainBalanceError }, meiAccountBalanceRefetch] = useMeichainBalance(meiAddress)
  const [{ data: cdpData, loading: cdpLoading, error: cdpError }, cdpRefetch] = useMeiCDP(meiAddress)
  const [{ data: priceData, loading: priceLoading, error: priceError }, priceRefetch] = usePrice()

  return (
    <Flex flexDirection="column" width="100%">
      {meichainBalanceData ? (
        <LoanStatus meiAddress={meiAddress} meichainBalance={meichainBalanceData} />)
        : "loading..."}
      {priceData && cdpData ? (
        <DebtMenu cdp={cdpData} price={priceData.cosmos.usd}/>)
        : "loading..."}
      {cdpData ? (
        <LockMenu cdp={cdpData} meichainBalance={meichainBalanceData}/>)
        : "loading..."}
    </Flex>
  )
}

export default ({ meiAddress, setMeiAddress }) => {

  return (
    <Card>
      {meiAddress ? (
        <LoggedInToMeiChain meiAddress={meiAddress}/>
      ) : (
        <Flex
          flexDirection="column"
          alignItems="center"
          justifyContent="center"
          width="100%"
          height="100%"
          p="1.8vw"
        >
          <Image src={ConnectCosmos} width="23vw" />
          <Button
            py="0.55vw"
            px="1vw"
            mt="2.22vw"
            background="#971e44"
            boxShadow="0px 4px 8px rgba(151, 30, 68, 0.25)"
            onClick={() => {
              const address = window.prompt('Insert MeiChain Address')
              setMeiAddress(address)
            }}
          >
            <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
              Connect To MeiChain
            </Text>
          </Button>
        </Flex>
      )}
    </Card>
  )
}
