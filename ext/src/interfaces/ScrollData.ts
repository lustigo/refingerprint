import MouseData from "./MouseData";

/**
 * Represents a Wheel Event
 */
export default interface ScrollData extends MouseData {
    deltaX: number;
    deltaY: number;
    deltaZ: number;
    deltaMode: number;
    // eslint-disable-next-line semi
}
