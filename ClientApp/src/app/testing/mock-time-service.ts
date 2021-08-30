import { Observable, of } from 'rxjs';

import { PaginatedTable } from 'src/app/models/paginated-table';
import { Tag } from 'src/app/models/tag';
import { TimeEntry } from 'src/app/models/time-entry';

export class MockTimeService {
  readonly API_URL: string = '/api/time';

  /**
   * Gets all tags for adding onto time entries
   */
  getTags(): Observable<Tag[]> {
    return of([]);
  }

  /**
   * Gets a page of time entries for the selected conditions
   * 
   * @param from date range start to filter on
   * @param to date range end to filter on
   * @param pageIndex index of the current page (starts at 0)
   * @param pageSize number of records per page
   * @param sort name of the column to sort by
   * @param sortDesc whether to sort descending or not
   * 
   * @returns array of the time entries
   */
  get(from: Date, to: Date, pageIndex: number, pageSize: number, sort: string, sortDesc: boolean): Observable<PaginatedTable<TimeEntry>> {
    return of({
        page: [],
        rowCount: 0
    });
  }

  /**
   * Creates a time entry for the logged in user from the passed in value
   * 
   * @param data values for the new time entry
   * 
   * @returns the newly created time entry
   */
  create(data: TimeEntry): Observable<TimeEntry> {
    return of(null);
  }

  /**
   * Patches the selected properties on the selected time entry
   * 
   * @param id id of the time entry to update
   * @param data values to update
   * 
   * @returns the full new time entry value
   */
  update(id: number, data: any): Observable<TimeEntry> {
    return of(null);
  }

  /**
   * Deletes the selected time entry 
   * 
   * @param id time entry id to delete
   * 
   * @returns true or false if successful
   */
  delete(id: number): Observable<boolean> {
    return of(false);
  }

}
