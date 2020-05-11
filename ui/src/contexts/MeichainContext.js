import React, { useContext, createContext, useState } from 'react'
import { ATOM_UNIT_SYMBOL, MEI_UNIT_SYMBOL, MEICHAIN_CHAIN_ID, getMeichainRestServer, convertSignMsg } from 'utils'
import { signAndBroadcastMessage } from 'helpers'

const MeichainContext = createContext()

export const MeichainProvider = ({ children}) => {
  const [privateKey, setPrivateKey] = useState('')
  const [meiAddress, setMeiAddress] = useState('')

  const cosmosjs = require("@cosmostation/cosmosjs")
  const meichain = cosmosjs.network(getMeichainRestServer(), MEICHAIN_CHAIN_ID)

  const isInitiateMeichain = () => {
    if (!meichain) {
      throw Error(`Error meichain js is not initialized`)
    }
  }

  const getMeichainAddress = (mnemonic) => {
    try {
      isInitiateMeichain()
      const address = meichain.getAddress(mnemonic)
      setMeiAddress(address)
      return address
    } catch (error) {
      console.log(error)
      throw Error(`Error cannot get meichain account from mnemonic: ${error.message}`)
    }
  }

  const setPrivateKeyFromMnemonic = (mnemonic) => {
    try {
      const privateKey = meichain.getECPairPriv(mnemonic)
      setPrivateKey(privateKey)
    } catch (error) {
      throw Error(`Error cannot set private key from mnemonic: ${error.message}`)
    }
  }

  const lockCollateral = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before lock atom`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "meichain/LockCollateral",
            value: {
              Amount: String(amount),
              Sender: meiAddress,
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })

  }

  const unlockCollateral = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before unlock atom`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "meichain/UnlockCollateral",
            value: {
              Amount: String(amount),
              Sender: meiAddress,
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })
  }

  const borrowDebt = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before borrow mei`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "meichain/BorrowDebt",
            value: {
              Amount: String(amount),
              Sender: meiAddress,
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })

  }

  const returnDebt = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before return mei`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "meichain/ReturnDebt",
            value: {
              Amount: String(amount),
              Sender: meiAddress,
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })
  }

  const liquidate = (cdpOwnerAddresss) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before liquidate debt`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "meichain/Liquidate",
            value: {
              CdpOwner: cdpOwnerAddresss,
              Liquidator: meiAddress,
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })

  }

  const sendMei = (amount, recipient) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before send mei`
    }

    meichain.getAccounts(meiAddress).then(data => {
      const msg = {
        msgs: [
          {
            type: "cosmos-sdk/MsgSend",
            value: {
              amount: [
                {
                  amount: String(amount),
                  denom: MEI_UNIT_SYMBOL
                }
              ],
              from_address: meiAddress,
              to_address: recipient
            }
          }
        ],
        chain_id: MEICHAIN_CHAIN_ID,
        fee: { amount: [], gas: String(200000) },
        memo: "",
        account_number: String(data.result.value.account_number),
        sequence: String(data.result.value.sequence)
      }
  
      signAndBroadcastMessage(meichain, msg, privateKey)
    })
  }

  return (
    <MeichainContext.Provider value={{ 
      getMeichainAddress,
      setPrivateKeyFromMnemonic,
      lockCollateral,
      unlockCollateral,
      borrowDebt,
      returnDebt,
      liquidate,
      sendMei,
     }}>
       {children}
    </MeichainContext.Provider>
  )
}

export const useMeichainContextState = () => useContext(MeichainContext)