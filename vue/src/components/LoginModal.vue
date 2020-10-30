<template>
<transition name="modal">
  <div class="modal-mask">
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <template v-if="true">
          <div class="has-text-left my-2">
            <!--<span class="tag is-primary">Enter Mnemonic</span>-->
            <h1 class="title has-text-black">Enter Mnemonic</h1>
          </div>
          <div class="field my-0">
            <div class="control">
              <textarea v-model="strMnemonic" :disabled="loggingIn" class="textarea is-primary" placeholder="Enter your nemonic here"></textarea>
            </div>
          </div>
          <div>{{ errorMessage }}</div>
          <div class="is-flex is-flex-direction-row my-4">
            <button :disabled="loggingIn" class="button is-primary is-light" @click="strMnemonic = ''">
              <span>Clear</span>
            </button>
            <div :style="{'width':'10px'}"></div>
            <button :disabled="loggingIn" class="is-flex-grow-1 button is-primary" @click="onLoginClicked">
              <span class="icon">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                  <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
                </svg>
              </span>
              <span>Login</span>
            </button>
          </div>
          <progress :style="{'visibility':loggingIn ? 'visible' : 'hidden'}" class="progress is-small is-primary" max="100"></progress>
        </template>
      </div>
    </div>
  </div>
</transition>
</template>

<script>
import * as bip39 from "bip39";

export default {
  props: {
    title: String,
    isVisible: Boolean,
    price: Number,
  },
  data() {
    return {
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
    async onLoginClicked() {
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
