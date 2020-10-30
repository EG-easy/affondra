<template>
<div class="container">
  <ListingModal v-if="isShowListingModal" @close="isShowListingModal = false;refreshItems();" />
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
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-start">
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
  <div id="listing-button" :disabled="!$store.getters.isLoggedIn" class="has-text-centered" @click="isShowListingModal = ($store.getters.isLoggedIn ? true : false)">
    <!--<span id="listing-button-text" class="mr-2" :style="{'font-size':'30px', 'line-height':'90px'}">List new item!</span>-->
    <span :style="{'font-size':'60px', 'line-height':'74px'}">+</span>
  </div>
</div>
</template>

<style lang="scss" scoped>
div.af-items {
  flex-basis: 25.00%;
  padding: 10px;
}

#listing-button{
  cursor: pointer;
  position: fixed;
  bottom: 5vh;
  right: 5vh;
	width: 80px;
	height: 80px;
	border-radius: 40px;
  background-color: rgb(13, 41, 163);
  text-align:center;
  line-height: 74px;
  font-size: 60px;
  color: #fff;
  box-shadow: 0 0 25px 0 rgba(256, 256, 256, 1);
  transition: all 0.2s;
  transform:scale(1,1);
  user-select: none;
}
#listing-button:hover {
  background-color: rgb(39, 61, 160);
  transform:scale(1.1,1.1);
  /*width: 300px;*/
  /*margin-right: 10px;*/
}
#listing-button[disabled] {
  cursor: no-drop;
  background-color: rgb(209, 209, 209);
  color: rgb(179, 179, 179);
}
#listing-button[disabled]:hover {
  transform:scale(1,1);
}

</style>

<script>
import ItemThumbnail from '@/components/ItemThumbnail.vue'
import ListingModal from '@/components/ListingModal.vue'

export default {
  components: {
    ItemThumbnail,
    ListingModal,
  },
  computed: {
    filterdItems: function () {
      return this.items.filter((v) => this.serachString.trim() === '' || v.title.indexOf(this.serachString.trim()) > -1);
    }
  },
  mounted: function(){
    this.refreshItems();
  },
  data() {
    return {
      isShowListingModal: false,
      serachString: "",
      items: [{
        title: "title1",
        imageUrl: "https://bulma.io/images/placeholders/256x256.png",
        price: "100$",
        visible: true,
      }]
    }
  },
  methods: {
    clearFilter: function () {
      this.serachString = "";
    },
    refreshItems: async function () {
      const storageRef = this.$_firebaseStorage.ref();
      const { result } = await this.$store.dispatch("getItemList")
      this.items = [];
      result.forEach(async (v) => {
        const tokenUri = JSON.parse(v.token_uri);
        this.items.push({
          title: tokenUri.name,
          denom: v.denom,
          description: v.description,
          imageUrl: await storageRef.child(tokenUri.imgurl).getDownloadURL(),
          price: `${v.price.amount} ${v.price.denom}`,
          visible: true,
        })
      })
    }
  }
};
</script>
