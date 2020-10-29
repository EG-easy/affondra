<template>
<transition name="modal">
  <div class="modal-mask">
    <div class="modal-wrapper" @click.self="$emit('close')">
      <div class="modal-container is-flex is-flex-direction-column has-text-black">
        <div class="is-flex is-flex-direction-row is-justify-content-center">
          <img id="loader" src="@/assets/loader.gif" alt="loader" width="64" height="64">
        </div>
        <div class="is-flex is-flex-direction-row is-justify-content-center">
          <h4 class="subtitle is-4 has-text-centered has-text-black ml-2 mt-2">{{message}}</h4>
        </div>
      </div>
    </div>
  </div>
</transition>
</template>

<style lang="scss">
#loader {
  width: max(10vh, 10vw);
  height: max(10vh, 10vw);
}
</style>

<script>
export default {
  props: {
    message: {
      type: String,
      require: false,
      default: 'proccessing...', 
    },
  },
}
</script>

<style lang="css">
.modal-mask {
  position: fixed;
  z-index: 9000;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(256, 256, 256, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
