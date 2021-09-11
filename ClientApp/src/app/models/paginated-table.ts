/**
 * Describes the returned details for a table with server-side 
 * sorting, filtering and pagination enabled
 */
export class PaginatedTable<T> {
    /**
     * The page of data
     */
    page: T[] = [];

    /**
     * Total available rows
     */
    rowCount: number = 0;
}
