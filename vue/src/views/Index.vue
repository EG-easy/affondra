<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
      <poll-form />
      <poll-list />
		<sp-type-form type="vote" :fields="['pollID', 'value', ]" />
		<sp-type-form type="poll" :fields="['title', 'options', ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
