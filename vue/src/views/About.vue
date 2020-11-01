<template>
<div class="container">
  <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
  <div class="columns is-multiline">
    <div class="column is-12">
      <img src="@/assets/presentation/1.jpg" alt="presentation page 1">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/2.jpg" alt="presentation page 2">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/3.jpg" alt="presentation page 3">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/4.jpg" alt="presentation page 4">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/5.jpg" alt="presentation page 5">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/6.jpg" alt="presentation page 6">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/7.jpg" alt="presentation page 7">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/8.jpg" alt="presentation page 8">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/9.jpg" alt="presentation page 9">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/10.jpg" alt="presentation page 10">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/11.jpg" alt="presentation page 11">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/12.jpg" alt="presentation page 12">
    </div>
    <div class="column is-12">
      <img src="@/assets/presentation/13.jpg" alt="presentation page 13">
    </div>
    <div class="column is-12 has-text-centered">
      <button :disabled="!isLoggedIn && notGotTokenYet" class="button is-rounded">
        <span class="icon">
          <svg style="width:24px;height:24px" viewBox="0 0 24 24">
            <path fill="currentColor" d="M2 12C2 9.2 3.6 6.8 6 5.7V3.5C2.5 4.8 0 8.1 0 12S2.5 19.2 6 20.5V18.3C3.6 17.2 2 14.8 2 12M15 3C10 3 6 7 6 12S10 21 15 21 24 17 24 12 20 3 15 3M20 13H16V17H14V13H10V11H14V7H16V11H20V13Z" />
          </svg>
        </span>
        <span>Get Tokens</span>
      </button>
    </div>
    <div class="column is-12 has-text-centered">
      <span v-if="!isLoggedIn">
        Please Login first for get token!
      </span>
    </div>
  </div>
</div>  
</template>

<script>
import LoaderModal from '@/components/LoaderModal.vue'

export default {
  components: {
    LoaderModal,
  },
  data() {
    return {
      isLoading: false,
      strLeaderMessage: 'Processing...',
      notGotTokenYet: true,
    }
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
  }
}
</script>