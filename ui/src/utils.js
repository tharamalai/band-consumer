import Big from 'big.js'

export const ATOM_UNIT_PER_ONE_ATOM = "1000000"

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
  return response.result.find(token => token.denom = "uatom")
}