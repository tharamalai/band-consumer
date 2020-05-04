import { Text } from 'rebass'
import React from 'react'
import Button from 'components/Button'

export default ({ onClick }) => (
  <Button
    py="0.55vw"
    width="6.528vw"
    background="linear-gradient(247.38deg, #971e44 9.86%, #d25c7d 89.2%)"
    boxShadow="0px 4px 8px rgba(151, 30, 68, 0.25)"
    onClick={onClick}
  >
    <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
      Borrow
    </Text>
  </Button>
)
