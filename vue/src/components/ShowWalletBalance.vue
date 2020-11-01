<template>
<span class="tag is-primary is-light">
  You have {{amountAffondollar === null ? '---' : amountAffondollar }} affondollars.
</span>
</template>

<script>
export default {
  computed: {
    amountAffondollar() {
      return this.$store.getters.amountAffondollar
    }
  },
};
</script>
