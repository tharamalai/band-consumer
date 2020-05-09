import { convertSignMsg } from "utils";


export const signAndBroadcastMessage = (cosmos, msg, privateKey) => {
  const stdSignMsg = cosmos.newStdMsg(msg);
  let signedTx = cosmos.sign(stdSignMsg, privateKey)
  signedTx = convertSignMsg(signedTx)
  cosmos.broadcast(signedTx).then(response => console.log(response))
}