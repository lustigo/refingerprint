/**
 * Describes the structure of the Checkbox-Widget
 */
export default interface CheckboxDescription {
    readonly question: string;
    readonly selection: string[]; // Labels of the buttons
    checked: boolean[];
    // eslint-disable-next-line semi
}
