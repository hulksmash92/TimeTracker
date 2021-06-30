import { AfterViewInit, Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';

import { TimeEntry } from 'src/app/models/time-entry';

@Component({
  selector: 'time-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.scss']
})
export class TableComponent implements AfterViewInit {
  @Input() set data(d: TimeEntry[]) {
    if (!d || d.length === 0) {
      this.dataSource.data = [];
    } else {
      this.dataSource.data = [...d];
    }
  }
  @Output() deleteItem: EventEmitter<number> = new EventEmitter<number>();
  @Output() updateItem: EventEmitter<TimeEntry> = new EventEmitter<TimeEntry>();
  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;
  dataSource: MatTableDataSource<TimeEntry> = new MatTableDataSource<TimeEntry>([]);
  displayedColumns: string[] = ['organisation', 'comments', 'created', 'updated', 'value', 'menu'];

  ngAfterViewInit(): void {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

}
