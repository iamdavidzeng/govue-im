<template>
  <div class="container">
    <input
      type="text"
      v-model="username"
      placeholder="Username"
      class="input"
    />
    <button @click="initData(username)" class="btn">Login</button>
    <DatePicker></DatePicker>
  </div>
</template>

<script>
import { DatePicker } from "ant-design-vue";
import { mapActions } from "vuex";
export default {
  name: "Login",
  components: {
    DatePicker,
  },
  data() {
    return {
      username: "",
    };
  },
  methods: {
    ...mapActions(["initData"]),
  },
};
</script>

<style lang="less" scoped>
.container {
  margin: 30% auto;
  padding: 10px;
  width: 300px;
  height: 100px;
  background-color: #fff;
  border-radius: 5px;
}
.container .input {
  padding: 12px 0;
  margin: 10px;
  border-radius: 3px;
  border: 2px solid transparent;
  text-align: center;
  width: 50%;
  font-size: 16px;
  transition: border 0.2s, background-color 0.2s;
  background-color: #ecf0f1;
}
.container .btn {
  background-color: #3498db;
  color: #fff;
  line-height: 25px;
  cursor: pointer;
  width: 50%;
  text-align: center;
}
</style>
