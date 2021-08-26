import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';

import { merge, Subscription } from 'rxjs';
import { debounceTime, distinctUntilChanged, skip } from 'rxjs/operators';
import { PaginatedTable } from 'src/app/models/paginated-table';

import { TimeEntry } from 'src/app/models/time-entry';
import { TimeService } from 'src/app/services/time/time.service';
import { TableComponent } from './table/table.component';

@Component({
  selector: 'app-time',
  templateUrl: './time.component.html',
  styleUrls: ['./time.component.scss']
})
export class TimeComponent implements OnInit, OnDestroy {
  @ViewChild('table') table: TableComponent;
  private filterSub: Subscription = new Subscription();
  showForm: boolean;
  data: TimeEntry[] = [];
  rowCount: number = 0;
  editing: TimeEntry;
  dateFrom: FormControl = new FormControl(null);
  dateTo: FormControl = new FormControl(null);

  get tableSort() {
    return this.table?.sort;
  }
  get tablePaginator() {
    return this.table?.paginator;
  }

  constructor(private readonly timeService: TimeService) {
    this.resetParams();
  }

  ngOnInit(): void {
    this.filterSub = merge(
      this.dateFrom.valueChanges.pipe(distinctUntilChanged(), debounceTime(200)),
      this.dateTo.valueChanges.pipe(distinctUntilChanged(), debounceTime(200)),
    )
    .subscribe(() => {
      this.get();
    });

    this.get();
  }

  ngOnDestroy(): void {
    this.filterSub.unsubscribe();
  }

  /**
   * Resets the table filter params
   */
  resetParams(): void {
    this.dateTo.setValue(new Date());

    const d = new Date();
    d.setDate(d.getDate() - 29);
    this.dateFrom.setValue(d);
  }

  /**
   * Gets time entries that a user has access to
   */
  get(): void {
    const dtF: Date = this.dateFrom.value;
    const dtT: Date = this.dateTo.value;
    
    if (!!dtF && !!dtT) {
      const { pageIndex, pageSize } = this.tablePaginator;
      const { active, direction } = this.tableSort;
      const sortDesc = direction !== 'asc';

      this.timeService.get(dtF, dtT, pageIndex, pageSize, active, sortDesc)
        .subscribe((res: PaginatedTable<TimeEntry>) => {
          this.data = res.page;
          this.rowCount = res.rowCount;
        });
    }
  }

  /**
   * Creates a new time entry with the selected values
   * 
   * @param data values for the new time entry
   */
  create(data: TimeEntry): void {
    this.timeService.create(data)
      .subscribe((res: TimeEntry) => {
        if (!!res) {
          this.data.push(res);
        }
      });
  }

  /**
   * Updates the selected time entry
   * 
   * @param id id of the item to update
   * @param newValues key-value pairs for the properties to update
   */
  update(id: number, newValues: any): void {
    if (id > 0) {
      this.timeService.update(id, newValues)
        .subscribe((entry: TimeEntry) => {
          if (!!entry) {
            const data = [...this.data];
            const index = data.findIndex(t => t.id === entry.id);
            if (index > -1) {
              data[index] = entry;
              this.data = [...data];
            }
          }
        });
    }
  }

  /**
   * Deletes the selected time entry
   * 
   * @param id id of the item to update
   */
  delete(id: number): void {
    if (id > 0) {
      this.timeService.delete(id)
        .subscribe((success: boolean) => {
          if (success) {
            const data = [...this.data];
            const index = data.findIndex(t => t.id === id);
            if (index > -1) {
              data.splice(index, 1);
              this.data = [...data];
            }
          }
        });
    }
  }

  resetView(): void {
    this.showForm = false;
    this.editing = null;
  }

}
