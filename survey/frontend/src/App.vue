<template>
  <div id="app">
    <Page v-if="isLoaded" />
    <p v-else-if="isError" class="center">Leider ist etwas schief gelaufen.</p>
    <v-progress-circular v-else size="64" indeterminate class="center" />
  </div>
</template>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.center {
  margin-top: 10%;
}
</style>

<script lang="ts">
import Vue from "vue";
import Page from "./components/Page.vue";
import SurveyDescription from "./interfaces/SurveyDescription";
export default Vue.extend({
  components: {
    Page
  },
  data: () => ({
    isLoaded: false,
    isError: false,
    structure: {}
  }),
  mounted() {
    /**
     * Regex for a UUIDv4
     * @src https://www.regextester.com/99148
     */
    const uuidRegex = new RegExp(
      "([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})"
    );

    // Get Survey Id
    if (uuidRegex.test(document.URL)) {
      const surveyIdRegexResult = uuidRegex.exec(document.URL);
      if (surveyIdRegexResult && surveyIdRegexResult.length > 1) {
        const surveyId = surveyIdRegexResult[1];

        //fetch the Surveydescription
        fetch(`https://localhost:8081/api/${surveyId}`)
          .then(async data => {
            if (!data.body) {
              this.isError = true;
              return;
            }
            this.structure = (await data.json()) as SurveyDescription;
            this.isLoaded = true;
          })
          .catch(err => {
            this.isError = true;
          });
      }
    } else {
      this.isError = true;
    }
  }
});
</script>
