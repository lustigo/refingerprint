<template>
  <span id="userId">
    {{ id }}
  </span>
</template>

<style>
#userId {
  visibility: hidden;
  display: none;
}
</style>

<script lang="ts">
import Vue from "vue";
import { UUID } from "../interfaces/UUID";
export default Vue.extend({
  props: ["value"],
  data: () => ({
    id: "" as UUID
  }),
  mounted() {
    this.id = this.value;
  }
});
</script>
