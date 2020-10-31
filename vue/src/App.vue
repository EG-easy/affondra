<template>
<section class="hero is-primary" :style="{'min-height':'100vh'}">
  <LoginModal v-if="isShowLoginModal" @close="isShowLoginModal = false" />
  <!-- Hero head: will stick at the top -->
  <div class="hero-head">
    <nav class="navbar">
      <v-snackbar top v-model="isShowSnackbar" :centered="false" color="#003C88" :timeout="2000"> <div class="has-text-centered is-size-5">{{strSnackbar}}</div> </v-snackbar>
      <div class="container">
        <div class="navbar-brand">
          <a class="navbar-item">
            <img src="@/assets/affondra.logo.png" alt="Logo">
          </a>
          <span class="navbar-burger burger" data-target="navbarMenuHeroA" @click="isNavOpen = !isNavOpen">
            <span></span>
            <span></span>
            <span></span>
          </span>
        </div>
        <div id="navbarMenuHeroA" class="navbar-menu" :class="{'is-active':isNavOpen}">
          <div class="navbar-end">
            <div v-if="isLoggedIn" class="navbar-item af-address">
              <div class="is-flex is-flex-direction-row">
                <span class="tag is-primary is-light" @click="copyToClipboard();" :style="{'user-select':'none'}">
                  {{address}}
                  <span class="icon is-small ml-1 has-text-primary">
                    <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                      <path fill="currentColor" d="M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z" />
                    </svg>
                  </span>
                </span>
                <input id="clipBoard" type="text" :value="address" :style="{height:'0'}" />
              </div>
            </div>
            <div v-if="isLoggedIn" class="navbar-item has-text-right">
              <ShowWalletBalance />
            </div>
            <!-- 
            <div v-if="isLoggedIn" class="navbar-item">
              <button class="button is-rounded">
                <span class="icon">
                  <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                    <path fill="currentColor" d="M2 12C2 9.2 3.6 6.8 6 5.7V3.5C2.5 4.8 0 8.1 0 12S2.5 19.2 6 20.5V18.3C3.6 17.2 2 14.8 2 12M15 3C10 3 6 7 6 12S10 21 15 21 24 17 24 12 20 3 15 3M20 13H16V17H14V13H10V11H14V7H16V11H20V13Z" />
                  </svg>
                </span>
                <span>Get Tokens</span>
              </button>
            </div>
            -->
            <div v-if="!isLoggedIn" class="navbar-item has-text-right">
              <button class="button is-rounded" @click="isShowLoginModal = true">
                <span class="icon">
                  <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                    <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
                  </svg>
                </span>
                <span>Login</span>
              </button>
            </div>
            <div v-if="isLoggedIn" class="navbar-item has-text-right">
              <button class="button is-rounded" @click="$store.dispatch('logout');showSnackbar('Logout successfully.');">
                <span class="icon">
                  <svg style="width:24px;height:24px" viewBox="0 0 24 24">
                    <path fill="currentColor" d="M10,17V14H3V10H10V7L15,12L10,17M10,2H19A2,2 0 0,1 21,4V20A2,2 0 0,1 19,22H10A2,2 0 0,1 8,20V18H10V20H19V4H10V6H8V4A2,2 0 0,1 10,2Z" />
                  </svg>
                </span>
                <span>Logout</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </div>

  <!-- Hero content: will be in the middle -->
  <div class="hero-body" :style="{'padding-top':'1rem'}">
    <div class="container has-background-white has-text-black">
      <router-view />
    </div>
  </div>

  <!-- Hero footer: will stick at the bottom -->
  <div class="hero-foot">
    <div class="container has-background-white" :style="{'padding':'6px 6px 0px 6px'}">
      <div class="columns is-gapless">
        <div class="column">
          <div class="content has-text-centered has-text-primary">
            <p>
              <strong>Affondra</strong> by <a href="https://github.com/EG-easy"><strong>EG-easy</strong></a> and <a herf="https://github.com/YukiKAJIHARA">YukiKAJIHARA</a> . The source code is licensed under the
              <a href="http://opensource.org/licenses/mit-license.php"><strong>MIT</strong></a>.
            </p>
          </div>
        </div>
        <div class="column is-narrow has-text-centered">
          <a class="github-button" href="https://github.com/EG-easy" data-size="large" data-show-count="true" aria-label="Follow @EG-easy on GitHub">Follow @EG-easy</a>
          <a class="github-button" href="https://github.com/EG-easy/affondra/subscription" data-icon="octicon-eye" data-size="large" data-show-count="true" aria-label="Watch EG-easy/affondra on GitHub">Watch</a>
          <a class="github-button" href="https://github.com/EG-easy/affondra" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star EG-easy/affondra on GitHub">Star</a>
        </div>
      </div>
    </div>
  </div>
</section>
</template>

<style lang="scss">
@charset "utf-8";

// Import a Google Font
@import url('https://fonts.googleapis.com/css?family=Nunito:400,700');

// Reduce radius of corners
$radius-small: 1px;
$radius: 2px;
$radius-large: 3px;

// Reduce gaps
$gap: 0px;

// Set your brand colors
$blue: #003C88;
$purple: #621894;
$brown: #757763;
$beige-light: #D0D1CD;
$beige-lighter: #EFF0EB;

// Update Bulma's global variables
$family-sans-serif: "Nunito",
sans-serif;
$grey-dark: $brown;
$grey-light: $beige-light;
$primary: $blue;
$link: $purple;
$widescreen-enabled: false;
$fullhd-enabled: false;

@import "../node_modules/bulma/bulma.sass";
</style>

<style lang="scss" scoped>
.af-address{
  max-width: 90vw;
  overflow-x: hidden;
}

.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
}

section.hero.is-primary{
  overflow-y: none;
}
</style>

<script>
import LoginModal from '@/components/LoginModal.vue'
import ShowWalletBalance from '@/components/ShowWalletBalance.vue'

export default {
  components: {
    LoginModal,
    ShowWalletBalance,
  },
  created() {
    this.$store.dispatch("init");
  },
  data() {
    return {
      isShowLoginModal: false,
      strSnackbar: '',
      isShowSnackbar: false,
      isNavOpen: false,
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
  },
  methods: {
    showSnackbar (message) {
      this.strSnackbar = message;
      this.isShowSnackbar = true;
    },
    copyToClipboard () {
      const clipBoard = document.getElementById('clipBoard');
      console.log(clipBoard);
      clipBoard.value = this.address;
      clipBoard.select();
      document.execCommand("copy");
      clipBoard.blur();
      this.showSnackbar('Your address was copied.')
    },
    async submit() {
      const payload = {
        type: "mint",
        body: {
          title: this.title,
          options: this.options.map(o => o.title)
        }
      };
      await this.$store.dispatch("entitySubmitFaucet", payload);
      await this.$store.dispatch("entityFetch", payload);
      await this.$store.dispatch("accountUpdate");
    }
  }
};
</script>
