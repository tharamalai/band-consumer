import { Text } from 'rebass'
import React from 'react'
import Button from 'components/Button'

export default ({ onClick }) => (
  <Button
    py="0.55vw"
    width="6.528vw"
    background="linear-gradient(247.38deg, #56459e 9.86%, #909fb9 89.2%)"
    boxShadow="0px 4px 8px rgba(86, 69, 158, 0.25)"
    onClick={onClick}
  >
    <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
      Send MEI
    </Text>
  </Button>
)
