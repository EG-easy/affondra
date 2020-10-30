<template>
<transition name="modal">
  <div class="modal-mask">
    <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <h1 class="title has-text-black">Enter Mnemonic</h1>
        </div>
        <div class="field my-0">
          <div class="control">
            <textarea v-model="strMnemonic" class="textarea is-primary" placeholder="Enter your nemonic here"></textarea>
          </div>
        </div>
        <div>{{ errorMessage }}</div>
        <div class="is-flex is-flex-direction-row my-2">
          <button class="button is-primary is-light" @click="strMnemonic = ''">
            <span>Clear</span>
          </button>
          <div :style="{'width':'10px'}"></div>
          <button class="is-flex-grow-1 button is-primary" @click="onLoginClicked">
            <span class="icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
              </svg>
            </span>
            <span>Login</span>
          </button>
        </div>
        <div class="has-text-left my-4">
          <!--<span class="tag is-primary">Enter Mnemonic</span>-->
          <h3 class="subtitle is-4 has-text-black has-text-centered">OR</h3 >
        </div>
        <button class="button is-info" @click="isLoading = true;onGenerateMnemonicClicked();">
          <span class="icon">
            <svg style="width:24px;height:24px" viewBox="0 0 24 24">
              <path fill="currentColor" d="M6.5,3C8.46,3 10.13,4.25 10.74,6H22V9H18V12H15V9H10.74C10.13,10.75 8.46,12 6.5,12C4,12 2,10 2,7.5C2,5 4,3 6.5,3M6.5,6A1.5,1.5 0 0,0 5,7.5A1.5,1.5 0 0,0 6.5,9A1.5,1.5 0 0,0 8,7.5A1.5,1.5 0 0,0 6.5,6M8,17H11V14H13V17H16V19H13V22H11V19H8V17Z" />
            </svg>
          </span>
          <span>Generate new mnemonic</span>
        </button>
      </div>
    </div>
  </div>
</transition>
</template>

<script>
import * as bip39 from "bip39";
import axios from "axios";
import { Secp256k1HdWallet, SigningCosmosClient } from "@cosmjs/launchpad";

const API = "http://lcd.affondra.com:8888";

export default {
  props: {
    title: String,
    isVisible: Boolean,
    price: Number,
  },
  data() {
    return {
      isLoading: false,
      strLeaderMessage: 'Processing...',
      strMnemonic: "spawn legend husband noble snap echo tennis dream cable pottery gauge share record clump only nerve play cute heavy edge inner wine tragic gap",
      loggingIn: false,
      errorMessage: "",
    }
  },
  computed: {
    isMnemonicValid() {
      return bip39.validateMnemonic(this.passwordClean);
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
    async onGenerateMnemonicClicked () {
      this.strLeaderMessage = 'Creating new mnemonic...';
      this.isLoading = true;
      await ( new Promise(resolve => { setTimeout(() => { resolve(0) }, 100)}) );
      const wallet = await Secp256k1HdWallet.generate(24).catch(e => {
        console.error(e);
        throw new Error (e);
      });
      const newMnemonic = wallet.secret.data;

      this.strLeaderMessage = 'Creating mnemonic textfile for download...';
      await this.sleep(1000);
      this.downloadText(newMnemonic, `affondra-mnemonic-${this.getRandomStr(8)}.txt`);
      this.strLeaderMessage = 'Save mnemonic text and DO NOT lose it.';
      await this.sleep(5000);
      this.strLeaderMessage = 'Logging in with new mnemonic';
      await this.sleep(1000);

      // Mutate account and client
      const [{ address }] = await wallet.getAccounts();
      const client = new SigningCosmosClient(API, address, wallet);
      const {
        data: {
          result: {
            value: account
          }
        }
      } = await axios.get(`${API}/auth/accounts/${address}`).catch(e => {
        console.error(e);
        throw new Error (e);
      })
      this.$store.commit("accountUpdate", { account });
      this.$store.commit("clientUpdate", { client });
      this.strLeaderMessage = 'Done';
      await this.sleep(1000);
      this.isLoading = false;
      this.$emit('close');
    },
    async onLoginClicked () {
      this.loggingIn = true;
      let strMnemonic = this.strMnemonic;
      console.log('bip39.validateMnemonic:', bip39.validateMnemonic(strMnemonic));
      if (!bip39.validateMnemonic(strMnemonic)) {
        this.errorMessage = 'Invalid mnemonic.';
        // this.loggingIn = false;
        return null;
      }
      await this.$store.dispatch("accountSignIn", {
        mnemonic: strMnemonic,
      }).catch((e) => {
        console.error(e)
        this.errorMessage = 'Login failed.';
        this.loggingIn = false;
        throw new Error(e)
      })
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
