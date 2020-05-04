import React from 'react'
import styled from 'styled-components'
import { Image, Flex, Text } from 'rebass'
import colors from 'ui/colors'
import { AbsoluteLink } from 'ui/common'

// images
import MeiLogo from 'images/mei-icon.svg'
import TwitterLogo from 'images/twitter.svg'
import TelegramLogo from 'images/telegram.svg'

const NavContainer = styled.nav`
  display: flex;
  height: 7vw;
  width: 84vw;
  margin-left: 8vw;
  margin-right: 8vw;
  align-items: center;
  top: 0px;
  z-index: 9999;
  justify-content: space-between;
  position: relative;
`

const LeftContainer = styled(Flex)`
  justify-content: center;
  align-items: center;
`

const RightContainer = styled(Flex).attrs(() => ({
  width: '4.5vw',
}))`
  justify-content: space-between;
  align-items: center;
`

const LogoOval = styled(Flex).attrs(() => ({
  justifyContent: 'center',
  alignItems: 'center',
  py: '0.35vw',
  pl: '1.2vw',
  pr: '0.54vw',
  bg: colors.white,
}))`
  border-radius: 10vw;
`

const SocialLink = ({ src, to, width }) => (
  <AbsoluteLink href={to}>
    <Image src={src} width={width} style={{ cursor: 'pointer' }} />
  </AbsoluteLink>
)

export default () => (
  <NavContainer>
    <LeftContainer>
      <LogoOval>
        <Image src={MeiLogo} width="1.9vw" />
        <Text
          fontSize={['1.667vw']}
          ml={'1.4vw'}
          color={colors.pink.dark}
          fontWeight={900}
        >
          MEI
        </Text>
      </LogoOval>
      <Text
        fontSize={['1.667vw']}
        ml={'0.6vw'}
        color={colors.white}
        fontWeight={900}
      >
        CHAIN
      </Text>
    </LeftContainer>
    <RightContainer>
      <SocialLink
        src={TwitterLogo}
        to="https://twitter.com"
        width={['1.9vw']}
      />
      <SocialLink
        src={TelegramLogo}
        to="https://telegram.org"
        width={['1.6vw']}
      />
    </RightContainer>
  </NavContainer>
)
