<template>
<div class="item-thumbnail is-flex is-flex-direction-column box pa-0"
  @click="$emit('click')">
    
  <div class="ribon">
    <figure class="image is-4by3">
      <img draggable="false" :src="imageUrl">
    </figure>
    <div v-if="!inSale" class="caption"><span>SOLD</span></div>
  </div>
  <div class="is-flex is-flex-direction-row my-1">
    <span class="tag is-primary">{{denom}}</span>
  </div>
  <div class="is-flex is-flex-direction-row">
    <div class="has-text-centered">{{title}}</div>
    <div class="is-flex-grow-1"></div>
    <div class="has-text-centered">{{price}}</div>
  </div>
</div>
</template>

<style lang="scss">

.item-thumbnail{
  cursor: pointer;
  user-select: none;
  transition: all 0.2s;
}

.item-thumbnail:hover {
  transform:scale(1.05,1.05);
}

//CSS
.ribon {
  position: relative;
  display: inline-block;
  overflow: hidden;
}

.ribon .caption {
  display: inline;
}

.ribon span {
  position: absolute;
  display: inline-block;
  top: 7px;
  left: -27px;
  width: 100px;
  padding: 2px 10px;
  background: #cc0c0c;
  color: #fff;
  text-align: center;
  transform: rotate(-45deg);
}

</style>

<script>
export default {
  props: {
    inSale: Boolean,
    title: String,
    isVisible: Boolean,
    price: String,
    denom: String,
    imageUrl:String,
    description: String,
  },
};
</script>
