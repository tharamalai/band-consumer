import React from 'react'
import { Flex, Image, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'
import Button from 'components/Button'

import Bg from 'images/bg-liquidate.svg'
import Thunder from 'images/thunder.svg'
import { liquidate } from 'cosmos/meichain'
import { useMeichainContextState } from 'contexts/MeichainContext'

const Card = styled(Flex).attrs(() => ({
  p: '1.8vw',
  pb: '2vw',
  width: '100%',
  flexDirection: 'column',
  alignItems: 'flexStart',
  mt: '2.5vw',
  height: '7.29vw',
}))`
  background: rgba(249, 251, 252, 0.9);
  box-shadow: 0px 16px 32px rgba(95, 106, 128, 0.1);
  border-radius: 0.56vw;
  position: relative;
`

const NotLogin = () => (
  <Text
    fontSize="0.83vw"
    fontWeight={500}
    lineHeight="1vw"
    color={colors.pink.light}
    mt="1.458vw"
    style={{ position: 'relative', zIndex: 2 }}
  >
    Please connect to Meichain before using this section
  </Text>
)

const Login = () => (
  <Button
    py="0.55vw"
    mt="1.04vw"
    width="16.67vw"
    background="linear-gradient(223.23deg, #d25c7d 9.86%, #f2918b 89.2%)"
    boxShadow="0px 4px 8px rgba(151, 30, 68, 0.25)"
    onClick={() => {
      const cdpOwner = window.prompt('Input CDP Owner Account')
      if (!cdpOwner) {
        return
      }
      liquidate(cdpOwner)
    }}
  >
    <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
      Liquidate Undercollateralized Loan
    </Text>
  </Button>
)

export default ({ meiAddress }) => {
  return (
    <Card>
      <Image
        src={Bg}
        width="8.4vw"
        style={{
          position: 'absolute',
          bottom: '0',
          right: '0',
          zIndex: 1,
        }}
      />
      <Image
        src={Thunder}
        width="4vw"
        style={{
          position: 'absolute',
          top: '-1.2vw',
          left: '-2vw',
          zIndex: 1,
        }}
      />
      <Text
        fontSize="0.83vw"
        fontWeight={500}
        lineHeight="1vw"
        color={colors.pink.light}
      >
        Dangerous Zone
      </Text>
      {meiAddress ? <Login/> : <NotLogin />}
    </Card>
  )
}
