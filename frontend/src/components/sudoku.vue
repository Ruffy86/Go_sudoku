<template>
    <div class="container">
      <h1>
      <table class="table-bordered" style="margin-left: auto; margin-right: auto;">
        <tbody>
        <tr v-for="item in message" v-bind:key="item" style="border-color: white; border-style: solid;">
          <td v-for="article in item" v-bind:key="article" style="border-color: white; border-style: solid;">
              <label v-if="article!=0">{{article}}</label>
              <input type="number" v-if="article == 0" v-bind="userInput" v-model="userInput" style="height: 30px; width: 20px"/>
          </td>
        </tr>
        </tbody>
      </table>
    </h1>

    <a @click="getMessage">Empezar!</a>

        <span>{{ userInput | json }}</span>

    </div>
</template>

<script>
    var container = new Vue({
        el: '...',
        data: {
            userInput: []
        }
    })
export default {
  data() {
    return {
      message: [],
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
