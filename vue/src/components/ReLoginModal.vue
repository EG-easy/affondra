<template>
<transition name="modal">
  <div class="modal-mask">
    <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="has-text-left my-2">
          <h1 class="title has-text-black">Welcome back</h1>
        </div>
        <div class="has-text-left my-2">
          <h4 class="subtitle is-4 has-text-black">Enter password</h4>
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
        <button :disabled="!isInputValid" class="button is-primary my-4" @click="onLoginClicked();">
          <span class="icon">
            <svg style="width:24px;height:24px" viewBox="0 0 24 24">
              <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
            </svg>
          </span>
          <span>Login</span>
        </button>
        <button class="button my-4" @click="$emit('close')">
          <span>CANCEL</span>
        </button>
      </div>
    </div>
  </div>
</transition>
</template>

<script>
import LoaderModal from '@/components/LoaderModal.vue'

export default {
  components: {
    LoaderModal,
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
    isLoggedIn () {
      return this.$store.getters.isLoggedIn
    },
    isInputValid: function () {
      return typeof this.password === 'string' && this.password.length >= 4;
    }
  },
  methods: {
    sleep (millsec) {
      return new Promise(resolve => { setTimeout(() => { resolve(0) }, millsec)})
    },
    async onLoginClicked() {
      if (this.isLoggedIn) return null;
      if (!this.isInputValid) return null;

      this.isLoading = true;
      this.strLeaderMessage = 'Decrypting stored wallet...'

      await this.sleep(100);
      await this.$store.dispatch("accountSignInWithSerializedWallet", {
        password: this.password,
      }).catch(e => {
        console.error(e)
        this.errorMessage = e;
        this.strLeaderMessage = 'Relogin failed. Click "Login" and enter your mnumonic.';
        setTimeout(() => {
          this.isLoading = false;
          this.$emit('close');
        }, 3000);
        throw new Error('RELOGIN_FAILED');
      })
      this.strLeaderMessage = 'Logging in...';
      await this.sleep(1000);
      this.strLeaderMessage = 'Done';
      setTimeout(() => {
        this.isLoading = false;
        this.$emit('close');
      }, 1000);
    },
  },
}
</script>
