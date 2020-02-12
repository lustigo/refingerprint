<template >
  <v-app>
    <CheckMark v-if="isFinished" />

    
    <div id="survey" v-else-if="isLoaded">
        <Header />

        <v-content>
            <Id v-model="uid" />
            <Page v-model="structure.pages[navigationOptions.currentPage]" v-on:completed="completedHandler"/>
        </v-content>
       
        <Footer v-model="this.navigationOptions" v-on:send="send"/>
    </div>


    <p v-else-if="isError" style="margin-top:10%">Leider ist etwas schief gelaufen.</p>
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
  margin-left: 45%;
}
</style>

<script lang="ts">
import Vue from "vue";
import uuid from "uuid";
import CheckMark from "./components/CheckMark.vue";
import Footer from "./components/Footer.vue";
import Header from "./components/Header.vue";
import Id from "./components/Id.vue";
import Page from "./components/Page.vue";
import DataSender from "./controller/DataSender";
import SurveyDescription from "./interfaces/SurveyDescription";

/**
 * Top-Level Component
 */
export default Vue.extend({
    components: {
        CheckMark,
        Footer,
        Header,
        Id,
        Page
    },
    data: () => ({
        isError: false,
        isFinished: false,
        isLoaded: false,

        navigationOptions: {
            currentPage: 0,
            canBack: false,
            canForward: false,
            canSend: false,
            pageAmount: 0,
        },

        structure: {} as SurveyDescription,
        uid: uuid(),
    }),
    methods: {
        /**
         * is being called, when a page is completed.
         * activates the Forward-Button
         */
        completedHandler: function(completed: boolean){
            this.navigationOptions.canForward = completed;
            this.navigationOptions.canSend = this.navigationOptions.currentPage == this.navigationOptions.pageAmount-1 && completed;
        },
        send: async function(){
            const sender = new DataSender(this.structure, this.uid);
            if(await sender.send()){
                this.isFinished = true; 
            } else {
                this.isError = true; 
                this.isLoaded = false;
            } 
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
                fetch(`/api/${surveyId}`)
                    .then(async data => {
                        if (!data.body) {
                            this.isError = true;
                            return;
                        }
                        this.structure = await data.json();
                        this.isLoaded = true;
                        this.navigationOptions.pageAmount = this.structure.pages.length;
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
