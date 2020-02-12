<template>
  <v-container>
    <v-container v-if="!structure.completed">
      <p>Wir möchten noch überprüfen, dass Sie kein Roboter sind <span style="color:red;">*</span>:</p>
      <div cols="3" style="margin-left:35%;" class="g-recaptcha" id="gotcha"></div>
    </v-container>
    <v-container v-else>
      <p>Sie sind kein Roboter 🤖</p>
    </v-container>
  </v-container>
</template>
<script lang="ts">
import Vue from "vue";
import CaptchaDescription from "../interfaces/CaptchaDescription";
declare const grecaptcha: any;
/**
 * Captcha-Widget
 * There is no need to set the required prop, it is always required.
 */
export default Vue.extend({
    props: ["value"],
    data: () => ({
        structure: {
            sitekey: "placeholder"
        } as CaptchaDescription
    }),
    methods: {
        captchaCreated: function(){
            const event = document.createEvent("Event");
            event.initEvent("captcha-created");
            document.getRootNode().dispatchEvent(event);
        },
        captchaSolved: function() {
            const event = document.createEvent("Event");
            event.initEvent("captcha-solved");
            document.getRootNode().dispatchEvent(event);
        }
    },
    mounted() {
        this.structure = this.value;
        if (!this.structure.completed) {
            this.structure.completed = false;
        }
        this.$emit("completed", this.structure.completed);

        if (!this.structure.completed) {
            // If the Captcha was not already solved, render it and notify extension
            this.captchaCreated();
            grecaptcha.render("gotcha", {
                sitekey: this.structure.sitekey,
                // After solving it, notify extension, verify it at the server and set the widget as completed
                callback: async (response: string) => {
                    this.captchaSolved();
                    const success = await fetch("https://localhost:8081/api/captcha", {
                        method: "POST",
                        body: JSON.stringify({
                            response: response
                        }),
                        headers: { "content-type": "application/json" }
                    });
                    this.structure.completed = success.status == 200;
                    this.$emit("completed", this.structure.completed);
                }
            });
        }
    }
});
</script>