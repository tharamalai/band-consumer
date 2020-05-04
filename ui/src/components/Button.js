import { Flex } from 'rebass'
import React from 'react'
import styled from 'styled-components'

const Button = styled(Flex).attrs(({ width, py, px, mt, mb }) => ({
  py,
  px,
  mt,
  mb,
  width,
}))`
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  border-radius: ${(props) => props.borderRadius || '3.472vw'};
  background: ${(props) => props.background || '#56459e'};
  box-shadow: ${(props) =>
    props.boxShadow || '0px 4px 8px rgba(86, 69, 158, 0.25)'};
  cursor: pointer;
`

export default ({
  width,
  mt,
  mb,
  py,
  px,
  borderRadius,
  background,
  boxShadow,
  onClick,
  children,
}) => (
  <Button
    width={width}
    mt={mt}
    mb={mb}
    py={py}
    px={px}
    borderRadius={borderRadius}
    background={background}
    boxShadow={boxShadow}
    onClick={onClick}
  >
    {children}
  </Button>
)
