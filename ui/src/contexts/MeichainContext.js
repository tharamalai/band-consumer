import React, { useContext, createContext, useState } from 'react'

const MeichainContext = createContext()

export const MeichainProvider = ({ children}) => {
  const [privateKey, setPrivateKey] = useState('')
  const [meiAddress, setMeiAddress] = useState('')
  const MEICHAIN_LCD_URL = "http://localhost:8010"
  const  MEICHAIN_CHAIN_ID = "meichain";

  const cosmosjs = require("@cosmostation/cosmosjs")
  const meichain = cosmosjs.network(MEICHAIN_LCD_URL, MEICHAIN_CHAIN_ID)
  meichain.setPath("m/44'/118'/0'/0/0")

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
      let stdSignMsg = meichain.newStdMsg({
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
      });
    
      const signedTx = meichain.sign(stdSignMsg, privateKey)
      meichain.broadcast(signedTx).then(response => console.log(response));
    })
  }

  const unlockCollateral = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before unlock atom`
    }
    meichain.getAccounts(meiAddress).then(data => {
      let stdSignMsg = meichain.newStdMsg({
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
      });
    
      const signedTx = meichain.sign(stdSignMsg, privateKey);
      meichain.broadcast(signedTx).then(response => console.log(response))
    })
  }

  const borrowDebt = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before borrow mei`
    }
    meichain.getAccounts(meiAddress).then(data => {
      let stdSignMsg = meichain.newStdMsg({
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
      });
    
      const signedTx = meichain.sign(stdSignMsg, privateKey);
      meichain.broadcast(signedTx).then(response => console.log(response))
    })
  }

  const returnDebt = (amount) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before return mei`
    }
    meichain.getAccounts(meiAddress).then(data => {
      let stdSignMsg = meichain.newStdMsg({
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
      });
    
      const signedTx = meichain.sign(stdSignMsg, privateKey);
      meichain.broadcast(signedTx).then(response => console.log(response))
    })
  }

  const liquidate = (cdpOwnerAddresss, meiAddress) => {
    isInitiateMeichain()
    if (!privateKey) {
      throw `Please connect wallet before liquidate debt`
    }
    meichain.getAccounts(meiAddress).then(data => {
      let stdSignMsg = meichain.newStdMsg({
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
      });
    
      const signedTx = meichain.sign(stdSignMsg, privateKey);
      meichain.broadcast(signedTx).then(response => console.log(response))
    })
  }

  return (
    <MeichainContext.Provider value={{ 
      MEICHAIN_CHAIN_ID,
      getMeichainAddress,
      setPrivateKeyFromMnemonic,
      lockCollateral,
      unlockCollateral,
      borrowDebt,
      returnDebt,
      liquidate
     }}>
       {children}
    </MeichainContext.Provider>
  )
}

export const useMeichainContextState = () => useContext(MeichainContext)