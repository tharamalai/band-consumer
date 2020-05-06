import React from 'react'
import { Flex, Text } from 'rebass'
import colors from 'ui/colors'
import Button from 'components/Button'
import { toAtom, toAtomUnit, findTokenBySymbol, getMeichainAtomSymbol } from 'utils'
import { useMeichainContextState } from 'contexts/MeichainContext'

export default ({ cdp, meichainBalance }) => { 
  const { lockCollateral, unlockCollateral } = useMeichainContextState()

  return (
    <Flex flexDirection="column" width="100%" p="1.8vw">
      <Flex flexDirection="row" justifyContent="space-between">
        <Flex flexDirection="row" alignItems="flex-end">
          <Text
            fontSize="1.867vw"
            fontWeight={400}
            lineHeight="2.114vw"
            color={colors.purple.dark}
          >
            {meichainBalance 
              ? `${toAtom(findTokenBySymbol(meichainBalance.result, getMeichainAtomSymbol()).amount)}`
              : "0" 
            }
          </Text>
          <Text
            fontSize="0.83vw"
            fontWeight={800}
            lineHeight="1vw"
            ml="0.2vw"
            mb="0.2vw"
            color={colors.purple.normal}
          >
            ATOM
          </Text>
          <Text
            fontSize="0.83vw"
            fontWeight={500}
            lineHeight="1vw"
            ml="0.3vw"
            mb="0.2vw"
            color={colors.purple.normal}
          >
            IN TOTAL
          </Text>
        </Flex>
        <Button
          py="0.55vw"
          width="8vw"
          boxShadow="0px 4px 8px rgba(86, 69, 158, 0.25)"
          background={colors.purple.dark}
          onClick={() => {
            const amount = window.prompt('Input Lock Atom Amount')
            if (!amount) {
              return
            }
            lockCollateral(toAtomUnit(amount))
          }}
        >
          <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
            LOCK ATOM
          </Text>
        </Button>
      </Flex>
      <Flex flexDirection="row" justifyContent="space-between" mt="0.8vw">
        <Flex flexDirection="row" alignItems="flex-end">
          <Text
            fontSize="1.867vw"
            fontWeight={400}
            lineHeight="2.114vw"
            color={colors.purple.dark}
          >
            {cdp 
              ? `${toAtom(cdp.result.collateralAmount)}`
              : "0" 
            }
          </Text>
          <Text
            fontSize="0.83vw"
            fontWeight={800}
            lineHeight="1vw"
            ml="0.2vw"
            mb="0.2vw"
            color={colors.purple.normal}
          >
            ATOM
          </Text>
          <Text
            fontSize="0.83vw"
            fontWeight={500}
            lineHeight="1vw"
            ml="0.3vw"
            mb="0.2vw"
            color={colors.purple.normal}
          >
            LOCKED
          </Text>
        </Flex>
        <Button
          py="0.55vw"
          width="8vw"
          boxShadow="0px 4px 8px rgba(86, 69, 158, 0.25)"
          background={colors.purple.normal}
          onClick={() => {
            const amount = window.prompt('Input Unlock Atom Amount')
            if (!amount) {
              return
            }
            unlockCollateral(toAtomUnit(amount))
          }}
        >
          <Text fontSize="0.83vw" fontWeight={500} lineHeight="1vw">
            UNLOCK ATOM
          </Text>
        </Button>
      </Flex>
    </Flex>
  )
}
