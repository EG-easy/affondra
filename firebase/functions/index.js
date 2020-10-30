const functions = require('firebase-functions');

const express = require('express');
const axios = require('axios');
const cors = require('cors');

const { Secp256k1HdWallet, SigningCosmosClient, makeCosmoshubPath } = require('@cosmjs/launchpad');

// get env variables
require('dotenv').config()
const env = process.env
const API_BASE_URL = env.API_BASE_URL;
const ADDRESS_PREFIX = env.ADDRESS_PREFIX;
const API_FAUCET_FROM = env.API_FAUCET_FROM;
const MNEMONIC = env.MNEMONIC;

const app = express();
app.use(cors({origin: true}));
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
          "from": API_FAUCET_FROM,
          "memo": "You are a part of Affondra🤪",
          "chain_id": "affondra",
          "gas": "auto"
        },
        "amount": [
          { "denom": "affondollar", "amount": "1000" },
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

app.get('/node_info', (req, res) => {

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'GET',
      url    : `/node_info`,
      headers: {'Content-Type': 'application/json'},
    });

  (async () => {
    const response = await sendTx;
    res.json(response.data);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});

app.get('/bank/balances', (req, res) => {

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'GET',
      url    : `/bank/balances/${req.body.address}`,
      headers: {'Content-Type': 'application/json'},
    });

  (async () => {
    const response = await sendTx;
    res.json(response.data);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});


app.post('/nfts/mint', (req, res) => {

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'POST',
      url    : `/nfts/mint`,
      headers: {'Content-Type': 'application/json'},
      data   : {
        "base_req": {
          "from": `${req.body.from}`,
          "chain_id": "affondra"
        },
        "recipient": `${req.body.from}`,
        "denom": `${req.body.denom}`,
        "id": `${req.body.id}`,
        "tokenURI": `${req.body.tokenURI}`,
      },
    });

  (async () => {
    const response = await sendTx;
    res.json(response.data);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});

app.post('/affondra/item', (req, res) => {

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'POST',
      url    : `/affondra/item`,
      headers: {'Content-Type': 'application/json'},
      data   : {
        "base_req": {
          "from": `${req.body.from}`,
          "chain_id": "affondra"
        },
        "creator": `${req.body.from}`,
        "denom": `${req.body.denom}`,
        "nftId": `${req.body.nftId}`,
        "price": `${req.body.price}`,
        "affiliate": `${req.body.affiliate}`,
        "inSale": `${req.body.inSale}`,
      },
    });

  (async () => {
    const response = await sendTx;
    res.json(response.data);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});

app.post('/affondra/item/buy', (req, res) => {

  const sendTx = axios({
      baseURL: API_BASE_URL,
      method : 'POST',
      url    : `/affondra/item/buy`,
      headers: {'Content-Type': 'application/json'},
      data   : {
        "base_req": {
          "from": `${req.body.from}`,
          "chain_id": "affondra"
        },
        "ID": `${req.body.ID}`,
        "receiver": `${req.body.from}`,
        "introducedBy": `${req.body.introducedBy}`,
      },
    });

  (async () => {
    const response = await sendTx;
    res.json(response.data);
  })().catch(error => {
    console.log('error', error);
    res.json({ error });
  });
});
// Expose Express API as a single Cloud Function:
exports.affondra = functions.region('asia-northeast1').https.onRequest(app);
