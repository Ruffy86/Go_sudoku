<template>
    <div class="container">
      <h1>
      <table class="table-bordered">
        <tbody>
        <tr v-for="item in message" v-bind:key="item">
          <td v-for="article in item" v-bind:key="article">
              <input type="text" value= {{article}} v-model="message">
          </td>
        </tr>
        </tbody>
      </table>
    </h1>
    <a @click="getMessage">Empezar!</a>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: []
    };
  },
  methods: {
    getMessage: function() {
      var self = this;

      window.backend.MyStruct.DisplayArray().then(mensaje => {
      self.message = mensaje;
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
    align-self: center;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
   .table-bordered{
       border-color: white;
   }
</style>
