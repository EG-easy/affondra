<template>
<transition name="modal">
  <div class="modal-mask">
    <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div v-if="inSale" class="has-text-left my-2">
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
        <!--
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
        -->
        <template v-if="inSale">
          <div class="has-text-left my-2">
            <h4 class="subtitle is-4 has-text-black">Address whose introduce you</h4>
          </div>
          <div class="field">
            <p class="control has-icons-left has-icons-right">
              <input v-model="addressIntroducedBy" class="input" type="text" placeholder="Address whose introduce you">
              <span class="icon is-small is-left">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                  <path fill="currentColor" d="M13,17V19H14A1,1 0 0,1 15,20H22V22H15A1,1 0 0,1 14,23H10A1,1 0 0,1 9,22H2V20H9A1,1 0 0,1 10,19H11V17H5V15.5C5,13.57 8.13,12 12,12C15.87,12 19,13.57 19,15.5V17H13M12,3A3.5,3.5 0 0,1 15.5,6.5A3.5,3.5 0 0,1 12,10A3.5,3.5 0 0,1 8.5,6.5A3.5,3.5 0 0,1 12,3Z" />
                </svg>
              </span>
              <span v-if="isValidAddress" class="icon is-small is-right">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24" :style="{'color':'green'}">
                  <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z" />
                </svg>
              </span>
              <span v-if="addressIntroducedBy.length > 0 && !isValidAddress" class="icon is-small is-right">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24" :style="{'color':'red'}">
                  <path fill="currentColor" d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z" />
                </svg>
              </span>
            </p>
          </div>
          <v-checkbox v-model="isConfirmPurchase" color="indigo">
            <template v-slot:label>
              <span class="is-size-6">
                Read the
                <span class="af-tip has-text-danger">
                  TERM
                  <div class="af-tip-description has-text-centered">Its DEMO so you will recieve only NFT🙃</div>
                </span>
                and confirm the purchase.</span>  
            </template>
          </v-checkbox>
          <div class="has-text-info">{{ errorMessage }}</div>
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
        </template>
      </div>
    </div>
  </div>
</transition>
</template>

<style lang="scss" scoped>
.af-tip{
    position: relative;
    cursor: pointer;
    display: inline-block;
}
.af-tip p{
    margin:0;
    padding:0;
}
.af-tip-description {
    display: none;
    position: absolute;
    padding: 10px;
    font-size: 12px;
    line-height: 1.6em;
    color: #fff;
    border-radius: 5px;
    background: #000;
    width: 250px;
}
.af-tip-description:before {
    content: "";
    position: absolute;
    top: 100%;
    left: 20%;
    border: 15px solid transparent;
    border-top: 15px solid #000;
    margin-left: -15px;
}
.af-tip:hover .af-tip-description{
    display: inline-block;
    top: -70px;
    left: -30px;
}
</style>

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
    inSale: Boolean,
  },
  data() {
    return {
      password: '',
      isLoading: false,
      strLeaderMessage: 'Processing...',
      errorMessage: '',
      isConfirmPurchase: false,
      addressIntroducedBy: '',
    }
  },
  computed: {
    isLoggedIn () {
      return this.$store.getters.isLoggedIn
    },
    isValidAddress () {
      return /^cosmos1[a-z0-9]{38}$/.test(this.addressIntroducedBy)
    },
    isInputValid: function () {
      return (
        this.isLoggedIn &&
        this.isConfirmPurchase &&
        (this.addressIntroducedBy === '' || this.isValidAddress)
      );
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
  },
  mounted () {
    if (!this.isLoggedIn) this.errorMessage = 'You need to login first for purchase.'
  },
  methods: {
    getRandomStr: function (len) {
      return Array(len).fill(0).map(() => Math.floor(Math.random() * Math.floor(62))).map(v => '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'.substr(v, 1)).join('')
    },
    async onBuyClicked() {
      if (!this.isLoggedIn) return null;
      this.isLoading = true;
      this.strLeaderMessage = 'Sending purchase transaction...'

      const resSendTxBuyItem = await this.$store.dispatch("sendTxBuyItem", {
        buyerAddress: this.address,
        id: this.id,
        addressIntroducedBy: this.isValidAddress ? this.addressIntroducedBy : 'cosmos1ffu8y09f8snw0730m4e4p668j88ds42jlpjenv',
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

<style lang="css" scoped>
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
