import { delay } from "../misc/Helper";
import frameListener from "../misc/FrameListener";
import MouseListener from "./MouseListener";
import ScrollData from "../interfaces/ScrollData";

/**
 * Module which collects Scrolling
 */
export default class ScrollListener extends MouseListener {
    /**
     * Name of the Module
     */
    public name = "ScrollEvents";

    /**
     * ScrollEvents
     */
    private events: Array<ScrollData> = [];

    /**
     * will be called when the Captcha is rendered
     * registers the Listeners
     */
    public start(): void {
        this.removeHandler.push(this.registerWheelListener(this.onWheel.bind(this)));

        this.registerOnFrames(this.onWheel.bind(this));
    }

    /**
     * will be called when the Captcha is solved
     * stops the tracking
     */
    public stop(): void {
        super.stopTracking();
    }

    /**
     * Registers the Listeners on the Frames
     * @param listener Listener to register
     */
    private async registerOnFrames(listener: (ev: WheelEvent) => any): Promise<void> {
        const t = true;

        while (t) {
            await delay(200);
            const captchaFrame = frameListener.getCaptchaFrame();
            if (captchaFrame && captchaFrame.contentWindow) {
                captchaFrame.contentWindow.addEventListener("wheel", listener);
                this.removeHandler.push((): void => { if (captchaFrame && captchaFrame.contentWindow) { captchaFrame.contentWindow.removeEventListener("wheel", listener); } });
                break;
            }
        }

        while (t) {
            await delay(200);
            const taskFrame = frameListener.getTaskFrame();
            if (taskFrame && taskFrame.contentWindow) {
                taskFrame.contentWindow.addEventListener("wheel", listener);
                this.removeHandler.push((): void => { if (taskFrame && taskFrame.contentWindow) { taskFrame.contentWindow.removeEventListener("wheel", listener); } });
                break;
            }
        }
    }

    /**
     * Registers the Listener to the Window
     * @param listener: Listener to call
     * @returns Method to remove the Listener
     */
    private registerWheelListener(listener: (ev: WheelEvent) => void): () => void {
        window.addEventListener("wheel", listener);
        return (): void => {
            window.removeEventListener("wheel", listener);
        };
    }

    /**
     * Will be called when the Scroll Event happens
     * Adds it to the Eventslist
     * @param ev  Scroll Event to add
     */
    private onWheel(ev: WheelEvent): void {
        this.events.push({
            x: ev.screenX,
            y: ev.screenY,
            time: new Date().getTime(),
            deltaMode: ev.deltaMode,
            deltaX: ev.deltaX,
            deltaY: ev.deltaY,
            deltaZ: ev.deltaZ,
        });
    }

    /**
     * returns the collected ScrollEvents
     * @returns ScrollData
     */
    public getCollectedData(): Array<ScrollData> {
        return this.events;
    }
    // eslint-disable-next-line semi
}
