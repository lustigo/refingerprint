<template >
  <v-app>
    <CheckMark v-if="isFinished" />

    
    <div id="survey" v-else-if="isLoaded">
        <Header />

        <v-content>
            <Id v-model="uid" />
            <Page v-model="structure.pages[currentPage]" v-on:completed="completedHandler"/>
        </v-content>
       
        <v-footer  absolute height="auto" class="font-weight-medium" >   
            <v-layout>
                <v-flex class="text-left">
                    <v-btn style="font-size:28px;" text v-bind:disabled="!canBack" v-on:click="pageBackward">
                        <v-icon>fa-arrow-left</v-icon>
                    </v-btn>
                </v-flex> 
                <v-flex class="text-right">
                    <v-btn style="font-size:28px;" text v-on:click="send" v-bind:disabled="!canSend" v-if="currentPage == structure.pages.length-1">
                        <v-icon>fa-paper-plane</v-icon>
                    </v-btn>
                    <v-btn style="font-size:28px;" text v-bind:disabled="!canForward" v-on:click="pageForward" v-else>
                        <v-icon>fa-arrow-right</v-icon>
                    </v-btn>
                </v-flex>
            </v-layout>
        </v-footer>
    </div>


    <p v-else-if="isError" class="center">Leider ist etwas schief gelaufen.</p>
    <v-progress-circular v-else size="64" indeterminate class="center" />
  </v-app>
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
import uuid from "uuid";
import Page from "./components/Page.vue";
import Id from "./components/Id.vue";
import Header from "./components/Header.vue";
import CheckMark from "./components/CheckMark.vue";
import SurveyDescription from "./interfaces/SurveyDescription";

/**
 * Top-Level Component
 */
export default Vue.extend({
    components: {
        Page,
        Id,
        Header,
        CheckMark
    },
    data: () => ({
        isLoaded: false,
        isError: false,
        canBack: false,
        canForward: false,
        canSend: false,
        isFinished: false,
        structure: {} as SurveyDescription,
        uid: uuid(),
        currentPage: 0
    }),
    methods: {
        /**
         * is being called, when a page is completed.
         * activates the Forward-Button
         */
        completedHandler: function(completed: boolean){
            this.canForward = completed;
            this.canSend = this.currentPage == this.structure.pages.length-1;
        },
        pageForward: function(){
            this.canForward = false;
            this.canBack = true;
            this.currentPage ++;
        },
        pageBackward: function(){
            this.canForward = true;
            this.currentPage --;
            this.canBack = this.currentPage > 0;
        },
        send: function(){
            //TODO
            this.isFinished = true;
        }
    },
    /**
   * Get the SurveyDescription from the server
   */
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
                        this.structure = await data.json();
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
