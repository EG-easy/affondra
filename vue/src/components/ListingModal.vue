<template>
<transition name="modal">
  <div class="modal-mask">
  <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <h1 class="title has-text-black">Listing</h1>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">1. Pick image</h4>
        </div>
        <div class="box has-text-centered">
          <ImageUploader :preview="true" :className="['fileinput', { 'fileinput--loaded': hasImage }]" capture="environment" :debug="1" :maxWidth="512" :maxHeight="512" :quality="0.8" :autoRotate="true" outputFormat="verbose" @input="setImage">
            <label class="button is-primary is-light" for="fileInput" slot="upload-label">
              <span class="file-icon">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                  <path fill="currentColor" d="M5,17L9.5,11L13,15.5L15.5,12.5L19,17M20,6H12L10,4H4A2,2 0 0,0 2,6V18A2,2 0 0,0 4,20H20A2,2 0 0,0 22,18V8A2,2 0 0,0 20,6Z" />
                </svg>
              </span>
              <span class="file-label">
                {{ hasImage ? "Replace" : "Click to upload" }}
              </span>
            </label>
          </ImageUploader>
        </div>
        <!--<input class="is-flex-grow-1 input is-rounded my-2" type="text" placeholder="Title">-->
        <!--<input class="is-flex-grow-1 input is-rounded my-2" type="number" placeholder="price">-->
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">2. Set category (denom)</h4>
        </div>
        <div>
          <div class="control has-icons-left">
            <div class="select is-flex-grow-1">
              <select v-model="categorySelected">
                <option>Default category</option>
                <option>Game Items</option>
                <option>Characters</option>
                <option>Vegetables</option>
                <option>Goods</option>
              </select>
            </div>
            <span class="icon is-left">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M11,13.5V21.5H3V13.5H11M12,2L17.5,11H6.5L12,2M17.5,13C20,13 22,15 22,17.5C22,20 20,22 17.5,22C15,22 13,20 13,17.5C13,15 15,13 17.5,13Z" />
              </svg>
            </span>
          </div>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">3. Set title</h4>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input v-model="title" class="input" type="text" placeholder="Title">
            <span class="icon is-small is-left">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M20.71,7.04C21.1,6.65 21.1,6 20.71,5.63L18.37,3.29C18,2.9 17.35,2.9 16.96,3.29L15.12,5.12L18.87,8.87M3,17.25V21H6.75L17.81,9.93L14.06,6.18L3,17.25Z" />
              </svg>
            </span>
            <span class="icon is-small is-right">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z" />
              </svg>
            </span>
          </p>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">4. Set Price</h4>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input v-model="price" class="input" type="number" placeholder="Price">
            <span class="icon is-small is-left">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M5,6H23V18H5V6M14,9A3,3 0 0,1 17,12A3,3 0 0,1 14,15A3,3 0 0,1 11,12A3,3 0 0,1 14,9M9,8A2,2 0 0,1 7,10V14A2,2 0 0,1 9,16H19A2,2 0 0,1 21,14V10A2,2 0 0,1 19,8H9M1,10H3V20H19V22H1V10Z" />
              </svg>
            </span>
            <span class="icon is-small is-right">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z" />
              </svg>
            </span>
          </p>
        </div>
        <div class="field has-addons">
          <p class="control">
            <button class="button is-success is-small" @click="setPrice(price+10000)">
              +10000
            </button>
          </p>
          <p class="control">
            <button class="button is-success is-small" @click="setPrice(price+1000)">
              +1000
            </button>
          </p>
          <p class="control">
            <button class="button is-success is-small" @click="setPrice(price+100)">
              +100
            </button>
          </p>
          <p class="control">
            <button class="button is-danger is-small" @click="setPrice(price-100)">
              -100
            </button>
          </p>
          <p class="control">
            <button class="button is-danger is-small" @click="setPrice(price-1000)">
              -1000
            </button>
          </p>
          <p class="control">
            <button class="button is-danger is-small" @click="setPrice(price-10000)">
              -10000
            </button>
          </p>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">5. Set affiliate reward rate</h4>
        </div>
        <div class="my-6">
          <VueSlider v-model="affiliateRewardRate" :min="0" :max="100" :interval="10" :marks="{'0':'0%','20':'20%','40':'40%','60':'60%','80':'80%','100':'100%'}">
          </VueSlider>
        </div>
        <div>{{ errorMessage }}</div>
        <div class="is-flex is-flex-direction-row my-4">
          <!--<button :disabled="isLoading" class="button is-primary is-light" @click="onClearClicked">
            <span>Clear</span>
          </button>
          <div :style="{'width':'10px'}"></div>-->
          <button :disabled=" !isInputValid || isLoading" class="is-flex-grow-1 button is-primary" @click="onListingClicked">
            <span class=" icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M13,2.03C17.73,2.5 21.5,6.25 21.95,11C22.5,16.5 18.5,21.38 13,21.93V19.93C16.64,19.5 19.5,16.61 19.96,12.97C20.5,8.58 17.39,4.59 13,4.05V2.05L13,2.03M11,2.06V4.06C9.57,4.26 8.22,4.84 7.1,5.74L5.67,4.26C7.19,3 9.05,2.25 11,2.06M4.26,5.67L5.69,7.1C4.8,8.23 4.24,9.58 4.05,11H2.05C2.25,9.04 3,7.19 4.26,5.67M2.06,13H4.06C4.24,14.42 4.81,15.77 5.69,16.9L4.27,18.33C3.03,16.81 2.26,14.96 2.06,13M7.1,18.37C8.23,19.25 9.58,19.82 11,20V22C9.04,21.79 7.18,21 5.67,19.74L7.1,18.37M12,7.5L7.5,12H11V16H13V12H16.5L12,7.5Z" />
              </svg>
            </span>
            <span>Listing</span>
          </button>
        </div>
        <progress :style="{'visibility':isLoading ? 'visible' : 'hidden'}" class="progress is-small is-primary" max="100"></progress>
      </div>
    </div>
  </div>
