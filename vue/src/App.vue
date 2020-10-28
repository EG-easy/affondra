<template>
<section class="hero is-primary">
  <!-- Hero head: will stick at the top -->
  <div class="hero-head">
    <nav class="navbar">
      <div class="container">
        <div class="navbar-brand">
          <a class="navbar-item">
            <img src="https://bulma.io/images/bulma-type-white.png" alt="Logo">
          </a>
          <span class="navbar-burger burger" data-target="navbarMenuHeroA">
            <span></span>
            <span></span>
            <span></span>
          </span>
        </div>
        <div id="navbarMenuHeroA" class="navbar-menu">
          <div class="navbar-end">
            <a class="navbar-item">
              Buy
            </a>
            <a class="navbar-item">
              Listing
            </a>
            <a class="navbar-item">
              About
            </a>
            <span class="navbar-item">
              <v-btn :loading="false" :disabled="false" :class="button" @click="submit">
                Get Token
                <v-icon right dark>
                  mdi-login
                </v-icon>
              </v-btn>
            </span>
            <span class="navbar-item">
              <v-btn :loading="false" :disabled="false" :class="button" @click="alert('login')">
                Login
                <v-icon right dark>
                  mdi-login
                </v-icon>
              </v-btn>
            </span>
          </div>
        </div>
      </div>
    </nav>
  </div>

  <!-- Hero content: will be in the middle -->
  <div class="hero-body">
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
        <div class="column is-narrow">
          <a class="github-button" href="https://github.com/EG-easy" data-size="large" data-show-count="true" aria-label="Follow @EG-easy on GitHub">Follow @EG-easy</a>
        </div>
        <div class="column is-narrow">
          <a class="github-button" href="https://github.com/EG-easy/affondra/subscription" data-icon="octicon-eye" data-size="large" data-show-count="true" aria-label="Watch EG-easy/affondra on GitHub">Watch</a>
        </div>
        <div class="column is-narrow">
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

<style>
.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("init");
  },
  methods: {
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
