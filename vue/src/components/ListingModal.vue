<template>
<transition name="modal">
  <div class="modal-mask">
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <!--<span class="tag is-primary">Enter Mnemonic</span>-->
          <h1 class="title has-text-black">Listing</h1>
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
        <div>{{ errorMessage }}</div>
        <div class="is-flex is-flex-direction-row my-4">
          <button :disabled="isLoading" class="button is-primary is-light" @click="strMnemonic = ''">
            <span>Clear</span>
          </button>
          <div :style="{'width':'10px'}"></div>
          <button :disabled="isLoading" class="is-flex-grow-1 button is-primary" @click="onListingClicked">
            <span class=" icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
              </svg>
            </span>
            <span>Login</span>
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
  props: {
    title: String,
    isVisible: Boolean,
    price: Number,
  },
  data() {
    return {
      isLoading: false,
      hasImage: false,
      image: null,
      errorMessage: "",
    }
  },
  methods: {
    getRandomStr: function (len) {
      return btoa(String.fromCharCode(...crypto.getRandomValues(new Uint8Array(len)))).substring(0, len)
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
