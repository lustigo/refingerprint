import Module from "../interfaces/Module";
import jsSHA from "jssha";

/**
 * Module which creates a Fingerprint based on the Web Audio API
 *
 * Freely adapted from Englehardt, Steven & Narayanan, Arvind. (2016). Online Tracking: A 1-million-site Measurement and Analysis. 1388-1401. 10.1145/2976749.2978313.
 */
export default class Audio implements Module {

    /**
    * Name of the module
    */
    public readonly name = "Audio";

    /**
     * The fingerprint generated by the Web Audio API
     */
    private fingerprint = "UNKNWON";

    /**
    * Converts an integer to a Hexadecimal value
    * @see https://stackoverflow.com/questions/39225161/convert-uint8array-into-hex-string-equivalent-in-node-js/39225475#39225475
    * @param i Number to convert
    */
    private i2hex(i: number): string {
        return ("0" + i.toString(16)).slice(-2);
    }

    /**
     * Computes the Hash of the Audio Data
     * @param data Audio Data
     * @returns SHA512 Hash in Base64
     */
    private computeHash(data: Uint8Array): string {
        const hasher = new jsSHA("SHA-512", "HEX");
        hasher.update(Array.from(data).map(this.i2hex).join(""));
        return hasher.getHash("B64");
    }


    /**
     * Generates the Fingerprint from the given AudioBuffer and stores it in this.fingerprint
     * @param event Generated Audiodata
     */
    private generateFingerprint(buff: AudioBuffer): void {
        const data = buff.getChannelData(0).buffer;
        const arr = new Uint8Array(data, 0, data.byteLength);
        this.fingerprint = this.computeHash(arr);
    }

    /**
     * Generates the Audio Data used as a Fingerprint
     */
    private generateAudioData(): void {
        // 3 Seconds with a bitrate of 44100bps
        const ctx = new OfflineAudioContext({
            numberOfChannels: 1,
            sampleRate: 44100,
            length: 44100 * 3,
        });

        const osi = ctx.createOscillator();
        const comp = ctx.createDynamicsCompressor();

        osi.start();

        osi.connect(comp);
        comp.connect(ctx.destination);

        ctx.startRendering()
            .then(buff => this.generateFingerprint(buff))
            .catch(e => {
                console.log(e);
                this.fingerprint = "ERROR";
            });
    }

    /**
    * Will be called, when the Captcha is rendered
    * Gets the Information and stores it.
    */
    public start(): void {
        this.generateAudioData();
    }

    /**
    * Will be called, when the Captcha is solved
    * does nothing
    */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public stop(): void { }

    /**
     * Returns the Fingerprint
     * 
     * @returns Fingerprint
     */
    public getCollectedData(): string {
        return this.fingerprint;
    }

}
