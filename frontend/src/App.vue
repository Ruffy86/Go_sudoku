<template>
  <div id="app">
    <sudoku/>
  </div>
</template>

<script>
import sudoku from "./components/sudoku.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    sudoku
  }
};
</script>
