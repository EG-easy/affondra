<template>
<div class="container">
  <ListingModal v-if="isShowListingModal" @close="isShowListingModal = false;refreshItems();" />
  <BuyModal v-bind="filterdItemsByRegExp[indexSelectedItem]" v-if="isShowBuyModal" @close="isShowBuyModal=false;refreshItems()" />
  <div class="is-flex is-flex-direction-column pa-2">
    <div class="columns is-multiline is-1">
      <div class="column is-flex is-flex-direction-row">
        <div class="field has-addons" :style="{'margin-bottom':'0'}">
          <p class="control">
            <button class="button is-small" :class="{'is-primary':selectedSellingStatusFilter === 'sale'}" @click="selectedSellingStatusFilter='sale'">
              <span>SALE</span>
            </button>
          </p>
          <p class="control">
            <button class="button is-small" :class="{'is-primary':selectedSellingStatusFilter === 'both'}" @click="selectedSellingStatusFilter='both'">
              <span>BOTH</span>
            </button>
          </p>
          <p class="control">
            <button class="button is-small" :class="{'is-primary':selectedSellingStatusFilter === 'sold'}" @click="selectedSellingStatusFilter='sold'">
              <span>SOLD</span>
            </button>
          </p>
        </div>
        <div class="is-flex-grow-1 is-flex is-flex-direction-row mx-2">
          <!--
          <div class="has-background-dark is-flex-grow-1">a</div>
          <div class="has-background-dark is-flex-grow-1">a</div>
          <div class="has-background-dark is-flex-grow-1">a</div>
          <div class="has-background-dark is-flex-grow-1">a</div>
          <div class="has-background-dark is-flex-grow-1">a</div>
          <div class="has-background-dark is-flex-grow-1">a</div>
          -->
        </div>
        <div class="control has-icons-left is-align-self-center">
          <input v-model="searchString" class="input is-small" type="email" placeholder="Search">
          <span class="icon is-small is-left">
            <i class="fa fa-search"></i>
          </span>
        </div>
      </div>
      <div class="column is-12" :style="{'padding':'10px 0.75rem 5px 0.75rem'}">
        <div :style="{'border-top':'1px solid #aaa'}">
          <strong class="has-text-primary">{{`${filterdItemsByRegExp.length} results for seraching.`}}</strong>
        </div>
      </div>
      <div v-for="(item, index) in filterdItemsByRegExp" :key="index" class="column is-3">
        <ItemThumbnail v-bind="item" @click="indexSelectedItem=index;isShowBuyModal=true;" />
      </div>
      <div v-if="filterdItemsByRegExp.length < 1" class="column is-12 has-text-centered" :style="{'flex-basis':'100%'}">
        <button class="button is-primary is-light" @click="clearFilter">
          No result. Click here to clear filters.
        </button>
      </div>
    </div>
  </div>
  <div id="listing-button" :disabled="!$store.getters.isLoggedIn" class="has-text-centered" @click="isShowListingModal = isLoggedIn">
    <!--<span id="listing-button-text" class="mr-2" :style="{'font-size':'30px', 'line-height':'90px'}">List new item!</span>-->
    <span :style="{'font-size':'60px', 'line-height':'74px'}">+</span>
  </div>
</div>
</template>

<style lang="scss" scoped>
/*
div.af-items {
  flex-basis: 25.00%;
  padding: 10px;
}
*/
#listing-button {
  cursor: pointer;
  position: fixed;
  bottom: 5vh;
  right: 5vh;
  width: 80px;
  height: 80px;
  border-radius: 40px;
  background-color: rgb(13, 41, 163);
  text-align: center;
  line-height: 74px;
  font-size: 60px;
  color: #fff;
  box-shadow: 0 0 25px 0 rgba(256, 256, 256, 1);
  transition: all 0.2s;
  transform: scale(1, 1);
  user-select: none;
}

#listing-button:hover {
  background-color: rgb(39, 61, 160);
  transform: scale(1.1, 1.1);
  /*width: 300px;*/
  /*margin-right: 10px;*/
}

#listing-button[disabled] {
  cursor: no-drop;
  background-color: rgb(209, 209, 209);
  color: rgb(179, 179, 179);
}

#listing-button[disabled]:hover {
  transform: scale(1, 1);
}
</style>

<script>
import ItemThumbnail from '@/components/ItemThumbnail.vue'
import ListingModal from '@/components/ListingModal.vue'
import BuyModal from '@/components/BuyModal.vue'

export default {
  components: {
    ItemThumbnail,
    ListingModal,
    BuyModal,
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    },
    filterdItemsByProps: function () {
      return this.items.filter((v) => {
        if (this.selectedSellingStatusFilter === 'sale' && !v.inSale) return false;
        if (this.selectedSellingStatusFilter === 'sold' && v.inSale) return false;
        return true
      });
    },
    filterdItemsByRegExp: function () {
      const searchString = this.searchString.trim();
      if (searchString === '') return this.filterdItemsByProps;
      const pattern = new RegExp(searchString, 'i');
      return this.filterdItemsByProps.filter((v) => pattern.test(v.title));
    }
  },
  mounted: function () {
    this.refreshItems();
    console.log(this.items[this.indexSelectedItem]);
  },
  data() {
    return {
      selectedSellingStatusFilter: 'both', // enum['sale', 'both', 'sold']
      isShowBuyModal: false,
      indexSelectedItem: 0,
      isShowListingModal: false,
      searchString: "",
      items: [{
        title: "title1",
        imageUrl: "https://bulma.io/images/placeholders/256x256.png",
        price: "100$",
        visible: true,
        denom: 'example',
      }]
    }
  },
  methods: {
    clearFilter: function () {
      this.searchString = "";
    },
    refreshItems: async function () {
      const storageRef = this.$_firebaseStorage.ref();
      const {
        result
      } = await this.$store.dispatch("getItemList")

      const prepareItem = function (v) {
        return new Promise((resolve, reject) => {
          (async () => {
            const tokenUri = JSON.parse(v.token_uri);
            return {
              title: tokenUri.name,
              denom: v.denom,
              id: v.id,
              inSale: v.inSale,
              description: v.description,
              imageUrl: (await storageRef.child(tokenUri.imgurl).getDownloadURL()),
              price: `${v.price.amount} ${v.price.denom}`,
              visible: true,
            }
          })().then((res) => resolve(res)).catch(err => reject(err));
        })
      }

      this.items = await Promise.all(result.map(prepareItem));

      console.log('Buy>refreshItems>items', this.items);
    }
  }
};
</script>
