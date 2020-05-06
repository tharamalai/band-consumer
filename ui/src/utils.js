import Big from 'big.js'

export const ATOM_UNIT_PER_ONE_ATOM = "1000000"

export const MEI_UNIT_PER_ONE_MEI = "1000000"

export const MEICHAIN_GAIA_TRANSFER_CHANNEL = "qnbghznznd"

export const GAIA_MEICHAIN_TRANSFER_CHANNEL = "izwkqheeij"

export const ATOM_UNIT_SYMBOL = "uatom"

export const MEI_UNIT_SYMBOL = "umei"

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
  let atomUnit = Big(atomUnitString)
  const atomUnitPerAtom = Big(ATOM_UNIT_PER_ONE_ATOM)
  atomUnit = atomUnit.div(atomUnitPerAtom)
  return atomUnit.toFixed(6)
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
    debtPercent = Big(0)
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