import React from 'react'
import { Box, Image, Text } from 'rebass'
import styled from 'styled-components'
import colors from 'ui/colors'

// images
import BgCircle from 'images/bg-about.svg'
import BgStar from 'images/bg-about2.svg'

const Card = styled(Box).attrs(() => ({
  p: '1.667vw',
  pb: '3.2vw',
  width: '100%',
}))`
  background: rgba(249, 251, 252, 0.9);
  box-shadow: 0px 16px 32px rgba(95, 106, 128, 0.1);
  border-radius: 0.56vw;
  position: relative;
`

export default () => (
  <Card>
    <Image
      src={BgCircle}
      width="10vw"
      style={{ position: 'absolute', top: 0, left: 0, zIndex: 0 }}
    />
    <Image
      src={BgStar}
      width="3vw"
      style={{
        position: 'absolute',
        bottom: '-0.8vw',
        right: '-1.2vw',
        zIndex: 4,
      }}
    />
    <Text
      color={colors.purple.dark}
      fontSize="1.25vw"
      fontWeight={900}
      lineHeight="1.53vw"
      style={{ position: 'relative', zIndex: 2 }}
    >
      About Mei Chain
    </Text>
    <Text
      color={colors.purple.dark}
      fontSize="0.833vw"
      fontWeight={500}
      lineHeight="1vw"
      my="0.64vw"
      style={{ position: 'relative', zIndex: 2 }}
    >
      An $ATOM-backed stablecoin build with Cosmos IBC and Band Protocol Price
      Oracle
    </Text>
  </Card>
)
