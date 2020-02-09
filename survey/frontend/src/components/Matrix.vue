<template>
  <div>
    <p class="font-regular font-weight-light">
      {{ structure.question }}
      <span v-if="req" style="color:red;">*</span>
    </p>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th scope></th>
            <th
              scope
              v-for="(answer, index) in structure.structure"
              class="text-center"
              v-bind:key="answer+index"
            >{{ answer }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(_, att) in structure.answers" v-bind:key="att">
            <td class="text-center">{{ att }}</td>
            <td v-for="(answer, key) in structure.structure" v-bind:key="att+key" class="text-center">
              <input
                class="v-input--selection-controls ripple"
                type="radio"
                v-bind:value="key"
                v-on:change="onDataChange"
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
        req: false,
        structure: {} as MatrixDescription,
    }),
    methods: {
        /**
         * Checks if the Matrix is completly filled out
         */
        isFilled: function(){
            for (const answer in this.structure.answers){
                if (this.structure.answers[answer] == -1){
                    return false;
                }
            }
            return true;
        },
        /**
         * Will be called when a radio button is clicked
         */
        onDataChange: function (){
            this.$emit("input",this.structure);
            this.$emit("completed", !this.req || this.isFilled());
        }
    },
    mounted() {
        this.structure = this.value;
        this.req = this.required;
        if(!this.req || this.isFilled()){
            this.$emit("completed",true);
        }
    }
});
</script>
