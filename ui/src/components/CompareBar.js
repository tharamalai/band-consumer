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
`

export default ({ debt, maxDebt, collateral }) => (
  <Flex width="100%" my="1.4vw">
    <Bar width="100%" height="2.083vw" bg={colors.gray.normal}>
      <Bar width="66.6%" height="1.389vw" bg={colors.pink.normal}>
        <Bar width="33.3%" height="0.694vw" bg={colors.pink.dark} />
      </Bar>
    </Bar>
  </Flex>
)
