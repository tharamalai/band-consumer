import { Image, Text } from 'rebass'
import Button from 'components/Button'
import React from 'react'
import faucet from 'images/faucet.svg'

export default ({ onClick }) => (
  <Button
    py="0.55vw"
    px="1vw"
    background="#5F6A80"
    boxShadow="0px 4px 8px rgba(95, 106, 128, 0.25)"
    onClick={onClick}
  >
    <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
      Faucet
    </Text>
    <Image src={faucet} width="0.9vw" ml="0.3vw" />
  </Button>
)
