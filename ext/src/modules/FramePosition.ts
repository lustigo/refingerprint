import Module from "../interfaces/Module";
import Rectangle from "../interfaces/Rectangle";
import frameListener from "../misc/FrameListener";
import { delay } from "../misc/Helper";

/**
 * Module which stores the Position of the Frames
 */
export default class FramePosition implements Module {

    /**
     * Position of the Captcha Frame
     */
    private captchaPos: Rectangle = {
        x: 0,
        y: 0,
        width: 0,
        height: 0,
    };

    /**
     * Position of the Task Frame
     */
    private taskPos: Rectangle = {
        x: 0,
        y: 0,
        width: 0,
        height: 0,
    };

    /**
     * Name of the module
     */
    public readonly name = "FramePosition";

    /**
     * Will be called when the Captcha is rendered
     * Tries to get the positions
     */
    public start(): void {
        this.getCaptchaPos();
        this.getTaskPos();
    }

    /**
     * Gets the Position of the Captcha Frame
     */
    private getCaptchaPos(): void {
        const frame = frameListener.getCaptchaFrame();
        if (!frame) {
            delay(10).then(this.getCaptchaPos.bind(this));
            return;
        }
        const rect = frame.getBoundingClientRect();
        this.captchaPos.x = rect.x;
        this.captchaPos.y = rect.y;
        this.captchaPos.height = rect.height;
        this.captchaPos.width = rect.width;
    }

    /**
     * Gets the Position of the Task Frame
     */
    private getTaskPos(): void {
        const frame = frameListener.getTaskFrame();
        if (!frame) {
            delay(10).then(this.getTaskPos.bind(this));
            return;
        }
        const rect = frame.getBoundingClientRect();

        //Frame has y position of -9999 if the TaskFrame is not popped up
        if (rect.y <= 0) {
            delay(10).then(this.getTaskPos.bind(this));
            return;
        }
        this.taskPos.x = rect.x;
        this.taskPos.y = rect.y;
        this.taskPos.height = rect.height;
        this.taskPos.width = rect.width;
    }

    /**
     * Will be called when the Captcha is solved
     * does nothing
     */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public stop(): void { }

    /**
     * Returns the position of the Frames.
     */
    public getCollectedData(): Record<string, any> {
        return {
            captcha: this.captchaPos,
            task: this.taskPos
        };
    }
}
