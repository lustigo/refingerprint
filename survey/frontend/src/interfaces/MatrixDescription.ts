import Keyed from "./Keyed";

export default interface MatrixDescription {
  readonly question: string;
  readonly structure: string[]; // Attributes that will be displayed on a column
  answers: Keyed<number>;
}
