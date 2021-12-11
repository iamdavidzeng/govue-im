<template>
  <el-main>
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Username">
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item label="Password">
        <el-input v-model="form.password" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Submit</el-button>
      </el-form-item>
    </el-form>
  </el-main>
</template>

<script>
import { mapActions } from "vuex";
export default {
  data() {
    return {
      form: {
        name: "",
        password: "",
      },
    };
  },
  methods: {
    ...mapActions(["initData"]),
    onSubmit() {
      if (!this.form.name || !this.form.password) {
        alert("Please input username and password");
        return;
      }
      this.initData(this.form.name);
    },
  },
};
</script>
