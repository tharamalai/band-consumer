import React from 'react'
import { Flex, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import BorrowBtn from 'components/BorrowBtn'
import ReturnBtn from 'components/ReturnBtn'
import SendMeiBtn from 'components/SendMeiBtn'
import CompareBar from 'components/CompareBar'
import { toAtom, toMei } from 'utils'

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

export default ({ cdp }) => (
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
        percent={60.51}
        valueInUSD={232420}
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
        valueInUSD={82420}
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
        valueInUSD={232420}
      />
    </Flex>
    <CompareBar />
    <Flex
      flexDirection="row"
      justifyContent="space-between"
      mt="0.4vw"
      width="100%"
    >
      <BorrowBtn onClick={() => alert('Borrow')} />
      <ReturnBtn onClick={() => alert('Return')} />
      <SendMeiBtn onClick={() => alert('Send MEI')} />
    </Flex>
  </Container>
)
