import { Component, Input, OnInit } from '@angular/core';
import { TimeEntry } from 'src/app/models/time-entry';

@Component({
  selector: 'app-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.scss']
})
export class TableComponent implements OnInit {
  @Input() data: TimeEntry[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
