import Big from 'big.js'

export const ATOM_UNIT_PER_ONE_ATOM = "1000000"

export const MEI_UNIT_PER_ONE_MEI = "1000000000000000000"

export const TRANSFER_CHANNEL = "tiawodbkqg"

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

export const findAtomAmount = (response) => {
  if (!response.result) {
    return {
      denom: "uatom",
      amount: "0",
    }
  } 
  return response.result.find(token => token.denom == "uatom")
}

// Meichain

export const toMei = (meiUnitString) => {
  try {
    let meiUnit = Big(meiUnitString)
    const meiUnitPerMei = Big(MEI_UNIT_PER_ONE_MEI)
    meiUnit = meiUnit.div(meiUnitPerMei)
    return meiUnit.toFixed(18)
  } catch (error) {
    throw "Error invalid mei amount string. Cannot convert mei amount"
  }
}

export const findMeiAmount = (response) => {
  if (!response.result) {
    return {
      denom: "umei",
      amount: "0",
    }
  } 

  const meiToken = response.result.find(token => token.denom == "umei")

  if (!meiToken) {
    return {
      denom: "umei",
      amount: "0",
    }
  }
  return meiToken
}

// Meichain atom is atom transfered from cosmoshub
export const getMeichainAtomSymbol = () => {
  return `transfer/${TRANSFER_CHANNEL}/uatom`
}

// export const toMeichainAtom = (tokenAmount) => {
//   const tokenSymbol = getMeichainAtomSymbol()
//   if (token.denim !== tokenSymbol) {
//     throw `Error invalid atom token in Meichain. denom must be ${tokenSymbol}.`
//   }
//   return toAtom(token.amount)
// }