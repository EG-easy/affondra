import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import app from "./app.js";
// import cosmos from "@tendermint/vue/src/store/cosmos.js";
import { Secp256k1HdWallet, SigningCosmosClient, makeCosmoshubPath  } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "https://api.affondra.com";
// const API = "http://affondra.com:1317";
const ADDRESS_PREFIX = "cosmos"

// axios.defaults.baseURL = 'http://192.168.1.160:1317';
axios.defaults.headers.common['Content-Type'] = 'application/json;charset=utf-8'; 
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';

export default new Vuex.Store({
  state: {
    app,
    account: null,
    chain_id: "",
    data: {},
    client: null,
  },
  getters: {
    isLoggedIn: (state) => {
      return (state.account !== null && typeof state.account === 'object' && 'address' in state.account)
    },
    amountAffondollar: (state) => {
      return (state.account !== null && typeof state.account === 'object' && 'coins' in state.account ? state.account.coins.filter(v => v.denom === 'affondollar').map(v => parseInt(v.amount, 10)).concat([0]).reduce((a, x) => a += x, 0) : null);
    }
  },
  mutations: {
    accountUpdate(state, { account }) {
      state.account = account;
    },
    chainIdSet(state, { chain_id }) {
      state.chain_id = chain_id;
    },
    entitySet(state, { type, body }) {
      const updated = {};
      updated[type] = body;
      state.data = { ...state.data, ...updated };
    },
    clientUpdate(state, { client }) {
      state.client = client;
    },
  },
  actions: {
    logout({ commit }) {
      commit("accountUpdate", { account: null });
      commit("clientUpdate", { account: null });
    },
    async init({ dispatch, state }) {
      await dispatch("chainIdFetch");
      state.app.types.forEach(({ type }) => {
        dispatch("entityFetch", { type });
      });
    },
    async chainIdFetch({ commit }) {
      const node_info = (await axios.get(`${API}/node_info`)).data.node_info;
      commit("chainIdSet", { chain_id: node_info.network });
    },
    async accountSignIn({ commit }, { mnemonic }) {
      const wallet = await Secp256k1HdWallet.fromMnemonic(mnemonic, makeCosmoshubPath(0), ADDRESS_PREFIX);
      const [{ address }] = await wallet.getAccounts();
      const url = `${API}/auth/accounts/${address}`;
      const acc = (await axios.get(url)).data;
      
      const account = acc.result.value;
      const client = new SigningCosmosClient(API, address, wallet);
      commit("accountUpdate", { account });
      commit("clientUpdate", { client });

      return account

      // if (acc.result.value.address === address) {
      //   const account = acc.result.value;
      //   const client = new SigningCosmosClient(API, address, wallet);
      //   commit("accountUpdate", { account });
      //   commit("clientUpdate", { client });
      //   return account
      // } else {
      //   console.error("Account doesn't exist.");
      //   return null;
      // }
    },
    async sendTxMintNft({ state }, payload) { // eslint-disable-line no-unused-vars
      
      const txAxiosConfig = {
        baseURL: API,
        method : 'POST',
        url    : `/nfts/mint`,
        data   : {
          "base_req": {
            "from": payload.address,
            "chain_id": "affondra"
          },
          "recipient": payload.address,
          "denom": payload.denom,
          "id": payload.nftId,
          "tokenURI": payload.tokenURI,
        },
      };

      console.log('sendTxMintNft>txAxiosConfig:', txAxiosConfig);

      const { data: rawTx } = await axios(txAxiosConfig);
      console.log('sendTxMintNft>rawTx:',rawTx)
      if (typeof rawTx !== 'object') throw new Error('INVALID_BODY');
      if (rawTx.type !== 'cosmos-sdk/StdTx') throw new Error('INVALID_TXN_TYPE');
      if (typeof rawTx.value !== 'object') throw new Error('INVALID_TXN_VALUE');

      const { msg, fee, memo } = rawTx.value;
      console.log('{ msg, fee, memo }', { msg, fee, memo })
      const logTx = await state.client.signAndBroadcast(msg, fee, memo);
      console.log('sendTxMintNft>logTx:', logTx);
      // If same nft-id is alread existed
      // {
      //   "height": "4483",
      //   "txhash": "D72C496EDFFCB735E08EFF8C7B665B25A68D48C4880A64FDBF026DD4205BA5D3",
      //   "codespace": "nft",
      //   "code": 5,
      //   "raw_log": "NFT already exists: NFT #2iZaPwXbfdiCXRbuzmQLPVRZ5Z88tv2z already exists in collection default: failed to execute message; message index: 0",
      //   "gas_wanted": "200000",
      //   "gas_used": "37594"
      // }
      if ('code' in logTx) throw new Error(`TXN_FAILED:${logTx.rawLog}`)

      // If successed
      // {
      //   "height": "4565",
      //   "txhash": "FC8281AE99C77A1BD8A2D1BC583FC465C6D6D6222CC1EFD9D3B850F88036BC0D",
      //   "raw_log": "[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"mint_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cosmos1sf273eh47xder2ql4aq4zzae8fh54vd2ap9477\"}]},{\"type\":\"mint_nft\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos1sf273eh47xder2ql4aq4zzae8fh54vd2ap9477\"},{\"key\":\"denom\",\"value\":\"default\"},{\"key\":\"nft-id\",\"value\":\"xdM05gO5QX0b2YwUxXqhvyMX9KAKHjl1\"},{\"key\":\"token-uri\",\"value\":\"{\\\"name\\\":\\\"My item\\\",\\\"imgurl\\\":\\\"images/e5P1fsANlO17528uws40TcVYXCbPGX.jpg\\\"}\"}]}]}]",
      //   "logs": [
      //     {
      //       "msg_index": 0,
      //       "log": "",
      //       "events": [
      //         {
      //           "type": "message",
      //           "attributes": [
      //             {
      //               "key": "action",
      //               "value": "mint_nft"
      //             },
      //             {
      //               "key": "module",
      //               "value": "nft"
      //             },
      //             {
      //               "key": "sender",
      //               "value": "cosmos1sf273eh47xder2ql4aq4zzae8fh54vd2ap9477"
      //             }
      //           ]
      //         },
      //         {
      //           "type": "mint_nft",
      //           "attributes": [
      //             {
      //               "key": "recipient",
      //               "value": "cosmos1sf273eh47xder2ql4aq4zzae8fh54vd2ap9477"
      //             },
      //             {
      //               "key": "denom",
      //               "value": "default"
      //             },
      //             {
      //               "key": "nft-id",
      //               "value": "xdM05gO5QX0b2YwUxXqhvyMX9KAKHjl1"
      //             },
      //             {
      //               "key": "token-uri",
      //               "value": "{\"name\":\"My item\",\"imgurl\":\"images/e5P1fsANlO17528uws40TcVYXCbPGX.jpg\"}"
      //             }
      //           ]
      //         }
      //       ]
      //     }
      //   ],
      //   "gas_wanted": "200000",
      //   "gas_used": "53808"
      // }

      return logTx

    },

    async getItemList({ state }, payload) { // eslint-disable-line no-unused-vars
      const { data: rawTx } = await axios({
        baseURL: API,
        method : 'GET',
        url    : `/affondra/item`,
      });
      console.log('getItemList>rawTx:',rawTx)
      return rawTx;
    },

    async sendTxCreateItem({ state }, payload) { // eslint-disable-line no-unused-vars

      const txCreateItemAxiosConfig = {
        baseURL: API,
        method : 'POST',
        url    : `/affondra/item`,
        data   : {
          "base_req": {
            "from": `${payload.address}`,
            "chain_id": "affondra"
          },
          "creator": `${payload.address}`,
          "denom": `${payload.denom}`,
          "nftId": `${payload.nftId}`,
          "price": `${payload.price}`,
          "affiliate": `${payload.affiliate}`,
          "inSale": `${payload.inSale}`,
          "description": `${payload.description}`,
        },
      };

      console.log('sendTxCreateItem>txAxiosConfig:', txCreateItemAxiosConfig);

      const { data: rawTx } = await axios(txCreateItemAxiosConfig);
      console.log('sendTxCreateItem>rawTx:',rawTx)
      if (typeof rawTx !== 'object') throw new Error('INVALID_BODY');
      if (rawTx.type !== 'cosmos-sdk/StdTx') throw new Error('INVALID_TXN_TYPE');
      if (typeof rawTx.value !== 'object') throw new Error('INVALID_TXN_VALUE');

      const { msg, fee, memo } = rawTx.value;
      console.log('{ msg, fee, memo }', { msg, fee, memo })
      const logTx = await state.client.signAndBroadcast(msg, fee, memo);
      console.log('sendTxCreateItem>logTx:', logTx);
      if ('code' in logTx) throw new Error(`TXN_FAILED:${logTx.rawLog}`)

      return logTx

    },
    async sendTxBuyItem({ state }, payload) { // eslint-disable-line no-unused-vars

      const txCreateItemAxiosConfig = {
        baseURL: API,
        method : 'POST',
        url    : `/affondra/item/buy`,
        data   : {
          "base_req": {
            "from": `${payload.buyerAddress}`,
            "chain_id": "affondra"
          },
          "receiver": `${payload.buyerAddress}`,
          "ID": `${payload.id}`,
          "introducedBy": `${payload.addressIntroducedBy}`,
        },
      };

      console.log('sendTxCreateItem>txAxiosConfig:', txCreateItemAxiosConfig);

      const { data: rawTx } = await axios(txCreateItemAxiosConfig);
      console.log('sendTxCreateItem>rawTx:',rawTx)
      if (typeof rawTx !== 'object') throw new Error('INVALID_BODY');
      if (rawTx.type !== 'cosmos-sdk/StdTx') throw new Error('INVALID_TXN_TYPE');
      if (typeof rawTx.value !== 'object') throw new Error('INVALID_TXN_VALUE');

      const { msg, fee, memo } = rawTx.value;
      console.log('{ msg, fee, memo }', { msg, fee, memo })
      const logTx = await state.client.signAndBroadcast(msg, fee, memo);
      console.log('sendTxCreateItem>logTx:', logTx);
      if ('code' in logTx) throw new Error(`TXN_FAILED:${logTx.rawLog}`)

      return logTx

    },
    async entityFetch({ state, commit }, { type }) {
      const { chain_id } = state;
      const url = `${API}/${chain_id}/${type}`;
      const body = (await axios.get(url)).data.result;
      commit("entitySet", { type, body });
    },
    async accountUpdate({ state, commit }) {
      const url = `${API}/auth/accounts/${state.client.senderAddress}`;
      const acc = (await axios.get(url)).data;
      const account = acc.result.value;
      commit("accountUpdate", { account });
    },
    async entitySubmit({ state }, { type, body }) {
      const { chain_id } = state;
      const creator = state.client.senderAddress;
			console.log(creator)
      const base_req = { chain_id, from: creator };
      const req = { base_req, creator, ...body };
      const { data } = await axios.post(`${API}/${chain_id}/${type}`, req);
      const { msg, fee, memo } = data.value;
      return await state.client.signAndBroadcast(msg, fee, memo);
    },
    async entitySubmitFaucet({ state }, { type, body }) {
      const { chain_id } = state;
      const creator = state.client.senderAddress;
			console.log(creator)
      const base_req = { chain_id, from: creator };
      const req = { base_req, creator, ...body };
      const { data } = await axios.post(`${API}/faucet/${type}`, req);
      const { msg, fee, memo } = data.value;
      return await state.client.signAndBroadcast(msg, fee, memo);
    },
  },
});
