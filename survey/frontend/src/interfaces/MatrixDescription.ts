import Keyed from "./Keyed";
/**
 * Description of the structure of the Matrixwidget
 */
export default interface MatrixDescription {
    readonly question: string;
    readonly structure: string[]; // Attributes that will be displayed on a column
    answers: Keyed<number>;
    // eslint-disable-next-line semi
}
