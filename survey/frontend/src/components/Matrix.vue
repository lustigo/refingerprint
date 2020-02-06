<template>
  <div>
    <p class="font-regular font-weight-light">{{ structure.question }}</p>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th scope>Frage</th>
            <th
              scope
              v-for="(answer, index) in structure.structure"
              class="text-left"
              :key="index"
            >
              {{ answer }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(_, att) in structure.answers" :key="att">
            <td>{{ att }}</td>
            <td v-for="(answer, key) in structure.structure" :key="key">
              <input
                class="v-input--selection-controls ripple"
                type="radio"
                :value="key"
                v-on:change="$emit('input', structure)"
                v-model="structure.answers[att]"
              />
            </td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MatrixDescription from "../interfaces/MatrixDescription";

/**
 * Matrix-Widget, so that a participant can answer multiple questions with the same answer-set
 */
export default Vue.extend({
  props: ["value", "required"],
  data: () => ({
    structure: {} as MatrixDescription,
    req: false
  }),
  mounted() {
    this.structure = this.value;
    this.req = this.required;
  }
});
</script>
