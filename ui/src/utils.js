import Big from 'big.js'

export const ATOM_UNIT_PER_ONE_ATOM = "1000000"

export const MEI_UNIT_PER_ONE_MEI = "1000000"

export const MEICHAIN_GAIA_TRANSFER_CHANNEL = "qnbghznznd"

export const GAIA_MEICHAIN_TRANSFER_CHANNEL = "izwkqheeij"

export const ATOM_UNIT_SYMBOL = "uatom"

export const MEI_UNIT_SYMBOL = "umei"
  
export const COSMOS_CHAIN_ID = "band-cosmoshub"

export const MEICHAIN_CHAIN_ID = "meichain"

export const COSMOS_GAIA_URL = "http://gaia-ibc-hackathon.node.bandchain.org"

export const MEICHAIN_URL = "http://13.250.187.211"

export const findTokenBySymbol = (tokens, tokenSymbol) => {
  if (!tokens) {
    return {
      denom: tokenSymbol,
      amount: "0",
    }
  }
  
  const token = tokens.find(token => token.denom == tokenSymbol)

  if (!token) {
    return {
      denom: tokenSymbol,
      amount: "0",
    }
  }

  return token
}


// Cosmos Hub
export const toAtom = (atomUnitString) => {
  try {
    let atomUnit = Big(atomUnitString)
    const atomUnitPerAtom = Big(ATOM_UNIT_PER_ONE_ATOM)
    const atom = atomUnit.div(atomUnitPerAtom)
    return atom.toFixed(6)
  } catch (error) {
      throw "Error invalid atom unit amount string. Cannot convert atom amount"
  }
}

export const toAtomUnit = (atomString) => {
  try {
    const atom = Big(atomString)
    const atomUnitPerAtom = Big(ATOM_UNIT_PER_ONE_ATOM)
    const atomUnit = atom.times(atomUnitPerAtom)
    return atomUnit.toString()
  } catch (error) {
    throw "Error invalid atom amount string. Cannot convert atom amount"
  }
}

export const convertAtomToUsd = (atomString, usdString) => {
  let atom = Big(atomString)
  const USD_PER_ATOM = Big(usdString)
  atom = atom.times(USD_PER_ATOM)
  return atom.toFixed(2)
}

// Meichain
export const toMei = (meiUnitString) => {
  try {
    let meiUnit = Big(meiUnitString)
    const meiUnitPerMei = Big(MEI_UNIT_PER_ONE_MEI)
    meiUnit = meiUnit.div(meiUnitPerMei)
    return meiUnit.toFixed(6)
  } catch (error) {
    throw "Error invalid mei unit amount string. Cannot convert mei amount"
  }
}

export const toMeiUnit = (meiString) => {
  try {
    const mei = Big(meiString)
    const meiUnitPerMei = Big(MEI_UNIT_PER_ONE_MEI)
    const meiUnit = mei.times(meiUnitPerMei)
    return meiUnit.toString()
  } catch (error) {
    throw "Error invalid mei amount string. Cannot convert mei amount"
  }

}

// Meichain atom is atom transfered from cosmoshub
export const getMeichainAtomSymbol = () => {
  return `transfer/${MEICHAIN_GAIA_TRANSFER_CHANNEL}/uatom`
}

export const calculateDebtPercent = (_debtInUSD, _collateralInUSD) => {
  const debtUSD = Big(_debtInUSD)
  const collateralUSD = Big(_collateralInUSD)
  let debtPercent
  if (collateralUSD.gt(0)) {
    debtPercent = debtUSD.div(collateralUSD).times(100)
  } else {
    debtPercent = debtUSD.times(100)
  }
  return debtPercent.toFixed(2)
}

export const calculateMaxDebtUSD = (_collateralInUSD) => {
  const allDebt = Big(2)
  const allCollateral = Big(3)
  const collateralUSD = Big(_collateralInUSD)
  let maxDebtUSD
  if (collateralUSD.gt(0)) {
    maxDebtUSD = allDebt.div(allCollateral)
    maxDebtUSD = maxDebtUSD.times(collateralUSD)
  } else {
    maxDebtUSD = Big(0)
  }
  return maxDebtUSD.toFixed(2)
}

export const getHost = () => {
  const host = window.location.host
  if (host && host.includes("localhost")) {
    return "localhost"
  }
  return host
}

export const getCosmosLcdUrl = () => {
  const host = getHost()
  if (host === "localhost") {
    return "http://localhost:8012"
  }
  return `${COSMOS_GAIA_URL}:8000`
}

export const getCosmosRestServer = () => {
  const host = getHost()
  if (host === "localhost") {
    return "http://localhost:8011"
  }
  return `${COSMOS_GAIA_URL}:1317`
}


export const getMeichainRestServer = () => {
  const host = getHost()
  if (host === "localhost") {
    return "http://localhost:8010"
  }
  return `${MEICHAIN_URL}:1317`
}

export const generateNewMnemonic = () => {
  try {
    const bip39 = require('bip39')
    const mnemonic = bip39.generateMnemonic()
    return mnemonic
  } catch (error) {
    throw `Error while generate new mnemonic: ${error}`
  }
}

export const safeAccess = (object, path) => {
  return object
    ? path.reduce(
        (accumulator, currentValue) => (accumulator && accumulator[currentValue] ? accumulator[currentValue] : null),
        object
      )
    : null
}

export const convertSignMsg = (signedMsg) => {
  for (const sig of signedMsg.tx.signatures) {
    // sha256("tendermint/PubKeySecp256k1") = f8ccea**eb5ae987**ea423e6cc0e94297a53bd6862df3b3a02a6f6fc89250308760
    // We use eb5ae987 AND 21 (= 21 base 16 = 33 = pubkey size)
    sig.public_key = Buffer.from(
      'eb5ae98721' + Buffer.from(sig.pub_key.value, 'base64').toString('hex'), 'hex'
    ).toString('base64')
    if (sig.pub_key) {
      delete sig.pub_key
    }
    signedMsg.tx.signatures = [sig]
    return signedMsg
  }
}

