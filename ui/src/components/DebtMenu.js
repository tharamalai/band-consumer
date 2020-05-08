import React from 'react'
import { Flex, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import BorrowBtn from 'components/BorrowBtn'
import ReturnBtn from 'components/ReturnBtn'
import SendMeiBtn from 'components/SendMeiBtn'
import CompareBar from 'components/CompareBar'
import { toAtom, toMei, toMeiUnit, convertAtomToUsd, calculateDebtPercent, calculateMaxDebtUSD, findTokenBySymbol, MEI_UNIT_SYMBOL, safeAccess } from 'utils'
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
  return calculateDebtPercent(toMei(safeAccess(cdp, ["result", "debtAmount"])), convertAtomToUsd(toAtom(safeAccess(cdp, ["result", "collateralAmount"])), price))
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
  return calculateMaxDebtUSD(convertAtomToUsd(toAtom(safeAccess(cdp, ["result", "collateralAmount"])), price))
}

const maxBorrow = (cdp, price) => {
  const cdpDebtAmount = Big(safeAccess(cdp, ["result", "debtAmount"]))
  const maxAmount = Big(toMeiUnit(maxDebt(cdp, price)))
  const max = maxAmount.minus(cdpDebtAmount)
  return max
}

export default ({ cdp, price, meichainBalance }) => {
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
                  ? `${toMei(safeAccess(cdp, ["result", "debtAmount"]))}  MEI`
                  : "0 MEI" 
                }
              </Text>
            </Flex>
          }
          percent={debtPercent(cdp, price)}
          valueInUSD={toMei(safeAccess(cdp, ["result", "debtAmount"]))}
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
                  ? `${toAtom(safeAccess(cdp, ["result", "collateralAmount"]))}  ATOM`
                  : "0 ATOM" 
                }
              </Text>
            </Flex>
          }
          valueInUSD={convertAtomToUsd(toAtom(safeAccess(cdp, ["result", "collateralAmount"])), price)}
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
          const amount = window.prompt('Input borrow debt amount')
          if (!amount) {
            return
          }

          let borrowAmount
          try {
            borrowAmount = Big(toMeiUnit(amount))
          } catch (error) {
            alert("Invalid amount")
            return
          }

          if (borrowAmount.lte(0)) {
            alert("Amount must more than 0")
            return
          }

          const maxBorrowAmount = maxBorrow(cdp, price)
          if (borrowAmount.gt(maxBorrowAmount)) {
            alert(`Max borrow amount is ${toMei(maxBorrowAmount)}`)
            return
          }

          borrowDebt(toMeiUnit(amount))
        }} />
        <ReturnBtn onClick={() => {
          const amount = window.prompt('Input return debt amount')
          if (!amount) {
            return
          }

          let returnAmount
          try {
            returnAmount = Big(toMeiUnit(amount))
          } catch (error) {
            alert("Invalid amount")
            return
          }

          if (returnAmount.lte(0)) {
            alert("Amount must more than 0")
            return
          }
          
          returnDebt(toMeiUnit(amount))
        }} />
        <SendMeiBtn onClick={() => {
          const amount = window.prompt('Input send amount')
          if (!amount) {
            return
          }

          let sendAmount
          try {
            sendAmount = Big(toMeiUnit(amount))
          } catch (error) {
            alert("Invalid amount")
            return
          }

          if (sendAmount.lte(0)) {
            alert("Amount must more than 0")
            return
          }

          const maxSendAmount = Big(findTokenBySymbol(safeAccess(meichainBalance, ["result"]), MEI_UNIT_SYMBOL).amount)
          if (sendAmount.gt(maxSendAmount)) {
            alert(`Max send amount is ${toMei(maxSendAmount)}`)
            return
          }

          const recipient = window.prompt('Input recipient')
          if (!recipient) {
            return
          }
          sendMei(toMeiUnit(amount), recipient)
        }} />
      </Flex>
    </Container>
  )
}
