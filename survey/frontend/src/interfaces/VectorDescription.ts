/**
 * Describes the structure of the Vector-Widget
 */
export default interface VectorDescription {
    readonly question: string;
    readonly selection: string[]; // Labels of the buttons
    answer: number;
    // eslint-disable-next-line semi
}
