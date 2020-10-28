<template>
<div class="container">
  <div class="is-flex is-flex-direction-column">
    <div class="is-flex is-flex-direction-row">
      <label class="checkbox is-align-self-center">
        <input type="checkbox">
        Now on sale
      </label>
      <div class="is-flex-grow-1"></div>
      <div class="control has-icons-left is-align-self-center">
        <input v-model="serachString" class="input is-small" type="email" placeholder="Search">
        <span class="icon is-small is-left">
          <i class="fa fa-search"></i>
        </span>
      </div>
    </div>
    <div :style="{'margin':'10px 0px 5px 0px','border-bottom':'1px solid #aaa'}"></div>
    <div class="is-flex is-flex-direction-row">
      <strong class="has-text-primary">{{`${filterdItems.length} results for seraching.`}}</strong>
      <div class="is-flex-grow-1"></div>
    </div>
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-center">
      <div v-for="(item, index) in filterdItems" :key="index" class="af-items">
        <ItemThumbnail v-bind="item" />
      </div>
      <div v-if="filterdItems.length < 1" class="has-text-centered" :style="{'flex-basis':'100%'}">
        <button class="button is-primary is-light" @click="clearFilter">
          No result. Click here to clear filters.
        </button>
      </div>
    </div>
  </div>
</div>
</template>

<style lang="scss" scoped>
div.af-items {
  flex-basis: 25.00%;
  padding: 10px;
}
</style>

<script>
import ItemThumbnail from '@/components/ItemThumbnail.vue'

export default {
  components: {
    ItemThumbnail
  },
  computed: {
    filterdItems: function () {
      return [...this.items, ...this.items, ...this.items, ...this.items, ...this.items].filter((v) => this.serachString.trim() === '' || v.title.indexOf(this.serachString.trim()) > -1);
    }
  },
  data() {
    return {
      serachString: "",
      items: [{
        title: "title1",
        price: "100$",
        visible: true,
      }, {
        title: "title2",
        price: "100$",
        visible: true,
      }, {
        title: "title3",
        price: "100$",
        visible: true,
      }, {
        title: "title4",
        price: "100$",
        visible: true,
      }]
    }
  },
  methods: {
    clearFilter: function () {
      this.serachString = "";
    }
  }
};
</script>
