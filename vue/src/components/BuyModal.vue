<template>
<transition name="modal">
  <div class="modal-mask">
    <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <h1 class="title has-text-black">Buy</h1>
        </div>
        <div class="box is-flex is-flex-direction-column">
          <figure class="image is-4by3">
            <img draggable="false" :src="imageUrl">
          </figure>
          <div class="is-flex is-flex-direction-row my-1">
            <span class="tag is-primary">{{denom}}</span>
          </div>
          <div class="is-flex is-flex-direction-row">
            <div class="has-text-centered">{{title}}</div>
            <div class="is-flex-grow-1"></div>
            <div class="has-text-centered">{{price}}</div>
          </div>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">Enter password to proceed</h4>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input v-model="password" class="input" type="password" placeholder="password">
            <span class="icon is-small is-left">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M7,14A2,2 0 0,1 5,12A2,2 0 0,1 7,10A2,2 0 0,1 9,12A2,2 0 0,1 7,14M12.65,10C11.83,7.67 9.61,6 7,6A6,6 0 0,0 1,12A6,6 0 0,0 7,18C9.61,18 11.83,16.33 12.65,14H17V18H21V14H23V10H12.65Z" />
              </svg>
            </span>
          </p>
        </div>
        <div>{{ errorMessage }}</div>
        <div class="is-flex is-flex-direction-row my-4">
          <!--<button :disabled="isLoading" class="button is-primary is-light" @click="onClearClicked">
            <span>Clear</span>
          </button>
          <div :style="{'width':'10px'}"></div>-->
          <button :disabled="!isInputValid" class="is-flex-grow-1 button is-primary" @click="onBuyClicked">
            <span class="icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M5.5,21C4.72,21 4.04,20.55 3.71,19.9V19.9L1.1,10.44L1,10A1,1 0 0,1 2,9H6.58L11.18,2.43C11.36,2.17 11.66,2 12,2C12.34,2 12.65,2.17 12.83,2.44L17.42,9H22A1,1 0 0,1 23,10L22.96,10.29L20.29,19.9C19.96,20.55 19.28,21 18.5,21H5.5M12,4.74L9,9H15L12,4.74M12,13A2,2 0 0,0 10,15A2,2 0 0,0 12,17A2,2 0 0,0 14,15A2,2 0 0,0 12,13Z" />
              </svg>
            </span>
            <span>Buy</span>
          </button>
        </div>
        <progress :style="{'visibility':isLoading ? 'visible' : 'hidden'}" class="progress is-small is-primary" max="100"></progress>
      </div>
    </div>
  </div>
</transition>
</template>

<script>

export default {
  props: {
    title: String,
    isVisible: Boolean,
    price: String,
    denom: String,
    imageUrl: String,
    description: String,
    id: String,
  },
  data() {
    return {
      password: '',
      isLoading: false,
      strLeaderMessage: 'Processing...',
      errorMessage: '',
    }
  },
  computed: {
    isInputValid: function () {
      return (
        typeof this.password === 'string' && this.password.length > 0
      );
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
  },
  methods: {
    getRandomStr: function (len) {
      return Array(len).fill(0).map(() => Math.floor(Math.random() * Math.floor(62))).map(v => '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'.substr(v, 1)).join('')
    },
    async onBuyClicked() {
      this.isLoading = true;
      this.strLeaderMessage = 'Sending purchase transaction...'

      const resSendTxBuyItem = await this.$store.dispatch("sendTxBuyItem", {
        buyerAddress: this.address,
        id: this.id,
        addressIntroducedBy: 'cosmos1ffu8y09f8snw0730m4e4p668j88ds42jlpjenv',
      }).catch((e) => {
        console.error(e)
        this.errorMessage = e;
        this.strLeaderMessage = 'Error has occured. Purchase was cancelled.';
        setTimeout(() => {
          this.isLoading = false;
          this.$emit('close');
        }, 3000);
        throw new Error('PURCHASE_FAILED');
      })

      console.log('onListingClicked>resSendTxBuyItem', resSendTxBuyItem);

      this.strLeaderMessage = 'Done';
      setTimeout(() => {
        this.isLoading = false;
        this.$emit('close');
      }, 1000);

    },
  },
}
</script>

<style lang="css">
#fileInput {
  display: none;
}

.modal-mask {
  position: fixed;
  z-index: 1000;
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
  max-height: 80vh;
  overflow: auto;
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
