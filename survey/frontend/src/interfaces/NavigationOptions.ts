/**
 * Options necessary for the navigation
 */
export default interface NavigationOptions {
    currentPage: number;
    canBack: boolean;
    canForward: boolean;
    canSend: boolean;
    pageAmount: number;
    // eslint-disable-next-line semi
}
