<template>
<transition name="modal">
  <div class="modal-mask">
    <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <h1 class="title has-text-black">Login</h1>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">1. Set mnumonic</h4>
        </div>
        <div class="box is-flex is-flex-direction-column" :style="{'border': '#000 solid 1px'}">
          <v-checkbox v-model="ischeckedGenerateNewMnumonic" color="indigo">
            <template v-slot:label>
              <span class="is-size-5">Generate new mnumonic</span>
            </template>
          </v-checkbox>
          <div class="has-text-left mb-4">
            <h3 class="subtitle is-4 has-text-black has-text-centered">OR</h3 >
          </div>
          <div class="field my-0">
            <div class="control">
              <textarea v-model="strMnemonic" :style="{'border': '#ddd solid 1px'}" :disabled="ischeckedGenerateNewMnumonic" class="textarea is-primary" placeholder="Enter your nemonic here"></textarea>
            </div>
          </div>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">2. Set password</h4>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input v-model="password" class="input" type="password" placeholder="Password">
            <span class="icon is-small is-left">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M7,14A2,2 0 0,1 5,12A2,2 0 0,1 7,10A2,2 0 0,1 9,12A2,2 0 0,1 7,14M12.65,10C11.83,7.67 9.61,6 7,6A6,6 0 0,0 1,12A6,6 0 0,0 7,18C9.61,18 11.83,16.33 12.65,14H17V18H21V14H23V10H12.65Z" />
              </svg>
            </span>
          </p>
        </div>
        <div>{{ errorMessage }}</div>
        <button :disabled="!isValid" class="button is-primary my-4" @click="onLoginClicked();">
          <span class="icon">
            <svg style="width:24px;height:24px" viewBox="0 0 24 24">
              <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
            </svg>
          </span>
          <span>Login</span>
        </button>
      </div>
    </div>
  </div>
</transition>
</template>

<script>
import * as bip39 from "bip39";
import axios from "axios";
import store from "store";
import { Secp256k1HdWallet, SigningCosmosClient } from "@cosmjs/launchpad";

const API = "https://api.affondra.com";
const API_FIREBASE = 'https://asia-northeast1-affondra.cloudfunctions.net';
axios.defaults.headers.common['Content-Type'] = 'application/json;charset=utf-8'; 
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';

export default {
  props: {
    title: String,
    isVisible: Boolean,
    price: Number,
  },
  data() {
    return {
      ischeckedGenerateNewMnumonic: true,
      isLoading: false,
      strLeaderMessage: 'Processing...',
      strMnemonic: "",
      loggingIn: false,
      errorMessage: "",
      password: "",
    }
  },
  computed: {
    isValid() {
      return (this.ischeckedGenerateNewMnumonic || bip39.validateMnemonic(this.strMnemonic)) && this.password.length >= 4;
    },
  },
  methods: {
    sleep (millsec) {
      return new Promise(resolve => { setTimeout(() => { resolve(0) }, millsec)})
    },	
    getRandomStr: function (len) {
      return Array(len).fill(0).map(() => Math.floor(Math.random() * Math.floor(62))).map(v => '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'.substr(v, 1)).join('')
    },
    downloadText: function (content, filename) {
      // referrence: https://memo.appri.me/programming/file-download-by-js
      // create a temporary "a" element.
      const a = document.createElement("a");
      document.body.appendChild(a);
      a.style = "display:none";
      const blob = new Blob([content], { type: "octet/stream" });
      const url = window.URL.createObjectURL(blob);
      a.href = url;
      a.download = filename;
      a.click();
      window.URL.revokeObjectURL(url); // release the used object.
      a.parentNode.removeChild(a); // delete the temporary "a" element
    },
    async onLoginClicked () {
      if (this.ischeckedGenerateNewMnumonic) {
        await this.generatedMnemonicAndLogin().catch(e => {
          throw new Error (e);
        })
      } else {
        await this.loginWithMnemonic().catch(e => {
          throw new Error (e);
        })
      }
      //store password
      store.set('user', { serializedWallet: await this.$store.state.client.signer.serialize(this.password) });
      console.log(store.get('user'));
    },
    async generatedMnemonicAndLogin () {
      axios.defaults.headers.common['Content-Type'] = 'application/json;charset=utf-8'; 
      axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
      
      this.strLeaderMessage = 'Creating new mnemonic...';
      this.isLoading = true;
      await ( new Promise(resolve => { setTimeout(() => { resolve(0) }, 100)}) );
      const wallet = await Secp256k1HdWallet.generate(24).catch(e => {
        console.error(e);
        this.strLeaderMessage = 'Failed!!';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error (e);
        })
      });
      const newMnemonic = wallet.secret.data;

      this.strLeaderMessage = 'Creating mnemonic textfile for download...';
      await this.sleep(1000);
      this.downloadText(newMnemonic, `affondra-mnemonic-${this.getRandomStr(8)}.txt`);
      this.strLeaderMessage = 'Save mnemonic text and DO NOT lose it.';
      await this.sleep(5000);

      const [{ address }] = await wallet.getAccounts();

      this.strLeaderMessage = 'Sending 1000 affondollar to your wallet...';
      
      const txAxiosConfig = {
        baseURL: API_FIREBASE,
        method : 'POST',
        url    : `/affondra/faucet`,
        data   : {
          to: address,
        },
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Content-Type': 'application/json;charset=utf-8',
        },
      };

      console.log('onGenerateMnemonicClicked>getFaucet:', txAxiosConfig);

      const rawTx = await axios(txAxiosConfig).catch(e => {
        console.error(e);
        this.strLeaderMessage = 'Failed!!';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error (e);
        })
      });

      console.log('onGenerateMnemonicClicked>getFaucet>rawTx:', rawTx);

      this.strLeaderMessage = 'Logging in with new mnemonic...';
      await this.sleep(1000);

      // Mutate account and client
      const client = new SigningCosmosClient(API, address, wallet);
      const {
        data: {
          result: {
            value: account
          }
        }
      } = await axios.get(`${API}/auth/accounts/${address}`).catch(e => {
        console.error(e);
        this.strLeaderMessage = 'Failed!!';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error (e);
        })
      })
      this.$store.commit("accountUpdate", { account });
      this.$store.commit("clientUpdate", { client });

      await this.sleep(1000);

      this.strLeaderMessage = 'Done';
      await this.sleep(1000);
      this.isLoading = false;
      this.$emit('close');
    },
    async loginWithMnemonic () {
      
      this.strLeaderMessage = 'Logging in...';
      this.isLoading = true;
      await this.sleep(100);

      // Validate mnemonic
      let strMnemonic = this.strMnemonic;
      console.log('bip39.validateMnemonic:', bip39.validateMnemonic(strMnemonic));
      if (!bip39.validateMnemonic(strMnemonic)) {
        this.strLeaderMessage = 'Invalid mnemonic.';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error ('INVALID_MNEMONIC');
        })
      }

      // Sign in
      await this.$store.dispatch("accountSignIn", {
        mnemonic: strMnemonic,
      }).catch(e => {
        console.error(e);
        this.strLeaderMessage = 'Login Failed!!';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error(e);
        })
      })
      await this.sleep(1000);
      this.strLeaderMessage = 'Done';
      await this.sleep(1000);

      // Close loader and modal
      this.loggingIn = false;
      this.$emit('close');
    },
  },
}
</script>

<style lang="css">
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 400px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
