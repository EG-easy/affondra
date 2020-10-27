<template>
<div>
  <router-view />
</div>
</template>

<style lang="scss">
@charset "utf-8";

// Import a Google Font
@import url('https://fonts.googleapis.com/css?family=Nunito:400,700');

// Reduce radius of corners
$radius-small: 1px;
$radius: 2px;
$radius-large: 3px;

// Set your brand colors
$purple: #8A4D76;
$pink: #FA7C91;
$brown: #757763;
$beige-light: #D0D1CD;
$beige-lighter: #EFF0EB;

// Update Bulma's global variables
$family-sans-serif: "Nunito",
sans-serif;
$grey-dark: $brown;
$grey-light: $beige-light;
$primary: $purple;
$link: $pink;
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
};
</script>
