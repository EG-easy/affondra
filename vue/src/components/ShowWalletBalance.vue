<template>
<span>
  You have {{amountAffondollar}} affondollars.
</span>
</template>

<script>
export default {
  computed: {
    account: function () {
      return this.$store.state.account;
    },
    amountAffondollar: function () {
      return ('coins' in this.account ? this.account.coins.filter(v => v.denom === 'affondollar').map(v => parseInt(v.amount, 10)).concat([0]).reduce((a, x) => a += x, 0) : '---');
    },
  }
};
</script>
