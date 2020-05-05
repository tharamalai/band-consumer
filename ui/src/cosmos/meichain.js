
const cosmosjs = require("@cosmostation/cosmosjs");

export const MEICHAIN_LCD_URL = "http://localhost:8010"
export const  MEICHAIN_CHAIN_ID = "meichain";

export let meichain = null
export let ecpairPriv = null

export const initiateMeichain = () => {
  meichain = cosmosjs.network(MEICHAIN_LCD_URL, MEICHAIN_CHAIN_ID);
  return meichain
}

export const isInitiateMeichain = () => {
  if (!meichain) {
    throw Error(`Error meichain js is not initialized`)
  }
}

export const getMeichainAddress = (mnemonic) => {
  try {
    isInitiateMeichain()
    meichain.setPath("m/44'/118'/0'/0/0");
    const address = meichain.getAddress(mnemonic);
    ecpairPriv = meichain.getECPairPriv(mnemonic);
    return address
  } catch (error) {
    console.log(error)
    throw Error(`Error cannot get meichain account from mnemonic: ${error.message}`)
  }
}

export const lockCollateral = (meiAddress, amount) => {
  isInitiateMeichain()
  if (!ecpairPriv) {
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
  
    const signedTx = meichain.sign(stdSignMsg, ecpairPriv);
    meichain.broadcast(signedTx).then(response => console.log(response));
  })
}

export const unlockCollateral = (meiAddress, amount) => {
  isInitiateMeichain()
  if (!ecpairPriv) {
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
  
    const signedTx = meichain.sign(stdSignMsg, ecpairPriv);
    meichain.broadcast(signedTx).then(response => console.log(response));
  })
}
