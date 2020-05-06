import React, { useState } from 'react'
import styled from 'styled-components'
import { Image, Flex } from 'rebass'

// components
import NavBar from 'components/NavBar'
import About from 'components/About'
import ConnectCosmos from 'components/CosmosAccount'
import Liquidate from 'components/Liquidate'
import MeiChain from 'components/MeiChain'

// images
import BgTopLeft from 'images/bg-top-left.svg'
import BgTopRight from 'images/bg-top-right.svg'
import BgBottomRight from 'images/bg-bottom-right.svg'
import BgCenter from 'images/bg-center.svg'

const Container = styled.div`
  position: relative;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
`

const InnerContainer = styled(Flex).attrs(() => ({
  mx: '8vw',
  width: '84vw',
  mt: '3vw',
}))`
  justify-content: space-between;
`

const LeftContainer = styled(Flex).attrs(() => ({
  width: '27vw',
  flexDirection: 'column',
}))``

const RightContainer = styled(Flex).attrs(() => ({
  width: '44vw',
}))``

const BG = ({ src, width, top, left, right, bottom }) => (
  <Image
    src={src}
    width={width}
    style={{ position: 'absolute', top, left, right, bottom, zIndex: -1 }}
  />
)

export default () => {
  const [cosmosAddress, setCosmosAddress] = useState('')
  const [meiAddress, setMeiAddress] = useState('')

  return (
    <Container>
      <BG src={BgTopLeft} width="35vw" top={0} left={0} />
      <BG src={BgTopRight} width="50vw" top={0} right={0} />
      <BG src={BgCenter} width="35vw" top="18vw" left="27vw" />
      <BG src={BgBottomRight} width="27vw" bottom="-6.5vw" right="-1.8vw" />
      <NavBar />
      <InnerContainer>
        <LeftContainer>
          <About />
          <ConnectCosmos
            cosmosAddress={cosmosAddress}
            setCosmosAddress={setCosmosAddress}
          />
          <Liquidate meiAddress={meiAddress} />
        </LeftContainer>
        <RightContainer>
          <MeiChain meiAddress={meiAddress} setMeiAddress={setMeiAddress} />
        </RightContainer>
      </InnerContainer>
    </Container>
  )
}
