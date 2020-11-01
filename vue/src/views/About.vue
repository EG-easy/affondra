<template>
<div class="container has-background-white has-text-black" :style="{'padding':'1rem'}">
  <LoaderModal v-if="isLoading" :message="strLeaderMessage" />
  <div class="columns is-multiline">
    <div class="column is-12 has-text-centered">
      <h1 class="title has-text-black">Presentations</h1>
    </div>
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
    <div class="column is-12 has-text-centered is-primary is-flex is-flex-direction-column is-align-items-center">
      <div v-if="notGotTokenYet" class="box has-background-primary is-narrow is-flex is-flex-direction-column">
        <h1 class="title">Get 1000 affondollar prize!</h1>
        <div class="is-flex is-flex-direction-row">
          <div class="field">
            <p class="control has-icons-left has-icons-right">
              <input v-model="targetAddress" class="input" type="text" placeholder="Your address">
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
              <span v-if="targetAddress.length > 0 && !isValidAddress" class="icon is-small is-right">
                <svg style="width:24px;height:24px" viewBox="0 0 24 24" :style="{'color':'red'}">
                  <path fill="currentColor" d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z" />
                </svg>
              </span>
            </p>
          </div>
          <button :disabled="!isValidAddress" class="button is-rounded ml-4" @click="notGotTokenYet=false;faucet1000Affondollar();">
            <span class="icon">
              <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                <path fill="currentColor" d="M2 12C2 9.2 3.6 6.8 6 5.7V3.5C2.5 4.8 0 8.1 0 12S2.5 19.2 6 20.5V18.3C3.6 17.2 2 14.8 2 12M15 3C10 3 6 7 6 12S10 21 15 21 24 17 24 12 20 3 15 3M20 13H16V17H14V13H10V11H14V7H16V11H20V13Z" />
              </svg>
            </span>
            <span>Get Tokens</span>
          </button>
        </div>
      </div>
      <div v-else class="column is-12 has-text-centered">
        <h1 class="title has-text-black">Enjoy!</h1>
      </div>
    </div>
    <!--
    <div class="column is-12 has-text-centered">
      <span v-if="!isLoggedIn">
        Please Login first for get token!
      </span>
    </div>
    -->
  </div>
</div>  
</template>

<script>
import axios from "axios";

import LoaderModal from '@/components/LoaderModal.vue'

const API_FIREBASE = 'https://asia-northeast1-affondra.cloudfunctions.net';

export default {
  components: {
    LoaderModal,
  },
  data() {
    return {
      isLoading: false,
      strLeaderMessage: 'Processing...',
      notGotTokenYet: true,
      targetAddress: '',
    }
  },
  computed: {
    isValidAddress () {
      return /^cosmos1[a-z0-9]{38}$/.test(this.targetAddress)
    },
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    },
  },
  methods:{
    sleep (millsec) {
      return new Promise(resolve => { setTimeout(() => { resolve(0) }, millsec)})
    },
    async faucet1000Affondollar () {
      this.isLoading = true;
      this.strLeaderMessage = 'Thank you for reading through documents!!\nSending 1000 affondollar to your wallet...';
       
      const txAxiosConfig = {
        baseURL: API_FIREBASE,
        method : 'POST',
        url    : `/affondra/faucet`,
        data   : {
          to: this.targetAddress,
        },
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Content-Type': 'application/json;charset=utf-8',
        },
      };

      console.log('About>faucet1000Affondollar>txAxiosConfig:', txAxiosConfig);

      const rawTx = await axios(txAxiosConfig).catch(e => {
        console.error(e);
        this.strLeaderMessage = 'Failed!!';
        this.sleep(1000).then(() => {
          this.isLoading = false;
          throw new Error (e);
        })
      });
      
      console.log('About>faucet1000Affondollar>rawTx:', rawTx);

      this.strLeaderMessage = 'Done';
      await this.sleep(1000);
      this.isLoading = false;
      this.$emit('close');
    }
  },
}
</script>