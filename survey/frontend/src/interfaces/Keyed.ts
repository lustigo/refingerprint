/**
 * Array with an string index
 */
export default interface Keyed<T> {
    [key: string]: T;
    // eslint-disable-next-line semi
}
