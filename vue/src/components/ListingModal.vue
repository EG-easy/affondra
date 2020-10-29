<template>
<transition name="modal">
  <div class="modal-mask">
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
          <h4 class="subtitle is-4 has-text-black">2. Set title</h4>
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
          <h4 class="subtitle is-4 has-text-black">3. Set Price</h4>
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
        <div>{{ errorMessage }}</div>
        <div class="is-flex is-flex-direction-row my-4">
          <!--<button :disabled="isLoading" class="button is-primary is-light" @click="onClearClicked">
            <span>Clear</span>
          </button>
          <div :style="{'width':'10px'}"></div>-->
          <button :disabled="isLoading" class="is-flex-grow-1 button is-primary" @click="onListingClicked">
            <span class=" icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
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

<script>
import ImageUploader from 'vue-image-upload-resize'

export default {
  components: {
    ImageUploader,
  },
  data() {
    return {
      isLoading: false,
      hasImage: false,
      image: null,
      errorMessage: "",
      price: 100,
      title: "",
    }
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
      console.log(this.$_firebaseStorage)
      let ref = this.$_firebaseStorage.ref().child(`images/${this.getRandomStr(30)}.jpg`);
      ref.putString(this.image.dataUrl, 'data_url').then(function () {
        console.log('Uploaded a blob or file!');
      });
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
