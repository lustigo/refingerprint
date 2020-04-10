import { delay } from "../misc/Helper";
import frameListener from "../misc/FrameListener";
import { TaskData, TaskType } from "../interfaces/TaskData";
import Module from "../interfaces/Module";
import Task from "./Task";

/**
 * Captcha Timeout is 120 Seconds
 */
const CAPTCHA_TIMEOUT = 120;

/**
 * Module which listens to Actions on the TaskFrame
 */
export default class TaskListener implements Module {

    /**
     * Name of the Module
     */
    public name = "TaskEvents";

    /**
     * If the Listener is stopped
     */
    private stopped = false;

    /**
     * Reference to the TaskFrame
     */
    private taskFrame: HTMLIFrameElement | null = null;

    /**
     * All processed Tasks
     */
    private tasks: Array<Task> = [];

    /**
     * Reference to Timeout
     */
    private timeout = -42;

    /**
     * Will be called when the Captcha is rendered
     * Checks if the TaskFrame is opened, starts the first task and registers the Button Listeners
     */
    public start(): void {
        const taskFrame = frameListener.getTaskFrame();
        if (!(taskFrame && taskFrame.contentDocument && taskFrame.contentDocument.body)) {
            delay(10).then(() => this.start());
        } else {
            this.listenToOpen();
        }
    }

    /**
     * Will be called when the Captcha is solved
     * Stops the Task Timer
     */
    public stop(): void {
        this.stopped = true;
        this.stopTask();
        if (this.tasks.length != 0) {
            this.tasks[this.tasks.length - 1].solved = true;
        }
        window.clearInterval(this.timeout);
    }

    /**
     * Returns the collected TaskData
     * @return collected TaskData
     */
    public getCollectedData(): Array<TaskData> {
        return this.tasks.map(t => t.getData());
    }

    /**
     * Creates and Starts a new Task
     */
    private addTask(): void {
        if (this.taskFrame) {
            this.tasks.push(new Task(this.taskFrame));
            this.timeout = window.setTimeout(this.onTimeout.bind(this), CAPTCHA_TIMEOUT * 1000);
            this.registerListeners();
        }
    }

    /**
     * Checks if the Task failed
     */
    private checkError(): void {
        if(this.taskFrame && this.taskFrame.contentDocument){
            const errorDiv = this.taskFrame.contentDocument.getElementsByClassName("rc-imageselect-incorrect-response");
            if(errorDiv.length > 0){
                if((errorDiv[0] as HTMLDivElement).style.display != "none"){
                    this.tasks[this.tasks.length - 2].failed = true;

                    // Object Identification tasks in series are treated as one task, so if the last one failes, all failed
                    if(this.tasks[this.tasks.length -2].type == TaskType.OBJ){
                        for(let i = this.tasks.length-3; i >= 0; i--){
                            if(this.tasks[i].type == TaskType.OBJ){
                                this.tasks[i].failed = true;
                            }
                        }
                    }
                }
            }
        }
    }

    /**
     * Checks if the Verification failed and no new task was started
     */
    private hasVerifyError(): boolean {
        if (this.taskFrame && this.taskFrame.contentDocument) {
            const selectMore = this.taskFrame.contentDocument.getElementsByClassName("rc-imageselect-error-select-more");
            if(selectMore.length > 0){
                if((selectMore[0] as HTMLDivElement).style.display != "none"){
                    return true;
                }
            }

            const dynamicMore = this.taskFrame.contentDocument.getElementsByClassName("rc-imageselect-error-dynamic-more");
            if(dynamicMore.length > 0){
                if ((dynamicMore[0] as HTMLDivElement).style.display != "none") {
                    this.tasks[this.tasks.length -1].type = TaskType.DYN;
                    return true;
                }
            }
        }
        return false;
    }

    /**
     * Waits until the TaskFrame is opened
     */
    private async listenToOpen(): Promise<void> {
        while (!this.stopped) {
            await delay(10);
            const taskFrame = frameListener.getTaskFrame();
            if (taskFrame && taskFrame.contentDocument && taskFrame.contentDocument.body) {
                const token = taskFrame.contentDocument.getElementById("recaptcha-token");
                if (token && (token as HTMLInputElement).value != "") {
                    // TaskFrame opened
                    this.taskFrame = taskFrame;
                    this.addTask();
                    break;
                }
            }
        }
    }

    /**
     * Will be called when the Task is switched to Audio
     * Stops the Task and waits until the Captcha is solved or it is again switched to Image Mode
     */
    private onAudio(): void {
        window.clearInterval(this.timeout);
        this.stopTask();
        if (this.taskFrame && this.taskFrame.contentDocument) {
            const image = this.taskFrame.contentDocument.getElementById("recaptcha-image-button");
            if (image) {
                image.addEventListener("click", this.addTask.bind(this));
            }
        }
    }

    /**
     * Will be called when the Captcha is Reloaded
     * Stops the Task and starts a new one if the Captcha is not solved
     */
    private async onEvent(): Promise<void> {
        window.clearInterval(this.timeout);
        this.stopTask();
        // Wait if the Captcha is solved
        await delay(200);
        if (!(this.stopped || this.hasVerifyError())) {
            this.addTask();
            // Wait for an error message to appear
            await delay(250);
            this.checkError();
        } else if(this.hasVerifyError()) {
            // No new Task started -> renew listeners
            this.registerListeners();
        }
    }

    /**
     * Will be called when the Captcha is timed out
     * Stops the Task and waits for the TaskFrame to be reopened
     */
    private onTimeout(): void {
        this.stopTask();
        frameListener.registerCaptchaFrameClickListener(this.addTask, this);
    }

    /**
     * Will be called when the Captcha is verified
     * Checks if the Task failed and calls onEvent
     */
    private onVerify(): void {
        if (this.taskFrame) {
            const resp = this.taskFrame.getElementsByClassName("rc-imageselect-incorrect-response");
            if (resp[0] && (resp[0] as HTMLDivElement).style.display != "none") {
                this.tasks[this.tasks.length - 1].failed = true;
            }
        }
        this.onEvent();
    }

    /**
     * Registers the Listeners on the TaskFrame
     */
    private registerListeners(): void {
        if (this.taskFrame && this.taskFrame.contentDocument) {
            const reload = this.taskFrame.contentDocument.getElementById("recaptcha-reload-button");
            const verify = this.taskFrame.contentDocument.getElementById("recaptcha-verify-button");
            const audio = this.taskFrame.contentDocument.getElementById("recaptcha-audio-button");

            if (reload) {
                reload.addEventListener("click", this.onEvent.bind(this), { once: true });
            }
            if (verify) {
                verify.addEventListener("click", this.onVerify.bind(this), { once: true });
            }
            if (audio) {
                audio.addEventListener("click", this.onAudio.bind(this), { once: true });
            }
        }
    }


    /**
     * Stops the latest Task
     */
    private stopTask(): void {
        if (this.tasks.length != 0) {
            this.tasks[this.tasks.length - 1].stop();
            this.tasks[this.tasks.length - 1].solved = false;
        }
    }

}
