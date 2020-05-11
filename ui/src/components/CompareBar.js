import React from 'react'
import styled from 'styled-components'
import { Flex } from 'rebass'
import colors from 'ui/colors'

const Bar = styled(Flex).attrs(({ width, bg, height }) => ({
  width,
  bg,
  height,
  alignItems: 'center',
}))`
  border-radius: ${(props) => props.height};
  position: absolute;
  left: 0;
  top: ${(props) => props.top || '0'};
`

const Container = styled(Flex).attrs(() => ({
  my: '1.4vw',
  width: '100%',
  height: '1.389vw',
}))`
  position: relative;
`

export default ({ debtPercent }) => (
  <Container>
    <Bar width="100%" height="2.083vw" bg={colors.gray.normal} />
    <Bar width="66.6%" height="1.389vw" bg={colors.gray.dark} top="0.3vw" />
    <Bar width={debtPercent + '%'} height="0.694vw" bg={colors.pink.dark} top="0.6vw" />
  </Container>
)