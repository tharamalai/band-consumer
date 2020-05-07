import React, { useContext, createContext, useState } from 'react'
import { MEICHAIN_GAIA_TRANSFER_CHANNEL, GAIA_MEICHAIN_TRANSFER_CHANNEL,  ATOM_UNIT_SYMBOL } from 'utils'

const CosmosHubContext = createContext()

export const CosmosHubProvider = ({ children }) => {
  const [privateKey, setPrivateKey] = useState('')
  const [cosmosAddress, setCosmosAddress] = useState('')

  const cosmosjs = require("@cosmostation/cosmosjs")
  const COSMOS_LCD_URL = "http://gaia-ibc-hackathon.node.bandchain.org:1317"
  const COSMOS_CHAIN_ID = "band-cosmoshub"
  
  const cosmos = cosmosjs.network(COSMOS_LCD_URL, COSMOS_CHAIN_ID)
  cosmos.setPath("m/44'/118'/0'/0/0");

  const isInitiateCosmos = () => {
    if (!cosmos) {
      throw Error(`Error cosmos js is not initialized`)
    }
  }

  const getCosmosAddress = (mnemonic) => {
    try {
      isInitiateCosmos()
      const address = cosmos.getAddress(mnemonic)
      setCosmosAddress(address)
      return address
    } catch (error) {
      console.log(error)
      throw Error(`Error cannot get cosmos account from mnemonic: ${error.message}`)
    }
  }

  const setPrivateKeyFromMnemonic = (mnemonic) => {
    try {
      const privateKey = cosmos.getECPairPriv(mnemonic)
      setPrivateKey(privateKey)
    } catch (error) {
      throw Error(`Error cannot set private key from mnemonic: ${error.message}`)
    }
  }

  const sendTokenToMeichain = (amount, receiver) => {
    try {
      isInitiateCosmos()
      if (!privateKey) {
        throw `Please connect wallet before send token to meichain`
      }
      cosmos.getAccounts(cosmosAddress).then(data => {
        let stdSignMsg = cosmos.newStdMsg({
          msgs: [
            {
              type: "ibc/transfer/MsgTransfer",
              value: {
                source_port: "transfer",
                source_channel: GAIA_MEICHAIN_TRANSFER_CHANNEL,
                dest_height: "10000000",
                amount: [
                  {
                    amount: String(amount),
                    denom: `transfer/${MEICHAIN_GAIA_TRANSFER_CHANNEL}/${ATOM_UNIT_SYMBOL}`
                  }
                ],
                sender: cosmosAddress,
                receiver: receiver
              }
            }
          ],
          chain_id: COSMOS_CHAIN_ID,
          fee: { amount: [], gas: String(200000) },
          memo: "",
          account_number: String(data.result.value.account_number),
          sequence: String(data.result.value.sequence)
        });
      
        const signedTx = cosmos.sign(stdSignMsg, privateKey)
        cosmos.broadcast(signedTx).then(response => console.log(response))
      })
    } catch (error) {
      throw Error(`Error cannot send token to meichain: ${error.message}`)
    }
  }


  return (
    <CosmosHubContext.Provider value={{ 
      COSMOS_CHAIN_ID,
      getCosmosAddress,
      setPrivateKeyFromMnemonic,
      sendTokenToMeichain
     }}>
       {children}
    </CosmosHubContext.Provider>
  )
}

export const useCosmosHubContextState = () => useContext(CosmosHubContext)