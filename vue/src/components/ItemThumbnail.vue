<template>
<div class="is-flex is-flex-direction-column">
  <figure class="image is-3by2">
    <img src="https://bulma.io/images/placeholders/256x256.png">
  </figure>
  <div class="is-flex is-flex-direction-row">
    <div class="has-text-centered">{{title}}</div>
    <div class="is-flex-grow-1"></div>
    <div class="has-text-centered">{{price}}</div>
  </div>
</div>
</template>

<script>
export default {
  props: {
    title: String,
    isVisible: Boolean,
    price: Number,
  },
};
</script>
