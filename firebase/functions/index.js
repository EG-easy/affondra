const functions = require('firebase-functions');

const express = require('express');
const axios = require('axios');
const { Secp256k1HdWallet, SigningCosmosClient, makeCosmoshubPath } = require('@cosmjs/launchpad');

const API_BASE_URL = 'http://117.102.198.165:1317';
const ADDRESS_PREFIX = "cosmos"
const API_FAUCET_FROM = 'cosmos1fahxf2qvave87ja9wze3f4ggmc4gzq8t6fqmkv';
const MNEMONIC = 'interest feature sauce youth voyage person gossip similar parrot drip melody pulse wide turn spice pond visit analyst once napkin arch sugar tumble torch'

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// function errorHandler (err, req, res, next) {

// }
// app.use(errorHandler);

app.post('/faucet', (req, res) => {
  // Type check
  if (typeof req.body.to === 'undefined' || typeof req.body.to !== 'string') throw new Error(`/faucet>INVALID_PARAMETER:typeof body.to: ${typeof req.body.to}`);
  if (!req.body.to.match(/^cosmos[a-z0-9]{39}$/)) throw new Error(`/faucet>INVALID_PARAMETER: body.to.match(/^cosmos[a-z0-9]{39}$/): false`);

  let tasks = [];
  
  const getClient = (async () => {
    const wallet = await Secp256k1HdWallet.fromMnemonic(MNEMONIC, makeCosmoshubPath(0), ADDRESS_PREFIX);
    // const [{ address }] = await wallet.getAccounts();
    const client = new SigningCosmosClient(API_BASE_URL, API_FAUCET_FROM, wallet);
    return { client }
  })();

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'POST',
      url    : `/bank/accounts/${req.body.to}/transfers`,
      headers: {'Content-Type': 'application/json'},
      data   : {
        "base_req": {
          "from": "cosmos1fahxf2qvave87ja9wze3f4ggmc4gzq8t6fqmkv",
          "memo": "You are a part of Affondra🤪",
          "chain_id": "affondra",
          "gas": "auto"
        },
        "amount": [
          { "denom": "affondollar", "amount": "50" },
          { "denom": "stake", "amount": "50" }
        ]
      }
    });

  (async () => {
    const [
      { client },
      { data: { value: { msg, fee, memo } } }
    ] = [await getClient, await sendTx];
    const response = await client.signAndBroadcast(msg, fee, memo);
    res.json(response);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});

// Expose Express API as a single Cloud Function:
exports.affondra = functions.region('asia-northeast1').https.onRequest(app);
