import React from 'react'
import { Flex, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import BorrowBtn from 'components/BorrowBtn'
import ReturnBtn from 'components/ReturnBtn'
import SendMeiBtn from 'components/SendMeiBtn'
import CompareBar from 'components/CompareBar'
import { toAtom, toMei, toMeiUnit, convertAtomToUsd, calculateDebtPercent, calculateMaxDebtUSD } from 'utils'
import { useMeichainContextState } from 'contexts/MeichainContext'
import Big from 'big.js'

const Circle = styled(Flex).attrs(({ color }) => ({
  width: '1.389vw',
  height: '1.389vw',
  bg: color,
}))`
  border-radius: 50%;
`

const Container = styled(Flex).attrs(() => ({
  height: '14.617vw',
  px: '1.8vw',
  pt: '1.2vw',
  pb: '1vw',
  flexDirection: 'column',
  mx: '-1.1vw',
}))`
  background: white;
  border-radius: 0.556vw;
  box-shadow: 0px 16px 32px rgba(151, 30, 68, 0.08);
`

const FeatureStat = ({ color, title, percent, valueInUSD }) => (
  <Flex>
    <Circle color={color} />
    <Flex flexDirection="column" ml="0.6vw">
      {title}
      {percent && (
        <Text
          fontSize="1.867vw"
          fontWeight={400}
          lineHeight="2vw"
          color={color}
          mt="0.4vw"
        >
          {percent} %
        </Text>
      )}
      <Text
        mt="0.3vw"
        fontSize="0.83vw"
        fontWeight={400}
        lineHeight="1vw"
        color={colors.purple.normal}
        style={{ fontStyle: 'italic' }}
      >
        ≈ {valueInUSD} USD
      </Text>
    </Flex>
  </Flex>
)

const debtPercent = (cdp, price) => {
  return calculateDebtPercent(toMei(cdp.result.debtAmount), convertAtomToUsd(toAtom(cdp.result.collateralAmount), price))
}

const debtPercentBar = (percent) => {
  let percentBar = Big(percent)
  const maxPercentBar = Big(100)
  if (percentBar.gt(maxPercentBar)) {
    percentBar = maxPercentBar
    return percentBar.toFixed(2)
  }
  return percentBar
}

const maxDebt = (cdp, price) => {
  return calculateMaxDebtUSD(convertAtomToUsd(toAtom(cdp.result.collateralAmount), price))
}

const isValidBorrowAmount = (cdp, price, borrowAmount) => {
  const cdpDebt = Big(cdp.result.debtAmount)
  const max = Big(toMeiUnit(maxDebt(cdp, price)))
  const borrow = Big(toMeiUnit(borrowAmount))
  const maxBorrow = max.minus(cdpDebt)
  return borrow.lte(maxBorrow)
}

export default ({ meiAddress, cdp, price }) => {
  const { borrowDebt, returnDebt, sendMei } = useMeichainContextState()

  return (
    <Container>
      <Flex justifyContent="space-between" width="93%">
        <FeatureStat
          color={colors.pink.dark}
          title={
            <Flex flexDirection="row" mt="0.14vw">
              <Text
                fontSize="0.83vw"
                fontWeight={500}
                lineHeight="1vw"
                color={colors.purple.dark}
              >
                Debt
              </Text>
              <Text
                fontSize="0.83vw"
                fontWeight={900}
                lineHeight="1vw"
                ml="0.2vw"
                color={colors.purple.dark}
              >
                {cdp 
                  ? `${toMei(cdp.result.debtAmount)}  MEI`
                  : "0 MEI" 
                }
              </Text>
            </Flex>
          }
          percent={debtPercent(cdp, price)}
          valueInUSD={toMei(cdp.result.debtAmount)}
        />
        <FeatureStat
          color={colors.pink.normal}
          title={
            <Text
              fontSize="0.83vw"
              fontWeight={500}
              lineHeight="1vw"
              color={colors.purple.dark}
            >
              Max Debt
            </Text>
            
          }
          percent={66.67}
          valueInUSD={maxDebt(cdp, price)}
        />
        <FeatureStat
          color={colors.gray.normal}
          title={
            <Flex flexDirection="row" mt="0.14vw">
              <Text
                fontSize="0.83vw"
                fontWeight={500}
                lineHeight="1vw"
                color={colors.purple.dark}
              >
                Collateral
              </Text>
              <Text
                fontSize="0.83vw"
                fontWeight={900}
                lineHeight="1vw"
                ml="0.2vw"
                color={colors.purple.dark}
              >
                {cdp 
                  ? `${toAtom(cdp.result.collateralAmount)}  ATOM`
                  : "0 ATOM" 
                }
              </Text>
            </Flex>
          }
          valueInUSD={convertAtomToUsd(toAtom(cdp.result.collateralAmount), price)}
        />
      </Flex>
      <CompareBar debtPercent={debtPercentBar(debtPercent(cdp, price))}/>
      <Flex
        flexDirection="row"
        justifyContent="space-between"
        mt="0.4vw"
        width="100%"
      >
        <BorrowBtn onClick={() => {
          const amount = window.prompt('Input Borrow Debt Amount')
          if (!amount) {
            return
          }

          if (!isValidBorrowAmount(cdp, price, amount)) {
            alert('Accumulated debt amount is more than max debt amount')
            return
          }

          borrowDebt(toMeiUnit(amount))
        }} />
        <ReturnBtn onClick={() => {
          const amount = window.prompt('Input Return Debt Amount')
          if (!amount) {
            return
          }
          returnDebt(toMeiUnit(amount))
        }} />
        <SendMeiBtn onClick={() => {
          const amount = window.prompt('Input Send Amount')
          if (!amount) {
            return
          }
          const recipient = window.prompt('Input Recipient')
          if (!recipient) {
            return
          }
          sendMei(toMeiUnit(amount), recipient)
        }} />
      </Flex>
    </Container>
  )
}
