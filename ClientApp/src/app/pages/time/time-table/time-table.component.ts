import { AfterViewInit, Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';

import { TimeEntry } from 'src/app/models/time-entry';

@Component({
  selector: 'time-table',
  templateUrl: './time-table.component.html',
  styleUrls: ['./time-table.component.scss']
})
export class TimeTableComponent implements AfterViewInit {
  @Input() set data(d: TimeEntry[]) {
    if (!d || d.length === 0) {
      this.dataSource = [];
    } else {
      this.dataSource = [...d];
    }
  }
  @Input() set rowCount(v: number) {
    this.count = v ?? this.dataSource.length;
  }
  @Output() deleteItem: EventEmitter<number> = new EventEmitter<number>();
  @Output() updateItem: EventEmitter<TimeEntry> = new EventEmitter<TimeEntry>();
  @Output() sortChange: EventEmitter<void> = new EventEmitter<void>();
  @Output() pageChange: EventEmitter<void> = new EventEmitter<void>();

  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;
  dataSource: TimeEntry[] = [];
  count: number = 0;
  displayedColumns: string[] = ['organisation', 'comments', 'created', 'updated', 'value', 'menu'];

  ngAfterViewInit(): void {
    this.paginator.page.subscribe({
      next: () => this.pageChange.emit()
    });

    this.sort.sortChange.subscribe({
      next: () => {
        if (this.paginator.pageIndex !== 0) {
          this.paginator.firstPage();
        } else {
          this.sortChange.emit();
        }
      }
    });
  }

}