</transition>
</template>

<style lang="scss">
.custom-mark {
  position: absolute;
  top: 10px;
  transform: translateX(-50%);
  white-space: nowrap;
}
</style>

<script>
// import axios from "axios";
import 'vue-slider-component/theme/default.css'

import ImageUploader from 'vue-image-upload-resize'
import VueSlider from 'vue-slider-component'
import LoaderModal from '@/components/LoaderModal.vue'

export default {
  components: {
    ImageUploader,
    LoaderModal,
    VueSlider,
  },
  data() {
    return {
      categorySelected: "Default category",
      isLoading: false,
      strLeaderMessage: 'Processing...',
      hasImage: false,
      image: null,
      errorMessage: "",
      price: 100,
      title: "My item",
      affiliateRewardRate: 10,
    }
  },
  computed: {
    isInputValid: function () {
      return (
        this.image !== null &&
        typeof this.title === 'string' && this.title.length > 0 &&
        typeof this.price === 'number' && this.price > 0
      );
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
  },
  methods: {
    setPrice: function (newV) {
      this.price = newV >= 100 ? newV : 100;
    },
    getRandomStr: function (len) {
      return Array(len).fill(0).map(() => Math.floor(Math.random() * Math.floor(62))).map(v => '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'.substr(v, 1)).join('')
    },
    setImage: function (output) {
      this.hasImage = true;
      this.image = output;
      console.log('Image selected:', output)
    },
    async onListingClicked() {
      this.isLoading = true;
      this.strLeaderMessage = 'Uploading image...'

      const imagePath = (await this.uploadImage()).metadata.fullPath;
      console.log('onListingClicked>imagePath:', imagePath);

      this.strLeaderMessage = 'Sending transaction...'

      const nftId = this.getRandomStr(32);

      const resSendTxMintNft = await this.$store.dispatch("sendTxMintNft", {
        address: this.address,
        denom: this.categorySelected,
        nftId: nftId,
        // id: '2iZaPwXbfdiCXRbuzmQLPVRZ5Z88tv2z',
        tokenURI: JSON.stringify({
          name: this.title,
          imgurl: imagePath,
        }),
      }).catch((e) => {
        console.error(e)
        this.errorMessage = e;
        this.strLeaderMessage = 'Error has occured. Listing was cancelled.';
        setTimeout(()=>{
          this.isLoading = false;
          this.$emit('close');
        }, 3000);
        throw new Error('LISTING_FAILED');
      })

      console.log('onListingClicked>resSendTxMintNft', resSendTxMintNft);

      this.strLeaderMessage = 'Creating your item from NFT...'      

      await this.$store.dispatch("sendTxCreateItem", {
        address: this.address,
        denom: this.categorySelected,
        nftId: nftId,
        price: `${this.price}affondollar`,
        affiliate: `${Math.floor(this.price * this.affiliateRewardRate / 100)}affondollar`,
        inSale: true,
        description: '*KZF*',
      }).catch((e) => {
        console.error(e)
        this.errorMessage = e;
        this.strLeaderMessage = 'Error has occured. Listing was cancelled.';
        setTimeout(()=>{
          this.isLoading = false;
          this.$emit('close');
        }, 3000);
        throw new Error('LISTING_FAILED');
      })

      this.strLeaderMessage = 'Done';
      setTimeout(()=>{
        this.isLoading = false;
        this.$emit('close');
      }, 1000);

    },
    async uploadImage() {
      // upload data
      let ref = this.$_firebaseStorage.ref().child(`images/${this.getRandomStr(30)}.jpg`);
      const res = await ref.putString(this.image.dataUrl, 'data_url');
      console.log('Uploading Image>res:', res);
      return res;
    },
    onClearClicked: function () {
      this.title = "";
      this.price = 100;
    }
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
